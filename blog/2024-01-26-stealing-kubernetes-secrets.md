![image](https://github.com/jreisinger/docs/assets/1047259/5fe480d5-b257-4fe3-975b-1bdcc9fe5ad2)

Kubernetes provides an object called [Secret](https://kubernetes.io/docs/concepts/configuration/secret) that is meant for storing sensitive data like passwords, tokens or keys. Secrets are decoupled (distinct) from Pods to decrease the risk of exposing sensitive data while creating, viewing and updating Pods. Containers in a Pod can access secrets via [environment](https://kubernetes.io/docs/tasks/inject-data-application/distribute-credentials-secure/#define-container-environment-variables-using-secret-data) variables or [files](https://kubernetes.io/docs/tasks/inject-data-application/distribute-credentials-secure/#create-a-pod-that-has-access-to-the-secret-data-through-a-volume) mounted through volumes.

Let's create a Secret named `mypassword` holding a key/value pair `password=s3cret!`:

```sh
$ kubectl create secret generic mypassword --from-literal password=s3cret!
```

NOTE: When you create secrets from command line they get persisted in your shell history file, e.g. `~/.bash_history`. To prevent this add space in front of the command. The secret is also visible in the processes listing, like `ps aux`. So it's best not to create production secrets from command line.

Ok, now, how secure is the secret we've created? It turns out that by default, not very. Let's have a look.

## Getting secrets from the API server

Anyone who has access to the Kubernetes API server can get the secret:

```sh
$ kubectl get secrets mypassword -o yaml
...
data:
  password: czNjcmV0IQ==
...
```

Oh, but we can't read it. Is it encrypted? No, it's just base64 decoded:

```sh
$ echo czNjcmV0IQ== | base64 -d -
s3cret!
```

## Getting secrets from etcd

Secrets, like other Kubernetes objects, are persisted in the etcd data store; by default unencrypted. So if we can access the data store, we can see the secrets. On a minikube cluster, we can do it like this:

```sh
$ minikube ssh
$ sudo -i
# cat << "EOF" | bash
export ETCDCTL_CACERT=/var/lib/minikube/certs/etcd/ca.crt
export ETCDCTL_CERT=/var/lib/minikube/certs/etcd/peer.crt
export ETCDCTL_KEY=/var/lib/minikube/certs/etcd/peer.key
export ETCDCTL_API=3

ETCDCTL_BIN=$(find / -name etcdctl | grep bin | head -1)

$ETCDCTL_BIN get /registry/secrets/default/mypassword
EOF
...
passwords3cret!▒Opaque▒"
...
```

## Getting secrets from within a Pod

Let's suppose that we can't access the API server or the etcd database because the cluster operator put some authorization in place. And that's the way to do it on a production cluster. The authorization mechanism in Kubernetes is called RBAC (Role Based Access Control). It's composed of the following primitives

- User represents a "normal user" connecting to the cluster (there's no API resource for User)
- ServiceAccount represents a program running in a pod and there is a pre-created `default` service account for each namespace assigned to each created pod
- Role defines a set of permissions on a namespace (or cluster) level
- RoleBinding maps roles to users or service accounts on a namespace (or cluster) level

Let's create a service account that is allowed to read (list allows for implicit reading) all secrets within the default namespace:

```sh
$ kubectl create serviceaccount secrets-reader
$ kubectl create role read-secrets --resource secrets --verb list
$ kubectl create rolebinding secrets-reader --serviceaccount default:secrets-reader --role read-secrets
```

Here's a pod using the service account we have created (instead of the default service account):

```sh
$ cat << EOF | k apply -f -
apiVersion: v1
kind: Pod
metadata:
  labels:
    run: nginx
  name: nginx
spec:
  serviceAccount: secrets-reader
  containers:
  - image: nginx
    name: nginx
EOF
```

Service account authenticates to the Kubernetes API via a JWT token that is mounted inside pod containers. If an attacker gains access to a container (for example by exploiting a vulnerability inside a web application or a web server) she can get all secrets in a namespace (or on the whole cluster if clusterrolebinding was used). Like this:

```sh
$ cat << 'EOF' | kubectl exec -i nginx -- bash | jq -r '.items[].data.password' | base64 -d
SAPATH=/var/run/secrets/kubernetes.io/serviceaccount
TOKEN=$(cat ${SAPATH}/token)
CACERT=${SAPATH}/ca.crt
URLPATH=https://kubernetes.default.svc/api/v1/namespaces/default/secrets
curl -s --cacert ${CACERT} --header "Authorization: Bearer ${TOKEN}" $URLPATH
EOF
```

## Finding risky Roles

There's a tool called [KubiScan](https://github.com/cyberark/KubiScan) that can find risky roles (and other objects) for you:

```sh
$ git clone git@github.com:cyberark/KubiScan.git
$ cd KubiScan
$ ./docker_run.sh ~/.kube/config

$ kubiscan --risky-roles # -r to show also rules (permissions)
...
+------------+
|Risky Roles |
+----------+------+-------------+------------------------------------+-----------------------------------+
| Priority | Kind | Namespace   | Name                               | Creation Time                     |
+----------+------+-------------+------------------------------------+-----------------------------------+
| CRITICAL | Role | default     | read-secrets                       | Sat Jan 27 16:30:21 2024 (0 days) |
| CRITICAL | Role | kube-system | system:controller:bootstrap-signer | Sat Jan 27 16:21:17 2024 (0 days) |
| CRITICAL | Role | kube-system | system:controller:token-cleaner    | Sat Jan 27 16:21:17 2024 (0 days) |
+----------+------+-------------+------------------------------------+-----------------------------------+
```
