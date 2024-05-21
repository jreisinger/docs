A generic name for configuration file that holds info on how to communicate with a cluster.

kubectl configuration loading order

1. --kubeconfig
2. KUBECONFIG
3. $HOME/.kube/config (default)

KUBECONFIG can hold multiple kubeconfigs (separated by : on Linux/Mac) that get merged.

To see your configuration
 
    kubectl config view

To see existing contexts (cluster/user/namespace combination)
    
    kubectl config get-contexts
    
To choose a context
    
    kubectl config use-context
