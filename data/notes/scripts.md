# Scripts

## Check TLS certificate validity start and end dates

```
export PORT=443

for FQDN in reisinge.net quote.reisinge.net quotes.reisinge.net wiki.reisinge.net www.reisinge.net; do
    echo "--> $FQDN"
    echo | \
    openssl s_client -servername $FQDN -connect $FQDN:$PORT 2>/dev/null | \
    openssl x509 -noout -dates
done
```
