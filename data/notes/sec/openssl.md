Check **generic** TLS certificate info

```
openssl s_client -servername reisinge.net -connect reisinge.net:443 | openssl x509 -noout -text
```

* most valuable information is in the `X509v3 extensions` section

Check TLS/SSL **version** (you should be using TLS)

```
openssl s_client -connect reisinge.net:443 2> /dev/null | egrep -i '(tls|ssl)'
```

Check TLS certificate **validity** start and end dates for multiple domains

```
export PORT=443

for FQDN in reisinge.net quote.reisinge.net quotes.reisinge.net wiki.reisinge.net www.reisinge.net; do
    echo "--> $FQDN"
    echo | \
    openssl s_client -servername $FQDN -connect $FQDN:$PORT 2>/dev/null | \
    openssl x509 -noout -dates
done
```

See also https://github.com/pete911/certinfo.
