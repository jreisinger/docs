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
- apiGroups: [""] # "" indicates the core API group
  resources: ["secrets"]
  verbs: ["get", "watch", "list"]
```

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

There are 4 different RBAC combinations possible and 3 valid ones:

1. Role + RoleBinding - available in single Namespace, applied in single Namespace
2. ClusterRole + ClusterRoleBinding - available cluster-wide, applied cluster-wide
3. ClusterRole + RoleBinding - available cluster-wide, applied in single Namespace
4. Role + ClusterRoleBinding - NOT POSSIBLE: available in single Namespace, applied cluster-wide

# kubectl commands

Find out whether RBAC is enabled on a cluster (one line for each control node):

```
$ k describe pod -n kube-system -l component=kube-apiserver | grep authorization
      --authorization-mode=Node,RBAC
      --authorization-mode=Node,RBAC
      --authorization-mode=Node,RBAC
```

See existing RoleBindings in all namespaces:

```
k get rolebindings.rbac.authorization.k8s.io -A -o wide
```

Basic user access management:

```
# create key
openssl genrsa -out jane.key 2048

# create CSR
openssl req -new -key jane.key -out jane.csr

# get cert
cat <<EOF | kubectl apply -f -
apiVersion: certificates.k8s.io/v1
kind: CertificateSigningRequest
metadata:
  name: jane
spec:
  request: <OUTPUT FROM: base64 -w 0 jane.csr>
  signerName: kubernetes.io/kube-apiserver-client
  expirationSeconds: 86400  # one day
  usages:
  - client auth
EOF

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

Basic service account access management:

```
k create serviceaccount api-access -n apps
k create clusterrole api-clusterrole --resource=pods --verb=watch,list,get
k create clusterrolebinding api-clusterrolebinding --clusterrole=api-clusterrole --serviceaccount=apps:api-access
```

NOTE: ClusterRoleBinding applies to all namespaces including future namespaces.

# Sources

* https://www.udemy.com/course/certified-kubernetes-security-specialist
* https://learning.oreilly.com/library/view/cloud-native-devops/9781492040750/ch11.html
* Benjamin Muschko: Certified Kubernetes Administrator (CKA) Study Guide (2022)
