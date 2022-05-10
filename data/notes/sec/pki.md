2011-09-16 (migrated from wiki.reisinge.net/PKI)

Public Key Infrastructure (PKI)

* binds public keys to entities
* enables other entities to verify public key bindings
* provides the services needed for ongoing management of keys in distributed environment

The value of CA:

* its brand
* ability to protect its private key

## PKI compontents

Certificates - digitally signed collection of information (2-4 KB)

* about user, computer, network device that holds the corresponding private key
* about the issuing CA
  * encryption and/or signing algorithms
  * list of X.509 v3 extensions
  * info for determining revocation status and validity of certificate

Certificate authority (CA)

* collection of HW, SW and people operating it
* known by two attributes: its name, its public key
* performs four basic PKI functions (some may be delegated to other components of PKI):
  * issues certificates (i.e., creates and signs them)
  * maintains certificate status information and issues CRLs
  * publishes its current (e.g., unexpired) certificates and CRLs
  * maintains archives of status info about the expired certificates that it issued
* may issue certificates to users, to other CAs, or both

Registration authority (RA)
 
* trusted by CA to register or vouch [ručiť] for the identity of users

Repository

* DB of active digital certificates for CA system

Certificate revocation list (CRL)

* certificate may be revoked [zrušiť, odobrať] because:
  * owner's private key has been lost
  * owner leaves the company/agency
  * owner's name changes

## Source

* [Introduction to Public Key Technology and the Federal PKI Infrastructure](http://csrc.nist.gov/publications/nistpubs/800-32/sp800-32.pdf)
* [An introduction to PKI (Video)](http://www.youtube.com/watch?v=EizeExsarH8)
