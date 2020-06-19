---
title: "Working with WAF containers"
date: 2020-03-16
categories: [prog,sec,sysadmin]
tags: [go,shell,docker,waf]
---

I've been working with Web Application Firewalls (WAFs) in the form of application containers for some time. To make my work easier I created couple of tools: [waf-runner](https://github.com/jreisinger/waf-runner) and [waf-tester](https://github.com/jreisinger/waf-tester). In this post I'll try to show you how they can be used.

My most common use cases boil down to:

1. Changing configuration of WAF components (upgrades, configuration changes).
2. Adding or modifying WAF rules. (Strictly speaking this is a subset of 1.)

For both use cases the workflow looks like this:

1. Build and run a WAF container.
2. Make temporary changes inside the WAF container.
3. Test the changes (sort of unit testing).
4. Make changes permanent by editing `Dockerfile` or WAF configuration files.
5. Test the changes again.
6. Commit and push the changes.
7. Run integration tests (CI/CD tests)

`waf-runner` can help with 1. and `waf-tester` with 3. and 5. (possibly with 7.).

## Running a WAF

`waf-runner` will build and run a WAF container based on supplied `Dockerfile` and related configuration files. For example:

```
$ waf-runner -s waf/nginx/modsecurity
--> Create directories for WAF logs in /tmp/var/log
--> Create temporary directory
/tmp/tmp.LceQagUPU5
--> Create /tmp/tmp.LceQagUPU5/docker-compose.yaml
--> Copy all needed files from waf/nginx/modsecurity to /tmp/tmp.LceQagUPU5
'waf/nginx/modsecurity/Dockerfile' -> '/tmp/tmp.LceQagUPU5/Dockerfile'
'waf/nginx/modsecurity/geoip2.sh' -> '/tmp/tmp.LceQagUPU5/geoip2.sh'
'waf/nginx/modsecurity/nginx.conf' -> '/tmp/tmp.LceQagUPU5/nginx.conf'
--> Run docker images for WAF and web server
Creating network "tmplceqagupu5_default" with the default driver
Building testing-waf
Creating testing-webserver ... done
Creating testing-waf       ... done
--> Check WAF is up and proxying requests
CONTAINER ID        IMAGE                       COMMAND                  CREATED             STATUS              PORTS                  NAMES
f97cf198b64e        tmplceqagupu5_testing-waf   "nginx -g 'daemon of…"   4 seconds ago       Up 3 seconds        0.0.0.0:80->80/tcp     testing-waf
cdc49e9ae687        nginx                       "nginx -g 'daemon of…"   5 seconds ago       Up 4 seconds        0.0.0.0:8080->80/tcp   testing-webserver
--> WAF container is up and running
--> Sleeping until Ctrl-C
```

If you want to make temporary changes to the WAF you can get into the container:

```
$ docker exec -it testing-waf /bin/sh
/ # vi /etc/nginx/nginx.conf
/ # nginx -s reload
/ # exit
```

To make permanent changes you modify the WAF's `Dockerfile` and/or related configuration files.

## Testing a WAF

`waf-tester` will run tests against a WAF (that is running on localhost in this case):

```
$ waf-tester -host localhost -tests waf_tests/generic/basic-tests.yaml -scheme http
OK	RCE                  GET       http://localhost/?exec=/bin/bash
OK	SQLi                 GET       http://localhost/?id=1'%20or%20'1'%20=%20'
OK	LFI                  GET       http://localhost/?page=/etc/passwd
OK	XSS                  GET       http://localhost/?<script>
OK	Scanner detection    GET       http://localhost/AppScan_fingerprint/MAC_ADDRESS_01234567890.html?9ABCDG1
OK	Session fixation     GET       http://localhost/foo.php?bar=blah%3Cscript%3Edocument.cookie=%22sessionid=1234;%20domain=.example.dom%22;%3C/script%3E
```

So what are these tests. They are basically HTTP requests and expected responses represented in YAML format. For example:

