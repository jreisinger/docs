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
* the victim will get unsollicited (often amplified) responses

## Combination attacks

* spoofed source IP address + query that will result in a large payload
* authoritative server provides the amplification, recursive server provides the reflection

# Cache poisoning

# Data exfiltration

# More

* https://www.infoblox.com/dns-security-resource-center/
