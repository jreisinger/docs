Nginx is a multitool: web server, load balancer, reverse proxy, WAF

# Key files and directories

* `/etc/nginx/nginx.conf` - sets up global setting for things like worker processes, performance tuning, logging, loading of dynamic modules + it `include`s config files in `/etc/nginx/conf.d`
* `/var/log/nginx/access.log` - entry for each request NGINX serves
* `/var/log/nginx/error.log` - errors and debug info (if debug module is enabled)

[Config file structure](https://nginx.org/en/docs/beginners_guide.html#conf_structure)

* simple directive vs block directive
* context

[Full example config](https://www.nginx.com/resources/wiki/start/topics/examples/full/)

* php/fastcgi
* reverse proxy
* load balancer

See [NGINXConfig](https://www.digitalocean.com/community/tools/nginx) for generating config files.

# Modules

NGINX consists of modules which are controlled by directives in the config file.

In version 1.9.11 (2016) Nginx added support for dynamic modules similar to Dynamic Shared Objects (DSO) of Apache HTTP server. Before that you needed to recompile Nginx if you updated a module. Now you can load/unload a module into NGINX at runtime.

To load a dynamic module:

```
load_module "modules/ngx_mail_module.so";
```

More:

* [Extending NGINX](https://www.nginx.com/resources/wiki/extending/)
* [3rd party modules](https://www.nginx.com/resources/wiki/modules/)

## Directives

The `if` directive

* https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#if
* https://www.nginx.com/resources/wiki/start/topics/tutorials/config_pitfalls/#using-if
* https://www.nginx.com/resources/wiki/start/topics/depth/ifisevil/
* https://agentzh.blogspot.com/2011/03/how-nginx-location-if-works.html

[ngx_http_proxy_module](https://nginx.org/en/docs/http/ngx_http_proxy_module.html)

* `proxy_redirect` - changes text in "Location" and "Refresh" headers of a proxied response

## Variables

* [all NGINX variables](http://nginx.org/en/docs/varindex.html) creation (declaration) happens while loading the configuration file (at "configuration time")
* on the other hand variables assignment happens when requests are actually being served (at "request time")

[ngx_http_core_module](http://nginx.org/en/docs/http/ngx_http_core_module.html)

* `$host` - in this order of precedence: hostname from the request line, or hostname from the `Host` request header field, or the server name matching a request
* `$scheme` - request scheme, "http" or "https"

[ngx_http_proxy_module](https://nginx.org/en/docs/http/ngx_http_proxy_module.html)

* `$proxy_host` - name and port (yes, also port if it's defined in `proxy_pass`) of a proxied server as specified in the `proxy_pass` directive

[ngx_http_map_module](https://www.digitalocean.com/community/tutorials/how-to-use-nginx-s-map-module-on-ubuntu-20-04)

* creates a new variable whose value depends on values of one or more source variables specified in the first parameter

```
http {
    # ...
    map $http_upgrade $connection_upgrade {
        default upgrade;
        ''      close;
    }
    server {
    # ...
        location / {
        # ...
            proxy_set_header Upgrade $http_upgrade;
            proxy_set_header Connection $connection_upgrade;
        }
    }
}
```

* [$http_upgrade](https://stackoverflow.com/questions/57898995/how-dose-nginx-get-the-value-of-http-upgrade) - source variable. It gets value from the request's `upgrade` header.
* $connection_upgrade - new variable whose value depends on $http_upgrade. If $http_upgrade exists and is an empty string it gets set to `close`. In all other cases it will be set to `upgrade`. 

# Reverse proxy

## Request headers

In proxied requests (`proxy_pass`) NGINX by default eliminates the header fields whose values are empty strings and redefines two header fields:

* `Host` gets set to `$proxy_host`
* `Connection` gets set to `Close`

More

* https://docs.nginx.com/nginx/admin-guide/web-server/reverse-proxy/#passing-request-headers

## Response headers

You might want to use the [proxy_redirect](http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_redirect) directive to change response "Location" header (used in 3xx redirects). E.g.:

```
# change host (fqdn) in response "Location" header
proxy_redirect ~*https?://[^/]+/(.+)$ https://$host/$1;
```

## Resolving

domain names statically configured in config file are only looked up once on startup (or configuration reload)

* https://forum.nginx.org/read.php?2,215830,215832#msg-215832
* https://www.nginx.com/blog/dns-service-discovery-nginx-plus/

# Traffic limiting

See also https://www.nginx.com/blog/mitigating-ddos-attacks-with-nginx-and-nginx-plus.

## Limiting connections

[limit_conn](http://nginx.org/en/docs/http/ngx_http_limit_conn_module.html) module limits the number of connections per the defined key, in particular, the number of connections from a single IP address

```
http {
    limit_conn_zone $binary_remote_addr zone=addr:10m;

    ...

    server {

        ...

        location /download/ {
            limit_conn addr 1;
        }
```

## Limiting rate

[limit_req](http://nginx.org/en/docs/http/ngx_http_limit_req_module.html) module is used to limit the request processing rate per a defined key, in particular, the processing rate of requests coming from a single IP address. The limitation is done using the “leaky bucket” method.

```
http {
    limit_req_zone $binary_remote_addr zone=one:10m rate=10r/s;

    ...

    server {

        ...

        location /search/ {
            limit_req zone=one burst=20 nodelay;
        }
```

* `10m` - in 10MB Nginx can hold cca 160,000 IP addresses
* `10r/s` really means 1 request in 100ms (Nginx tracks requests at millisecond granularity)
* `burst=20` allows a client to make 20 requests in excess of the specified rate
* `nodelay` forwards queued requests immediately instead of spacing them (1 request every 100ms in our case)

More

* https://www.nginx.com/blog/rate-limiting-nginx/
* https://www.freecodecamp.org/news/nginx-rate-limiting-in-a-nutshell-128fe9e0126c/

## Limiting bandwidth

* [limit_rate](http://nginx.org/en/docs/http/ngx_http_core_module.html#limit_rate) directive (from the core module) limits the rate of response transmission to a client
* enables NGINX to share its upload bandwidth across all of the clients in a manner you specify
* specified in bytes per second
* the limit is set per a request, so you may want to institute a connection limit as well

```
location /download/ {
    limit_rate_after 10m; # After 10 megabytes are dowloaded
    limit_rate 1m;        # limit rate to 1 megabyte per second.
}
```

* [proxy_limit_rate](https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_limit_rate) - limits the speed of response reading from the proxied server, response buffering must be enabled
* https://docs.nginx.com/nginx/admin-guide/security-controls/controlling-access-proxied-http/#limiting-the-bandwidth

# Tips and tricks

Useful commands:

```
nginx -t # test configuration
nginx -T # test configurtation and print it
nginx -s <signal> # stop (immediately)
                  # quit (wait to finish processing requests)
                  # reload (configuration gracefully, i.e. without dropping packets!)
                  # reopen (log files) 
```

Verify installation (works also inside a container):

```
nginx -v
ps -ef | grep nginx
curl localhost -I
```

Get values of variables (debugging):

```
# /etc/nginx/nginx.conf
add_header X-mine "$upstream_addr";

curl localhost -v # look for X-mine header
```

# Resources

* [NGINX Cookbook](https://learning.oreilly.com/library/view/nginx-cookbook/9781492049098/) (2019)
* [agentzh's Nginx Tutorials](https://openresty.org/download/agentzh-nginx-tutorials-en.html) (2019)
