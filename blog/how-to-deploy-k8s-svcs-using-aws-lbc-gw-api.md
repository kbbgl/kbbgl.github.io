---
slug: how-to-deploy-k8s-svcs-using-aws-lbc-gw-api
title: Hot to Deploy Kubernetes Services using Gateway API/AWS Load Balancer Controller
description: Hot to Deploy Kubernetes Services using Gateway API/AWS Load Balancer Controller
authors: [kgal-akl]
tags: [k8s, kubernetes, gateway, api, aws, load_balancer]
---

This tutorial contains a working example of exposing **TCP services** (LDAP/LDAPS + SSH) from a **single-node k3s** cluster running on an EC2 instance, using:

- **Kubernetes Gateway API**
- **AWS Load Balancer Controller (LBC)** for:
  - **NLB** (L4) via `TCPRoute`
  - **ALB** (L7) via `HTTPRoute`/`GRPCRoute` (example file included)

The key implementation detail for k3s-on-EC2 with the default overlay networking (flannel): **use `instance` targets + NodePorts** for L4 routes. ClusterIP + pod IP targets won’t work unless pods are VPC-routable (AWS VPC CNI).

## Environment

- **Node OS**: Ubuntu 24.04.4 LTS
- **Kernel**: 6.14.0-1018-aws
- **k3s / server**: v1.33.4+k3s1
- **containerd**: 2.0.5-k3s2
- **kubectl client**: v1.35.0 (warning: > +/-1 minor skew vs server)
- **Gateway API**: 1.3.0
- **AWS Load Balancer Controller**
  - **Helm chart**: `aws-load-balancer-controller-3.0.0` (app version v3.0.0)
  - **Controller image**: `public.ecr.aws/eks/aws-load-balancer-controller:v2.17.0`
  - **Flags**:
    - `--feature-gates=ALBGatewayAPI=true,NLBGatewayAPI=true`
    - `--enable-manage-backend-security-group-rules=true`

## Prerequisites

- An EC2 instance with k3s installed and reachable via SSH (example host alias: `$DOMAIN_NAME`)
- AWS credentials able to create/modify:
  - ELBv2 load balancers/listeners/target groups
  - EC2 security groups and tags
  - (plus whatever IAM is required by the controller)
- `kubectl`, `helm`, `aws` CLI installed locally

## Files

Gateway API standard CRDs:
```bash
wget --output-document standard_crds_1.3.0.yaml https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.3.0/standard-install.yam
```

Gateway API experimental CRDs (needed for some L4 routes depending on version)
```bash
wget --output-document experimental_crds_1.3.0.yaml https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.3.0/experimental-install.yaml
```

AWS LBC Gateway API CRDs (`gateway.k8s.aws/*`):
```bash
wget --output-document aws_lbc_gateway-crds_3.0.yaml https://raw.githubusercontent.com/kubernetes-sigs/aws-load-balancer-controller/refs/heads/release-3.0/config/crd/gateway/gateway-crds.yaml
```

Helm values used to install/upgrade AWS LBC:
```bash
helm show values eks/aws-load-balancer-controller > aws-lbc.yaml
```

`GatewayClasses` definitions for both the ALB and NLB:
```yaml title="00-gatewayclasses.yaml"
apiVersion: "gateway.networking.k8s.io/v1"
kind: GatewayClass
metadata: 
  name: aws-lbc-alb
spec:
  controllerName: gateway.k8s.aws/alb
  parametersRef:
    group: gateway.k8s.aws
    kind: LoadBalancerConfiguration
    name: alb-public-config
    namespace: default
---
apiVersion: "gateway.networking.k8s.io/v1"
kind: GatewayClass
metadata: 
  name: aws-lbc-nlb
spec:
  controllerName: gateway.k8s.aws/nlb
  parametersRef:
    group: gateway.k8s.aws
    kind: LoadBalancerConfiguration
    name: nlb-public-config
    namespace: default
```

