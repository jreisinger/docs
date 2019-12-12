* Network File System
* lets you share files among computers
* nearly transparent to users (you need to mount it)
* no information lost when an NFS server crashes
* introduced in 1984 by Sun Microsystems

Transport protocols

* v3 - choice of UDP or TCP
* v4 - only TCP (typically port 2049)

State

* v3 - stateless (secret cookies)
* v4 - stateful (both client and server maintain info on open files an locks)

Security

* NFS is security mechanism independent, supporting mulitple "flavors" of authentication
 * AUTH_NONE - no authentication
 * AUTH_SYS (traditional, weak security) - UNIX-style user and group access control
 * RPCSEC_GSS (optional in v3, mandatory in v4) - integrity and privacy (confidentiality) in addition to authentication
* identity mapping plays no role in authentication or access control!
* `krb5i` in /etc/exports (on server) or /etc/fstab (on client) means Kerberos authentication and integrity

Sources
* ULSAH, 4th
