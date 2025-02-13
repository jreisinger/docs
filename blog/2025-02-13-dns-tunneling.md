## DNS tunneling

A technique called DNS tunneling allows data to be transmitted covertly using the Domain Name System (DNS). Attackers can leverage DNS tunneling for:

- *Data exfiltration* – Stealing sensitive information by embedding it in DNS queries.
- *Malware communication* – Enabling malware to reach its command and control (C2) servers without detection.

In these scenarios, attackers break data into small chunks, disguise them as DNS queries, and send them to malicious DNS servers. These servers then reconstruct the original data, enabling undetected data leakage.

Here are the steps an attacker needs to take to exfiltrate data:

![image](https://github.com/user-attachments/assets/d10b0ba7-8f75-4156-a7b2-5174b0094240)

And this is a proof of concept that demonstrates these steps:

![Screen Recording 2025-02-13 at 15 54 42](https://github.com/user-attachments/assets/5b83a9a4-34de-4a86-ade8-cb57e423ef2e)

You can find the code [here](https://github.com/jreisinger/pocs/tree/main/dns/exfil).

## Detecting DNS Tunneling Using DNS Server Logs

Detecting DNS tunneling is challenging because not all tunneling activity is inherently malicious. If an attacker exfiltrates data slowly over time using only a few DNS packets, detection becomes even harder. However, network defenders can analyze DNS traffic for anomalies.

### FQDN Entropy

The higher the entropy, the higher the likelihood of a malicious Fully Qualified Domain Name (FQDN).

```
ENTROPY   FQDN
2.646439  google.com
2.646439  golang.org
2.721928  amazon.com
4.016876  asdlfkjasdflwerjka.t1.security.local
```

Legitimate domains, such as those used by CDNs and social media platforms, may also have high entropy, so additional context is required.

### FQDN Length

Abnormally long domain names can indicate tunneling activity. Anomalies are often identified when domain name lengths exceed *three times the average request length*, although thresholds can vary between organizations.

### Unusual DNS Record Types

Certain DNS record types can signal suspicious activity:

- *AAAA (IPv6 DNS Records):* While normal in IPv6 environments, unexpected use could indicate tunneling.
- *AXFR (Zone Transfers):* Unless explicitly allowed, zone transfers suggest an attacker may be footprinting the network.
- *DNSSEC Records:* These are used for DNS authentication, but unexpected DNSSEC traffic may indicate tampering or misuse.

### High Volume of DNS Requests

DNS tunnels generate *large volumes of DNS queries* when transferring files. A sudden spike in DNS request traffic could indicate tunneling activity.

### Geographic Anomalies

If an organization does not conduct business in a specific country, but DNS requests are frequently being sent there, it may signal suspicious activity.
