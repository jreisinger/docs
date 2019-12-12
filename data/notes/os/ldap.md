(Up-to-date [source](https://github.com/jreisinger/blog/blob/master/posts/ldap.md) of this post.)

## Concepts and terms

* protocol for querying and modifying a X.500-based directory service running over TCP/IP
* current version - LDAPv3 (defined in [RFC4510](http://tools.ietf.org/html/rfc4510))
* Debian uses [OpenLDAP](http://www.openldap.org/) implementation (`slapd` package - recent versions are compiled with GnuTLS instead of OpenSSL due to licencing concerns)
* can be used for network authentication (login), similarly like Kerberos, Windows NT domains, NIS, AD ("LDAP + Kerberos")
 * replacement for `useradd`, `usermod`, `passwd`, `/etc/passwd`, `/etc/shadow`
* good for bigger networks

LDAP directory 

* hierarchical DB, more often read than written
* tree of *entries* or directory information tree (DIT)
* LDAP directory root = *base*

LDAP entry

* consists of set of *attributes*
* an attribute has a *type* (a name/description) and one or more *values*
* every attribute has to be defined in a at least one *objectClass* (a special kind of attribute)
* attributes and objeclasses are defined in *schemas*
* each entry has a unique identifier: *distinguished name* (DN) = relative distinguished name (RDN) + parent entry's DN
 * DN: "cn=John Doe,dc=example,dc=com"
 * RDN: "cn=John Doe"
 * parent DN: "dc=example,dc=com"
* DN in not an attribute, i.e. no part of the entry itself

## Preparing system to use LDAP (Debian 6.0.7)

Set [FQDN](http://wiki.debian.org/HowTo/ChangeHostname) if not already set (it is used by `slapd` for initial configuration):

* `/etc/hostname` (this file should only contain the hostname and not the full FQDN):

        ldap
        
* `/etc/hosts`:

        127.0.0.1       ldap.example.com ldap localhost
        # ... IPv6 stuff skipped ...
        1.2.3.4         ldap.example.com ldap
        
* Restart networking:

        invoke-rc.d hostname.sh start
        invoke-rc.d networking force-reload

* check hostname and FQDN

        $ hostname
        $ hostname -f

Install packages:

    aptitude install slapd ldap-utils
    
Configure `ldap-utils` (client programs):

    cp -p /etc/ldap/ldap.conf{,.orig}
    
    cat << EOF > /etc/ldap/ldap.conf
    # LDAP base - usually domain name
    BASE        dc=example,dc=com
    # ldap://, ldaps://
    URI         ldaps://ldap.example.com
    # certificate file (encryption)
    TLS_CACERT  /etc/ldap/ssl/certs/slapd-cert.crt
    EOF

## LDAP + TLS (Debian 6.0.7)

Configure [TLS](https://help.ubuntu.com/12.04/serverguide/openldap-server.html#openldap-tls):

* Create private key for certificate authority (CA):

        certtool --generate-privkey --outfile /etc/ssl/private/ca.<example.com>.key

* Create the template file `/etc/ssl/ca.info` to define the CA:

        cn = <Example Company>
        ca
        cert_signing_key

* Create self-signed CA certificate:

        certtool --generate-self-signed \
        --load-privkey /etc/ssl/private/ca.<example.com>.key \
        --template /etc/ssl/ca.info \
        --outfile /etc/ssl/certs/ca.<example.com>.cert

* Make a private key for the server:

        certtool --generate-privkey \
        --bits 1024 \
        --outfile /etc/ssl/private/ldap.<example.com>.key

* Create template file `/etc/ssl/ldap.info`:

        organization = <Example Company>
        cn = ldap.<example.com>
        tls_www_server
        encryption_key
        signing_key
        expiration_days = 3650

* Create the server's certificate:

        certtool --generate-certificate \
        --load-privkey /etc/ssl/private/ldap.<example.com>.key \
        --load-ca-certificate /etc/ssl/certs/ca.<example.com>.cert \
        --load-ca-privkey /etc/ssl/private/ca.<example.com>.key \
        --template /etc/ssl/ldap.info \
        --outfile /etc/ssl/certs/ldap.<example.com>.cert

* Create configuration file `/etc/ssl/certinfo.ldif`:

        dn: cn=config
        add: olcTLSCACertificateFile
        olcTLSCACertificateFile: /etc/ssl/certs/ca.<example.com>.cert
        -
        add: olcTLSCertificateFile
        olcTLSCertificateFile: /etc/ssl/certs/ldap.<example.com>.cert
        -
        add: olcTLSCertificateKeyFile
        olcTLSCertificateKeyFile: /etc/ssl/private/ldap.<example.com>.key

* Add configuration to LDAP:

        ldapmodify -Y EXTERNAL -H ldapi:/// -f /etc/ssl/certinfo.ldif

* Set up ownership and permissions of private key:

        adduser openldap ssl-cert
        chgrp ssl-cert /etc/ssl/private/ldap.<example.com>.key
        chmod g+r /etc/ssl/private/ldap.<example.com>.key
        chmod o-r /etc/ssl/private/ldap.<example.com>.key

* Edit `/etc/default/sldapd` ([Ubuntu says](https://help.ubuntu.com/12.04/serverguide/openldap-server.html#openldap-tls) it's not needed):

        SLAPD_SERVICES="ldap://127.0.0.1:389/ ldaps:/// ldapi:///"

* Restart and check LDAP

        service slapd restart
        
        netstat -tlpn | grep slapd
        tcp        0      0 0.0.0.0:636             0.0.0.0:*               LISTEN      16161/slapd
        tcp        0      0 127.0.0.1:389           0.0.0.0:*               LISTEN      16161/slapd
        tcp6       0      0 :::636                  :::*                    LISTEN      16161/slapd

## Populate LDAP via LDIF files

Create LDIF (LDAP Data Interchange Format) file with basic tree structure (`/var/tmp/tree.ldif`):

    # Account directory
    dn: ou=People,dc=example,dc=com
    ou: People
    objectClass: organizationalUnit

    # Group directory
    dn: ou=Group,dc=example,dc=com
    ou: Group
    objectClass: organizationalUnit

Create LDIF file with user account information (`/var/tmp/acct.ldif`):

    # User data (equivalent to /etc/passwd)
    dn: uid=jlebowski,ou=people,dc=example,dc=com
    uid: jlebowski
    uidNumber: 1010
    gidNumber: 100
    cn: Jeffrey
    sn: Lebowski
    displayName: JeffreyLebowski
    mail: the.dude@example.com
    objectClass: top
    objectClass: person
    objectClass: posixAccount
    objectClass: shadowAccount
    objectClass: inetOrgPerson
    loginShell: /bin/bash
    homeDirectory: /home/jlebowski

    # Group data (equivalent to /etc/group)
    dn: cn=users,ou=Group,dc=example,dc=com
    objectClass: posixGroup
    objectClass: top
    cn: users
    gidNumber: 100
    memberUid: jlebowski

Adding information from LDIF files to LDAP:

    ldapadd -c -x -D cn=admin,dc=example,dc=com -W -f /var/tmp/tree.ldif
    ldapadd -c -x -D cn=admin,dc=example,dc=com -W -f /var/tmp/acct.ldif

* `-c` -- continue even if errors are detected
* `-x` -- use a simpler authentication rather than the default SASL
* `-D` -- binds to the directory using the specified distinguished name (DN)
* `-W <binddn>` -- prompt for authentication
* `-f <file>` -- read LDIF records from <file>

## Accounts management

Changing password - one of:

* use `slappasswd`, cut/paste the hash into the LDIF file and run it through `ldapmodify`:

        ldappasswd -D cn=admin,dc=example,dc=com -W -S uid=jlebowski,ou=People,dc=example,dc=com

 * `-S` -- prompts for the new password
* if PAM is configured correctly, user can use `passwd` on an LDAP client

Deleting accounts:

        ldapdelete -c -x -D cn=admin,dc=example,dc=com -W uid=jlebowski,ou=people,dc=example,dc=com

## Querying a server about accounts

### Linux

GUI

* [JXplorer](http://jxplorer.org/)

See also [Querying Active Directory with Unix LDAP tools](http://jrwren.wrenfam.com/blog/2006/11/17/querying-active-directory-with-unix-ldap-tools/).

* `getent` returns info from various sources, including local account DB

        getent passwd jlebowski

* you can use filters (conceptually similar to regexes):

        ldapsearch -x uid=jlebowski
        
 * `(&(uid=jlebowski)(!(ou=Accounting)))` -- search for `jlebowski` who is _not_ a member of the Accounting department
 
* when `ldapsearch` sees UTF-8 encoding it displays it as base64, so you need to convert it:

        ldapsearch -x -h ldap.company.com -b 'dc=company,dc=com' -s sub -D 'user@company.com' -S 'employeeID' \
        -W '(&(objectClass=person)(employeeID>=0)(employeeID<=20416))' employeeID title sn | PERL_UNICODE=S \
        perl -MMIME::Base64 -MEncode=decode -n -00 -e 's/\n //g;s/(?<=:: )(\S+)/decode("UTF-8",decode_base64($1))/eg;print' \
        > ldap.out

### Windows

Use GUI tools like [ADExplorer](http://technet.microsoft.com/en-us/sysinternals/bb963907.aspx) or [LDAP Admin](http://www.ldapadmin.org/).

## More

* [Ubuntu LDAP guide](https://help.ubuntu.com/12.04/serverguide/openldap-server.html)
* [Querying LDAP from Perl](https://github.com/jreisinger/audit/blob/master/orsr/lib/My/Ldap.pm)
* Perl Cookbook: 18.8 Accessing an LDAP Server (p. 738)
* The Otter Book: Ch. 9 Directory Services (p. 313)
* ULSAH: 19.3 LDAP (p. 728)
