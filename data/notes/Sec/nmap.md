Show supported authentication methods:

```
$ nmap -p 22 --script ssh-auth-methods scanne.nmap.com
Starting Nmap 7.91 ( https://nmap.org ) at 2021-04-26 18:15 CEST
Nmap scan report for scanne.nmap.com (45.33.49.119)
Host is up (0.16s latency).
rDNS record for 45.33.49.119: ack.nmap.org

PORT   STATE SERVICE
22/tcp open  ssh
| ssh-auth-methods:
|   Supported authentication methods:
|     publickey
|     gssapi-keyex
|     gssapi-with-mic
|_    password

Nmap done: 1 IP address (1 host up) scanned in 3.22 seconds
```
