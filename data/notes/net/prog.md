# Sockets

- a method for IPC
- endpoints for communication
- allow processes to communicate on a host or over a network

## Socket types

1) `SOCK_STREAM` - similar to pipes

- bidirectional --> full duplex (*simultaneous* communication in both directions)
- connection-oriented = like a phone call
- byte-stream = no concept of message boundaries
- reliable = data will be either delivered exactly as transmitted or we'll get notification of a failure

2) `SOCK_DGRAM` - only garantee that message borders will be preserverd when read
   (but lower overhead)

## Domains

* a.k.a. Protocol (or Address) Families
* communication range and address format

1) `PF_INET` - socket is identified by host (IP address) and port

2) `PF_UNIX` - .. by filename (ex. `/tmp/mysock`)

Domains and types are identified by symbolic names above (that are mapped to
numeric constants) which are constants exported by `Socket` and `IO::Socket`.

## Protocols 

There's rarely more than one protocol for the given domain and type of socket.

1) `tcp`

2) `udp`

Protocols have names that correspond to numbers used by the OS.
`getprotobyname()` (built into Perl) returns these numbers:

``` perl
$ perl -le 'print "$_ -> ", scalar getprotobyname $_ for qw(tcp udp)'
tcp -> 6
udp -> 17
```

## Perl's built-in functions

- low-level direct access to every part of the system
- on error return `undef` and set `$!`
- `socket()` - make a socket
- `bind()` - give a socket a local name by binding it to an address
- `connect()` - connect a local socket to a (possibly remote) one
- `listen()` - ready a socket for connections from other sockets
- `accept()` - receive the connections one by one creating new sockets
- use `print` and `<>` or `syswrite` or `sysread` to communicate over a stream
  socket
- .. `send` and `recv` for datagram socket

## Workflows

Typical SERVER

1. socket()
2. bind() and listen()
3. loop in a blocking accept() waiting for incoming connections

Typical CLIENT

1. socket() and connect()

.. datagram clients don't need to connect(); they specify the destination as argument to send()

UDP client `bind()` vs `connect()`

* bind() - grab a particular port
* connect() - limit received replies so they come only from a particular server

# Sources

- [Foundation of Python Network Programming](https://github.com/brandon-rhodes/fopnp) (2014)
- Black Hat Python (2014)
- The Linux Programming Interface, Ch. 56-61 (2010)
- Perl Cookbook, Ch. 17 Sockets (2003)
- Network Programming with Perl (2001)

