# Sharing files

Prerequisites

* A registered public domain name (e.g. myserver.example.com)
* External access to ports 80 and 443
* caddy [binary](https://github.com/caddyserver/caddy/releases) (see Assets)

Caddyfile

```
myserver.example.com
file_server browse
root * /data/public
basicauth /* {
        # password hash created with `caddy hash-password`
        Bob JDJhJDE0JENVblJPOVljYml4a3phTHNVelpmYk90MVVEdXV2aVFjYkl2ODJENDFEaG1KU29TRGNCUHp5
}
log {
        output file /home/ubuntu/access.log
}
```

```
sudo ./caddy run
```

# More

* https://caddyserver.com/docs/