`LoadBalancerConfiguration` for public ALB and NLB:

```yaml title="12-lb-public-lbc.yaml"
apiVersion: gateway.k8s.aws/v1beta1
kind: LoadBalancerConfiguration
metadata:
  name: alb-public-config
  namespace: default
spec:
  scheme: internet-facing
  loadBalancerSubnets:
    - identifier: subnet-1
    - identifier: subnet-2
---
apiVersion: gateway.k8s.aws/v1beta1
kind: LoadBalancerConfiguration
metadata:
  name: nlb-public-config
  namespace: default
spec:
  scheme: internet-facing
  loadBalancerAttributes:
    - key: load_balancing.cross_zone.enabled
      value: "true"
  loadBalancerSubnets:
    - identifier: subnet-3
    - identifier: subnet-4
    - identifier: subnet-5
```

ALB HTTP 443 `Gateway`:
```yaml title="10-gateway-alb.yaml"
apiVersion: "gateway.networking.k8s.io/v1"
kind: Gateway
metadata:
  name: public-alb-gw
  namespace: default
spec:
  gatewayClassName: aws-lbc-alb
  listeners:
    - name: https
      protocol: HTTPS
      port: 443
      hostname: "*.$DOMAIN_NAME"
      tls:
        mode: Terminate
        certificateRefs:
          - kind: Secret
            group: ""
            name: tls-$DOMAIN_NAME-crt
      allowedRoutes:
        namespaces:
          from: All
```

:::note

You need to make sure that there's a TLS secret named `tls-$DOMAIN_NAME-crt` as referenced by the Gateway that has SAN that includes the hostname provided above. To create it:

```bash
kubectl create secret tls tls-$DOMAIN_NAME-crt --cert=$DOMAIN_NAME.crt.pem --key=deploy/ingress/certs/$DOMAIN_NAME.key.pem -n default
```

:::

NLB TCP 636 and 2222 `Gateway`:
```yaml title="11-gateway-nlb.yaml"
apiVersion: "gateway.networking.k8s.io/v1"
kind: Gateway
metadata:
  name: public-nlb-gw
  namespace: default
spec:
  gatewayClassName: aws-lbc-nlb
  listeners:
    - name: ldaps
      protocol: TCP
      port: 636
      allowedRoutes:
        namespaces:
          from: All

    - name: ssh
      protocol: TCP
      port: 2222
      allowedRoutes:
        namespaces:
          from: All
```

LDAP `TCPRoute`:

```yaml title="20-route-ldap-tcp.yaml"
apiVersion: "gateway.networking.k8s.io/v1alpha2"
kind: TCPRoute
metadata:
  name: ldap-ldaps-route
  namespace: ldap
spec:
  parentRefs:
    - name: public-nlb-gw
      namespace: default
      sectionName: ldaps
  rules:
    - backendRefs:
      - name: ldap
        port: 636
```

SSH `TCPRoute`:
```yaml title="21-route-ssh-tcp.yaml"
apiVersion: "gateway.networking.k8s.io/v1alpha2"
kind: TCPRoute
metadata:
  name: ssh-tcp-route
  namespace: ssh
spec:
  parentRefs:
    - name: public-nlb-gw
      namespace: default
      sectionName: ssh
  rules:
    - backendRefs:
        - name: ssh
          port: 2222
```

Example `HTTPRoute`:
```yaml title="30-route-http-example.yaml"
# 30-route-http-example.yaml (optional example)
# Replace namespace/service/port with your actual HTTP app
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: example-web-route
  namespace: default
spec:
  parentRefs:
    - name: public-alb-gw
      namespace: default
      sectionName: https
  hostnames:
    - app.$DOMAIN_NAME
  rules:
    - backendRefs:
        - name: example-web-svc
          port: 80
```

## Deployment

### 1. Set Convenience Variables

