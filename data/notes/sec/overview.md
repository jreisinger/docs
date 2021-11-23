> Practical Cloud Security (2019)

Assets management

* data identification and classification
* compute, storage, network inventory

Identity and access management

Secrets management

Vulnerability management

Network security

Detecting, responding to, and recovering from security incidents

> Securing DevOps (2018)

Continuos security

* DevOps (and [Agile](http://agilemanifesto.org/) and [Deming](https://deming.org/explore/fourteen-points)) focuse is on shipping better products to *customers* faster
* many security teams focus on compliance with a security standard, number of security incidents and vulnerabilities on production systems
* both are valid but different goals; this creates conflict and hurts the organization
* solution => "continuos security" (or SecDevOps): allign goals by switching focus of security team from defending only the infrastructure to protecting the entire organization by improving it continuously
* techniques from traditional security, such as vulnerability scanning, intrusion detection, and log monitoring, should be reused and adapted to fit in the DevOps pipeline

![image](https://user-images.githubusercontent.com/1047259/141968423-133c5f24-6c1e-4eaf-89e0-167fae88c31e.png)

* (1) TDS - security testing should be handled the same way application tests are handled in the CI and CD pipelines: automatically and all the time
* (2) Monitoring and responding to attacks - 2nd phase of continuous security; fraud and intrusion detection, digital forensics, and incident response
* (3) Assessing risks and maturing security - go beyond the technology and look at the organization’s security posture from a high altitude; risk management
