Check SSL/TLS protocol (you should be using TLS) and version

```
openssl s_client -connect reisinge.net:443 2> /dev/null | grep Protocol
```

Check generic TLS certificate info

```
openssl s_client -connect reisinge.net:443 -servername reisinge.net | openssl x509 -noout -text
```

* most valuable information is in the `X509v3 extensions` section
* you can get certificate also from a file, e.g. `cat file.crt | openssl x509 -noout -text`

See also 

* ULSAH 5th, ch 27.6
* https://github.com/pete911/certinfo
* https://go.dev/blog/tls-cipher-suites
