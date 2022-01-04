# Areas

Security governance

* what are your (most valuable) assets => data and compute resources identification and classification ([asset management](https://danielmiessler.com/blog/continuous-asset-management-security/))
* what are most probable attacks and attackers (threat actors) targeting those assets
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

# Concepts

CIA triad - security goals

* Confidentiality - seeks to prevent unauthorized read access to data
* Integrity - seeks to prevent unauthorized write access to data
  * data integrity (modification of data in DB)
  * system integrity (malicious SW to open "back door" to OS) 
* Availability - ensures that information and service is available when needed

Least privilege

* concerns people and automated tools 
* should be able to access only what they need to do their job, and no more
* deny by default

Defense in depth

* acknowledgement that almost any security control can fail
* thus you need multiple layers of overlapping security controls

Risk management

* risk is something bad that could happen
* its level is based on its likelihood to happen and its impact
 
# Approach

The ultimate security objectives don't change with the adaption of a new paradigm (e.g. cloud services or DevOps). Security teams must still focus on reducing business risk from attacks and work to get confidentiality, integrity, and availability (C.I.A.) security controls built into information systems and data. How those objectives are achieved will change.

## Basic steps

Think about what you need to protect (assets: VMs, containers, DBs) and who is most likely to cause problems (threat actors: criminals, hacktivists, script kiddies, inside attackers, state actors).

Understand what areas you need to secure - this depends on the cloud model you are using and whether you are a consumer or provider:

<img src="https://user-images.githubusercontent.com/1047259/138699080-24091008-c78f-48c1-bcc9-e9ac6afd0f8d.png" style="max-width:100%;height:auto;"> 

Figure out what needs to talk to what in your application. You should first secure places where line crosses a trust boundary:

<img src="https://user-images.githubusercontent.com/1047259/138698724-4a6ecae8-fe54-4d45-b7a8-3b35dfab50e1.png" style="max-width:100%;height:auto;"> 

Know your risks (have at least a spreadsheet) and how you approach them:

* avoid the risk - turn off the system, benefits < risk
* mitigate the risk - apply some security measures
* transfer the risk - pay someone else to manage the risk (e.g. insurance)
* accept the risk - benefits > risk

## Continuos security

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