TLS

* new name for SSL
* **identity** of one or both communication parties
* **encryption** and **integrity** of the communication
* kind of "secure TCP"

Phases

1. Create TCP connection
2. Encrypt HTTP connection with TLS (skip if not HTTPS)
3. HTTP connection

(see [details](https://speakerdeck.com/lizrice/a-go-programmers-guide-to-secure-connections?slide=10))

Asymmetric cryptography keys

* public - used to encrypt a message or verify a signature
* private (secret) - used to decrypt or sign a message

X.509 Certificates

* basically a public key signed with private key of a trusted third party called certificate authority (CA)
* can be verified using the public key of the CA
* it proves the identity of the public key holder
* it contains these fields:
  * subject name
  * subject public key
  * issuer (CA) name
  * validity
* certs should use Subject Alternative Names (SANs), Common name (CN) was deprecated in 2000
* binary format (ASN.1): .der
* text format (Base64 encoded): .pem
* there is not consistency in file naming, you can see:
  * .key, .crt - extension denotes information type (private key, certificate)
  * .pem - extension denotes file format (PEM)

CLI tools

* openssl - does not easily support SANs
* cfssl
* mkcert
* minica

Error messages
 
* `connection refused` - wrong port (or lack of server resources)
* `certificate signed by unknown authority` - received a cert but it's not trusted
* `remote error` - it's the other end that's complaining

See also:

* <https://github.com/jreisinger/docs/tree/master/blog/gosec>
* <https://github.com/jreisinger/pocs>

