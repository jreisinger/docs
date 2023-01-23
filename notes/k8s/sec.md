# Kubernetes attack vectors

```
                  +---------------------------------+
                  | Cluster                         |
Access to         |   +--------------------------+  |
machines or VMs   |   | Control plane            |  |
------------------+-->|           +------+       |  |   Access to etcd API
                  |   |           | etcd |<------+--+---------------------
                  |   |           +------+       |  |
                  |   |                          |  |
Access via        |   |    +----------------+    |  |   Intercept/modify/inject
K8s API or proxy  |   |    |  Control plane |    |  |   control-plane traffic
------------------+---+--->|  components    +----+--+--------------------------
                  |   |    +----------------+    |  |
                  |   +--------------------------+  |
                  |                                 |
                  |   +--------------------------+  |
Access via        |   | Nodes                    |  |
Kubelet API       |   |         +-----------+    |  |
------------------+---+-------->| Kubelet   |    |  |
                  |   |         +-----------+    |  |
                  |   |                          |  |
                  |   | +----------------------+ |  |
                  |   | |Pod                   | |  |
                  |   | | +------------------+ | |  |
Escape container  |   | | | Container        | | |  |  Intercept/modify/inject
to host through   |   | | |  +-------------+ | | |  |  application traffic
<-----------------+---+-+-+--+ Application +-+-+-+--+-------------------------
vulnerability or  |   | | |  +-------------+ | | |  |
volume mount      |   | | |   ^              | | |  |
                  |   | | +---+--------------+ | |  |
                  |   | |     |                | |  |
                  |   | +-----+----------------+ |  |
                  |   |       |                  |  |
                  |   +-------+------------------+  |
                  |           |                     |
                  +-----------+---------------------+
                              |
                              | Exploit vulnerability
                              | in application code
```

Possible damage

* unauthorized compute resources usage (e.g. Tesla cryptojacking)
* disrupting your existing services
* accessing your data

# Securing the cluster

* in the early days Kubernetes default settings were insecure
* different installation tools/ways configure cluster in different ways

## API server

* REST API for controlling the cluster
* user with full permissions on API == root access to all nodes
---
* check `--insecure-port` is not enabled:
```
root@kind-control-plane:/# curl localhost:8080
curl: (7) Failed to connect to localhost port 8080 after 0 ms: Connection refused
```
* this means API can be accessed only over TLS
* set `--anonymous-auth=false` to allow only authenticated users to access API
* enable RBAC and RBAC for kubelets (`Node`): `--authorization-mode=Node,RBAC`

## Kubelet

