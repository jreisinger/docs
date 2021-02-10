# DDoS attacks

## Aplification attacks

* a small query can trigger a large response
* this query is 48 bytes and the response is 3390 bytes (as seen in Wireshark)

```
$ dig @ns1.isc.org. any isc.org. +norec +dnssec | grep -i size
;; MSG SIZE  rcvd: 3390
```

* so a bot with 1Mbps connection can make a DNS server to generate cca 70Mb each second (3390/48)

## Reflection attacks

# Cache poisoning

# Data exfiltration

# More

* https://www.infoblox.com/dns-security-resource-center/
