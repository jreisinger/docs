[Source](https://github.com/lizrice/secure-connections).

TLS

* new name for SSL
* **identity** of one or both communcation parties
* **encryption** of the communication
* HTTPS = HTTP over TLS

HTTP(S) runs over TCP

1. Create TCP connection
2. TLS - encrypt TCP connection (skip if HTTP)
3. HTTP connection

(see [details](https://speakerdeck.com/lizrice/a-go-programmers-guide-to-secure-connections?slide=10))

Keys

* public key - used to **encrypt** a message or verify a signature
* private (secret) key - used to decrypt or sign a message

X.509 Certificates

* basically a public key signed with private key of a trusted third party called certificate authority (CA)
* can be verified using the public key of the CA
* it proves the **identity** of the public key holder
* it contains these fields:
  * subject name
  * subject's public key
  * issuer (CA) name
  * validity
* certs should use Subject Alternative Names (SANs), Common name (CN) was deprecated in 2000
* binary format (ASN.1) - `.der`
* Base64 encoding to represent it as text - `.pem`:

```
-----BEGIN CERTIFICATE-----
MIIC7jCCAlegAwIBAgIBATANBgkqhkiG9w0BAQQFADCBqTELMAkGA1UEBhMCWFkx
FTATBgNVBAgTDFNuYWtlIERlc2VydDETMBEGA1UEBxMKU25ha2UgVG93bjEXMBUG
A1UEChMOU25ha2UgT2lsLCBMdGQxHjAcBgNVBAsTFUNlcnRpZmljYXRlIEF1dGhv
cml0eTEVMBMGA1UEAxMMU25ha2UgT2lsIENBMR4wHAYJKoZIhvcNAQkBFg9jYUBz
bmFrZW9pbC5kb20wHhcNOTgxMDIxMDg1ODM2WhcNOTkxMDIxMDg1ODM2WjCBpzEL
MAkGA1UEBhMCWFkxFTATBgNVBAgTDFNuYWtlIERlc2VydDETMBEGA1UEBxMKU25h
a2UgVG93bjEXMBUGA1UEChMOU25ha2UgT2lsLCBMdGQxFzAVBgNVBAsTDldlYnNl
cnZlciBUZWFtMRkwFwYDVQQDExB3d3cuc25ha2VvaWwuZG9tMR8wHQYJKoZIhvcN
AQkBFhB3d3dAc25ha2VvaWwuZG9tMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKB
gQDH9Ge/s2zcH+da+rPTx/DPRp3xGjHZ4GG6pCmvADIEtBtKBFAcZ64n+Dy7Np8b
vKR+yy5DGQiijsH1D/j8HlGE+q4TZ8OFk7BNBFazHxFbYI4OKMiCxdKzdif1yfaa
lWoANFlAzlSdbxeGVHoT0K+gT5w3UxwZKv2DLbCTzLZyPwIDAQABoyYwJDAPBgNV
HRMECDAGAQH/AgEAMBEGCWCGSAGG+EIBAQQEAwIAQDANBgkqhkiG9w0BAQQFAAOB
gQAZUIHAL4D09oE6Lv2k56Gp38OBDuILvwLg1v1KL8mQR+KFjghCrtpqaztZqcDt
2q2QoyulCgSzHbEGmi0EsdkPfg6mp0penssIFePYNI+/8u9HT4LuKMJX15hxBam7
dUHzICxBVC1lnHyYGjDuAMhe396lYAn8bCld1/L4NMGBCQ==
-----END CERTIFICATE-----
```

* unfortunately some of these file extensions ^ are also used for other data such as private keys
* there is not consitency in file naming, you can see:
  * .key, .crt - information type (private key, certificate)
  * .pem - file format (PEM)

CLI tools

* openssl - does not easily support SANs
* cfssl
* mkcert
* minica

Error messages
 
* `connection refused` - wrong port (or lack of server resources)
* `certificate signed by uknown authority` - reveived a cert but it's not trusted
* `remote error` - it's the other end that's complaining
