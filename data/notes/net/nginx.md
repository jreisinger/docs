Nginx is a multitool: web server, load balancer, reverse proxy, WAF

# General

## Key files and directories

* `/etc/nginx/nginx.conf` - sets up global setting for things like worker processes, tuning, logging, loading of dynamic modules + it references config files in `/etc/nginx/conf.d`
* `/var/log/nginx/access.log` - entry for each request NGINX serves
* `/var/log/nginx/error.log` - errors and debug info (if debug module is enabled)

Useful commands:

```
nginx -t # test configuration
nginx -T # test configurtation and print it
nginx -s <signal> # stop (immediately)
                  # quit (wait to finish processing requests)
                  # reload (configuration gracefully, i.e. without dropping packets!)
                  # reopen (log files) 
```

[Config file structure](https://nginx.org/en/docs/beginners_guide.html#conf_structure) (simple directive vs block directive, context)

[Sample config](https://www.nginx.com/resources/wiki/start/topics/examples/full/)

## Modules

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

* `$proxy_host` - name ~~and port~~ (error in upstream docs?) of a proxied server as specified in the `proxy_pass` directive

## Tips and tricks

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

# Traffic limiting

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
    limit_req_zone $binary_remote_addr zone=one:10m rate=1r/s;

    ...

    server {

        ...

        location /search/ {
            limit_req zone=one burst=5;
        }
```

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

# Resources

* [NGINX Cookbook](https://learning.oreilly.com/library/view/nginx-cookbook/9781492049098/) (2019)
* [agentzh's Nginx Tutorials](https://openresty.org/download/agentzh-nginx-tutorials-en.html) (2019)
