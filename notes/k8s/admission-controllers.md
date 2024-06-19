[admission controller](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) is a piece of code that intercepts API requests to the `kube-apiserver` before they are stored into `etcd`

types
- validating
- mutating
- both

can limit requests to: create, delete, modify objects
can't limit requests to read: get, watch, list

out of the [list](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/#what-does-each-admission-controller-do) of admission controllers compiled into `kubec-apiserver` two are special
- MutatingAdmissionWebhook
- ValidatingAdmissionWebhook
these execute [admission control webhooks](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/#admission-webhooks) which are configured (registered) in the API at run time