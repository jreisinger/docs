
Pod's status field contains `phase` field that provides high-level summary of where the Pod is in its lifecycle.

1. `Pending`
2. `Running` - if at least one of its primary containers starts OK
3. `Succeeded` or `Failed` if any container terminated in failure

`Unknown` - typically when there's a comm problem with the node running the Pod

Don't confuse kubectl/k9s `Status` (meant for user's intuition) with the pod's phase (explicit part of K8s data model and API). Status can show:

- `CrashLoopBackOff` - pod fails to start repeatedly
- `Terminating` - pod is being deleted

Container states

- `Waiting`
- `Running`
- `Terminated`

---

Source: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/
