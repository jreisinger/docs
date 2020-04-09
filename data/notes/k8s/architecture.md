# Cluster components

[Source](https://kubernetes.io/docs/concepts/overview/components).

<img src="https://d33wubrfki0l68.cloudfront.net/7016517375d10c702489167e704dcb99e570df85/7bb53/images/docs/components-of-kubernetes.png" style="width:100%;"> 

* nodes - worker machines (VMs or physicals) that run containerized applications - Pods
* control plane - manages nodes and Pods

Control plane components

* kube-apiserver - exposes the API; scales horizontally
* etcd - consistent and highly-available key value store for all cluster data
* kube-scheduler - watches for newly created Pods and selects a node for them to run on
* kube-contoller-manager - runs controller processes; logically each controller is a separate process but they are all compiled into a single binary and run in a single process (node controller, replication controller, endpoints controller, service account & token controllers)
* cloud-controller-manager - runs controllers that interact with the underlying cloud providers (noder controller, route controller, service controller, volume controller)

Node components

* kubelet - takes a set od PodSpecs and ensures that the containers described in those PodSpecs are running (in a Pod) and healthy
* kube-proxy - implements part of the Service concept by maintaining network rules on nodes (using OS packet filtering layer or forwarding traffic by itself)
* container runtime - software responsible for running containers (Docker, containerd, CRI-O)
