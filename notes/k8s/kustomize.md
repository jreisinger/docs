Kubernetes YAML files are complicated, verbose and repetitive. Kustomize allows you to start with a base YAML manifests, and use *overlays* to patch manifests for different environments and configurations.

```
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
$ cat kustomization.yaml 
namespace: t012
resources:
  - nginx.yaml
```

```
k kustomize  # print the transformed manifest to stdout
k apply -k . # apply the transformed manifest
```

See https://github.com/cloudnativedevops/demo/tree/main/hello-kustomize/demo for more.