# Lessons learned

* (information) security == infosecurity == cybersecurity
* There's no 100% secure system. There are just more or less secure systems.
* In many environments development and operations are prioritized over security. Until a security incident happens.
* Security is a long term systematic process.
* "Ensure your doors are locked sucurely before putting bars on your second-store windows!"

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

Data, asset and risk management

* what are your (most valuable) data and (compute, storage, network) assets
* risk = possibility of something bad happening
* risk level = likelihood x impact
* threat = a path to that risk occuring
* what are most probable threats against your data and assets and their impact 
* how to handle these threats
* encrypt (some) data at rest (secrets mngt)

Identity and access management (IAM)

* if an attacker has (admin!) credentials all patches and firewalls won't help
* life cycle of authn (identities) and authz (access rights, roles)

Vulnerability mngt - detect security issues in

* network and systems from outside (Nexpose)
* applications from outside - DAST (OWASP ZAP)
* application and infrastructure code - SAST (SonarQube, tfsec)
* application image and container (trivy)
---
* code reviews
* penetration testing

Network security

* if you can't talk to a component, you can't compromise it
* WAF
* antiDDoS
* IDS/IPS
* honeypots
* encryption in motion (PKI, TLS)

Security monitoring (SIEM)

* detecting and responding to security incidents
* unfortunately you won't always be susccefull at protecting your assets
* in 2022, it took an average of 9 months to identify and contain a [breach](https://www.ibm.com/reports/data-breach)

Compliance

* proving your security to a 3rd party
* much easier if you have actually secured your systems and data

See also [CISSP](https://en.wikipedia.org/wiki/Certified_Information_Systems_Security_Professional) domains and cyberseek [roles](https://www.cyberseek.org/pathway.html).

# First steps

(1) Understand the business of the organization you are trying to protect.

(2) Think about what you need to protect (data, assets) and who is most likely to cause problems (criminals, hacktivists, script kiddies, inside attackers, state actors).

(3) Understand what is your responsibility - this depends on the cloud model you are using and whether you are a consumer or provider:

<img src="https://user-images.githubusercontent.com/1047259/138699080-24091008-c78f-48c1-bcc9-e9ac6afd0f8d.png" style="max-width:100%;height:auto;"> 

(4) Know your risks (have at least a spreadsheet) and how you approach them:

* avoid the risk - turn off the system, benefits < risk
* mitigate the risk - apply some security measures
* transfer the risk - pay someone else to manage the risk (e.g. insurance)
* accept the risk - benefits > risk

(5) Figure out trust (or security) boudaries. Draw what needs to talk to what:

<img src="https://user-images.githubusercontent.com/1047259/207269071-8fffd922-7fe5-4bdd-8172-944cc5a470a7.png" style="max-width:100%;height:auto;">

Anything inside a trust boundary can trust, at least to some level, anything else inside that boundary but requires verification before trusting anything outside that boundary. If an attacker gets into a part of a trust boundary, she'll probably have control over all trust boundary eventually. So getting through each trust boundary should take some effort. An example of a trust boundary is an application container (docker).

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

# Sources and more

* Practical Cloud Security (2019)
* Securing DevOps (2018)
* Attack tactics and techniques: https://attack.mitre.org/
