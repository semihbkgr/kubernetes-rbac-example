# Kubernetes RBAC Example

[RBAC](./rbac.yaml)

docker image `semihbkgr/kubernetes-rbac-example:1.0.0`

apply rbac:

```shell
kubectl apply -f rbac.yaml
```

to create pod in kubernetes cluster:

```shell
kubectl apply -f pod.yaml
```

to get logs:

```shell
kubectl logs kubernetes-rbac-example
```
