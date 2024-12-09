# Pod Scheduling in Kubernetes: Taints, Tolerations, and Node Affinity

In Kubernetes, you can control where pods get scheduled by using two powerful mechanisms:

1. **Taints and Tolerations**: These allow nodes to repel specific pods unless the pods explicitly "tolerate" the node's taints.  
2. **Node Affinity**: A property that attracts pods to specific nodes based on matching labels.

Let’s explore how to use these features in practice.

## Taints and Tolerations

Taints let you mark a node to repel certain pods unless the pods declare tolerations for those taints.

### Step 1: Start a Kubernetes Cluster

First, launch a Kubernetes cluster using Minikube:

```bash
❯ minikube start
```

### Step 2: Taint a Node

Taint the Minikube node with a key-value pair and an effect (`NoSchedule` in this case):

```bash
❯ kubectl taint node minikube key1=value1:NoSchedule
```

The `NoSchedule` effect ensures that no pod gets scheduled on this node unless it tolerates the taint.

### Step 3: Deploy and Observe Scheduling

Create a deployment and check the status of its pods:

```bash
❯ kubectl create deployment nginx --image=nginx
❯ kubectl get deployments.apps 
NAME    READY   UP-TO-DATE   AVAILABLE   AGE
nginx   0/1     1            0           109s
```

The pods will remain unscheduled. Investigate further by describing the pod:

```bash
❯ kubectl describe pod nginx-676b6c5bbc-lxswz
...
Events:
Type     Reason            Age   From               Message
----     ------            ----  ----               -------
Warning  FailedScheduling  29s   default-scheduler  0/1 nodes are available: 1 node(s) had untolerated taint {key1: value1}.
```

### Step 4: Add Tolerations

To allow the pods to be scheduled, edit the deployment to add tolerations:

```bash
❯ kubectl edit deployments.apps nginx
```

Update the spec as follows:

```yaml
spec:
  template:
    spec:
      tolerations:
      - key: key1
        operator: Equal
        value: value1
        effect: NoSchedule
```

Once applied, the pods will tolerate the taint and get scheduled.

### Step 5: Cleanup

When done, clean up the deployment:

```bash
❯ kubectl delete deployments.apps nginx
```

## Node Affinity

Node affinity provides a way to attract pods to nodes based on matching labels. Let’s see this in action.

### Step 1: Add Cluster Nodes

Add two more nodes to your cluster:

```bash
❯ minikube node add
❯ minikube node add
```

### Step 2: Create a Deployment

Deploy an application with multiple replicas and verify where the pods are scheduled:

```bash
❯ kubectl create deployment nginx --image=nginx --replicas=3
❯ kubectl get pods -o wide
```

By default, the scheduler will distribute the pods across all nodes.

### Step 3: Apply Node Affinity

To constrain pods to a specific node, label one of the nodes:

```bash
❯ kubectl label nodes minikube-m02 key2=value2
```

Then, edit the deployment to add a node affinity rule:

```bash
❯ kubectl edit deployments.apps nginx
```

Update the spec as follows:

```yaml
spec:
  template:
    spec:
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: key2
                operator: In
                values:
                - value2
```

With this configuration, the pods will only be scheduled on nodes labeled with `key2=value2`.

### Step 4: Cleanup

Once you’re done, clean up the cluster:

```bash
❯ minikube delete
```

## Conclusion

Taints, tolerations, and node affinity are essential tools for fine-tuning pod placement in Kubernetes. By combining these features, you can achieve more predictable and efficient scheduling, especially in clusters with diverse workloads.  

Happy Kubernetes-ing! 🚀