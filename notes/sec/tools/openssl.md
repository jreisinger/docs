Check TLS/SSL **version** (you should be using TLS)

```
openssl s_client -connect reisinge.net:443 2> /dev/null | egrep -i '(tls|ssl)'
```

Check **generic** TLS certificate info

```
openssl x509 -in jane.crt -text -noout
```

```
openssl s_client -connect reisinge.net:443  -servername reisinge.net | openssl x509 -noout -text
```

* most valuable information is in the `X509v3 extensions` section

Check TLS certificate **validity** start and end dates for multiple domains

```
for FQDN in reisinge.net quote.reisinge.net quotes.reisinge.net wiki.reisinge.net www.reisinge.net; do
    echo "--- $FQDN ---"
    echo | \
    openssl s_client -connect $FQDN:443 -servername $FQDN 2>/dev/null | \
    openssl x509 -noout -dates | perl -wlpe 's/=/\t/'
done
```

See also 

* ULSAH 5th, ch 27.6
* https://github.com/pete911/certinfo
* https://go.dev/blog/tls-cipher-suites
