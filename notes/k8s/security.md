# Kubernetes attack vectors

```
                  +---------------------------------+
                  | Cluster                         |
Access to         |   +--------------------------+  |
machines or VMs   |   | Node                     |  |
------------------+-->|           +------+       |  |   Access to etcd API
                  |   |           | etcd |<------+--+----------------------
                  |   |           +------+       |  |
                  |   |                          |  |
Access via        |   |    +----------------+    |  |   Intercept/modify/inject
K8s API or proxy  |   |    |  Control plane |    |  |   control-plane traffic
------------------+---+--->|  components    +----+--+--------------------------
                  |   |    +----------------+    |  |
                  |   +--------------------------+  |
                  |                                 |
                  |   +--------------------------+  |
Access via        |   | Node                     |  |
Kubelet API       |   |         +-----------+    |  |
------------------+---+-------->| Kubelet   |    |  |
                  |   |         +-----------+    |  |
                  |   |                          |  |
                  |   | +----------------------+ |  |
                  |   | |Pod                   | |  |
                  |   | | +------------------+ | |  |
Escape container  |   | | | Container        | | |  |  Intercept/modify/inject
to host through   |   | | |  +-------------+ | | |  |  application traffic
<-----------------+---+-+-+--+ Application +-+-+-+--+------------------------
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
$ k run -it --rm jumpod --restart=Never --image=alpine -- sh
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

# Securing container images

* software you run on a cluster gets there in the form of container images
* images must not include known critical [vulnerabilities](https://nvd.nist.gov/)
* images must be the ones you intended to use and mustn't been manipulated

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

A failed scan in CI/CD can
* result in a failed build
* prevent the image from being pushed to registry
* prevent the image from being deployed
* send an alert (if the image is already in production)

<img width="545" alt="image" src="https://user-images.githubusercontent.com/1047259/205614637-6f42337a-83cb-40ee-9e51-f9d830cf85a0.png">

## Storage

## Correct versions

## Trust and supply chain

## Minimizing attack surface

# Running containers securely

# Secrets management

# Sources and more

* Kubernetes Security (2018)
* https://kubernetes-security.info/
