Last reviewed: *2023-03-12*

* despite its name it can do much more than capturing TCP headers
* can sniff traffic on many network types (including 802.1Q VLAN)
* de facto standard for command line packet analysis in Unix environment

# Useful options

`-D` -- list available interfaces

`-i INTERFACE` -- listen on INTERFACE (default: lowest numbered interface)

`-w FILE` -- write raw packets to FILE

`-r FILE` -- read packets from FILE

`-nn` -- turn off host and protocol name resolution (to avoid generating DNS
packets)

`-s0` -- set snaplength to 0, i.e. read the whole packet not just first 68
bytes (default if version >= 4.0)

`-t` -- turn off timestamp entries

`-c COUNT` -- capture COUNT packets and stop

Examples:

```sh
tcpdump -nni any -w packets.pcap
tcpdump -nnr packets.pcap
```

# Output format

will vary based upon what protocols are in use ...

TCP:

```plain
timestamp L3_protocol sIP.sPort > dIP.dPort: TCP_flags,
TCP_sequence_number, TCP_acknowledgement_number, TCP_windows_size,
data_length_in_bytes
```

UDP:

```plain
timestamp L3_protocol sIP.sPort > dIP.dPort: L4_protocol, data_length
```

## Output options

* use up to `-vvv` to provide more information on headers
* use `-x` to get entire packets (including data not just headers) in hex format
* use `-A` to get entire packets in ASCII format
* use `-X` to get entire packets in hex and ASCII format

# Packet Filtering

* utilizes the Berkeley Packet Filter (BPF) format
* added to the end of the command (recommended to use single quotes)

```sh
tcpdump -nnr packets.pcap 'tcp dst port 8080' -w packets_tcp8080.pcap
tcpdump -nnr packets.pcap -F known_good_hosts.bpf
```

## BPF

```plain
           operator
 primitive   |      primitive
     |       |         |
+---------+  | +----------------+
|         |  | |                |
udp port 53 && dst host 192.0.2.2
 |        |
 |        value
qualifier
```

Qualifiers

* host
* net - network in CIDR notation
* port
* src - communication source
* dst - communication destination
* ip - IP protocol
* tcp - TCP protocol
* udp - UDP protocol
* greater 1000 - data length is greater than 1000 bytes

Logical operators

* && - true when both conditions are true
* || - true when either condition is true
* ! - true when a condition is NOT met

Examples

* `host 192.0.2.100` -  match traffic to/from 192.0.2.100
* `dst host 2001:db8:85a3::8a2e:370:7334` - match traffic to the IPv6 address
* `ether host 00:50:56:98:60:92` - match traffic to the specified MAC address
* `!port 22` - match any traffic not to/from port 22
* `icmp` - match all ICMP traffic
* `!ip6` - match everything that is not IPv6

# Cookbook

[Show](https://serverfault.com/questions/504431/human-readable-format-for-http-headers-with-tcpdump) HTTP Host header:

```sh
stdbuf -oL -eL /usr/sbin/tcpdump -nn -A -s 10240 \
"tcp port 80 and (((ip[2:2] - ((ip[0]&0xf)<<2)) - ((tcp[12]&0xf0)>>2)) != 0)"
```

---

# Resources

* Applied Network Security Monitoring
