* speficication of how groups of pods can communicate with each other and other network endpoints
* in other words, they are the firewall rules of a k8s cluster
* by default, all pods in a cluster are non-isolated and accept traffic from any source
* as soon as you have a network policy that selects some pods, those pods become isolated and reject any traffic that is not allowed by a network policy
* network policies are additive, i.e. they don't conflict and evaluation order has no effect
* network policies are effective only within a namespace

*WARNING*: To use network policies, you must be using a [network plugin](https://kubernetes.io/docs/concepts/extend-kubernetes/compute-storage-net/network-plugins/) which supports NetworkPolicy. Creating a NetworkPolicy resource without a controller that implements it will have no effect.

Disallowing all traffic:

```
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: default-deny-all
spec:
  podSelector: {} # all pods in NS
  policyTypes:    # types of traffic
  - Ingress
  - Egress
```

To verify a network policy:

```
# run nginx pod
kubectl run nginx --image=nginx

# try to reach the nginx pod
kubectl run busybox --image=busybox -it -- /bin/sh
/ # wget 192.168.141.249 -O-
```

More

* https://kubernetes.io/docs/concepts/services-networking/network-policies/
* https://docs.giantswarm.io/guides/limiting-pod-communication-with-network-policies/
