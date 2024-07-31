Kubernetes is a distributed operating system. Applications running on it are called cloud native. See [What is Kubernetes?](https://kubernetes.io/docs/concepts/overview/what-is-kubernetes/) for more.

Config
======

`kubectl` cluster configuration

```bash
kubectl config view
```

Context

* to manage different `namespace`s, `cluster`s and `user`s

```bash
# list contexts
kubectl config get-contexts

# switch context
kubectl config use-context <context-name>
```

Basic objects
=============

* everything in Kubernetes is represented by a RESTful resource or a Kubernetes object ([resources vs objects](https://stackoverflow.com/questions/52309496/difference-between-kubernetes-objects-and-resources))
* each object exists at a unique HTTP path
* the `kubectl` makes requests to these URLs to access the objects
* `get` is conceptually similar to `ps`

```sh
# view Kubernetes objects
kubectl get all [-l app=nginx] # all objects [with a label app=nginx]
kubectl get <type>             # all objects of given type
kubectl get <type> <object>    # specific object

# details about an object
kubectl describe <type> <object>

# create, update objects
kubectl apply -f obj.yaml

# delete objects
kubectl delete -f obj.yaml  # no additional prompting!
kubectl delete <type> <object>
```

Namespace
---------

* virtual cluster
* group of objects in a cluster
* similar to a filesystem folder
* see [Namespaces](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/) for more 

```bash
# all namespaces in a cluster, need admin rights
kubectl get namespaces
```

Pod
---

* atomic unit of work in Kubernetes cluster
* one or more containers working together symbiotically
* all containers in a Pod always land on the same node
* once scheduled to a node, Pods don't move
* each container runs 
    * runs its *own* cgroup
    * but they *share* hostname, IP address + port space and filesystem
* containers can communicate via IPC, loopback interface (`localhost` hostname) and can share files
* like a logical host
* if you want to persist data across multiple instances of a Pod, you need to use `PersistentVolumes`

```sh
# Pod manifest - just a text-file representation of the Kubernetes API object
$ cat kuard-pod.yaml
apiVersion: v1
kind: Pod
metadata:
  name: kuard
spec:
  containers:
    - image: gcr.io/kuar-demo/kuard-amd64:1
      name: kuard
      ports:
        - containerPort: 8080
          name: http
          protocol: TCP
```

```sh
# Creating a Pod
kubectl apply -f kuard-pod.yaml
```

What should I put into a single pod?

* "Will these containers work correctly if they land on different machines?"
* should go into a Pod: web server + git synchronizer - they communicate via filesystem
* should go into separate Pods: Wordpress + DB - they can communicate over net

Deployment
----------

* object of type controller
* manages replicasets/pods

One way to manage a deployment:

```bash
kubectl create deployment quotes-prod --image=reisinge/quotes
kubectl scale deployment quotes-prod --replicas=3
kubectl label deployment quotes-prod ver=1 env=prod
```

Service
-------

* object that solves the service discovery problem (i.e. finding things in K8s cluster)
* a way to create a named label selector (see `kubectl get service -o wide`)
* a service is assigned a VIP called a *cluster IP* -> load balanced across all the
  pods identified by the selector
* good for identifying services inside a cluster

One way to create a service:

```bash
kubectl expose deployment quotes-prod --port=80 --target-port=5000
```

A service is accessible from within the cluster:

```bash
kubectl run busybox --image=busybox --rm -it --restart=Never -- wget quotes-prod/api/v1/random -qO-
```

Exposing services outside of the cluster
========================================

Service of type NodePort

* it enhances a service
* in addition to a cluster IP, a service gets a port (user defined or picked by
    the system)
* every node in the cluster forwards traffic to that port to the service
* if you can reach any node in the cluster you can get to the service
* this can be intergrated with HW/SW load balancers to expose the service even furher

Ingress

* for HTTP or HTTPS

Service of type LoadBalancer

* for other ports than HTTP(S)

Resources
=========

* Kubernetes: Up and Running (2017, 2019)
* Cloud Native DevOps with Kubernetes (2019)
* Managing Kubernetes (2018)
