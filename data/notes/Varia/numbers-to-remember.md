# Programmer

## Latency numbers (~2012)

```
L1 cache reference                           0.5 ns
Branch mispredict                            5   ns
L2 cache reference                           7   ns                      14x L1 cache
Mutex lock/unlock                           25   ns
Main memory reference                      100   ns                      20x L2 cache, 200x L1 cache
Compress 1K bytes with Zippy             3,000   ns        3 us
Send 1K bytes over 1 Gbps network       10,000   ns       10 us
Read 4K randomly from SSD*             150,000   ns      150 us          ~1GB/sec SSD
Read 1 MB sequentially from memory     250,000   ns      250 us
Round trip within same datacenter      500,000   ns      500 us
Read 1 MB sequentially from SSD*     1,000,000   ns    1,000 us    1 ms  ~1GB/sec SSD, 4X memory
Disk seek                           10,000,000   ns   10,000 us   10 ms  20x datacenter roundtrip
Read 1 MB sequentially from disk    20,000,000   ns   20,000 us   20 ms  80x memory, 20X SSD
Send packet CA->Netherlands->CA    150,000,000   ns  150,000 us  150 ms
```

Source: https://gist.github.com/jboner/2841832

# Sysadmin

## Media (~2010)

    +-------------------+------------+----------------------+-------------------+-------------+----------+
    | What              | Size       | Sequential Read Speed| Random Read Speed | Write Speed | Cost     |
    +-------------------+------------+----------------------+-------------------+-------------+----------+
    | LTO-3 write speed |            |                      |                   |     80 MB/s | $0.06/GB |
    | LTO-4 write speed |            |                      |                   |    120 MB/s | $0.05/GB |
    | HD                | Terrabytes |             100 MB/s |            2 MB/s |             | $0.10/GB |
    | SSD               | Gigabytes  |             250 MB/s |          250 MB/s |             | $3.00/GB |
    '-------------------+------------+----------------------+-------------------+-------------+----------'

Source: ULSAH, p. 210, 204

## Networking

IPv4 Address Classes

    .--------------------------------------------------------------------------------------------------------.
    | Class | First octet | Networks                  | Number of networks | Purpose                         |
    +-------+-------------+---------------------------+--------------------+---------------------------------+
    | A     | 1   - 126   | 1.0.0.0   - 126.0.0.0     | 2^7 - 2 = 126      | Unicast (large networks)        |
    | B     | 128 - 191   | 128.0.0.0 - 191.255.0.0   | 2^14 = 16,384      | Unicast (medium-sized networks) |
    | C     | 192 - 223   | 192.0.0.0 - 223.255.255.0 | 2^21 = 2,097,152   | Unicast (small networks)        |
    | D     | 224 - 239   |                           |                    | Multicast                       |
    | E     | 240 - 255   |                           |                    | Experimental                    |
    '-------+-------------+---------------------------+--------------------+---------------------------------'

[RFC 1918](https://tools.ietf.org/html/rfc1918) Private Address Space

    .----------------------------------------------------.
    | Class | Networks                      | Prefix     |
    +-------+-------------------------------+------------+
    | A     |    10.0.0.0 - 10.255.255.255  | 10/8       |
    | B     |  172.16.0.0 - 172.31.255.255  | 172.16/12  |
    | C     | 192.168.0.0 - 192.168.255.255 | 192.168/16 |
    '-------+-------------------------------+------------'

IPv4 private networks

```
# rfc 1918
10.0.0.0/8
172.16.0.0/12
192.168.0.0/16
# https://en.wikipedia.org/wiki/Reserved_IP_addresses
100.64.0.0/10
192.0.0.0/24
198.18.0.0/15
```
