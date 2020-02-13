* a multitool: web server, load balancer, reverse proxy, WAF

Verify installation (works also inside a container):

```
nginx -v
ps -ef | grep nginx
curl localhost -I
```

Resources

* [NGINX Cookbook](https://learning.oreilly.com/library/view/nginx-cookbook/9781492049098/) (2019)
