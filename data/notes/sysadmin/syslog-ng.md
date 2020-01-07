Basics
======

Sofware for sending and receiving of log messages.

Understands these log formats
* RFC 3164 (BSD syslog)
* RFC 5424 (IETF syslog) - since v 3.0

Sources
-------

* where syslog-ng receives log messages

Syntax

    source <identifier> { source-driver(params); source-driver(params); ... };
    
    source s_demo_tcp { tcp(ip(10.1.2.3) port(1999)); };
    
    source s_demo_two_drivers {
           tcp(ip(10.1.2.3) port(1999));
           udp(ip(10.1.2.3) port(1999)); };

* define a source only once
* sources (and destinations) are initialized only when they are used in a log statement

When receiving messages using the UDP protocol, increase the size of the UDP receive buffer (`so_rcvbuf()`) on the receiver host (that is, the syslog-ng OSE server or relay receiving the messages).

Troubleshooting
===============

[Possible causes of losing log messages](https://www.balabit.com/documents/syslog-ng-ose-latest-guides/en/syslog-ng-ose-guide-admin/html/concepts-losing-messages.html) (latest version)

[Statistics](https://www.balabit.com/documents/syslog-ng-ose-latest-guides/en/syslog-ng-ose-guide-admin/html/chapter-log-statistics.html)

    syslog-ng-ctl stats

Resources
=========

* https://www.balabit.com/documents/syslog-ng-ose-3.3-guides/en/syslog-ng-ose-v3.3-guide-admin-en/html-single/index.html
