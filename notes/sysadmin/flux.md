Flux deploys changes to Kubernetes by polling a Git repository from flux pods inside the same Kubernetes (flux-system namespace).

Install flux:

    eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)" # on Linux
    brew install fluxcd/tap/flux

Create a GitHub personal token: https://github.com/settings/tokens

Check prerequisites:

    export GITHUB_TOKEN=...
    export GITHUB_USER=...
    flux check --pre

Install pods and create a GitHub repo:

    flux bootstrap github \
      --owner $GITHUB_USER \
      --repository flux-demo \
      --branch main \
      --path ./clusters/demo-cluster \
      --personal

Create new deployment:

    git clone git@github.com:$GITHUB_USER/flux-demo.git
    cd flux-demo
    mkdir clusters/demo-cluster/flux-demo
    cd clusters/demo-cluster/flux-demo

    k create namespace flux-demo -o yaml --dry-run=client > namespace.yaml
    k create deployment flux-demo -n flux-demo -o yaml --dry-run=client --image=cloudnatived/demo:hello > deployment.yaml

    git add namespace.yaml deployement.yaml
    git commit -m "create flux-demo deployment"
    git push

Other flux functionalities:

* can manage Helm releases
* can manage kustomize manifests
* can poll a container registry, deploy new image and make the image version commit back to the repo
