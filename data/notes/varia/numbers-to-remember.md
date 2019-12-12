## Programmer

    .--------------------------------------------------------------------------------.
    |               Source: https://gist.github.com/hellerbarde/2843375              |
    +------------------------------------+---------------------+---------------------+
    | What                               | Time                | Human Readable Time |
    +------------------------------------+---------------------+---------------------+
    | L1 cache reference                 |              0.5 ns |                     |
    | L2 cache reference                 |              7   ns |                     |
    |                                    |                     |                     |
    | Main memory reference              |            100   ns |                     |
    | Read 1 MB sequentially from memory |        250,000   ns | 250 µs              |
    |                                    |                     |                     |
    | Compress 1K bytes with Zippy       |          3,000   ns | 3 µs                |
    |                                    |                     |                     |
    | Send 2K bytes over 1 Gbps network  |         20,000   ns | 20 µs               |
    | Round trip within same datacenter  |        500,000   ns | 0.5 ms              |
    | Send packet CA->Netherlands->CA    |    150,000,000   ns | 150 ms              |
    |                                    |                     |                     |
    | SSD random read                    |        150,000   ns | 150 µs              |
    | Read 1 MB sequentially from SSD*   |      1,000,000   ns | 1 ms                |
    | Disk seek                          |     10,000,000   ns | 10 ms               |
    | Read 1 MB sequentially from disk   |     20,000,000   ns | 20 ms               |
    '------------------------------------+---------------------+---------------------'

`*` Assuming ~1GB/sec SSD

## Sysadmin

Media

    .----------------------------------------------------------------------------------------------------.
    |                                     Source: ULSAH, p. 210, 204                                     |
    +-------------------+------------+----------------------+-------------------+-------------+----------+
    | What              | Size       | Sequential Read Speed| Random Read Speed | Write Speed | Cost     |
    +-------------------+------------+----------------------+-------------------+-------------+----------+
    | LTO-3 write speed |            |                      |                   |     80 MB/s | $0.06/GB |
    | LTO-4 write speed |            |                      |                   |    120 MB/s | $0.05/GB |
    | HD                | Terrabytes |             100 MB/s |            2 MB/s |             | $0.10/GB |
    | SSD               | Gigabytes  |             250 MB/s |          250 MB/s |             | $3.00/GB |
    '-------------------+------------+----------------------+-------------------+-------------+----------'

 * [Write Throughput Average for Enterprise HDDs](http://www.tomshardware.com/charts/enterprise-hdd-charts/-04-Write-Throughput-Average-h2benchw-3.16,3376.html)

IPv4 Address Classes

    .--------------------------------------------------------------------------------------------------------.
    | Class | First octet | Networks                  | Number of networks | Purpose                         |
    +-------+-------------+---------------------------+--------------------+---------------------------------+
    | A     | 1 - 126     | 10.0.0.0 - 126.0.0.0      | 2^7 - 2 = 126      | Unicast (large networks)        |
    | B     | 128 - 191   | 128.0.0.0 - 191.255.0.0   | 2^14 = 16,384      | Unicast (medium-sized networks) |
    | C     | 192 - 223   | 192.0.0.0 - 223.255.255.0 | 2^21 = 2,097,152   | Unicast (small networks)        |
    | D     | 224 - 239   |                           |                    | Multicast                       |
    | E     | 240 - 255   |                           |                    | Experimental                    |
    '-------+-------------+---------------------------+--------------------+---------------------------------'

RFC 1918 Private Address Space

    .------------------------------------------------------------------.
    | Network class | Networks                    | Number of networks |
    +---------------+-----------------------------+--------------------+
    | A             |   10.0.0.0 - 10.255.255.255 |                  1 |
    | B             |    172.16.0.0. - 172.31.0.0 |                 16 |
    | C             | 192.168.0.0 - 192.168.255.0 |                256 |
    '---------------+-----------------------------+--------------------'
