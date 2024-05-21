Helm is a templating engine and package manager for the Kubernetes "operating system". The main concepts are:

* Chart: a Helm package
* Release: a running instance of a Chart
* Repository: place where Charts are kept (like CPAN)

Chart looks like

    $ tree
    .
    ├── Chart.yaml  # describes the chart (name, description, version, author)
    ├── templates   # K8s manifests with templating directives
    │   ├── deployment.yaml
    │   └── service.yaml
    └── values.yaml # defaults

(You can try out the following commands on https://github.com/cloudnativedevops/demo/tree/main/hello-helm3/k8s/demo)

Template locally and display on console

    helm template .

Bundle .tgz archive file

    helm package .

Add a Repository

    helm search hub wordpress # or https://artifacthub.io
    helm repo add bitnami https://charts.bitnami.com/bitnami
    helm repo update
    helm repo list

Install Release with default values

    helm search repo wordpress
    helm install happy-panda bitnami/wordpress

Install another Release with customized values

    echo '{mariadb.auth.database: user0db, mariadb.auth.username: user0}' > my-values.yaml
    helm install --generate-name -f my-values.yaml bitnami/wordpress

Upgrade existing Release

    helm upgrade happy-panda -f my-values.yaml bitnami/wordpress
    helm get values happy-panda

Rollback existing Release

    helm rollback happy-panda 1

Get status and list of Releases

    helm status happy-panda
    helm list --all

Uninstall Release

    helm uninstall happy-panda

MORE: Cloud Native DevOps with Kubernetes, ch 12.
