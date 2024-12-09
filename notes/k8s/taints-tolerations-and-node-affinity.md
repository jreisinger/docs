There are two ways how to constrain on which nodes get the pods scheduled:

- taints and tolerations
- node affinity

Taints allow a node to repel some pods - those that don't tolerate node's taints. Node affinity is a property of pods that attracts them to some nodes. Let's see these in action.

Taints and tolerations
----------------------

Start a kubernetes cluster:

    ❯ minikube start

Taint the node:

    ❯ k taint node minikube key1=value1:NoSchedule # NoSchedule is called effect

Create a deployment and see whether it gets scheduled:

    ❯ k create deployment nginx --image=nginx
    ❯ k get deployments.apps 
    NAME    READY   UP-TO-DATE   AVAILABLE   AGE
    nginx   0/1     1            0           109s
    ❯ k describe pod nginx-676b6c5bbc-lxswz
    ...
    Events:
    Type     Reason            Age   From               Message
    ----     ------            ----  ----               -------
    Warning  FailedScheduling  29s   default-scheduler  0/1 nodes are available: 1 node(s) had untolerated taint {key1: value1}. preemption: 0/1 nodes are available: 1 Preemption is not helpful for scheduling.

Add tolerations for the deployment pods (spec -> template -> spec) so they get scheduled:

    ❯ k edit deployments.apps nginx
    ...
    tolerations:
    - effect: NoSchedule
      key: key1
      operator: Equal # default
      value: value1   # optional
    ...

Cleanup

    ❯ k delete deployments.apps nginx

Node affinity
-------------

Add two more cluster nodes:

    ❯ minikube node add
    ❯ minikube node add

Create a deployment and see that the pods get scheduled on both new nodes:

    ❯ k create deployment nginx --image=nginx --replicas=3
    ❯ k get pods -o wide

Now make sure that the pods get scheduled to only one of the nodes:

    ❯ k label nodes minikube-m02 key2=value2
    ❯ k edit deployments.apps nginx
    ...
    affinity:
      nodeAffinity:
        requiredDuringSchedulingIgnoredDuringExecution:
          nodeSelectorTerms:
          - matchExpressions:
            - key: key2
              operator: In
              values:
              - value2
    ...

Cleanup

    ❯ minikube delete