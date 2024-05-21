Ideally, developers should be able to push code to the source control repository and all of the build, test, and deploy phases happen automatically in a centralized pipeline.

Continuos integration (CI) is the automatic integration and testing of developer's changes against the mainline branch.

Continuos deployment (CD) is the automatic deployment of successful builds to production. Triggered by

* pushing a button
* merging a merge request
* pushing a Git release tag

Pipeline is the CI/CD machinery.

Some hosted CI/CD tools

* GitHub actions
* GitLab CI

Some self-hosted CI/CD tools

* Jenkins
* Argo - continuously "pulls" in changes from the Git repo and applies them from within the cluster

More: Cloud Native Devops with Kubernetes, Ch. 14
