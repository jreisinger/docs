## Intro

What

* stateless, connectionless, reliable protocol
* HTTP versions: 0.9 ('91), 1.0 ('96), **1.1** ('99), 2 ('15)
* HTTP < 1.1 needed a separate TCP connection for each HTTP request
* HTTP < 2 are text based, HTTP 2 is binary

What for

* used to fetch network resources (documents) based on their hostname and path
* user agents actually get a *representation* of the given resource (e.g. a copy of a static file)
* dominant document type - WWW of hypertext documents
* dominant version - HTTP/1.1

Python libraries

* `urllib` - client in Standard Library, for simple use
```python
from urllib.request import urlopen
import urllib.error
r = urlopen('http://httpbin.org/headers')
# you need to instruct urllib how to turn raw bytes into text
print(r.read().decode('ascii'))
```
* `requests` - the go-to tool
```python
import requests
r = requests.get('http://httpbin.org/headers')
print(r.text)
```

## HTTP message

HTTP message format (both request `>` and response `<`)

* line of information (method, resource, protocol, status code)
* zero or more of `name: value` [headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers)
* blank line - server/client calls `recv()` until `CR-LF-CR-LF` to find out the headers' end
* optional body - its length (framing) is defined by the `Content-Lenght` header

```plain
$ curl reisinge.net -v
* Rebuilt URL to: reisinge.net/
*   Trying 109.230.20.210...
* Connected to reisinge.net (109.230.20.210) port 80 (#0)
> GET / HTTP/1.1
> Host: reisinge.net
> User-Agent: curl/7.47.0
> Accept: */*
>
< HTTP/1.1 302 Moved Temporarily
< Server: nginx/1.12.2
< Date: Thu, 11 Jan 2018 07:51:16 GMT
< Content-Type: text/html
< Content-Length: 161
< Connection: keep-alive
< Location: http://jreisinger.github.io
<
<html>
<head><title>302 Found</title></head>
<body bgcolor="white">
<center><h1>302 Found</h1></center>
<hr><center>nginx/1.12.2</center>
</body>
</html>
* Connection #0 to host reisinge.net left intact
```

The client can't issue another request over the same socket until the response
is finished.

For requests, the body can include parameters (for POST or PUT requests) or the contents of a file to upload. For responses, the body is the payload of the resource being requested (e.g. HTML, image data, or query results). The message body is not necessarily human readable, since it can contain images or other binary data. The body can also empty, as for GET requests or most error messages.

## URL

* specifies how and where to access a resource
* not HTTP specific, e.g. mobile OSs use URLs to communicate between apps
* stick to URL, forget about URI and URN

```
scheme://[username:password@]hostname[:port][/path][?query][#anchor]
```

* `path` used to represent a file system path, now it can be anything
* `query` section can contain multiple `key=value` pairs separated by `&`
* `anchor` is also called `fragment`

```
http://example.com:81/a/b.html?user=Alice&year=2020#p2
```

## Methods (verbs)

* actions - what the server should do

GET

* "read", fetch a resource
* cannot include a body
* can only modify the document being returned (ex. `?q=python` or `?result=10`)

POST

* "write", update resource on a server with request data
* the result of POST can't be cached
* can't be retried automatically if the response does not arrive

## Encoding

HTTP transfer encoding <-> content encoding

Transfer encoding (Content-Length or chunked encoding, raw or compressed)

* wrapper used for data delivery, not a change in the underlying data itself

```plain
GET / HTTP/1.1
Accept-Encoding: gzip
...

HTTP/1.1 200 OK
Content-Length: 3913
Transfer-Encoding: gzip
...
```

Content type - what format will be selected to represent a given resource

* `application/octet-stream` - a plain sequence of bytes for which server can guarantee no more specific interpretation
* `text/html`

Content encoding - if the content type ^ is text, what encoding will be used to turn text code points into bytes

* `charset=utf-8`

```plain
Content-Type: text/html; charset=utf-8
```

## Authentication and cookies

Basic Auth (HTTP-mediated authentication)

* poor design initially, then fixed by SSL/TLS but still ugly
* still used by simple APIs
* Cookies are much more used for authentication today

TLS/SSL

* server authentication and transfer encryption layer around HTTP
* the only stuff that's NOT encrypted: IP addresses, ports (and probably DNS traffic)

[Cookies](https://tools.ietf.org/html/rfc6265)

* every request is independent of all other requests (from the point of the
    view of the protocol)
* -> authentication info must be carried in every request
* cookie = key-value pair sent by the server and then inserted in all
    further requests:

```plain
(initial request)
POST /login HTTP/1.1
...

(initial response)
HTTP/1.1 200 OK
Set-Cookie: session-id=d41d8cd98f00b204e9800998ecf8427e; Path=/
...

(all subsequent requests to the same host)
GET /login HTTP/1.1
Cookie: session-id=d41d8cd98f00b204e9800998ecf8427e
...

```
* cookie should be opaque - random UUID mapped to the username on the
    server or encrypted string that server alone can decrypt
* some servers give you cookie simply for visiting to track how you move
    through the site
* later when you log in, your browsing history can be copied into your 
    permanent account history :-/

## Keep-Alive

* the three-way TCP handshake can be avoided if a connection is already open
* HTTP/1.1 - default to keep HTTP connection open after a request to re-use a
    single TCP connection
* server/client can specify `Connection: close` if they plan on hanging up once
    a request is completed
* web browsers often create four or more simultaneous TCP connections per site
    to get the resources in parallel

## [Status codes](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status)

Returned by a server with each response.

1xx - Informational (Hold on)

2xx - Success (Here you go)

* 200 OK

3xx - Redirects (Go the other way)

* not expected to carry a body
* new location is in the `Location` header
* 301 Moved Permanently - resource has a new permanent URL

```python
>>> r = requests.get('http://httpbin.org/status/301', allow_redirects=False)
>>> (r.status_code, r.url, r.headers['Location'])
(301, 'http://httpbin.org/status/301', '/redirect/1')
```

* 302 Found - resource temporarily resides at different URL
* 304 Not Modified - the resource has not been modified so the client can use the cached version

4xx - Client errors (You messed up)

* client request is unintelligible or illegal
* 400 Bad Request - the server cannot or will not process the request due to a client error (e.g. malformed request syntax)
* 401 Unauthorized - it really means "unauthenticated"
* 403 Forbidden - the client has no access rights to the content; that is it's unauthorized (e.g. blocked by a WAF)
* 404 Not Found - the resource is not at the given URL

5xx - Server errors (I messed up)

* 500 Internal server error - a generic "catch-all" response
* 502 Bad Gateway - the server is a proxy and it cannot contact the upstream server (the server behind the proxy)
* 503 Service unavailable - server down for maintenance or overloaded
* 504 Gateway timeout - the server is a proxy and it did not get a response from the upstream server in order to complete the request

## Various

Minimally correct request nowadays (otherwise 404):

```plain
GET /html/rfc7230 HTTP/1.1
Host: tools.ietf.org
```

Caching headers

* allow client to cache and reuse resources locally
* let server skip redelivering an unchanged resource

## Sources

* Foundations of Python Network Programming (2014) - [ch. 9](https://www.safaribooksonline.com/library/view/foundations-of-python/9781430258551/9781430258544_Ch09.xhtml)
* Network Programming with Go - ch. 8
* ULSAH 5th - ch. 19
* https://web.stanford.edu/class/cs253/