```bash
AWS_PROFILE=dev
AWS_REGION=$AWS_REGION
```

### 2. Install Gateway API CRDs and AWS LBC Gateway CRDs

Apply the CRDs from this directory:

```bash
kubectl apply -f standard_crds_1.3.0.yaml
kubectl apply -f experimental_crds_1.3.0.yaml
kubectl apply -f aws_lbc_gateway-crds_3.0.yaml
```

Verify the `TCPRoute` CRD is present and served:

```bash
kubectl get crd tcproutes.gateway.networking.k8s.io
kubectl get crd tcproutes.gateway.networking.k8s.io -o jsonpath='{.spec.versions[*].name}{"\n"}'
```

### 3. Install / upgrade AWS Load Balancer Controller (Helm)

This repo uses the EKS chart with explicit values in `aws-lbc.yaml`.

```bash
helm repo add eks https://aws.github.io/eks-charts
helm repo update

helm upgrade --install aws-lbc eks/aws-load-balancer-controller \
  -n kube-system \
  -f aws-lbc.yaml \
  --version 3.0.0
```

Confirm the controller flags include the Gateway API feature gates:

```bash
kubectl  -n kube-system get deploy aws-lbc-aws-load-balancer-controller \
  -o jsonpath='{.spec.template.spec.containers[0].args}{"\n"}'
```

You should see:

- `--feature-gates=ALBGatewayAPI=true,NLBGatewayAPI=true`

For k3s/flannel + NLB instance targets, it’s helpful to also run:

- `--enable-manage-backend-security-group-rules=true`

### 4. Apply LBC `LoadBalancerConfiguration`, `GatewayClass`, and `Gateway`

```bash
kubectl apply -f 12-lb-public-lbc.yaml
kubectl apply -f 00-gatewayclasses.yaml
kubectl apply -f 11-gateway-nlb.yaml
```

Important bits from `11-gateway-nlb.yaml` (listeners define the NLB ports):

```yaml
spec:
  gatewayClassName: aws-lbc-nlb
  listeners:
    - name: ldaps
      protocol: TCP
      port: 636
    - name: ssh
      protocol: TCP
      port: 2222
```

Verify the NLB Gateway gets an address:

```bash
kubectl  -n default get gateway public-nlb-gw -o wide
kubectl  -n default get gateway public-nlb-gw -o yaml | rg -n "addresses:|value:|listeners:|attachedRoutes" -n
```

### 5. Deploy LDAP + SSH Workloads and `NodePort` Services

For L4 NLB on k3s/flannel, **Services must be `NodePort`**, because the controller will create **instance** target groups that point to the node’s NodePorts.

Important snippets (from `../ldap/service.yaml` and `../ssh/service.yaml` in this repo):

```yaml
spec:
  type: NodePort
  ports:
  - name: ldaps
    port: 636
    targetPort: 636
    nodePort: 30636
```

```yaml
spec:
  type: NodePort
  ports:
  - name: ssh
    port: 2222
    targetPort: 2222
    nodePort: 32222
```

Apply your workloads/services (paths in this repo):

```bash
kubectl  apply -f ../ldap/deployment.yaml
kubectl  apply -f ../ldap/service.yaml

kubectl  apply -f ../ssh/namespace.yaml
kubectl  apply -f ../ssh/configmap.yaml
kubectl  apply -f ../ssh/deployment.yaml
kubectl  apply -f ../ssh/service.yaml
```

Verify endpoints exist:

```bash
kubectl  -n ldap get endpointslices -l kubernetes.io/service-name=ldap
kubectl  -n ssh  get endpointslices -l kubernetes.io/service-name=ssh
```

### 6. Create `TCPRoute`s to attach Services to Gateway listeners

```bash
kubectl  apply -f 20-route-ldap-tcp.yaml
kubectl  apply -f 21-route-ssh-tcp.yaml
```

Important route bits:

