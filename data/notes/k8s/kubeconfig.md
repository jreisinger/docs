* a generic name for configuration file that holds info to access a cluster
* contains info about one or more clusters, users, namespaces and authentication mechanisms
* cluster/user/namespace combination is a context
* to see to contexts - `kubectl congig get-contexts`
* by default, `kubectl` uses parameters from the current context to communicate with the cluster
* to choose the current context - `kubectl config use-context`
* `kubectl` loading [order](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#config) 1. `--kubeconfig` 2. `KUBECONFIG` 3. `$HOME/.kube/config` (default)
* `KUBECONGIG` can hold multiple kubeconfigs (separated by `:` on Linux/Mac) that get [merged](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/#merging-kubeconfig-files)
* to see your configuration - `kubectl config view`
