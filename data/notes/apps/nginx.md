Nginx is a multitool: web server, load balancer, reverse proxy, WAF

## Basics

Key files and directories:

* `/etc/nginx/nginx.conf` - sets up global setting for things like worker processes, tuning, logging, loading of dynamic modules + references config files in `/etc/nginx/conf.d`
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

[Config file structure](https://nginx.org/en/docs/beginners_guide.html#conf_structure) (simple vs block directive, context)

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

## The if directive

More:

* https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#if
* https://www.nginx.com/resources/wiki/start/topics/tutorials/config_pitfalls/#using-if

## Reverse proxy

In proxied requests (`proxy_pass`) NGINX by default eliminates the header fields whose values are empty strings and redefines two header fields:

* `Host` gets set to `$proxy_host`
* `Connection` gets set to `Close`

More

* https://docs.nginx.com/nginx/admin-guide/web-server/reverse-proxy/#passing-request-headers

## Variables

* all NGINX variables creation (declaration) happens while loading the configuration file (at "configuration time")
* on the other hand variables assignment happens when requests are actually being served (at "request time")

All NGINX [variables](http://nginx.org/en/docs/varindex.html).

* `$proxy_host` [ngx_http_proxy_module](https://nginx.org/en/docs/http/ngx_http_proxy_module.html) - name and port of a proxied server as specified in the `proxy_pass` directive
* `$host` [ngx_http_core_module](http://nginx.org/en/docs/http/ngx_http_core_module.html) - in this order of precedence: hostname from the request line, or hostname from the `Host` request header field, or the server name matching a request

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

## Resources

* [NGINX Cookbook](https://learning.oreilly.com/library/view/nginx-cookbook/9781492049098/) (2019)
* [agentzh's Nginx Tutorials](https://openresty.org/download/agentzh-nginx-tutorials-en.html) (2019)
