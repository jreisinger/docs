> My view of security areas to cover (2021-11 PN, coming to an end :-)

(This is not how to do security but what to cover)

Risk assessment and threat modelling

* what are most valuable assets, most probable attacks and attackers (threat actors)
* estimate where are the major risks (impact x probability)
* public cloud brings shared responsibility - identify your responsibilities

Identity and access management (IAM)

* life cycle of identity and access
* authentication and authorization

Vulnerability management

* Network vulnerability scanning
* Application scanning (DAST)
* Code scanning (SAST)
* Image/container scanning

Network security

* encryption in motion (PKI, TLS)
* WAF
* AntiDDoS
* IDS/IPS

Detecting (monitoring) and responding to security incidents (attacks)

* logs, metrics, SIEM
* SIRP, SOC

> M$ Cloud Adoption Framework - [Secure](https://docs.microsoft.com/en-us/azure/cloud-adoption-framework/secure/)

The ultimate objectives for a security organization don't change with the adaption of cloud services. How those objectives are achieved will change. Security teams must still focus on reducing business risk from attacks and work to get confidentiality, integrity, and availability security controls built into all information systems and data.

<img src="https://user-images.githubusercontent.com/1047259/143010672-671723ff-85ee-4b9a-a6e1-e45ad50eef97.png" style="max-width:100%;height:auto;"> 

* static security processes can't keep up with the pace of change in cloud platforms, the threat environment and the evolution of security technologies
* security must shift to a continuosly evolving approach (continuos security) to match this pace

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

* DevOps (and [Agile](http://agilemanifesto.org/) and [Deming](https://deming.org/explore/fourteen-points)) focuses on shipping better products to *customers* faster
* many security teams focus on compliance with a security standard, number of security incidents and vulnerabilities on production systems
* both are valid but different goals; this creates conflict and hurts the organization
* solution => "continuos security" (or SecDevOps): allign goals by switching focus of security team from defending only the infrastructure to protecting the entire organization by improving it continuously
* techniques from traditional security, such as vulnerability scanning, intrusion detection, and log monitoring, should be reused and adapted to fit in the DevOps pipeline

<img src="https://user-images.githubusercontent.com/1047259/141968423-133c5f24-6c1e-4eaf-89e0-167fae88c31e.png" style="max-width:100%;height:auto;"> 

* (1) TDS - security testing should be handled the same way application tests are handled in the CI and CD pipelines: automatically and all the time
* (2) Monitoring and responding to attacks - 2nd phase of continuous security; fraud and intrusion detection, digital forensics, and incident response
* (3) Assessing risks and maturing security - go beyond the technology and look at the organization’s security posture from a high altitude; risk management
