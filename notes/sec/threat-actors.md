Actor, Motivation, Capability, Sample attacks.

```
A: Vandal: script kiddie.
M: Curiosity, personal fame from bringing down service or exfiltrating data of a high-profile company.
C: Uses public tools (nmap, Metasploit, CVE PoCs), some experimentation. Attacks are poorly concealed. Low level of targeting.
S: Small-scale DoS. Plants trojans. Launches prepackaged exploits for access, crypto mining.
```

```
A: Motivated individual: political activist, thief, terrorist.
M: Gain from selling exfiltrated data for fraud. Personal kudos or spreading political messages by DDoSing or defacing large public-facing services.
C: May combine publicly available exploits in a targeted fashion. Modify open source supply chains. Concealing attacks of minimal concern.
S: Phishing. DDoS. Exploit known vulnerabilities. Compromise open source projects to embed code to exfiltrate environment variables and secrets when code is run by users. Exfiltrated values are used to gain system access and perform crypto mining.
```

```
A: Insider: employee, external contractor, temporary worker.
M: Discontent, profit. Personal gain from selling large amounts of personal data for fraud, or making small alterations to the integrity of data in order to bypass authentication for fraud. Ecnrypt data volumes for ransom.
C: Detailed knowledge of the system, understands how to epxloit it, conceals actions.
S: Uses privileges to exfiltrate data (to sell on). Misconfiguration or codebombs to take service down as retribution.
```

```
A: Organized crime: sindicates, state-affiliated groups.
M: High motivation for crypto-ransomware, mass extraction of PII/credentials/PCI data, manipulation of transactions for financial gain. 
C: Ability to devote considerable resources, hire "authors" to write specialized tools and exploits. Some ability to bribe/coerce/intimidate individuals. Level of targeting varies. Conceals until goals are met.
S: Social engineering/phishing. Ransomware (becoming more targeted). Cryptojacking. RATs (in decline). Coordinated atacks using multiple exploits, possibly using a single zero-day or assisted by a rogue individual to pivot through infrastructure (e.g. Carbanak).
```

```
A: Cloud service insider: employee, external contractor, temporary worker.
M: Personal gain, curiosity. Uknown level of motivation.
C: Depends on cloud provider's segregation of duties and technical controls.
S: Access to or manipulation of datastores.
```

```
A: Foreign Intelligence Service (FIS): nation states.
M: Intelligence gathering, disrupt critical national infrastructure, unknown. May steal intellectual property, access sensitive systems, mine personal data en masse, or track down individuals through location data held by the system.
C: Disrupt of modify HW/SW supply chains. Ability to infiltrate organizations/suppliers, call upon research programs, develop multiple zero-days. Highly targeted. High level of concealment.
S: Stuxnet (multiple zero-days, infiltration of 3 organizations including 2 PKIs with offline root CAs), SUNBURST (targeted supply chain attack, infiltration of hundreds of organizations).
```

NOTE: threat actors can be hybrids of different categories.

Source: Hacking Kubernetes (2021)
