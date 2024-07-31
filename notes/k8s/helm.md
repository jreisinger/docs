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

Install 3rd party chart

    # 1. Search chart (e.g. wordpress) on https://artifacthub.io
    # 2. Add repository containing the chart:
    helm repo add bitnami https://charts.bitnami.com/bitnami
    # 3. Install the chart: 
    helm install my-wordpress bitnami/wordpress --version 23.0.11

Install another release with customized values

    echo '{mariadb.auth.database: user0db, mariadb.auth.username: user0}' > my-values.yaml
    helm install --generate-name -f my-values.yaml bitnami/wordpress

Upgrade existing release

    helm upgrade my-wordpress -f my-values.yaml bitnami/wordpress
    helm get values my-wordpress

Rollback existing release

    helm rollback my-wordpress 1

List releases and get status

    helm list --all
    helm status my-wordpress

Uninstall release

    helm uninstall my-wordpress

MORE: Cloud Native DevOps with Kubernetes, ch 12.
