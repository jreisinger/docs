# What is it

Cert-manager creates and renews TLS certificates for your Kubernetes workloads. It can obtain certificates from a variety of issuers (CAs). With the Certificate resource the private key (`tls.key`) and the certificate (basically a public key signed with private key of a CA; `tls.crt`) are stored in a Kubernetes secret which is mounted by a Pod or used by an Ingress controller.

<img width="756" height="438" alt="image" src="https://github.com/user-attachments/assets/3e82e74a-f240-44a2-bc53-0a48085581c9" />

It's implemented as a set of CRDs and controllers.

# Installation

```
# CRDs
$ k apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.19.0/cert-manager.yaml
customresourcedefinition.apiextensions.k8s.io/challenges.acme.cert-manager.io created
customresourcedefinition.apiextensions.k8s.io/orders.acme.cert-manager.io created
customresourcedefinition.apiextensions.k8s.io/certificaterequests.cert-manager.io created
customresourcedefinition.apiextensions.k8s.io/certificates.cert-manager.io created
customresourcedefinition.apiextensions.k8s.io/clusterissuers.cert-manager.io created
customresourcedefinition.apiextensions.k8s.io/issuers.cert-manager.io created
<...>

# controllers
$ k get deploy -n cert-manager
NAME                      READY   UP-TO-DATE   AVAILABLE   AGE
cert-manager              1/1     1            1           119s
cert-manager-cainjector   1/1     1            1           119s
cert-manager-webhook      1/1     1            1           119s
```

# Resources

Issuer (or ClusterIssuer)
- represents CA able to sign a certitficate in response to certificate (signing) request
- see https://cert-manager.io/docs/configuration/selfsigned for a practical example

CertificateRequest
- used to request X.509 certificates from an Issuer
- see https://cert-manager.io/docs/usage/certificaterequest for more

Certificate
- human readable definition of a certificate request
- cert-manager uses this input to generate a private key and CertificateRequest to obtain a signed certificate from an Issuer
- the signed certificate and private key are then store in the specified Secret
- cert-manager makes sure the certificate is auto-renewed before it expires

# Ingress certificates

- simply add `cert-manager.io/cluster-issuer: nameOfClusterIssuer` annotation to an Ingress metadata
- see https://cert-manager.io/docs/usage/ingress for details
