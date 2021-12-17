# Areas

Security governance

* what are most valuable assets (data and compute resources identification and classification)
* what are most probable attacks and attackers (threat actors)
* estimate where are the major risks - impact x probability (risk assessment)
* public cloud brings shared responsibility - identify your responsibilities
* educate other engineers - they should know the security basics

Identity and access management (IAM)

* life cycle of identity (authentication) and access rights (authorization)
* concerns both humans and programs

Vulnerability management

* Network vulnerability scanning
* Application scanning (DAST)
* Code scanning (SAST)
* Image/container scanning

Network security

* encryption in motion (PKI, TLS)
* encryption at rest (secrets management)
* WAF
* antiDDoS
* IDS/IPS, honeypots

Detecting (monitoring) and responding to security incidents (attacks)

* logs, metrics => SIEM
* SIRP, SOC

# Approach

The ultimate objectives don't change with the adaption of cloud services. Security teams must still focus on reducing business risk from attacks and work to get confidentiality, integrity, and availability (C.I.A.) security controls built into information systems and data. How those objectives are achieved will change.

Static security processes can't keep up with the pace of change in cloud platforms, the threat environment and the evolution of security technologies. Security must shift to a continuosly evolving approach to match this pace.

Continuos security

* DevOps (and [Agile](http://agilemanifesto.org/) and [Deming](https://deming.org/explore/fourteen-points)) focuses on shipping better products to *customers* faster
* many security teams focus on compliance with a security standard, number of security incidents and vulnerabilities on production systems
* both are valid but different goals; this creates conflict and hurts the organization
* solution => "continuos security" (or DevSecOps): allign goals by switching focus of security team from defending only the infrastructure to protecting the entire organization by improving it continuously
* techniques from traditional security, such as vulnerability scanning, intrusion detection, and log monitoring, should be reused and adapted to fit in the DevOps pipeline

<img src="https://user-images.githubusercontent.com/1047259/141968423-133c5f24-6c1e-4eaf-89e0-167fae88c31e.png" style="max-width:100%;height:auto;"> 

* (1) TDS - security testing should be handled the same way application tests are handled in the CI and CD pipelines: automatically and all the time
* (2) Monitoring and responding to attacks - 2nd phase of continuous security; fraud and intrusion detection, digital forensics, and incident response
* (3) Assessing risks and maturing security - go beyond the technology and look at the organization’s security posture from a high altitude; risk management

# More

* Practical Cloud Security (2019)
* Securing DevOps (2018)
* M$ Cloud Adoption Framework - security: https://docs.microsoft.com/en-us/azure/cloud-adoption-framework/secure/
* Attack tactics and techniques: https://attack.mitre.org/
