L4

```
# try to connect to <host> and <port> with timeout 2 seconds
for i in {1..5}; do echo -n "connection $i: "; nc -vz -w 2 <host> <port>; done
```

L7

```
# try to get <url> with timeout 2 seconds
for i in {1..5}; do echo -n "connection $i: "; curl <url> -I --max-time 2 && true; done
```

Check connectivity to the Internet:

```
while true; do curl https://ifconfig.me; echo; done
```
