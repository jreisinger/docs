Security is a neverending process of preventing and handling security incidents.

There's no secure system. There are just more or less secure systems.

# Goals

CIA triad represents the traditional (since [1977](https://nvlpubs.nist.gov/nistpubs/Legacy/SP/nbsspecialpublication500-19.pdf)), security goals:

> The protection of system data and resources from accidental and deliberate threats to confidentiality, integrity and availability.

* Confidentiality: no unauthorized read access to data or systems
* Integrity: no unauthorized write access to data or systems
* Availability: data and service available when needed

Sometimes non-repudation is added to these three.

The ultimate security goals don't change with the adaption of a new paradigm (e.g. cloud services or DevOps). Security teams must still focus on reducing business risk from attacks and work to get confidentiality, integrity, and availability (CIA) security controls built into information systems and data. How those goals are achieved will change.

# Principles

These security principles will help you to increase your security:

* Simplicity - [complexity](https://www.schneier.com/blog/archives/2022/08/security-and-cheap-complexity.html) is the worst enemy of security
* Minimal attack surface - minimize possible ways a system can be attacked
* Least privilege - deny by default to limit the blast radius of a compromise
* Defense in depth - since any security control can fail, multiple overlapping layers are needed

# Areas

Asset and risk management

* what are your most valuable assets 
* what are most probable attacks and their impact

Identity and access management (IAM) - life cycle of 

* identities (authn)
* access rights (authz)

Data encryption

* in motion (PKI, TLS)
* at rest (secrets mngt)

Vulnerability mngt - detect security issues in

* network and systems from outside (Nexpose)
* applications from outside - DAST (OWASP ZAP)
* application and infrastructure code - SAST (SonarQube, tfsec)
* application image and container (trivy) scanning

Network security

* WAF
* antiDDoS
* IDS/IPS
* honepots

See also [CISSP](https://en.wikipedia.org/wiki/Certified_Information_Systems_Security_Professional) domains cyberseek [roles](https://www.cyberseek.org/pathway.html).

# Basic steps

Understand the business of the organization you are trying to protect.

Think about what you need to protect (assets: VMs, containers, DBs) and who is most likely to cause problems (threat actors: criminals, hacktivists, script kiddies, inside attackers, state actors).

Understand what is your responsibility - this depends on the cloud model you are using and whether you are a consumer or provider:

<img src="https://user-images.githubusercontent.com/1047259/138699080-24091008-c78f-48c1-bcc9-e9ac6afd0f8d.png" style="max-width:100%;height:auto;"> 

Figure out what needs to talk to what in your application. You should first secure places where line crosses a trust boundary:

<img src="https://user-images.githubusercontent.com/1047259/138698724-4a6ecae8-fe54-4d45-b7a8-3b35dfab50e1.png" style="max-width:100%;height:auto;"> 

Anything inside a trust boundary can trust, at least to some level, anything else inside that boundary but requires verification before trusting anything outside that boundary. If an attacker gets into a part of a trust boundary, she'll probably have control over all trust boundary eventually. So getting through each trust boundary should take some effort.

Know your risks (have at least a spreadsheet) and how you approach them:

* avoid the risk - turn off the system, benefits < risk
* mitigate the risk - apply some security measures
* transfer the risk - pay someone else to manage the risk (e.g. insurance)
* accept the risk - benefits > risk

# Continuous security

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
