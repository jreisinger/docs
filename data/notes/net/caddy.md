# Sharing files

Prerequisites

* A registered public domain name (e.g. myserver.example.com)
* External access to ports 80 and 443
* [caddy](https://github.com/caddyserver/caddy/releases) binary (e.g. `wget -O caddy_2.4.6_linux_amd64.tar.gz https://github.com/caddyserver/caddy/releases/download/v2.4.6/caddy_2.4.6_linux_amd64.tar.gz`)
* Password hash created with `./caddy hash-password`

Caddyfile

```
myserver.example.com
file_server browse
root * /data/public
basicauth /* {
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
