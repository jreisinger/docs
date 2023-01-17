You can do k8s access control via:

* managing access per cluster (anyone with access can do anything) - ok for testing, small deployments
* RBAC (Role Based Access Control) - production setup

# RBAC primitives

<img width="638" alt="image" src="https://user-images.githubusercontent.com/1047259/167874790-755953d0-2f25-467e-911e-6f4703c52500.png">

User ("normal user")

* every time you connect to a cluster API you do so as a specific user
* user represents a real person
* it's someone with cert (signed by cluster CA and distributed by cluster admin) and key
* there is no API resource User (`k api-resources | grep -i user`)
* users can have different sets of permissions - governed by `roles`

Service account

* represents a program running in a pod (there is `default` service account for each namespace)
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

Role

* a specific set of permissions
* `Role` - defines permissions on a namespace level
* `ClusterRole` - defines permissions accross the whole cluster


<img width="638" alt="image" src="https://user-images.githubusercontent.com/1047259/167877282-9a9e4e0c-d68b-4b25-9710-f4a11cd0c8b6.png">

or you can create your own:

```
# role that grants read access to secrets in any namespace
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: secret-reader
rules:
- apiGroups: [""] # "" indicates the core API group
  resources: ["secrets"]
  verbs: ["get", "watch", "list"]
```

Before creating your own Role consider using [default ones](https://kubernetes.io/docs/reference/access-authn-authz/rbac/#default-roles-and-role-bindings):

* User facing roles: cluster-admin, admin (for namespaces), edit and view
* For core cluster components: system:kube-controller-manager, system:node, ...
* For other cluster components: system:persistent-volume-provisioner, ...

RoleBinding

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

Verbs (actions on resources)

* get, list (read-only)
* create, update, patch, delete, deletecollection (read-write)

Rules of thumb:

* to grant access to a resource in a single namespace use Role + RoleBinding
* to reuse a role in a couple of namespaces define ClusterRole + RoleBinding
* to grant access to cluster-wide resources (like nodes) or to namespaces resources across all namespaces, use ClusterRole + ClusterRoleBinding

# Demo (kubectl commands)

Find out whether RBAC is enabled on a cluster (one line for each control node):

```
$ k describe pod -n kube-system -l component=kube-apiserver | grep authorization
      --authorization-mode=Node,RBAC
      --authorization-mode=Node,RBAC
      --authorization-mode=Node,RBAC
```

See existing RoleBindings in all namespaces:

```
k get rolebindings -A -o wide --sort-by=".metadata.creationTimestamp"
```

## User access management

```
# create key
openssl genrsa -out jane.key 2048

# create CSR
openssl req -new -key jane.key -out jane.csr

# create CSR K8s resource
cat <<EOF | kubectl apply -f -
apiVersion: certificates.k8s.io/v1
kind: CertificateSigningRequest
metadata:
  name: jane
spec:
  request: $(base64 -w 0 jane.csr)
  signerName: kubernetes.io/kube-apiserver-client
  expirationSeconds: 86400  # one day
  usages:
  - client auth
EOF

# approve CSR, i.e. issue cert
k certificate approve jane

# get cert
k get csr jane -ojson | jq -r .status.certificate | base64 -d > jane.crt

# use crt and key
k config set-credentials jane --client-key=jane.key --client-certificate=jane.crt --embed-certs
k config set-context jane --user=jane --cluster=kind-kind
k config use-context jane
```

```
# check the permissions assigned to user johndoe
k auth can-i list pods --as johndoe

# assign new permissions to user johndoe
k create role pod-reader -n default --resource=pods --verb=watch,list,get
k create rolebinding read-pods -n default --role=pod-reader --user=johndoe
```

NOTE: In Kubernetes, permissions are additive; users start with no permissions, and you can add permissions using Roles and RoleBindings. You can’t subtract permissions from someone who already has them.

## Service account access management

You have an app that needs access to pod info. The default `view` role is too much (`kubectl describe clusterrole view`).

```
k create ns myapp
k create serviceaccount myappid -n myapp
k create role podview --resource=pods --verb=list,get -n myapp
k create rolebinding podviewer --serviceaccount=myapp:myappid --role podview -n myapp
k auth can-i --as=system:serviceaccount:myapp:myappid list pods -n myapp
```

NOTE: ClusterRoleBinding applies to all namespaces including future namespaces.

# Sources

* https://www.udemy.com/course/certified-kubernetes-security-specialist
* https://learning.oreilly.com/library/view/cloud-native-devops/9781492040750/ch11.html
* Benjamin Muschko: Certified Kubernetes Administrator (CKA) Study Guide (2022)
