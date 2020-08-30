Python CA certificates

```
export ALL_CA_CERTIFICATES="/usr/local/share/ca-certificates/cacert.pem"

# NOTE: we are appending so don't run these commands multiple times
sudo bash -c "cat $(python3 -m certifi) >> $ALL_CA_CERTIFICATES"
sudo bash -c "cat your-ca.crt >> $ALL_CA_CERTIFICATES"

# Put this into ~/.bashrc
export REQUESTS_CA_BUNDLE=$ALL_CA_CERTIFICATES
```

Screenshot (print screen)

1. Command (⌘) + Shift + 4
1. hold down Control and make your selection
1. Command (⌘) +  V
