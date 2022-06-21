You can do k8s access control via:

* managing access per cluster (anyone with access can do anything) - ok for testing, small deployments
* RBAC (Role Based Access Control) - production setup

Find out whether RBAC is enabled on a cluster (one line for each control node):

```
$ k describe pod -n kube-system -l component=kube-apiserver | grep authorization
      --authorization-mode=Node,RBAC
      --authorization-mode=Node,RBAC
      --authorization-mode=Node,RBAC
```

Check the pods list permissions assigned to user johndoe:

```
$ k auth can-i list pods --as johndoe
no
```

# Overview

* every time you connect to a cluster API you do so as a specific user
* service account - user account associated with a program running in a pod
* there is `default` service account for each namespace
* users can have different sets of permissions - governed by `roles`

<img width="639" alt="image" src="https://user-images.githubusercontent.com/1047259/167849081-bc128f2b-5757-4d4c-82e1-c19f71836cee.png">

Groups and users

* represent real persons
* distributed by cluster admin
* not an API resource

Service account

* represents a program
* assigned to a pod (if not, `default` service account is used)

```
apiVersion: v1
kind: ServiceAccount
metadata:
  name: build-bot
---
apiVersion: v1
kind: Pod
metadata:
  name: build-observer
spec:
  serviceAccountName: build-bot
...
```

Authentication depends on the cluster provider.

<img width="631" alt="image" src="https://user-images.githubusercontent.com/1047259/167849397-a4aa7317-1e6b-4f9d-beb8-c4ed1edd28dd.png">

# RBAC primitives

<img width="638" alt="image" src="https://user-images.githubusercontent.com/1047259/167874790-755953d0-2f25-467e-911e-6f4703c52500.png">

In Kubernetes, permissions are additive; users start with no permissions, and you can add permissions using Roles and RoleBindings. You can’t subtract permissions from someone who already has them.

## Role

* a specific set of permissions
* `Role` - defines permissions on a namespace level
* `ClusterRole` - defines permissions accross the whole cluster

There are some [defaults](https://kubernetes.io/docs/reference/access-authn-authz/rbac/#default-roles-and-role-bindings):

<img width="638" alt="image" src="https://user-images.githubusercontent.com/1047259/167877282-9a9e4e0c-d68b-4b25-9710-f4a11cd0c8b6.png">

or you can create your own:

```
# role that grants read access to secrets in any namespace
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: secret-reader
rules:
- apiGroups: [""]
  resources: ["secrets"]
  verbs: ["get", "watch", "list"]
```

## RoleBinding

* associates a user with a role
* also here you can have RoleBinding or ClusterRoleBinding

```
# daisy can edit stuff in demo namespace only
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: daisy-edit
  namespace: demo
subjects:
- kind: User
  name: daisy
roleRef:
  kind: ClusterRole
  name: edit
  apiGroup: rbac.authorization.k8s.io
```

* to see existing RoleBindings:

```
kubectl get rolebindings.rbac.authorization.k8s.io --all-namespaces
```

# Sources

* https://learning.oreilly.com/library/view/cloud-native-devops/9781492040750/ch11.html
* Benjamin Muschko: Certified Kubernetes Administrator (CKA) Study Guide (2022)
