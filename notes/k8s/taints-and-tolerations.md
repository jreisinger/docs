
Allow for scheduling pods onto appropriate nodes by using taints (opposite of node affinity) on nodes and tolerations on pods.

Add taint to node:

```
                                          taint
                                 +-----------------------+
k taint nodes docker-for-desktop dedicated=true:NoSchedule
                                 +-------+ +--+ +--------+
                                    key     val   effect
```

Add toleration to pod:

```yaml
apiVersion: v1
kind: Pod
...
spec:
  tolerations:
  - key: "dedicated"
    operator: "Equal" # default
    value: "true"     # optional
    effect: "NoSchedule"
```

Remove taint from node:
```
k taint nodes docker-for-desktop dedicated=true:NoSchedule-
```
