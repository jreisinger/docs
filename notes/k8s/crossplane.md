# Crossplane

![](https://docs.crossplane.io/content/media/crossplane-intro-diagram_hud9dc847ee0e2ab0b53319b680d79d1fd_55780_1200x0_resize_q75_h2_box_3.webp)

---

## Kubernetes components

![](https://kubernetes.io/images/docs/components-of-kubernetes.svg)

---

## kube-apiserver

Kubernetes API represents Kubernetes resources as HTTP endpoints

```
k get pods -v6
...
I1021 07:19:46.292131 1415509 round_trippers.go:553] GET https://127.0.0.1:6443/api/v1/namespaces/default/pods?limit=500 200 OK in 10 milliseconds
...
```

---

## Kubernetes resources

* `k api-resources`
* definition - some structured data: `k explain pods`
* controller - infinite loop moving current state to desired state (like thermostat)

---

## Kubernetes API is extensible

* custom resources -> CRDs
* custom controllers

---

## Crossplane

- framework to transform your Kubernetes into universal control plane
- you can manage external resources through standard Kubernetes APIs
- it uses CRDs and custom controllers underneath

---

## But why

- also external resources managed via Kubernetes 
- complexity hidden from users (developers)
- security or compliance enforcement across all resources

---

## Installation

```
helm install crossplane --namespace crossplane-system --create-namespace crossplane-stable/crossplane
kubectl get pods -n crossplane-system
k get crds
```

---

## Main Crossplane components

* Provider - creates and manages additional CRDs representing external service
* ProviderConfig - configuration of the provider (authn)
* Managed Resource - Kubernetes resource representing external resource

---

## AWS provider

```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-aws-s3
spec:
  package: xpkg.upbound.io/upbound/provider-aws-s3:v1.1.0
EOF

k get providers | grep aws
k get crds | grep aws
```

---

## Managed resource

```
cat <<EOF | kubectl create -f -
apiVersion: s3.aws.upbound.io/v1beta1
kind: Bucket
metadata:
  generateName: crossplane-bucket-
spec:
  forProvider:
    region: us-east-2
  providerConfigRef:
    name: default
EOF
```

---

## ProviderConfig

```
cat <<EOF | kubectl apply -f -
apiVersion: aws.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: crossplane-system
      name: aws-secret
      key: creds
EOF
```

---

## Where do providers come from

* use existing ones: https://marketplace.upbound.io/providers
* write your own: https://github.com/grantgumina/provider-pizza
* template: https://github.com/crossplane/provider-template
