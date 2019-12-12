`traceroute` shows the route (and the transit delays) the packets have to take to get to a destination host across an IP network. For example:

    $ traceroute sdf.lonestar.org
    traceroute to sdf.lonestar.org (192.94.73.15), 30 hops max, 60 byte packets
     1  192.168.1.1 (192.168.1.1)  5.475 ms  6.020 ms  6.647 ms
     2  st-static-srk231.87-197-192.telecom.sk (87.197.192.231)  8.832 ms  15.973 ms  15.933 ms
    < ... >
    20  ge8-7.distb1.sea2.hopone.net (209.160.60.194)  186.286 ms  186.246 ms  175.897 ms
    21  SDF.ORG (192.94.73.15)  174.879 ms  174.283 ms  174.816 ms

But what does the output mean exactly and how does `traceroute` work?

It displays the sequence of gateways (showing the name and the IP address) through which an IP packet travels to reach its destination. The three numbers are the round trip times for each gateway. You can sometimes see the following instead of the number of miliseconds:

 * `*` -- one response (i.e. error) packet not received => congestion or ICMP packet was dropped because it has a low priority
 * `* * *` -- no "time exceed" messages received at all => gateway is down, firewall discards the packets or packets are slow to return
 * one of `!N`, `!H`, `!P` -- "network unreachable", "host unreachable", "protocol unreachable" - in any of these cases usually this is the last gateway you can get to => routing problem or a broken network link

`traceroute` works by sending three packets to each gateway on its route. In Linux, UDP packets are used by default; ICMP echo request or TCP SYN packets can also be used. These packets have artificially low TTL field (actually "hop count to live") set. The first three packets have TTL of 1. When they reach the gateway their TTL is decreased and when it reaches 0 the gateway discards the packet and sends back an ICMP "time exceeded" message. The originating host exctracts the gateway's IP address from the header of the error packet and resolves it to a name by using the DNS. This process repeats until the destination is reached or the gateway number limit (30) is exceeded.

![traceroute in Wireshark](https://raw.github.com/jreisinger/blog/master/files/wireshark-traceroute.png "traceroute in Wireshark")
