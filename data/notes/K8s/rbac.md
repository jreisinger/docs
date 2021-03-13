You can do k8s access control via:

* RBAC (Role Based Access Control)
* managing access by cluster (w/o RBAC anyone with access to cluster can do anything)

# RBAC

Find out whether RBAC is enabled on a cluster (one line for each control node):

```
> kubectl describe pod -n kube-system -l component=kube-apiserver | grep authorization
      --authorization-mode=Node,RBAC
      --authorization-mode=Node,RBAC
      --authorization-mode=Node,RBAC
```

In Kubernetes, permissions are additive; users start with no permissions, and you can add permissions using Roles and RoleBindings. You can’t subtract permissions from someone who already has them.

## User (`serviceaccounts`)

* every time you connect to a cluster you do so as a specific user
* there is `default` service account for each namespace
* service account - user account associated with automated system
* authentication depends on the cluster provider (e.g. `gcloud` uses a token per cluster)
* users can have different sets of permissions - governed by `roles`

## Role (`roles`)

* a specific set of permissions
* `Role` - defines roles on a namespace level
* `ClusterRole` - defines roles accross the whole cluster

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

## RoleBinding (`rolebindings`)

* associate a user with a role
* also here you can have RoleBinding or ClusterRoleBinding

```
# daisy can edit stuff in demo namespace only
king: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: daisy-edit
  namespace: demo
subjects:
- kind: User
  name: daisy
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: ClusterRole
  name: edit
  apiGroup: rbac.authorization.k8s.io
```

* see existing RoleBindings:

```
kubectl get rolebindings.rbac.authorization.k8s.io --all-namespaces
```

# Sources

* https://learning.oreilly.com/library/view/cloud-native-devops/9781492040750/ch11.html
