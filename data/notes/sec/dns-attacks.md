# DDoS attacks

## Amplification attacks

* a small query can trigger a large response
* this query is 48 bytes (seen in Wireshark) and the response is 3390 bytes

```
$ dig @ns1.isc.org. any isc.org. +norec +dnssec | grep -i size
;; MSG SIZE  rcvd: 3390
```

* so a bot with 1Mbps connection can make a DNS server to generate cca 70Mb each second (3390/48)
* a botnet with 15 such bots can make the DNS server saturate a 1Gbps network (15*70)

## Reflection attacks

* queries with spoofed (victim's) source IP address
* the victim will get unsolicited (often amplified) responses

## Combination attacks

* spoofed source IP address + query that will result in a large payload
* authoritative server provides the amplification, recursive server provides the reflection

# Cache poisoning

* corrupting the cached answers on the recursive name servers
* either through software bugs (vendor dependent) or protocol weaknesses
* protocol weakness 1: UDP is lightweight but it's easier to spoof than TCP
* protocol weakness 2: the only answer field that's not easy to spoof is Query ID (aka TXID) but it has not enough randomness
* TXID is 16 bits large and thus can be guessed

<img src="/static/dns-cache-poisoning.jpeg" style="max-width:100%;width:320px">

1. The attacker has prior knowledge of the target domain and sends a query to the recursive DNS server for a name that does not exist, such as q0001xxx.example.com
2. Because this is a name that does not exist, the recursive DNS server must traverse the DNS namespace to find it.
3. The attacker can beat the legitimate NXDOMAIN response from the authoritative name server, by sending a lot of spoofed responses that look like they are coming from the legitimate example.com authoritative name server. In the spoofed response, attacker claims www.example.com is the NS record of the domain, to trick the recursive name server into accepting www.example.com and its IP address.
4. By the laws of probability, the attacker’s spoofed response may be accepted by the recursive server, and the bad answer www.example.com is now stored in its cache.
5. Unsuspecting client queries for the name www.example.com A record.
6. The recursive server provides the answer from the now-poisoned cache with the forged answer from the attacker.

See Cloudflare's [article](https://www.cloudflare.com/learning/dns/dns-cache-poisoning/) for more.

# Data exfiltration via DNS tunneling

DNS tunneling allows for

* getting free airport WiFi
* use SSH over DNS to get through corporate FW
* stealing sensitive information and malware finding command and control points

In the last case, clients evade detection by breaking data down into
query-sized chunks, disguising sensitive data as DNS queries, and sending
them to malicious DNS servers on the far end who can unpack these queries and
reconstruct the data.

<img src="/static/dns-data-exfiltration.jpeg" style="max-width:100%;width:640px">

# More

* https://www.infoblox.com/dns-security-resource-center/
* https://medium.com/@alex.birsan/dependency-confusion-4a5d60fec610 - exfiltrating data via DNS (2021)
