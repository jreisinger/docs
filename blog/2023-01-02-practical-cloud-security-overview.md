![](https://user-images.githubusercontent.com/1047259/222765764-e826697a-0f33-4240-892f-db39265adbbc.png)

IT industry has been undergoing seismic shifts regularly. The latest one goes by the name cloud computing. However, the security goals, principles and areas don't change with the adaption of a new paradigm. Terminology, methods and tools change, though.

# Why we need security and what are its goals

> The art of war teaches us to rely not on the likelihood of the enemy’s not coming, but on our own readiness to receive him; not on the chance of his not attacking, but rather on the fact that we have made our position unassailable. -- Sun Tzŭ: The Art of War

The need for security stems from the fact that people, for various reasons, make suboptimal decisions and mistakes - in design, implementation, configuration, operations - creating vulnerabilities in systems. And, some other people (threat actors or groups) try to exploit these vulnerabilities for various reasons. Since it's impossible to prevent all the vulnerabilities in complex systems, you can't achieve perfect security. What is your north star, then?

Information security, or cybersecurity, is a never ending systematic effort to control security risks. A risk is the possibility of something bad happening. And threat is a path to the risk occurring. In case you don't know about your risks and threats or you don't manage the risks, you can run into [troubles](https://www.hackmageddon.com/). Your data can be stolen or encrypted for ransom. Your infrastructure or applications can be shut down, misused for crypto-mining or launching attacks against other targets. Your customers or employees can get scammed. All these lead to operational problems (you have to handle security incidents instead of normal business), reputation and financial losses (you lose customers and/or get fined). Obviously, you want to avoid this at least to a certain degree.

More formally, CIA triad represents the traditional (since [1977](https://nvlpubs.nist.gov/nistpubs/Legacy/SP/nbsspecialpublication500-19.pdf)) security goals: "The protection of system data and resources from accidental and deliberate threats to confidentiality, integrity, and availability."

* confidentiality: no unauthorized access to data or resources
* integrity: no unauthorized changes to data or resources
* availability: data and resources available when needed

# Security principles and areas

Ok, so how do we manage the security risks and protect against threats? First, there are some principles everybody should keep in mind and apply whenever possible:

* simplicity - the worst enemy of security is [complexity](https://www.schneier.com/blog/archives/2022/08/security-and-cheap-complexity.html), it makes attacks easier and defense harder
* minimal attack surface - minimize possible ways a system can be attacked
* least privilege - deny by default to limit the blast radius of a compromise
* segmentation - create boundaries between systems to limit the blast radius of a compromise
* defense in depth - since any security control can fail have multiple overlapping layers of controls

Second, you have to care about many things because the weakest link in the chain of interconnected systems can get exploited. Not to get overwhelmed one might create some abstractions in the form of distinct areas to cover.

## Governance and risk management

> If you know the enemy and know yourself, you need not fear the result of a hundred battles. -- Sun Tzŭ: The Art of War

You should get at least a rough understanding of your organization's business and products. Find out your [responsibility](https://docs.aws.amazon.com/wellarchitected/latest/security-pillar/shared-responsibility.html) boundaries and what data and compute/storage/network resources you need to protect. This depends on the service model you use or provide (IaaS, PaaS or SaaS). Get some idea [who](https://github.com/jreisinger/docs/blob/master/notes/sec/threat-actors.md) is most likely to cause problems and [how](https://attack.mitre.org).

Risk level is the likelihood of a risk times its impact. You can approach each risk in one of these ways:

* avoid it - don't build the system in the first place or turn it off if benefits are lower than risks
* transfer it - pay someone else to manage the risk (e.g. SaaS, insurance)
* mitigate it - apply some security measures (controls)
* accept it - if benefits are higher than risks (this should be conscious)

You might need to prove your security to a 3rd party; this is called [compliance](https://aws.amazon.com/compliance/).

## Identity and access management

If an attacker gets credentials all patches and firewalls won't help. Manage user and program identities (authn) and access rights (authz) in as few places as possible. Have process of removing users that left the company. Make sure that strong passwords and MFA are used. Use a password manager (1Password) and don't commit unencrypted passwords or API keys to repositories (gitleaks). Access rights (roles, policies) should follow the least privilege principle.

## Vulnerability management

Detect and remediate security bugs and misconfigurations in application (SonarQube, ZaP, trivy) and infrastructure (tfsec) code, systems and networks (Nexpose). Before ([SAST, SCA](https://github.blog/2022-09-09-sca-vs-sast-what-are-they-and-which-one-is-right-for-you/)) and after (DAST) deployment. Important point to emphasize here is to make sure that the vulnerabilities found by the scanners are also remediated not only reported. So detect them as soon as possible in the process of developing and deploying code and infrastructure. First handle only the critical ones. Code reviews and penetration testing is helpful but expensive. Regularly upgrade (patch) your systems and dependencies.

## Security monitoring

You want to know what's going on and then do something about. Detect threats and security incidents, and respond to them. You do this by first collecting and parsing logs and metrics in a central place (Splunk, Graylog, Datadog). Then you create alerts (log/metric queries with a threshold) and handle them when they get triggered. Find a good balance between too many and too few alerts. Prefer simplicity and quality over cleverness and quantity to avoid alert fatigue.

## Network security

If you can't talk to a component, you can't compromise it. Use network policies, ACLs, VPNs, WAFs, antiDDoS, IDS/IPS when it makes sense. Try to create trust boundaries. Anything inside a trust (or security) boundary can trust, at least to some level, anything else inside that boundary but requires verification before trusting anything outside that boundary. Also (almost always) encrypt data in motion using [TLS](https://github.com/jreisinger/docs/blob/master/blog/gosec/2023-09-26-go-for-cybersecurity-learning.md#what-is-tls---learning-by-reading).
