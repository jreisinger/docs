Security is the ability to resist attack.

Security is elusive and hard to measure.

There's no secure system. There are just more or less secure systems.

You need some level of security. To achieve it you need patience, vigilance, knowledge and persistence.

Security is a neverending process.

# Goals

CIA triad represents the traditional (since [1977](https://nvlpubs.nist.gov/nistpubs/Legacy/SP/nbsspecialpublication500-19.pdf)) security goals:

![image](https://user-images.githubusercontent.com/1047259/148757897-f51c3d58-8f26-46d8-973b-854ae47d84d4.png)

Confidentiality

* seek to prevent unauthorized read access to data

Integrity

* seek to prevent unauthorized write access to data
* data integrity (modification of data in DB)
* system integrity (malicious SW opens a "back door")
 
Availability

* ensure that information and service is available when needed

Sometimes [non-repudation](https://en.wikipedia.org/wiki/Non-repudiation) is added to these three.

# Principles

These security principles will help you to increase your security.

Clarity and simplicity

* obscurity and complexity in code, processes and communication hide (security) problems
* easy to understand and easy to use security controls and tools are more likely to be used (correctly)

Limit attack surface

* attack surface is all possible ways a system can be attacked (through user, system, network)
* higher complexity means bigger attack surface
* bigger attack surface means higher probability of getting attacked
* apply to exposed systems, unnecessary services, underused systems

Defense in depth

* acknowledgement that almost any security control can fail
* thus you need multiple layers of overlapping security controls
* adding another layer is also more cost effective than perfecting a single layer

Least privilege

* people and programs (and their modules) should be able to access only what they need to do their job, and no more
* this limits the blast radius when system gets compromised 
* deny by default
* apply to any access control or authorization situation (e.g. FW rules, user groups, file permissions) 

# Areas

Asset and risk management

* what are your (most valuable) data and compute/storage/network resources ([asset management](https://danielmiessler.com/blog/continuous-asset-management-security/))
* what are most probable attacks and attackers (threat actors) targeting those assets
* estimate where are the major risks (something bad that could happen) - impact x probability (risk assessment)

Identity and access management (IAM)

* life cycle of identity (authentication) and access rights (authorization)
* concerns both humans and programs

Data encryption

* encryption in motion (PKI, TLS)
* encryption at rest (secrets management)

Vulnerability management

* Network vulnerability scanning (Nexpose)
* Application scanning, DAST (OWASP ZAP)
* Code scanning, SAST (SonarQube - code, tfsec - IaC)
* Image/container scanning (trivy)

Network security

* WAF
* antiDDoS
* IDS/IPS, honeypots

Security monitoring

* detecting and responding to security incidents
* security incidents = attacks
* SIEM (logs, metrics), SIRP, SOC

See also [CISSP](https://en.wikipedia.org/wiki/Certified_Information_Systems_Security_Professional) domains.

# Approach

The ultimate security goals don't change with the adaption of a new paradigm (e.g. cloud services or DevOps). Security teams must still focus on reducing business risk from attacks and work to get confidentiality, integrity, and availability (CIA) security controls built into information systems and data. How those goals are achieved will change.

## Basic steps

Understand the business of the organization you are trying to protect.

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

## Continuous security

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

* Full Stack Python Security (2021)
* Practical Cloud Security (2019)
* Securing DevOps (2018)
* Kubernetes Security (2018)
* ULSAH 5th (2017)
* Attack tactics and techniques: https://attack.mitre.org/
