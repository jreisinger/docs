# Practical cloud security overview

![](https://user-images.githubusercontent.com/1047259/222765764-e826697a-0f33-4240-892f-db39265adbbc.png)

Information security, or cybersecurity, is a never ending systematic effort to manage security risks. The risks mostly stem from the fact that people tend to make suboptimal decisions and mistakes - in design, processes, code, configurations, operations - creating vulnerabilities in systems. Some other people try to exploit these vulnerabilities. Probably the most popular way to handle the computing infrastructure and applications these days goes by the name cloud (native) computing. The following security principles and goals don't change with the adaption of a new paradigm, like cloud computing or DevOps.

The **principles** that help mitigate the security risks:

* [simplicity](https://www.schneier.com/essays/archives/1999/11/a_plea_for_simplicit.html) - the worst enemy of security is [complexity](https://www.schneier.com/blog/archives/2022/08/security-and-cheap-complexity.html), it makes attacks easier and defense harder
* minimal attack surface - minimize possible ways a system can be attacked
* least privilege - deny by default to limit the blast radius of a compromise
* segmentation - create boundaries between systems to limit the blast radius of a compromise
* defense in depth - since any security control can fail have multiple overlapping layers
* ability to restore data and systems - backups, IaC, documentation, fire drills

CIA triad represents the traditional (since [1977](https://nvlpubs.nist.gov/nistpubs/Legacy/SP/nbsspecialpublication500-19.pdf)) security **goals**: "The protection of system data and resources from accidental and deliberate threats to confidentiality, integrity, and availability."

* confidentiality: no unauthorized read access to data or systems
* integrity: no unauthorized write access to data or systems
* availability: data and service available when needed

To not get overwhelmed one might create some abstractions in the form of distinct areas to cover.

## Governance and risk management

You should get at least a rough understanding of your organization's business. Understand what is your responsibility (IaaS, PaaS, SaaS) and what data and compute/storage/network resources you are protecting. Get some idea which of these assets are the most critical (valuable) and [who](https://attack.mitre.org/groups) is most likely to cause problems and [how](https://attack.mitre.org).

Risk is the possibility of something bad happening. Risk level is the likelihood of a risk times its impact. Threat is a path to the risk occurring. You can approach each risk in one of these ways:

* avoid it - don't build or turn off the system if benefits < risk
* transfer it - pay someone else to manage the risk (e.g. SaaS, insurance)
* mitigate it - apply some security measures (controls)
* accept it - if benefits > risk (this should be conscious)

You might need to prove your security to a 3rd party; this is called compliance.

## Identity and access management

If an attacker gets credentials all patches and firewalls won't help. Manage user and program identities (auth) and access rights (authz) in as few places as possible. Have process of removing users that left the company. Make sure that strong passwords and MFA are used. Use a password manager (1password) and don't commit unencrypted passwords or API keys to repositories. Access rights (roles, policies) should follow the least privilege principle.

## Vulnerability management

Detect and remediate security bugs and misconfigurations in application (Sonar[Qube|Cloud], ZaP, trivy) and infrastructure (tfsec) code, systems and networks (Nexpose). Before ([SAST, SCA](https://github.blog/2022-09-09-sca-vs-sast-what-are-they-and-which-one-is-right-for-you/)) and after (DAST) deployment. Important point to emphasize here is to make sure that the vulnerabilities found by the scanners are also remediated not only reported. So detect them as soon as possible in the process of developing and deploying code and infrastructure. First handle only the critical ones. Code reviews and penetration testing is helpful but expensive.

## Security monitoring

You want to know what's going on and then do something about. Detect threats and security incidents, and respond to them. You do this by first collecting and parsing logs and metrics in a central place. Then you create alerts (a log/metric query with a threshold) and handle them when they get triggered. Find a good balance between too many and too few alerts. Prefer quality over quantity to avoid alert fatigue.

## Network security

If you can't talk to a component, you can't compromise it. Use network policies, ACLs, WAFs, antiDDoS, IDS/IPS, honeypots when it makes sense. Try to create trust boundaries. Anything inside a trust (or security) boundary can trust, at least to some level, anything else inside that boundary but requires verification before trusting anything outside that boundary. Also (almost always) encrypt data in motion using TLS.
