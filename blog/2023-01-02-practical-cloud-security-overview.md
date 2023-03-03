# Practical cloud security overview

![image](https://user-images.githubusercontent.com/1047259/222765764-e826697a-0f33-4240-892f-db39265adbbc.png)

Information security, or cybersecurity, is a never ending systematic effort to reduce security risks. Security risks stem from the fact that people make mistakes (in code, configurations, processes, architecture). Some other people try to exploit these vulnerabilities.

# Goals

CIA triad represents the traditional (since [1977](https://nvlpubs.nist.gov/nistpubs/Legacy/SP/nbsspecialpublication500-19.pdf)) security goals: "The protection of system data and resources from accidental and deliberate threats to confidentiality, integrity and availability."

* Confidentiality: no unauthorized read access to data or systems
* Integrity: no unauthorized write access to data or systems
* Availability: data and service available when needed

Sometimes non-repudation is added to these three.

The ultimate security goals don't change with the adaption of a new paradigm, like cloud computing or DevOps. Security teams must still focus on reducing business risk from attacks and work to get confidentiality, integrity, and availability (CIA) security controls built into information systems and data. How those goals are achieved will change.

# Areas

To not get overwhelmed one needs to create some abstractions in the form of distinct areas. The most popular way to handle the computing infrastructure and applications these days goes by the name cloud (native) computing and can be visualized as [4Cs](https://kubernetes.io/docs/concepts/security/overview/#the-4c-s-of-cloud-native-security). These four layers define tech stack areas relevant to the cloud computing. But I'll talk about security areas that you need to consider.

## Risk management and governance

Have at least basic idea about what data and compute/storage/network resources you are protecting. Think (and write down) which of these assets are the most critical (valuable) and what are the greatest risks against them. You don't need to spend too much time on this especially if you have only a small security team. However, if you skip this area altogether you are just guessing.

You can't expect all employees to be security experts but you should make sure they have basic security awareness; e.g. produce regular Slack posts about relevant topics. All engineers should adhere to basic security principles that help to decrease risks:

* simplicity - [complexity](https://www.schneier.com/blog/archives/2022/08/security-and-cheap-complexity.html) is the worst enemy of security; it makes attacks easier and defense harder
* minimal attack surface - minimize possible ways a system can be attacked
* least privilege - deny by default to limit the blast radius of a compromise
* segmentation - create boundaries between systems to limit the blast radius of a compromise
* defense in depth - since any security control can fail have multiple overlapping layers
* ability to restore data and systems - backups, IaC, documentation, fire drills

You might need to prove your security to a 3rd party; this is called compliance. It is much easier if you have actually secured your systems and data :-).

## Identity and access management

If an attacker has (admin!) credentials all patches and firewalls won't help. Manage user and program identities (auth) and access rights (authz) in as few places as possible. Have process of removing users that left the company. Make sure that strong passwords and MFA are used (especially for admin accounts). Use a password manager (1password) and don't commit unencrypted secrets (passwords, API keys) to repositories. Access rights (roles, policies) should follow the least privilege principle.

## Vulnerability management

Detect and remediate security bugs and misconfigurations in application (Sonar[Qube|Cloud], ZaP, trivy) and infrastructure (tfsec) code, systems and networks (Nexpose). Before (SAST) and after (DAST) deployment. Important point to emphasize here is to make sure that the vulnerabilities found by the scanners are also remediated not only reported. So detect them as soon as possible in the process of developing and deploying code and infrastructure and (at first) handle only the critical and high ones. Code reviews and penetration testing is helpful but hard.

## Network security

If you can't talk to a component, you can't compromise it. Use network policies, ACLs, WAFs, antiDDoS, IDS/IPS, honeypots when it makes sense. Also (almost always) encrypt data in motion using TLS.

## Security monitoring

You want to know what's going on and then do something about. Detect threats and security incidents, and respond to them. You do this by first collecting and parsing logs and metrics in a central place. Then you create alerts (a log/metric query with a threshold) and handle them when they get triggered. Find a good balance between too many and too few alerts. Prefer quality over quantity to avoid alert fatigue.

# First steps

It can be overwhelming to know where to start. Because the weakest link matters, plus there are dependencies between systems that can be chain-exploited. Ensure your doors are locked securely before putting bars on your second-store windows!

Understand the business of the organization you are trying to protect.

Think about what you need to protect, who (threat actors and [groups](https://attack.mitre.org/groups)) is most likely to cause problems and how ([threat](https://attack.mitre.org) [models](https://microsoft.github.io/Threat-Matrix-for-Kubernetes)).

Understand what is your responsibility. This depends on the cloud model you are using (IaaS, PaaS, SaaS) and whether you are a consumer or provider.

Know your risks and how you approach them:

* avoid the risk - turn off the system, benefits < risk
* transfer the risk - pay someone else to manage the risk (e.g. insurance, SaaS)
* mitigate the risk - apply some security measures
* accept the risk - benefits > risk

When looking at a system figure out trust boundaries. Draw what needs to talk to what. Anything inside a trust (or security) boundary can trust, at least to some level, anything else inside that boundary but requires verification before trusting anything outside that boundary. An example of a trust boundary is an application container (docker).

# Vocabulary 

Risk - possibility of something bad happening

Threat - a path to the risk occurring

Risk level = likelihood x impact
