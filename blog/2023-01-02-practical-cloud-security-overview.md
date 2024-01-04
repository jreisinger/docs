![](https://user-images.githubusercontent.com/1047259/222765764-e826697a-0f33-4240-892f-db39265adbbc.png)

IT industry has been undergoing seismic shifts regularly. The latest one goes by the name cloud computing. However, the security goals, principles and areas don't change with the adaption of a new paradigm. Terminology, methods and tools change, though.

# Why we need security

The need for security stems from the fact that people make suboptimal decisions and mistakes - in design, implementation, configuration, operations - creating vulnerabilities in systems. And some other people (threat actors or groups) try to exploit these vulnerabilities for various reasons.

In case you don't know or don't care about your vulnerabilities and threats, you can run into trouble. Your data can be stolen, altered, deleted or encrypted for ransom. Your infrastructure or applications can be shut down, misused for crypto-mining or for launching attacks against other targets. Your customers or employees can get scammed.

This creates all sorts of problems. Operational problems, when employees have to handle security incidents instead of normal business. Loss of reputation and consequently of employees and customers. The company can get fined or someone (like a CISO) can even go to jail. All this gets reflected also in terms of financial losses. Obviously, you want to avoid this at least to a certain degree.

# Security goals

It's impossible to not create any vulnerabilities and you can't make the bad actors disappear. So what can we do then? Well, we can avoid at least some of the vulnerabilities and we can handle or contain at least some of the exploits.

More formally, cybersecurity is a never ending effort to manage risks of computer systems by increasing their resiliency to threats. A threat is the possibility of something bad happening; it's the combination of a vulnerability and a threat actor. A risk is the quantified refinement of a threat; it's defined by its likelihood and impact.

CIA triad represents the traditional (since [1977](https://nvlpubs.nist.gov/nistpubs/Legacy/SP/nbsspecialpublication500-19.pdf)) security goals: "The protection of system data and resources from accidental and deliberate threats to confidentiality, integrity, and availability."

* confidentiality: no unauthorized access to data or resources
* integrity: no unauthorized changes to data or resources
* availability: data and resources available when needed

# Security principles

OK, so how do we achieve the security goals? First, there are some principles everybody should keep in mind and apply whenever possible:

* simplicity - the worst enemy of security is [complexity](https://www.schneier.com/blog/archives/2022/08/security-and-cheap-complexity.html) because it makes attacks easier and defense harder
* minimal attack surface - minimize possible ways a system can be attacked
* least privilege - deny by default to limit the blast radius of a compromise
* segmentation - create boundaries between systems to limit the blast radius of a compromise
* defense in depth - since any security control can fail have multiple overlapping layers of controls
* secure by design - security shouldn't be an afterthought because then it's much more expensive

# Security areas

Second, you have to care about many things because the weakest link in the chain of interconnected systems can get exploited (security is a systems property). Not to get overwhelmed one might create some abstractions in the form of distinct areas to cover.

## Governance and risk management

> If you know the enemy and know yourself, you need not fear the result of a hundred battles. -- Sun Tzŭ: The Art of War

You should get at least a rough understanding of your organization's business and products. Find out your [responsibility](https://docs.aws.amazon.com/wellarchitected/latest/security-pillar/shared-responsibility.html) boundaries and what data and compute/storage/network resources you need to protect. This depends on the service model you use or provide (IaaS, PaaS or SaaS). Get some idea [who](https://github.com/jreisinger/docs/blob/master/notes/sec/threat-actors.md) is most likely to cause problems and [how](https://attack.mitre.org).

Risk level is the likelihood of a future problem times its impact. You can approach each risk in one of these ways:

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