```
  - test_title: RCE
    stages:
      - stage:
          input:
            method: "GET"
            uri: "?exec=/bin/bash"
          output:
            status: 403
```

This means that `waf-tester` will make a GET request with this URL `<scheme>://<host>/?exec=/bin/bash` (`scheme` and `host` are command line flags) and it will expect status 403 in the response. If this expectation proves true the test passed (`OK`). Otherwise it failed (`FAIL`).

There exist additional `input` and `output` fields. You can use different methods, insert custom headers or instead of checking the request code you can check the WAF logs:

```
  - test_title: 911100-5
    stages:
      - stage:
          input:
            method: "TEST"
            port: 80
            headers:
              User-Agent: "ModSecurity CRS 3 Tests"
              Host: "localhost"
          output:
            log_contains: "id \"911100\""
```

The YAML format is based on [FTW](https://github.com/CRS-support/ftw/blob/master/docs/YAMLFormat.md) but some fields are missing. Others, like `dest_addr`, are ignored. See the [code](https://github.com/jreisinger/waf-tester/blob/master/yaml/types.go) for details.

## Practical example: adding a WAF rule

WAF rules (or signatures) is what helps WAFs to distinguish between legitimate and malicious requests. Let's say that I want to run some tests against a [NAXSI](https://github.com/nbs-system/naxsi) WAF with default rules. To run this WAF locally in a container I execute this:

```
$ waf-runner -s waf/nginx/naxsi
<...snip...>
--> Sleeping until Ctrl-C
```

`waf/nginx/naxsi` folder contains the `Dockerfile` and all needed files like `nginx.conf`, `naxsi.conf` and `naxsi_core.rules` that get copied into the container.

Then I open another terminal and run basic tests against the WAF:

```
$ waf-tester -host localhost -scheme http -tests waf_tests/generic/basic-tests.yaml
FAIL	RCE                  GET       http://localhost/?exec=/bin/bash
OK	    SQLi                 GET       http://localhost/?id=1'%20or%20'1'%20=%20'
OK	    LFI                  GET       http://localhost/?page=/etc/passwd
OK	    XSS                  GET       http://localhost/?<script>
FAIL	Scanner detection    GET       http://localhost/AppScan_fingerprint/MAC_ADDRESS_01234567890.html?9ABCDG1
OK	    Session fixation     GET       http://localhost/foo.php?bar=blah%3Cscript%3Edocument.cookie=%22sessionid=1234;%20domain=.example.dom%22;%3C/script%3E
```

You can see `FAIL` status for two tests. This means that the WAF didn't block these two requests. Let's have a more detailed view on the first failing test:

```
$ waf-tester -host localhost -scheme http -tests waf_tests/generic/basic-tests.yaml -title RCE -verbose
FAIL	RCE                  GET       http://localhost/?exec=/bin/bash
  DESC       
  FILE       waf_tests/generic/basic-tests.yaml
  STATUS     200 OK
  CODE       200
  EXP_CODES  [403]
  EXP_LOGS   
  EXP_NOLOGS 
  EXP_ERR    false
  ERROR      <nil>
  DATA       []
  HEADERS    
    waf-tester-id: cb0338a4-d88e-bba1-7177-cb9f026ae7f8
  LOGS   
```

As you can see the test expected 403 code (`EXP_CODES`) but got 200 (`CODE`). That's why it failed.

Let's try to add (a rather naïve) custom WAF rule to fix the failing test:

```
# nginx/naxsi/naxsi_custom.rules
BasicRule "str:/bin/bash" "msg:RCE" "mz:ARGS" "s:$UWA:4" id:10001;
```

Now I rebuild the WAF container: I hit `Ctrl-C` in the first terminal and run `waf-runner -s waf/nginx/naxsi` again. When I re-run the test you can see it's `OK` now:

```
$ waf-tester -host localhost -scheme http -tests waf_tests/generic/basic-tests.yaml -title RCE
OK	RCE                  GET       http://localhost/?exec=/bin/bash
```
