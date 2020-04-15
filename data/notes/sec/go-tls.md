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

Keys and certificates

* public key - can be freely distributed, used to encrypt a message or verify a signature
* private key - must be kept secret, used to decrypt or sign a message

to share a keys you might need a trusted 3rd pary -> certificate authority (CA)