* agent that interacts with container runtime
* launches pods and reports node and pod status and metrics
* operates an API through which components asks e.g. to start/stop pods
* unauthorized access can lead to [owning the cluster](https://medium.com/handy-tech/analysis-of-a-kubernetes-hack-backdooring-through-kubelet-823be5c3d67c)
---
* disable anonymous access: `--anonymous-auth=false` (that's why API server needs to use `--kubelet-client-certificate` and `--kubelet-client-key`)
* ensure requests are authorized by setting `--authorization-mode` to something else than `AlwaysAllow`
* limit permissions of kubelets to its own Node by setting `--enable-admission-plugins=NodeRestriction` on kubernetes API
* set `--read-only-port=0` to disable anonymous access to info about workloads

```
root@kind-control-plane:/# curl -sk https://localhost:10250/pods/
Unauthorized
```

* kubelet needs client certificate to access API; certs are rotated automatically

## etcd

* distributed key/value store
* all k8s configuration and state
* anyone who can write to etcd can control the cluster
* read access provide hints for an attacker
---
* set `--cert-file` and `--key-file` to enable HTTPS
* set `--client-cert-auth=true` to require authentication
* set `--trusted-ca-file` to specify the CA that signed the client certificates
* set `--peer-client-cert-auth=true` to make etcd nodes communicate with each other securely
* set `--etcd-cafile` on the API server to the CA the signed etcd's cerfificate
* consider encrypting etcd data stored on disk; especially if you are storing secrets in etcd rather than an external secrets store
* consider using network firewalling; only control-plane components have any business talking to etcd

## Kubernetes dashboard

* powerful tool historically used to gain control of clusters
* make sure it's not an easy entry point for attackers
---
* allow only authenticated users
* use RBAC
* don't expose it to public internet
* see https://github.com/kubernetes/dashboard/tree/master/docs for details

## Validating configuration

CIS

* https://www.cisecurity.org/benchmark/kubernetes -> automated via https://github.com/aquasecurity/kube-bench

Penetrations testing

* hire a pen-tester company/specialist
* https://github.com/aquasecurity/kube-hunter
```
k apply -f https://raw.githubusercontent.com/aquasecurity/kube-hunter/main/job.yaml
k describe job kube-hunter
k logs kube-hunter-ID
```

# Authentication

* IAM (both in AWS and K8s) lets you define access to resources for users and services
* first step the K8s API has to do is to verify who or what (program) is issuing the request
* to establish the identity of the caller, or in other words, to authenticate the caller

## Identity

Normal users

* K8s doesn't have a first-class notion of a human user
* K8s assumes users are managed outside via directory services (like LDAP) or SSO (like SAML or Kerberos) 
* user accounts are cluster-wide, so usernames must be unique across namespaces

Service accounts

* for applications that need to communicate with the API
* a namespaced resource
* if you don't specify `spec.serviceAccountName` in pod default SA is used

```
$ k run -it --rm somepod --restart=Never --image=alpine -- sh
/ # ls /var/run/secrets/kubernetes.io/serviceaccount/
ca.crt     namespace  token
```

* `token` is as JWT token that you can decode at jwt.io

## Authentication concepts

Flow

1. Client presents its credentials to the API server.
2. The API server uses one of the configured authentication plug-ins (see below) to establish the identity with an identity provider.
3. The identity provider verifies the credentials (including username, group).
4. If the credentials are ok, the API server moves on the the authorization. Otherwise, HTTP `401 Unauthorized` is returned.

## Authentication strategies

* represented by authentication plug-ins:

Static password or token file 

* client provides identity via HTTP header `Authorization` with value `Basic base64($USER:PASSWORD)` or `Bearer $TOKEN`
* inflexible, not recommended to production

X.509 certificates

* every user has their own X.509 client certificate
* the API server validates the cert via a configured CA
* CN of the subject is used as the username and any defined organizations are used as groups
* certificates are issued by cluster admin
* there is no certificate revocation -> good reason to use SSO if possible

OpenID Connect (OIDC)

* identity layer on top of the OAuth 2.0
* caller provides `id-token` in the form of JSON Web Token after using provider's (e.g. Google, AD) login page

Bootstrap tokens

* experimental, targeting the cluster setup phase

To integrate with other authentication protocols such as LDAP, SAML, and Kerberos:

* authenticating proxy - API idetifies users from headers such as `X-Remote-User`; you need to setup and run the proxy
* webhook token authentication - hook verifying bearer tokens

## Tooling and good practices

Depending on where you run you clusters, you migh use:

* Keycloak - open source IAM with support for existing LDAP servers
* Dex - defers authentication to other identity providers, like LDAP, SAML, GitHub, Google, AD
* AWS IAM authenticator for Kubernetes - uses IAM credentials to authenticate to a Kubernetes cluster
* Guard - webhook allowing you to usr various identity providers, like GitHub, Google, LDAP

Best practices

* use 3rd party providers, like Azure, Google, or GitHub
* don't use static files; if you can't use 3rd party use X.509 certs
* have identity lifecycle; when people leave, invalidate their credentials

# Authorization

* assigning and enforcing permissions (to create pods or get secrets) to users and applications
* by default permissions are denied, unless explicitly allowed by a policy

## Authorization concepts

```
                (1)   +--------+   (3)
             +--------+ Client +--------+
             |        +--------+        |path
 +---+   +---v---+                      |resource
 |401|<--+ Authn |                      |verb
 +---+   +---+---+                      |namespace
             | (2)                      |
             | user                     |...
             | group                    |
+--------+   |                          |
| Authz  +---v--------------------------v--+
| modules|          Authorization          |
+--------+---+--------------------------+--+
             |                          |(4)
             |                          |
           +-v-+              +---------v--+
           |403|              |Admission   |
           +---+              |controllers |
                              +------------+
```

## Authorization modes

* Node - special purpose; grants permissions to kubelets based on the pods they are scheduled to run
* ABAC - access rights granted through policies which combine attributes together
* RBAC - regulates access based on the roles of individuals users; see `rbac.md`
* Webhook - HTTP callback (POST); allows for integration with external authorizers

## Tooling and good practices

Tooling - see https://kubernetes-security.info/#authorization

Use RBAC (`--authorization-mode=RBAC`)

Disable automounting of the default service account token (`automountServiceAccountToken: false`)
* especially important if you're not using RBAC
* most apps don't need to talks to the API server
```
$ kubectl patch serviceaccount default -p $'automountServiceAccountToken: false'
```

Use dedicated service accounts
* if a pod is compromised, attacker can access serviceaccount associated with that pod
* if your app needs API access create a dedicated service account and RBAC for it

## EKS authentication and authorization

Steps 1. - 4. -> authentication; 5., 6. -> authorization

1. `kubectl` runs `aws eks get-token ...` (to see it do `cat $KUBECONFIG`) to get a bearer token.
2. The token is passed to the kube-apiserver which forwards it to the authentication webhook.
3. The webhook calls the pre-signed URL that is base64-encoded in the token’s body.
4. The URL validates the signature and returns user info (user’s account, ARN, user ID).
5. The kube-apiserver reads the `aws-auth` cm to associate user with an RBAC group.
6. RBAC groups are referenced in [Cluster]RoleBindings.

# Securing container images

* software you run on a cluster gets there in the form of container images
* images must not include known critical [vulnerabilities](https://nvd.nist.gov/)
* images must be the ones you intended to use and mustn't have been manipulated

# Scanning

* to detect vulns you need to use container image scanner
* rescan (every 24 hours) because new vulnerabilities are found all the time
* some registries can do the scanning 
 
[Image sccanners](https://kubernetes-security.info/#securing-your-container-images) do the following
* at a minimum, look at the installed (yum, apt) packages
* examine files installed (ADD, COPY, or RUN) during build time
* detect known malware (e.g., viruses)
* detect sensitive data (like passwords and tokens)

## Patching

* when a vuln is found update the container with a fixed version of the package
* rebuild and redeploy the new container image (usually done via CI/CD)

## CI/CD

<img width="545" alt="image" src="https://user-images.githubusercontent.com/1047259/205614637-6f42337a-83cb-40ee-9e51-f9d830cf85a0.png">

A failed scan in CI/CD can
* prevent the image from being pushed to registry
* prevent the image from being deployed
* send an alert

Sample image policy: fail all images with high-severity vulns

## Storage

* unless you are pulling public images, you need to grant registry access to your cluster
* use read-only accounts if your cluster doesn't need to push images to registry

## Correct versions

* pod spec references container image by using registry, owner, repo and version, e.g. `gcr.io/myname/myimage:1.0.1`
* tags (`1.0.1`) are mutable - the same tag can be moved to refer to a different image
* images can have multiple tags
* you can use unique digest instead of tag (gcr.io/myname/myimage@sha256:4a5573037f358b6cdfa2...) but then you need to update YAML whenever there's a new version
* it's much more common to *use semver tags*
* if you supply neither tag nor digest, the image tagged `latest` is used (should be avoided because you won't know what version you are running)
* use `AlwaysPullImages` admission controller to ensure you have most recent version for the given tag

## Trust and supply chain

* in high-risk or high-security environments you need to make sure that the pulled image is the genuine, intended code
* see https://kubernetes-security.info/#running-containers-securely for tools

## Minimizing attack surface

* as a general rule, the smaller the image, the smaller the attack surface
* you rarely need to [include ssh](https://jpetazzo.github.io/2014/06/23/docker-ssh-considered-evil/)
* if you exclude cat, more etc. the attacker can't read (easily) the credentials
* if you exclude shells the attacker can't do anything  (easily)
* on the other hand, troubleshooting will be harder

# Running containers securely

* use least privilege to carry out the task at hand
* do only minimal host mounts necessary
* limit communication between apps, and to/from outside word

## Say no to root

There's [little](https://opensource.com/article/18/3/just-say-no-root-containers) [need](https://youtu.be/ltrV-Qmh3oY) to run containers as root:

* container needs to modify the host system (e.g. kernel configurations)
* container needs to bind to privileged ports (below 1024) -> can be avoided via port mappings and service abstraction
* installing packages at container runtime -> bad practice anyway, because you can't scan the packages

Use `USER` command in Dockerfile.

## Admission control

After authn and authz API server perfmorms admission control before persisting the request to etcd. From the admission controllers included in the API these are security relevant:

AlwaysPullImages
* modifies every new pod to Always pull policy
* by default when an image is pulled to a node it can be accessed by other pods on the node bypassing registry credentials check

DenyEscalatingExec
* denies exec and attach commands
* prevents attackers from launching interactive shells in (possibly privileged) containers

PodSecurityPolicy
* determines whether a pod should be admitted based on the security context and policies

LimitRange, ResourceQuota
* ensures resource constraints
* preventiing DoS attacks

NodeRestriction
* limits permissions of each kubelet

## Security boundaries

[Security boundary](https://cloud.google.com/blog/products/gcp/exploring-container-security-isolation-at-different-layers-of-the-kubernetes-stack/) - set of controls to provent a process affecting other processes or accessing other users' data.

Cluster
* network isolation
* you might prefer different cluster for each team or environment

Node
* resource isolation
* use nodeSelector or even better node or pod affinity to assign pods to specific nodes

Namespace
* virtual cluster
* basic unit of authorization

Pod
* groups containers; they are scheduled on the same node
* you can isolate pods via security context or network policies

Container
* combination of cgoups, namespaces and copy-on-write filesystems that manages the application-level dependencies
* unless you're using runtime sandboxing, no strong isolation beyond the kernel-level security ones

## Policies

Kubernetes offers two pod-level security-policy mechanisms that can restrict what processes can do within a pod and how pods can communicate.

Security context

* defines privilege and access control settings on pod or container level

```
# all containers must run as 1001
# in webserver container prevent setuid binaries changing the effective user ID as well as prevent files from enabling extra capabilities
apiVersion: v1
kind: Pod
metadata:
  name: securepod
spec:
  securityContext:
    runAsUser: 1001
  containers:
  - name: webserver
    image: quay.io/mhausenblas/pingsvc:2
    securityContext:
      allowPrivilegeEscalation: false
  - name: shell
    image: centos:7
    command:
    - "/bin/bash"
    - "-c"
    - "sleep 10000"
```

Security policies

* allow cluster or namespace admin to enforce security context
* cluster-wide resource used by admission controllers (PodSecurityPolicy plugin must be enabled)
* NOTE: replaced by pod security admission or a 3rd party admission plugin

Network policies

* see network-policy.md
* public clouds pass config info to nodes via Metadata API that can cause security issues like the one disclosed in [Shopify bug bounty](https://hackerone.com/reports/341876)

# Secrets management

* your application often needs access to secrets like credentials

## Applying the principle of least privilege

* containers should be able to read only secrets they need
* use different secrets for different environments (dev credentials can be shared with more people than prod credentials)

## Secret encryption

* [at rest](https://kubernetes.io/docs/tasks/administer-cluster/encrypt-data/) - so attacker with access to the FS cannot simply read them from a file
* in transit (TLS) - so attacker snooping on the network cannot read them as they pass

etcd

* the default storage
* base64 encoded (not encrypted!)

Using 3rd party stores (like cloud provider's KMS or HashiCorp Vault) is considered by many a more secure solution.

## Passing secrets into containerized code

(1) build them into the image itself - don't do this
* anyone who has access to the image can get secrets 
* if you want to change secret, like DB credentials, you need to rebuild the image -> possible downtime
* it's probably [git versioned](https://www.infoworld.com/article/3064355/how-you-might-be-leaking-your-secrets-onto-github.html)

(2) passing them as environment variables
* [12-factor app](https://12factor.net/config) has tought us to separate config from code
* it's helpful when you need to run the same code in different scenarios (laptop, dev, prod)
* however it's easy to leak environment via logs or command line (contemplate if this is an issue for your application or organization)
  * a crashed process often dumps entire environment
  * `kubectl describe pod ...` shows environment
  * `docker inspect ...` shows environment

(3) passing them in files
* via volume mounted into the container
* if the volume is a temporary FS it's even better because it stays in memory (it's good practice to never store secrets in plain text)

## Secret rotation and revocation

* the longer a given secret remains valid, the more likely it has been compromised
* it's [no longer considered best practice](https://www.ncsc.gov.uk/blog-post/problems-forcing-regular-password-expiry) but not for secrets used by machines
* but you should have a way to change secrets when compromised
* depending on how your code is written, you may need to restart a pod (e.g. when read from envvar/file as part of initialization)
* application might reread the secret on regular basis or in response to a failure with currently held value -> Kubernetes can update file secret without restarting pod (this is not true for environment secrets)

## Secret access from within the container and from kubelet

* if attacker gains execution access to a container there's high likelihood they accessed secrets
* mitigations: runtime protection, slim images (without cat, more, less, bash, sh)
* --enable-admission-plugins=NodeRestriction - kubelet can access only secrets of pods scheduled to its node

# Sources and more

* Kubernetes Security (2018)
* https://kubernetes-security.info/
