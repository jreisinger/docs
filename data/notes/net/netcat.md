(Up-to-date <a href="https://github.com/jreisinger/blog/blob/master/posts/netcat.md">source</a> of this post.)

<img src="https://raw.github.com/jreisinger/blog/master/files/knife.jpg" alt="Knife" height="63" width="109" align="right">

TCP/IP swiss army knife. Simple (yet powerful!) Unix utility that reads and writes data across network connections, using TCP or UDP.

Netcat as a Client
==================

Connect to some port of some host:

    nc <host> <port>

* your STDIN is sent to the host
* anything that comes back across network is sent to your STDOUT
* this continues indefinitely, until the network side closes (not until EOF on STDIN like many other apps)

Test remote HTTP server:

    nc google.com 80
    GET / HTTP/1.0

(press Enter two times after the `GET` line)

Check UDP port is open (`telnet` does not work for UDP ports) but keep in mind that you never know for sure what's the state if you don't get anything back (https://serverfault.com/questions/416205/testing-udp-port-connectivity):

    $ netcat -vu vpn.ist.ac.at 1194
    Connection to vpn.ist.ac.at 1194 port [udp/openvpn] succeeded!

Make sure no data (zero) is sent to the port you connect to:

    nc -v -z host.tld 21-25

Change source port / address (ex. to evade a FW):

    nc -p 16000 host.tld 22      # 16000 is the local port
    nc -s 1.2.3.4 host.tld 8181  # 1.2.3.4 is the local source address

Netcat as a Server
==================

Listen for an incoming connection on some port:

    nc -l -p <port>

Send a directory over the network:

.. host A (receiving data)

    nc -l -p 1234 | tar xvf -

.. host B (sending data)

    tar cf - </some/dir> | nc -w 3 <hostA> 1234

Send a whole partition over the network:

.. host A (receiving data)

    nc -l -p 1234 | dd of=backup_sda1

.. host B (sending data)

    dd if=/dev/sda1 | nc -w 3 <hostA> 1234

Run a command (potentially dangerous!); ex. open a shell access:

.. host A (server)

    nc -l -p 9999 -e /bin/bash

.. host B (client)

    nc hostA 9999

More
====

* <http://mylinuxbook.com/linux-netcat-command/>
