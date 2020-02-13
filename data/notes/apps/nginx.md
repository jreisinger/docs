Nginx is a multitool: web server, load balancer, reverse proxy, WAF

Verify installation (works also inside a container):

```
nginx -v
ps -ef | grep nginx
curl localhost -I
```

Key files and directories:

* `/etc/nginx/nginx.conf` - sets up global setting for things like worker processes, tuning, logging, loading of dynamic modules + references config files in `/etc/nginx/conf.d`
* `/var/log/nginx/access.log` - entry for each request NGINX serves
* `/var/log/nginx/error.log` - errors and debug info (if debug module is enabled)

Resources

* [NGINX Cookbook](https://learning.oreilly.com/library/view/nginx-cookbook/9781492049098/) (2019)
