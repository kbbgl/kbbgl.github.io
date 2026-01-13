---
title: kubectl Cheatsheet
slug: kubectl-cheatsheet
app2or: kgal-akl
tags: [devops, k8s, kubernetes, kubectl, cheatsheet]
---

## `kubectl`

`kubectl` is used to configure and manage the cluster.

It will use the `~/.kube/config` file as the configuration file which includes the cluster IP, the credentials and the context.

A **context** is a combination of a cluster and user credentials.
The context can be changed using:

```bash
kubectl config use-context $SOME_CONTEXT
```

## List Cluster Roles Service Accounts

```bash
kubectl get rolebindings,clusterrolebindings -A -o custom-columns='KIND:kind,NAMESPACE:metadata.namespace,NAME:metadata.name,SERVICE_ACCOUNTS:subjects[?(@.kind=="ServiceAccount")].name'
```
```
KIND                 NAMESPACE       NAME                                                            SERVICE_ACCOUNTS
RoleBinding          ingress-nginx   ingress-nginx                                                   ingress-nginx
RoleBinding          kube-public     system:controller:bootstrap-signer                              bootstrap-signer
RoleBinding          kube-system     k3s-cloud-controller-manager-authentication-reader              <none>
RoleBinding          kube-system     metrics-server-auth-reader                                      metrics-server
RoleBinding          kube-system     system::extension-apiserver-authentication-reader               <none>
RoleBinding          kube-system     system::leader-locking-kube-controller-manager                  kube-controller-manager,leader-election-controller
RoleBinding          kube-system     system::leader-locking-kube-scheduler                           kube-scheduler
RoleBinding          kube-system     system:controller:bootstrap-signer                              bootstrap-signer
RoleBinding          kube-system     system:controller:cloud-provider                                cloud-provider
RoleBinding          kube-system     system:controller:token-cleaner                                 token-cleaner
ClusterRoleBinding   <none>          cluster-admin                                                   <none>
ClusterRoleBinding   <none>          clustercidrs-node                                               <none>
```

## Check Permissions for User

```bash
kubectl auth can-i create secrets -n default --as=system:serviceaccount:default:app
yes

kubectl auth can-i list secrets -n default --as=system:serviceaccount:default:app
yes
```