ECS
- no control plane to manage, no patching or upgrading
- no cost for cluster, only for compute and other infra you use to run containers

task
- defines image location, amount of CPU or RAM, logging
- doesn't launch a container, it's just configuration (like pod YAML)

service
- running container
- defines number of container replicas (copies)
- defines AZs to deploy container to
- set up ALB to forwards requests from the Internet to your service

![](https://d1.awsstatic.com/getting-started-guides/gsg-build-ecs-1.ad1b412728f68e0293191d7a298111e7387f953c.png)