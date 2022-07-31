# K8S RBAC Example

[RBAC](./rbac.yaml)

docker image `semihbkgr/k8s-rbas-example:1.0.3`

apply rbac:

```shell
kubectl apply -f rbac.yaml
```

to create pod in k8s cluster:

```shell
kubectl apply -f pod.yaml
```

to get logs:

```shell
kubectl logs k8s-rbac-exampl
```
