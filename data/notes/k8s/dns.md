* k8s is geared towards operating a microservices architecture
* microservices implement distinct, self-contained functionality
* microservices communicate with one another to complement each other
* k8s runs CoreDNS (an implementation of DNS) that maps a Service name to an IP address

```
k get pods -n kube-system | grep dns     # CoreDNS pods
k get cm -n kube-system coredns -o yaml  # configuration of CoreDNS pods
```

Resolving a Service from the same namespace

```
k create ns namespace
k run echoserver --image=k8s.gcr.io/echoserver:1.10 --restart=Never --port=8080 --expose -n namespace
k run busybox --image=busybox --rm -it --restart=Never -n namespace -- wget echoserver:8080
```

Resolving a Service from different namespace

```
k create ns other
k run busybox --image=busybox --rm -it --restart=Never -n other -- wget echoserver.namespace:8080
```