```yaml
spec:
  parentRefs:
    - name: public-nlb-gw
      namespace: default
      sectionName: ldaps   # matches listener name
  rules:
    - backendRefs:
      - name: ldap
        port: 636
```

Check route status:

```bash
kubectl  -n ldap get tcproute ldap-ldaps-route -o yaml
kubectl  -n ssh  get tcproute ssh-tcp-route -o yaml
```

You want `status.parents[*].conditions` to include `Accepted=True` and `ResolvedRefs=True`.

### 7. Verify AWS Resources (listeners + target groups + health)

Get the NLB hostname from the Gateway:

```bash
kubectl  -n default get gateway public-nlb-gw -o jsonpath='{.status.addresses[0].value}{"\n"}'
```

Resolve and test from your machine:

```bash
NLB_HOST=$(kubectl -n default get gateway public-nlb-gw -o jsonpath='{.status.addresses[0].value}')
dig +short "$NLB_HOST"
nc -G 2 -vz "$NLB_HOST" 636
nc -G 2 -vz "$NLB_HOST" 2222
```

Confirm listeners exist in AWS:

```bash
aws elbv2 describe-listeners \
  --load-balancer-arn "arn:aws:elasticloadbalancing:$AWS_REGION:$AWS_ACCOUNT_ID:loadbalancer/net/<name>/<id>" \
  --profile "$AWS_PROFILE" --region "$AWS_REGION" --no-cli-pager
```

List target groups and check target health:

```bash
aws elbv2 describe-target-groups \
  --load-balancer-arn "arn:aws:elasticloadbalancing:$AWS_REGION:$AWS_ACCOUNT_ID:loadbalancer/net/<name>/<id>" \
  --profile "$AWS_PROFILE" --region "$AWS_REGION" --no-cli-pager

aws elbv2 describe-target-health \
  --target-group-arn "<target-group-arn>" \
  --profile "$AWS_PROFILE" --region "$AWS_REGION" --no-cli-pager
```

### 8. Validate `NodePorts` locally on the EC2 node (via SSH)

This is a fast way to separate “Kubernetes service routing” from “AWS/NLB connectivity”.

```bash
ssh $DOMAIN_NAME 'nc -vz -w 2 127.0.0.1 30636; nc -vz -w 2 127.0.0.1 32222'
```

### 9. (Recommended) NLB cross-zone + subnet/AZ notes for single-node clusters

For a **single node in one AZ**, an internet-facing NLB in multiple AZs can resolve to IPs in AZs that have **no healthy targets**. You have two options:

- **Enable cross-zone load balancing** (recommended)
- Or constrain the NLB to only the node’s subnet/AZ

This repo enables cross-zone in `12-lb-public-lbc.yaml`:

```yaml
spec:
  loadBalancerAttributes:
    - key: load_balancing.cross_zone.enabled
      value: "true"
```

Verify in AWS:

```bash
aws elbv2 describe-load-balancer-attributes \
  --load-balancer-arn "<nlb-arn>" \
  --profile "$AWS_PROFILE" --region "$AWS_REGION" --no-cli-pager \
  --query 'Attributes[?Key==`load_balancing.cross_zone.enabled`]'
```

## Troubleshooting

During the deployment, I found the following commands helpful to use troubleshoot what was wrong with the deployment:

### A) Gateway has an address, but NLB has **no listeners**

Check controller logs; the exact message that pointed to the root cause was:

```bash
kubectl  -n kube-system logs deploy/aws-lbc-aws-load-balancer-controller --since=2h | rg -n "Skipping listener creation|public-nlb-gw"
```

If you see:

- `Skipping listener creation due to no backend references`

…it usually means the controller couldn’t materialize a usable backend (commonly: ClusterIP-only backends on overlay networking).

### B) Confirm `TCPRoute` attachment and backend resolution

```bash
kubectl  -n ldap get tcproute ldap-ldaps-route -o yaml
kubectl  -n ssh  get tcproute ssh-tcp-route -o yaml
kubectl  -n default describe gateway public-nlb-gw
```

