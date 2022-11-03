Simple to setup and run, nice logging (JSON), automatic TLS certificate management!

# Prerequisites

* A registered public domain name (e.g. something.example.com)
* External access to ports 80 and 443
* caddy [binary](https://github.com/caddyserver/caddy/releases) (see Assets)

# Configuration

`Caddyfile`:

```
example.com something.example.com

# Redirect
redir https://example.org

# Sharing files
file_server browse
root * /data/public
basicauth /* {
        # password hash created with `caddy hash-password`
        Bob JDJhJDE0JENVblJPOVljYml4a3phTHNVelpmYk90MVVEdXV2aVFjYkl2ODJENDFEaG1KU29TRGNCUHp5
}

# Logging
log {
        format json
        output file /home/ubuntu/access.log
}
```

# Running

```
sudo -E ./caddy run
```

# Parsing logs

```
# continually get remote IP addresses accessing the server
tail -F access.log | jq -r '.request.remote_ip'

# get remote IP addresses excluding 4xx response status 
cat access.log | jq -r '"\(.request.remote_ip)\t\(.status)"' | grep -vE ' 4..$' | cut -d':' -f1 | sort | uniq
```

Or use recent version of [goaccess](https://goaccess.io/).

# More

* https://caddyserver.com/docs/
