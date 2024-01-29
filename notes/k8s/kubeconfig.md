* a generic name for configuration file that holds info on how to communicate with a cluster
* contains one or more clusters, users, namespaces and authentication mechanisms
* kubectl configuration loading [order](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#config): 1. `--kubeconfig` 2. `KUBECONFIG` 3. `$HOME/.kube/config` (default)
* `KUBECONFIG` can hold multiple kubeconfigs (separated by `:` on Linux/Mac) that get [merged](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/#merging-kubeconfig-files)
* to see your configuration - `kubectl config view`
* cluster/user/namespace combination is a context
* to see to contexts - `kubectl config get-contexts`
* by default, kubectl uses parameters from the current context
* to choose the current context - `kubectl config use-context`
