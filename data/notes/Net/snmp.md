SNMP
====

Network management protocol:

* discover device configuration, health, network connections
* modify some configuration
* agents (managed devices, server program) can send *traps* (notification
    messages) to management stations (client program)

SNMP organization
=================

SNMP defines a hierarchical namespace of variables. The naming hierarchy is
made up of *MIB*s - structured text files that describe the data accessible
through SNMP. *OID*s - names for a specific managed piece of data. For example
the OID that refers to the uptime of the system is 1.3.6.1.2.1.1.3.

The current basic MIB for TCP/IP is MIB-II (defined by RFC1213). In addition to
basic MIB, there are MIBs for
* various hardware interfaces
* various protocols
* individual vendors

A MIB is *just a schema* for naming management data. To be useful, there must be
an agent-side program that maps between the SNMP namespace and the device's
actual state. SNMP agents (like NET-SNMP) come with built-in support for MIB-II
and can be extended to support supplemental MIBs.

SNMP operations
===============

1. get -- read data
2. get-next -- step through MIB hierarchy or read table
3. set -- write data
4. trap -- unsolicited, asynchronous notification from server (agent)

Since SNMP messages could modify data some security is needed. The simplest
version is based on *community strings* (for reading and writing) which is
another name for passwords. SNMP v3 brings access control with higher security.

NET-SNMP
========

http://net-snmp.sourceforge.net

Authoritative free implementation of SNMP for UNIX and Linux. Includes:
* agent (deb: `snmpd`)
* command line tools (deb: `snmp`)
* server (for receiving traps)
* library for developing SNMP-aware applications (deb: `libsnmp15`)
* MIBs for network interfaces, memory, disk, processes and CPU (deb:
    `libsnmp-base`)

The agent is *easily extensible* - it can execute any command and return its
output as an SNMP response (you can monitor almost anything).

NET-SNMP tools
--------------

snmpget - query for a specific OID value

    snmpget -v1 -c commnunity-string my-device .1.3.6.1.2.1.1.1.0

* `-v1` -- use simple authentication
* `.1.3.6.1.2.1.1.1.0` -- OID for system description

snmpwalk - list available OIDs

    snmpwalk -v1 -c community-string my-device


Source
======

* ULSAH
