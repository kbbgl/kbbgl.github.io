---
title: How to Create an EKS Cluster using eksctl
slug: how-to-create-eks-cluster-ekctl
app2or: kgal-akl
tags: [devops, k8s, kubernetes, kubectl, eksctl, eks, iam, aws]
---


## Log Into AWS Using CLI

```bash
aws sso login --profile $AWS_PROFILE
```

## Create Cluster
```bash
eksctl create cluster \
--name $EKS_CLUSTER_NAME \
--region $AWS_REGION \
--nodegroup-name "kbbgl-nodegroup" \
--node-type "t3.medium" \
--nodes 1 \
--managed
```


## Enable CloudWatch Logging on the EKS for authentication events

```bash
aws eks update-cluster-config \
--name $EKS_CLUSTER_NAME \
--region $AWS_REGION \
--logging '{"clusterLogging":[{"types":["authenticator"],"enabled":true}]}'
```


## Interacting With Cluster

In cases when we have an existing EKS cluster that we want to interact with once the AWS access token has expired, we need to first reauthenticate with AWS:

```bash
aws sso login --profile $AWS_PROFILE
```

Once we're authenticated, we can run the following command to let `kubectl` know that it needs to retrieve the cluster access token using the AWS CLI:

```bash
aws eks update-kubeconfig --region $AWS_REGION --name $EKS_CLUSTER_NAME --profile $AWS_PROFILE
```

This will update the `.kubeconfig` file to include the necessary AWS command to retrieve the secret and allow interaction with the EKS cluster:
```yaml
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS[redacted]
    server: https://abcdefg.hi2.$AWS_REGION.eks.amazonaws.com
  name: arn:aws:eks:$AWS_REGION:$AWS_ACCOUNT_ID:cluster/$EKS_CLUSTER_NAME
contexts:
- context:
    cluster: arn:aws:eks:$AWS_REGION:$AWS_ACCOUNT_ID:cluster/$EKS_CLUSTER_NAME
    user: arn:aws:eks:$AWS_REGION:$AWS_ACCOUNT_ID:cluster/$EKS_CLUSTER_NAME
  name: arn:aws:eks:$AWS_REGION:$AWS_ACCOUNT_ID:cluster/$EKS_CLUSTER_NAME
current-context: arn:aws:eks:$AWS_REGION:$AWS_ACCOUNT_ID:cluster/$EKS_CLUSTER_NAME
kind: Config
users:
- name: arn:aws:eks:$AWS_REGION:$AWS_ACCOUNT_ID:cluster/$EKS_CLUSTER_NAME
  user:
    exec:
      apiVersion: client.authentication.k8s.io/v1beta1
      args:
      - --region
      - $AWS_REGION
      - eks
      - get-token
      - --cluster-name
      - $EKS_CLUSTER_NAME
      - --profile
      - $AWS_PROFILE
      command: aws
```

We can then run `kubectl` commands.

## Deploying AWS Load Balancer in EKS

First thing we need to do is to install the AWS Load Balancer Controller in the cluster:

```bash
helm repo add eks https://aws.github.io/eks-charts
helm repo update
helm install aws-load-balancer-controller eks/aws-load-balancer-controller \
-n kube-system \
-f values.yaml
```

The helm chart will deploy a `Pod` and a `ServiceAccount` named `aws-lb-controller` in the `kube-system` namespace (among other things). Per the [AWS documentation](https://kubernetes-sigs.github.io/aws-load-balancer-controller/latest/deploy/installation/#configure-iam):

> The controller runs on the worker nodes, so it needs access to the AWS ALB/NLB APIs with IAM permissions.
The IAM permissions can either be setup using IAM roles for service accounts (IRSA), Pod Identity, or can be attached directly to the worker node IAM roles. The best practice is using IRSA if you're using Amazon EKS.

To create the AWS Load Balancer Controller IAM policy, we need to download the official policy from [GitHub](https://github.com/kubernetes-sigs/aws-load-balancer-controller) and apply it:

```bash
curl -o /tmp/iam_policy.json https://raw.githubusercontent.com/kubernetes-sigs/aws-load-balancer-controller/v3.0.0/docs/install/iam_policy.json

aws iam create-policy \
--policy-name AWSLoadBalancerControllerIAMPolicy \
--policy-document file:///tmp/iam_policy.json \
--profile $AWS_PROFILE
```

The response will contain the `PolicyArn` that we will need later.

Now that we have the AWS Load Balancer Controller IAM policy, we can create the IAM role and trust policy.

We first need to retrieve the cluster OIDC issuer, AWS account ID an set the AWS Load Balancer Kubernetes `ServiceAccount` and namespace:

```bash
OIDC_ISSUER=$(aws eks describe-cluster --name $EKS_CLUSTER_NAME --query "cluster.identity.oidc.issuer" --output text --profile $AWS_PROFILE)
# e.g. https://oidc.eks.us-east-1.amazonaws.com/id/EXAMPLED539D4633E53DE1B716D3041E

ACCOUNT_ID=$(aws sts get-caller-identity --query Account --output text --profile $AWS_PROFILE)
OIDC_ID=$(echo $OIDC_ISSUER | sed -e 's|https://oidc.eks.'"$AWS_REGION"'.amazonaws.com/id/||')
# EXAMPLED539D4633E53DE1B716D3041E
SA_NAME=aws-lb-controller
NAMESPACE=kube-system

cat > /tmp/trust-policy.json << EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Federated": "arn:aws:iam::${ACCOUNT_ID}:oidc-provider/oidc.eks.${AWS_REGION}.amazonaws.com/id/${OIDC_ID}"
      },
      "Action": "sts:AssumeRoleWithWebIdentity",
      "Condition": {
        "StringEquals": {
          "oidc.eks.${AWS_REGION}.amazonaws.com/id/${OIDC_ID}:sub": "system:serviceaccount:${NAMESPACE}:${SA_NAME}",
          "oidc.eks.${AWS_REGION}.amazonaws.com/id/${OIDC_ID}:aud": "sts.amazonaws.com"
        }
      }
    }
  ]
}
EOF
```

And we create the role:

```bash
aws iam create-role \
--role-name eks-${EKS_CLUSTER_NAME}-aws-lb-controller \
--assume-role-policy-document file:///tmp/trust-policy.json \
--description "IRSA for AWS Load Balancer Controller on $EKS_CLUSTER_NAME" \
--profile $AWS_PROFILE
  ```

And attach the policy to the role:

```bash
aws iam attach-role-policy \
--role-name eks-${EKS_CLUSTER_NAME}-aws-lb-controller \
--policy-arn arn:aws:iam::${ACCOUNT_ID}:policy/AWSLoadBalancerControllerIAMPolicy \
--profile $AWS_PROFILE
```

We then annotate the AWS Load Balancer Kubernetes `ServiceAccount` with new role ARN:

```bash
ROLE_ARN=$(aws iam get-role --role-name eks-$EKS_CLUSTER_NAME-aws-lb-controller --query "Role.Arn" --output text --profile $AWS_PROFILE)

kubectl annotate serviceaccounts -n $NAMESPACE aws-lb-controller eks.amazonaws.com/role-arn=$ROLE_ARN --overwrite
```

Restart the controller to pick up the new role:

```bash
kubectl rollout restart deployment -n $NAMESPACE -l app.kubernetes.io/name=aws-load-balancer-controller
```

