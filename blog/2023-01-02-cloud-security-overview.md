# Practical cloud security overview

Information security is a never ending systematic effort to reduce security risks. Security risks stem from the fact that humans make mistakes and other humans want to exploit these mistakes. This means that information security is relevant everywhere. But it's not practical to secure everything so you must define areas you want to cover.

I have some experience securing computing infrastructure and applications. These days it goes by the name of cloud (native) computing and can be visualized as [4Cs](https://kubernetes.io/docs/concepts/security/overview/#the-4c-s-of-cloud-native-security). These four layers define tech stack areas relevant to the cloud computing. But I'll talk about security areas that you need to consider.

## Risk management and governance

Have at least basic idea about what data and compute/storage/network resources you are protecting. Think (and write down) which of these assets are the most critical and what are the greatest risks. You don't need to spend too much time on this especially if you have only a small security team. However, if you skip this area altogether you are just guessing.

You can't expect all employees to be security experts but you should make sure they have basic security awareness, e.g. produce regular Slack posts about relevant topics. All engineers should adhere to basic security principles:

* minimal attack surface - minimize possible ways a system can be attacked
* least privilege - deny by default to limit the blast radius of a compromise
* defense in depth - since any security control can fail have multiple overlapping layers
* simplicity (usually not considered a security principle) - complexity makes attacks easier and defense harder

You might need to prove your security to a 3rd party; this is called compliance. It is much easier if you have actually secured your systems and data.

## Identity and access management (authn, authz)

Manage user and program identities (usernames) and access rights in as few places as possible. Have process of removing users that left the company. Make sure that strong passwords and MFA are used. Use a password manager (1password) and don't commit unencrypted secrets (passwords, API keys) to repositories. Access rights (roles, policies) should follow the least privilege principle.

## Vulnerability management

Detect and remediate security bugs and misconfigurations in application (SonarQube, ZaP, trivy) and infrastructure (tfsec) code, systems and networks (Nexpose). Important point to emphasize here is to make sure that the vulnerabilities found by the scanners are also remediated not only reported. So detect them as soon as possible in the process of developing and deploying code and infrastructure and (at first) handle only the critical and high ones.

## Network security

If you can't talk to a component, you can't compromise it. Use network policies, ACLs, WAFs when it makes sense. Also (almost always) encrypt data in motion using TLS.

## Security monitoring (SIEM)

You want to know what's going on. Detect and respond to security incidents. You do this by first collecting and parsing logs and metrics in a central place. Then you create alerts (a log/metric query with a threshold) and handle them when they get triggered. Find a good balance between too many and too few alerts. Prefer quality over quantity to avoid alert fatigue.
