DNS database
============

Zone

* a portion of the DNS namespace
* a domain name that has been delegated to other servers/administrators

Zone files

* a zone's DNS database
* set of text files maintained by the sysadmin on the zone's master name server

Zone files have two types of entries:

* parser commands
* resource records

Parser commands (directives)
----------------------------

`$ORIGIN <domain-name>` - sets/changes the origin, i.e. the default domain (defaults to the domain name specified in the name server's config file)

`$INCLUDE <filename> [origin]` - allows you to separate records into files or to keep cryptographic keys in a file with restricted permissions

`$TTL <default-ttl>` - must be the first line of the zone file

Resource records
----------------

Special characters in resource records

    ;   comment
    @   the current zone name
    ()  allows data to span lines
    *   wild card (`name` field only)

Syntax

    [name] [ttl] [class] type data
    
name

* host or domain (if ommitted referes to the previous entity)
* relative or absolute (ends with a dot)

ttl

* how long can data be cached and still considered valid
* defaults to `$TTL` directive
* once records are cached outside you local network, you cannot force them to be discarded

class

* network type: IN (default), HS, CH (ChaosNet :-)

type

    SOA    Start of Authority (Defines a DNS zone)
    NS     Name Server
    A      IPv4 Address (Name-to-address translation)
    AAAA   IPv6 Address
    PTR    Pointer (Address-to-name translation)
    MX     Mail Exchanger (Controls email routing)
    DNSKEY Public Key (Public key for a DNS name - used for DNSSEC)
    CAA    Certification Authority Authorization
    SPF    Sender Policy Framework (Identifies mail servers, inhibits forging)
    DKIM   DomainKeys Identified Mail (Signature system for email - verify sender and message integrity)
    CNAME  Canonical Name (Nicknames or aliases for a host)
    SRV    Services (Gives locations of well-known services)
    TXT    Text (Comments or untyped information; used for trying out new ideas)

See [Cloudflare article](https://www.cloudflare.com/learning/dns/dns-records/) for more types.

*SOA*

* each zone has exactly one SOA record. 
* the SOA record includes the name of the zone, the primary name server for the zone, a technical contact, and various timeout values

*NS* records

* identify the servers that are authoritative fot a zone (all master and slave servers)
* delegate subdomains to other organizations

*CNAME*

* RFC1033 denies CNAME at zone apex, i.e. for FQDN that is the same name as the zone name (aka bare, naked or root domain)
 * for this scenario, you can use `ALIAS` record type
* if a domain name has a CNAME record, it is not allowed to have any other records according to the DNS standards

*CAA*

* a new type to indicate to CAs whether they are authorized to issue digital certificates for a particular domain name
* CAA can't coexist with CNAME

*SPF*

* allows administrators to specify which hosts are allowed to send email on behalf of a given domain

*DKIM* - powered by asymmetric cryptography

1. The sender's Mail Transfer Agent (MTA) signs every outgoing message with a
private key.
2. The recipient retrieves the public key from the sender's DNS records and
verifies if the message body and some of the header fields were not altered
since the message signing took place.

DNS query process
=================

<img src="https://www.cs.nmsu.edu/~istrnad/cs480/lecture_notes/dns_query.png" style="max-width:100%;height:auto;"> 

* `lair` is a client (lair.cs.colorado.edu)
* `ns.cs.colorado.edu` is the local nameserver for lair
* the answer is not cached at `ns.cs.colorado.edu` when `lair` is doing the query

Common return statuses

* `NOERROR` - the query returned a response without notable errors
* `NXDOMAIN` - the requested name does not exist (or isn't registered)
* `SERVFAIL` - usually a configuration error on the name server

Name server taxonomy
====================

authoritative - an official representative of a zone

* master, primary - gets data from a disk file
* slave, secondary - copies data from the master

non-authoritative - answers queries from cache; doesn't know if the data is still valid

* caching - caches data from previous queries; usually has no local zones
* forwarder - performs queries on behalf of many clients; builds a large cache

recursive - queries on your behalf until it returns either an answer or an error

non-recursive - refers you to another server if it can't answer a query

resolver (meaning 1) - client side software (library) doing lookups 

* Resolver libraries do not understand referrals. Any local nameserver listed in a client's `resolv.conf` file must be recursive.
 
resolver (meaning 2) - local nameserver (like that one you put in `/etc/resolv.conf`) for doing lookups

* some people call it recursor or recursive resolver

Tips and tricks
===============

host
----

    host name|addr [server]
 
nslookup
--------

(Cricket Liu doesn't like it :-))

    nslookup [name|addr] [server]

dig
---

    dig [@server] [-x addr] [name] [type] [+trace] [+short]

* queries the name servers configured in `/etc/resolv.conf` by default (use `@<nameserver>` to override)

The pseudo-type `any` is a bit sneaky: instead of returning all data associated with a name, it returns all cached data associated with the name. So, to get all records, you might have to do `dig domain NS` followed by `dig @ns1.domain domain any`. (Authoritative data counts as cached in this context.)

Find out the names of authoritative nameservers for a domain

```
dig ist.ac.at ns
```

Find out master nameserver

```
dig ist.ac.at soa
```

Find out the version of a bind nameserver (can be concealed in some cases)

```
dig @ns1.ist.ac.at version.bind txt chaos
```

Open resolver
-------------

* recursive, caching name server that accepts and answers queries from anyone
 on the Internet
* resources consumption
* resolver's cache poisoning
* amplification of DDoS attacks

Checking for open resolvers: http://dns.measurement-factory.com/tools/ => open resolver test

Client side
-----------

Find DNS server used by your system:

```
# Ubuntu 16.04
nmcli device show | grep IP4.DNS
```
 
Sources and more
================

* [Cloudflare](https://www.cloudflare.com/learning/dns/what-is-dns/)
* [ULSAH](http://ulsah.com/) ch. 17. DNS: The Domain Name System
* [SPF](https://www.digitalocean.com/community/tutorials/how-to-use-an-spf-record-to-prevent-spoofing-improve-e-mail-reliability)
* [DKIM](https://www.digitalocean.com/community/tutorials/how-to-install-and-configure-dkim-with-postfix-on-debian-wheezy)
