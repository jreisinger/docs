If you can't talk to a component at all, it's difficult to compromise.

If everything were configured absolutely perfectly, you could safely have no network controls at all.

IP whitelisting is still important as long as it isn't the primary defense or auth method:

* whitelist - list of things that are allowed, with everything else denied
* blacklist - list of things that are denied, with everything else allowed

You place simpler, less-trusted components in DMZ (proxy, LB, static content web server); if that component is compromised, it should not provide large advantage to the attacker.

Proxies (forward, reverse) are useful for both functional and security requirements. Although you can have proxy for almost any protocol there are usually HTTP/HTTPS proxies.

Network features virtualisation (NFV) or virtual network functions (VNFs) reflect the idea you no longer need a dedicated box for FW, router or IDS/IPS.

Overlay network is a virtual network often accomplished by encapsulation (VxLAN, GRE, IP-in-IP).

Virtual Private Cloud (VPC) despite its name generally deals only with network isolation. Allows you to keep the majority of your app in private area reachable only by you. Usually implemented via SDN and/or overlay network.

NAT is used heavily in cloud. Originally to combat shortage of IP addresses by using the same RFC1918 ipaddr in many parts of the internet and translating them to publicly routable addresses. NAT is not security but it implies existence of a FW.

IPv6 security improvements: mandatory support for IPsec transport security, cryptographically generated addresses, larger address space difficult to scan.

**Prioritize** network controls in the following order (don't put bars on second-story windows before your doors have a lock).

1) Encryption in motion

* use TLS (1.3) for all communication that crosses a switch (physical or virtual)
* not needed for components within OS or pod
* you will need to include new ciphersuites and remove compromised ones - does not happen that often
* TLS loses most of its effect if you don't authenticate the other end by certificate checking (MITM attack)
* this means you need to create a separate keypair and getting a certificate signed - painful and difficult to automate -> getting easier (HashiCorp Vault, Istio)

2) FW implementations in the cloud

* Virtual FW appliances - largely lift-and-shift model from on-prem environments
* Network access control list (NACLs) - instead of operating a FW you simply define rules
* Security groups - similar for NACLs but at per-OS or per-pod level

The first FW control you should design is a perimeter of some form.

3) Admin access

* bastion hosts
* VPN (site-to-site, client-to-site)

4) WAF and RASP

* extra layer of protection against common programming errors
* smart proxy
* you need to setup the rules properly and look at the alerts
* make sure all traffic passes through WAF (whitelisting, but can be difficult)

5) Anti-DDoS

* too many fake requests or too much useless traffic deny service for legitimate users
* consider if anyone is going to care to knock you off the Internet and how big a problem is it for you if they do

6) IDS/IPS

* detect/block traffic matching a signature or behavior (e.g. connections to a lot of ports -> scanning)
* network or host based

Egress filtering protects against data exfiltration and watering hole attack (e.g. you pull update from a malicious server).

Data loss prevention (DLP) watches for sensitive data that are improperly stored or are leaving the environment.

Source: Practical Cloud Security (2019)
