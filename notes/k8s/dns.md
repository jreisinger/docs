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
k create ns there
k run nginx --image=nginx --expose --port 80 -n there
k run busybox --image=busybox --rm -it --restart=Never -n there -- wget -O- nginx:80
```

Resolving a Service from different namespace

```
k create ns here
k run busybox --image=busybox --rm -it --restart=Never -n here -- wget -O- nginx.there:80
```
