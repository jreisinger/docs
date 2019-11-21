# Check TLS certificate validity start and end dates

```
export FQDN=reisinge.net
export PORT=443

echo | \
openssl s_client -servername $FQDN -connect $FQDN:$PORT 2>/dev/null | \
openssl x509 -noout -dates
```
