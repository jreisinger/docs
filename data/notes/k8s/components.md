# Cluster components

[Source](https://kubernetes.io/docs/concepts/overview/components).

<img src="https://d33wubrfki0l68.cloudfront.net/2475489eaf20163ec0f54ddc1d92aa8d4c87c96b/e7c81/images/docs/components-of-kubernetes.svg" style="max-width:100%;height:auto;"> 

* Node - worker machine (virtual or physical) that runs containerized applications via Pods
* Control plane - manages Nodes and Pods

## Control plane components

kube-apiserver

* exposes the API
* front end for the K8s control plane
* scales horizontally

cloud-controller-manager

* embeds cloud specific control logic
* links K8s cluster with cloud provider's API
* runs controllers specific to the underlying cloud provider
* these controllers can have cloud provider dependencies: node controller, route controller, service controller

kube-controller-manager

* runs controller processes
* logically each controller is a separate process but they are all compiled into a single binary and run in a single process
* node controller, replication controller, endpoints controller, service account & token controllers

etcd

* consistent and highly-available key value store for all cluster data

kube-scheduler
 
* watches for newly created Pods and selects a node for them to run on

## Node components

kubelet

* takes a set od PodSpecs and ensures that the containers described in those PodSpecs are running (in a Pod) and healthy

kube-proxy

* implements part of the Service concept by maintaining network rules on nodes (using OS packet filtering layer or forwarding traffic by itself)

container runtime

* software responsible for running containers
* e.g. Docker, containerd, CRI-O

## Addons

* use K8s resources (DaemonSet, Deployment, etc) to implement cluster features
* namespaced addon resources belong within `kube-system` namespace

Selected addons:

* DNS - should be in all K8s clusters; containers automatically include this DNS server in their DNS searches (via `/etc/resolv.conf`)
* WebUI (Dashboard) - web based UI
* Container Resource Monitoring - generic time-series metrics
* Cluster-level Logging

# Kubernetes primitives

<img width="687" alt="image" src="https://user-images.githubusercontent.com/1047259/170937963-dd5927c8-3624-4874-95dd-4a3906caef6c.png">
