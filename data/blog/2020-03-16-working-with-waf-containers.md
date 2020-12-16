# Working with WAF containers

I've been working with Web Application Firewalls (WAFs) in the form of application containers for some time. To make my work easier I created couple of tools: [waf-runner](https://github.com/jreisinger/waf-runner) and [waf-tester](https://github.com/jreisinger/waf-tester). In this post I'll try to show you how they can be used.

My most common use cases boil down to:

1. Changing or upgrading WAF components (e.g. Nginx, ModSecurity, CRS)
2. Adding or modifying WAF rules (strictly speaking this is a subset of 1.)

For both use cases the workflow looks like this:

1. Build and run a WAF container
2. Make changes to the WAF container
3. Test the WAF container
4. Commit and push the changes

## Build and run a WAF container

[waf-runner](https://github.com/jreisinger/waf-runner) will build and run a WAF container based on supplied `Dockerfile` and related configuration files. For example:

```
$ waf-runner waf/nginx/modsecurity
--> Create directories for WAF logs in /tmp/var/log
--> Create temporary directory
/var/folders/8d/49xspl216vqf52y6b_s_y9x00000gn/T/tmp.i3XRl9PA
--> Create /var/folders/8d/49xspl216vqf52y6b_s_y9x00000gn/T/tmp.i3XRl9PA/docker-compose.yaml
--> Copy recursively all files from waf/nginx/modsecurity to /var/folders/8d/49xspl216vqf52y6b_s_y9x00000gn/T/tmp.i3XRl9PA
--> Run docker images for WAF and web server
Creating network "tmpi3xrl9pa_default" with the default driver
Building testing-waf
Creating testing-webserver ... done
Creating testing-waf       ... done
--> Check WAF is up and proxying requests
CONTAINER ID        IMAGE                     COMMAND                  CREATED             STATUS              PORTS                  NAMES
c815458a6eef        tmpi3xrl9pa_testing-waf   "nginx -g 'daemon of…"   4 seconds ago       Up 3 seconds        0.0.0.0:80->80/tcp     testing-waf
bcb445f3a683        nginx                     "/docker-entrypoint.…"   4 seconds ago       Up 3 seconds        0.0.0.0:8080->80/tcp   testing-webserver
--> WAF container is up and running (hit Ctrl-C to quit)

==> /tmp/var/log/modsec_audit.log <==

==> /tmp/var/log/nginx/access.log <==
192.168.240.1 - - [29/Aug/2020:15:36:31 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.64.1" "-"

==> /tmp/var/log/nginx/error.log <==
```

`waf-runner` will keep on `tail`ing the logs until you hit Ctrl-C.

## Make changes to the WAF container

If you want to make temporary changes to the WAF you can get into the container:

```
$ docker exec -it testing-waf /bin/sh
/ # vi /etc/nginx/nginx.conf
/ # nginx -s reload
/ # exit
```

To make permanent changes you modify the WAF's `Dockerfile` and/or related configuration files.

## Test the WAF container

[waf-tester](https://github.com/jreisinger/waf-tester) will run tests against a WAF (that is running on localhost in this case):

```
$ waf-tester -tests waf_tests/generic/basic-tests.yaml -print OK
OK	RCE                  GET       http://localhost/?exec=/bin/bash
OK	SQLi                 GET       http://localhost/?id=1'%20or%20'1'%20=%20'
OK	OS file access       GET       http://localhost/?page=/etc/passwd
OK	Path traversal       GET       http://localhost/get-files?file=/../../../../etc/shadow
OK	XSS                  GET       http://localhost/?<script>
OK	Session fixation     GET       http://localhost/foo.php?bar=blah%3Cscript%3Edocument.cookie=%22sessionid=1234;%20domain=.example.dom%22;%3C/script%3E
```

So what are these tests. They are basically HTTP requests and expected responses represented in YAML format. For example:

```
$ waf-tester -template
tests:
- test_title: SQLi
  stages:
  - stage:
      input:
        headers:
          User-Agent: waf-tester
        method: GET
        uri: ?id=1'%20or%20'1'%20=%20'
        data: []
      output:
        status:
        - 403
<...snip...>
```

This means that `waf-tester` will make a GET request with this URL `<scheme>://<host>/?id=1'%20or%20'1'%20=%20'` (`scheme` and `host` are command line flags) and it will expect status 403 in the response. If this expectation proves true the test passed (`OK`). Otherwise it failed (`FAIL`).

There exist additional `input` and `output` fields. You can use different methods, insert custom headers or instead of checking the request status code you can check the WAF logs.

The YAML format is based on [FTW](https://github.com/CRS-support/ftw/blob/master/docs/YAMLFormat.md) but some fields are missing. Others, like `dest_addr`, are ignored. See the [code](https://github.com/jreisinger/waf-tester/blob/master/yaml/types.go) for details.

## Practical example: adding a WAF rule

WAF rules (or signatures) is what helps WAFs to distinguish between legitimate and malicious requests. Let's say that I want to run some tests against a [NAXSI](https://github.com/nbs-system/naxsi) WAF with default rules. To run this WAF locally in a container I execute this:

```
$ waf-runner waf/nginx/naxsi
<...snip...>
```

`waf/nginx/naxsi` folder contains the `Dockerfile` and all needed files like `nginx.conf`, `naxsi.conf` and `naxsi_core.rules` that get copied into the container.

Let's adapt one of the [FTW](https://github.com/coreruleset/ftw) tests

```
$ cat 913120-2.yaml 
tests:
- test_title: 913120-2
  desc: IBM fingerprint from (http://www-01.ibm.com/support/docview.wss?uid=swg21293132)
  stages:
  - stage:
      input:
        uri: /AppScan_fingerprint/MAC_ADDRESS_01234567890.html?9ABCDG1
      output:
        status: [403]
```
and run the test against the WAF:

```
$ waf-tester -tests 913120-2.yaml -verbose
FAIL	913120-2                       http://localhost/AppScan_fingerprint/MAC_ADDRESS_01234567890.html?9ABCDG1
  DESC       IBM fingerprint from (http://www-01.ibm.com/support/docview.wss?uid=swg21293132)
  FILE       913120-2.yaml
  STATUS     404 Not Found
  CODE       404
  EXP_CODES  [403]
  EXP_LOGS   
  EXP_NOLOGS 
  EXP_ERR    false
  ERROR      <nil>
  DATA       []
  HEADERS    
    waf-tester-id: 64502ec0-7468-b872-4793-b63bdb8bc9e9
  LOGS       
```

You can see a `FAIL` status. This means that the WAF didn't block this request. As you can see the test expected 403 code (`EXP_CODES`) but got 404 (`CODE`). That's why it failed.

Let's try to add (a rather naïve) custom WAF rule to fix the failing test:

```
# nginx/naxsi/naxsi_custom.rules
BasicRule "str:9ABCDG1" "msg:FTW 913120-2" "mz:ARGS" "s:$UWA:4" id:10002;
```

Now rebuild the WAF container: hit `Ctrl-C` in the first terminal and run `waf-runner waf/nginx/naxsi` again. When we re-run the test we can see it's `OK` now.