### C) Verify Services/EndpointSlices

```bash
kubectl  -n ldap get svc ldap -o yaml
kubectl  -n ssh  get svc ssh  -o yaml
kubectl  -n ldap get endpointslices -l kubernetes.io/service-name=ldap -o yaml
kubectl  -n ssh  get endpointslices -l kubernetes.io/service-name=ssh  -o yaml
```

### D) Inspect AWS listeners / target groups / target health

```bash
aws elbv2 describe-listeners --load-balancer-arn "<nlb-arn>" --profile "$AWS_PROFILE" --region "$AWS_REGION" --no-cli-pager
aws elbv2 describe-target-groups --load-balancer-arn "<nlb-arn>" --profile "$AWS_PROFILE" --region "$AWS_REGION" --no-cli-pager
aws elbv2 describe-target-health --target-group-arn "<tg-arn>" --profile "$AWS_PROFILE" --region "$AWS_REGION" --no-cli-pager
```

### E) Security groups (NodePort/instance targets)

For `instance` targets, NLB health checks hit the **node’s NodePort**. If the node SG blocks that port, targets become unhealthy and connections time out.

Commands to inspect SG rules:

```bash
aws ec2 describe-security-groups --group-ids "<node-sg>" --profile "$AWS_PROFILE" --region "$AWS_REGION" --no-cli-pager
```

To allow the NLB/backend SG to reach NodePorts (example):

```bash
aws ec2 authorize-security-group-ingress \
  --group-id "<node-sg>" \
  --ip-permissions \
    'IpProtocol=tcp,FromPort=30636,ToPort=30636,UserIdGroupPairs=[{GroupId=<backend-sg>}]' \
    'IpProtocol=tcp,FromPort=32222,ToPort=32222,UserIdGroupPairs=[{GroupId=<backend-sg>}]' \
  --profile "$AWS_PROFILE" --region "$AWS_REGION" --no-cli-pager
```

Also, the controller’s worker-node SG selection expects a SG tagged like:

- `kubernetes.io/cluster/<cluster-name> = owned|shared`

### F) Subnets / AZ mismatch

If the node is in `$AWS_REGIONd` but the NLB only has subnets in `1b/1c`, targets can’t be registered/used correctly.

```bash
kubectl  get nodes -o wide
kubectl  get nodes -o jsonpath='{range .items[*]}{.metadata.name}{"\t"}{.spec.providerID}{"\n"}{end}'

aws ec2 describe-instances --instance-ids "<i-...>" --profile "$AWS_PROFILE" --region "$AWS_REGION" --no-cli-pager
aws ec2 describe-subnets --subnet-ids "<subnet-1>" "<subnet-2>" --profile "$AWS_PROFILE" --region "$AWS_REGION" --no-cli-pager

aws elbv2 describe-load-balancers --load-balancer-arns "<nlb-arn>" --profile "$AWS_PROFILE" --region "$AWS_REGION" --no-cli-pager \
  --query 'LoadBalancers[].AvailabilityZones[].{Zone:ZoneName,SubnetId:SubnetId}'
```

### G) Multi-AZ NLB resolves to multiple IPs (some work, some time out)

```bash
dig +short "$NLB_HOST"
for ip in $(dig +short "$NLB_HOST"); do
  nc -G 2 -vz "$ip" 2222 || true
done
```

Fix: enable cross-zone (`load_balancing.cross_zone.enabled=true`) or restrict subnets.

### Notes / gotchas

- **ClusterIP services + overlay pod IPs**: You’ll often get `Accepted=True` on routes, but the controller still can’t create listeners/targets (or will skip them) because it can’t reach pod IPs from the NLB.
- **kubectl version skew**: in this environment, client is newer than server by >1 minor. It still works for these resources, but keep it in mind if you hit strange client-side issues.

