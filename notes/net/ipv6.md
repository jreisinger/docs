* origninally to solve IP address exhaustion problem
* 128 bits long
* boundary between network and host portion is fixed at /64
* so there is no more subnetting (but "subnet" lives as synonym for "local network")
```
2607:f8b0:000a:0806:0000:0000:0000:200e
2607:f8b0:a:   806:    0:   0:   0:200e # whitespace added for readability
2607:f8b0:a:   806:               :200e # whitespace added for readability
```
* each 16-bit group is represented by 4 hexadecimal digits
* in IPv4 notation each byte (8 bits) is represented by a decimal number
* loopback address (anologous to 127.0.0.1 in IPv4): `::1`