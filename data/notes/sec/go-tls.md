[Source](https://github.com/lizrice/secure-connections).

TLS

* new name for SSL
* identity of one or both communcation parties
* encryption of the communication
* HTTPS = HTTP over TLS

HTTP(S) runs over TCP

1. Create TCP connection
2. TLS - encrypt TCP connection (skip if HTTP)
3. HTTP connection

(see [details](https://speakerdeck.com/lizrice/a-go-programmers-guide-to-secure-connections?slide=10))

Error messages

* `connection refused` = wrong port (or lack of server resources)

Keys

* public key - used to encrypt a message or verify a signature
* private (secret) key - used to decrypt or sign a message

Certificate

* basically a public key signed with private key of a trusted third party called certificate authority (CA)
* can be verified using the public key of the CA
* it proves the identity of the public key holder
* it contains these fields:
  * ff
