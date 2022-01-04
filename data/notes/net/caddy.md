# Sharing files

Prerequisites

* A registered public domain name (e.g. myserver.example.com)
* External access to ports 80 and 443
* caddy [binary](https://github.com/caddyserver/caddy/releases) (see Assets)

Configuration

```
# Caddyfile
myserver.example.com
file_server browse
root * /data/public
basicauth /* {
        # password hash created with `caddy hash-password`
        Bob JDJhJDE0JENVblJPOVljYml4a3phTHNVelpmYk90MVVEdXV2aVFjYkl2ODJENDFEaG1KU29TRGNCUHp5
}
log {
        format json
        output file /home/ubuntu/access.log
}
```

Running

```
sudo ./caddy run
```

Parsing logs

```
# get remote IP addresses
cat access.log | jq -r '.request.remote_addr' | cut -d':' -f1 | sort | uniq

# get remote IP addresses excluding 4xx response status 
cat access.log | jq -r '"\(.request.remote_addr) \(.status)"' | grep -vE ' 4..$' | cut -d':' -f1 | sort | uniq
```

Or use recent version of [goaccess](https://goaccess.io/).

# More

* https://caddyserver.com/docs/
