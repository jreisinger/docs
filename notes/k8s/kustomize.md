[kustomization](https://kubernetes.io/docs/tasks/manage-kubernetes-objects/kustomization/) is the customization of k8s objects through a `kustomization.yaml` file.

```
$ ls
kustomization.yaml  nginx.yaml

$ cat kustomization.yaml 
namespace: t012
resources:
  - nginx.yaml

$ cat nginx.yaml 
apiVersion: v1
kind: Pod
metadata:
  name: nginx
spec:
  containers:
  - image: nginx:1.21.1
    name: nginx
```

```
k kustomize  # print the transformed manifest to stdout
k apply -k . # apply the transformed manifest
```
