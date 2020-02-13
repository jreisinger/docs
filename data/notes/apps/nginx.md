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
