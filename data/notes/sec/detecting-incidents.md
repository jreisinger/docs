It's important to identify a breach ASAP. According to some studies the mean time is around 200 days.

* log or event - record of specific thing that happened (e.g. log about an HTTP request)
* metric - time based number giving info about something (e.g. number of HTTP requests per second)
* alert - event or metric where the system decides it's worth notifying someone

Logs are more difficult/costly to store and process.

Each log should contain when, what, and who.

# What (logs and metrics) to watch

* depends on your threat model (what assets you have and who is most likely to attack them) and what logs/metrics come out of your systems
* you should have mutliple detection layers

Privileged user access - to detect unauthorized person pretending to be an admin

Logs from defensive tooling (WAFs, Anti-DDoS, FWs, IDS/IPS, Antivirus, Honeypots)

Services, OSs, middleware

* CPU usage metrics: increase -> ransomware encryption or cryptomining
* Network logs/metrics: traffic spike -> DoS attack or attacker stealing data
* Storage I/O metrics: increase -> ransomware, DoS, attacker stealing data
* DB, message queue metrics: increase -> attacker stealing data or sends messages to other components

Secret server

* you should log all access to secrets
* unusual activity: authn and authz failures, high secrets retrieval, use of administrative credentials

Honeypots

* distract and slow down attackers and alert you
* advanced technique; after you have logging, monitoring, alerting and response running effectively

# How to watch

These steps may be all done by a single product (e.g. SIEM), or by multiple products/services acting together.

Logging and alerting chain:

![image](https://user-images.githubusercontent.com/1047259/139064744-9e542ce1-8e1b-437c-a136-f9fd2dab8a78.png)

Unfortunately, there are thousands of different log formats. A few common formats: syslog (atlhough the body/message is free-form), CLF, ELF, CEF, CADF.

Search/correlation examples: search for all login failures during certain time period, all successful logins with a VPN, malware detection followed by a login.

Alerting is where the art lies in log analysis. You need a balance between too many false positives and no alerts at all. You need a feedback loop to know whether to modify (increase threshold) or remove an alert. Consider running periodic tests that will generate alerts.

There are some alerts that you should always follow up; e.g. multiple login failures for privileged users, malware detected. When logs stop flowing is a security issue too!

Automated responses have potential to disrupt your business. Also can be leveraged by attackers - an easy DoS attack using a simple port scanner or a few failed logins.

Rotate different individuals in and out. You need some way to ensure that an alert is acknowledged within a certain amount of time or escalated to someone else to handle.

In many cases, organizations use a hybrid model where some of the lower-level monitoring and alerting is performed by a MSSP (managed security service provider)/SOC, and the more important alerts are escalated to in-house staff.

Sample SIEM alerts:

* “Database traffic is up 200% from the monthly average. Maybe the application is just really popular right now, or is someone systematically stealing our data?”
* “We just saw an outbound connection to an IP address that has been used by a known threat actor recently, according to this threat intelligence feed. Is that a compromised system talking to a command-and-control server?”
* “There were 150 failed login attempts on an account, followed by a success. Is that a successful brute-force attack?”
* “We saw a single failed login attempt on 300 different accounts, followed by a success on account #301. Is that a successful password spraying attack?”
* “A port scan was followed by a lot of traffic from a port that hasn’t been used in months. Port scans happen all the time, but perhaps a vulnerable service was found and compromised?”
* “John doesn’t normally log in at 3:00 AM ET, or from that country. Maybe that’s not really John?”
* “Three different accounts logged in from the same system over the course of 30 minutes. It seems unlikely all of those people are actually using that system, so maybe the system and those accounts are compromised?”
* “A new administrative account was just created outside of normal business hours. Maybe someone’s working late, but maybe there’s an issue?”
* “Someone was just added to the administrator group. That’s a rare event, so shouldn’t we check on it?”
* “Why are there firewall denies with an internal system as the source? Either something is misconfigured or there’s an unauthorized user trying to move around the network.”

Source: Practical Cloud Security (2019)
