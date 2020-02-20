* speficications of how groups of pods can communicate with each other and other network endpoints
* in other words, they are the firewall rules of a k8s cluster
* by default, all pods in a cluster are non-isolated and accept traffic from any source
* as soon as you have a network policy that selects some pods, those pods become isolated and reject any traffic that is not allowed by a network policy
* network policies are additive, i.e. they don't conflict and evaluation order has no effect
* network policies are effective only within a namespace

More

* https://kubernetes.io/docs/concepts/services-networking/network-policies/
* https://docs.giantswarm.io/guides/limiting-pod-communication-with-network-policies/
