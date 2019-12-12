# LAN Switching

## Concepts

* Hub -- one collision domain
* Bridge -- two collision domains
* Switch -- one collision domain per interface (microsegmentation)

Switching logic

1. forward or filter frame depending on the MAC address
1. learn MAC addresses by seeing frames' source MAC address
1. avoid loops by using STP

Switch's MAC address table = switching table = bridging table = CAM = forwarding table

Flooding

* send frames out of all interfaces except the incoming interface
* switch does this when there is no matching entry in MAC address table

Inactivity timer

* goes from 0 up (zeroed every time the same source MAC address arrives)
* when MAC address table is full, oldest entries (with largest timer) are removed

STP

* allows building redundant switching topology without loops
* makes interfaces into blocking or forwarding state
* without STP frames for turned off hosts would loop forever (in a circular three switch topology)
 * the same is true for broadcasts

Switch internal processing

* store-and-forward (most used on recent switches) -- receive all bits before forwarding a frame (can check FCS and discard faulty frames)
* cut-through -- forward as soon as possible (lower latency)
* fragment-free -- forward after receiving the first 64 bytes (can discard frames errored due to a collision)

Switch features

* if a single device is connected to a port => dedicated bandwidth to that device (full duplex, doubling the bandwidth)
* multiple simultaneous conversations between devices on different ports
* rate adoption for devices with different Ethernet speeds

## LAN design

* collision domain (seprated by switch) -- set of NICs whose frames can collide (NICs share the available bandwidth)
* broadcast domain (separated by router) -- set of NICs who all receive a broadcast frame sent be one of them

VLANs

(LAN = all devices in the same broadcast domain)

* group users by department
* create multiple broadcast domains
* reduce workload for STP
* better security - keep sensitive data on separate VLAN
* separate IP phone traffic from the traffic of PC connected to the phone

Switch types

* access -- connects directly to end-user device
* distribution -- aggregation point for access switches (less cabling, better performance)
* core -- even more aggregation

Ethernet LAN media

* 10BASE-T -- copper, CAT3 or better, two pair, max. 100m
* 100BASE-TX -- copper, CAT5 UTP or better, two pair, max. 100m
* 1000BASE-T -- copper, CAT5e UTP or better, four pair, max. 100m
* 1000BASE-LX -- fiber, max. 5km

## Cisco switches

Two types

1. Catalyst -- for Enterprises (core switch 6500 can run Cisco IOS or Cat OS)
2. Linksys -- for home use

CLI can be accessed via

* console (rollover cable - DB-9 to RJ-45, bright blue) -- physical access
 * 9600 bits/second
 * No HW flow control
 * 8-bit ASCII
 * No parity bits
 * 1 stop bit
* Telnet
* SSH

CLI modes

.. user EXEC mode (user mode)

.. enable mode

    > enable

.. configuration changes in enable mode affect the active config (RAM) after pressing Enter!

Configuration modes

.. global (`hostname(config)#`)

    configure terminal

.. line (`hostname(config-line)#`)

    line console 0
    line vty 0 15

.. interface (`hostname(config-if)#`)

    interface <type> <number>

Exiting modes

* `end`, <Ctrl-z> -- go back to privileged EXEC mode
* `exit` -- go one configuration mode up

Configuration files

* Startup-config (`#show startup-config`) -- initial configuration, set after reload
* Running-config (`#show running-config`) -- current configuration, dynamically changed by configuration commands

Config files storage

* RAM (or DRAM) -- Running-config
* ROM -- bootstrap (or boothelper) - first loaded program that subsequently load Cisco IOS into RAM
* Flash memory (chip or removable memory card) -- Cisco IOS image, backups of config files, other files
* NVRAM -- Startup-config

SW initialization

* Cisco: reload
* PC: reboot, restart

Managing config files

    copy {tftp | runnning-config | startup-config} {tftp | running-config | startup-config}

.. file => NVRAM or file => TFTP -- file replaces the original one

.. file => RAM -- merge

.. save configuration changes

    copy running-config startup-config

.. revert changes in running-config

    copy startup-config running-config  # not 100% reliable
    reload                              # 100% reliable
 
.. erase NVRAM

    erase nvram:  # new, recommended
    write erase
    erase startup-config
 
.. erase running config -- erase NVRAM + `reload`

IFS (IOS File System) alternative names

* startup-config = nvram: = nvram:startup-config
* running-config = system:running-config

Setup mode -- initial switch configuration via questions (System configuration dialog)

## Switch configuration

### Features in common with routers

Password + hostname

    #configure terminal
    (config)#enable secret cisco  # hide (via MD5 hashing) clear text passwords in running-config
    (config)#hostname Emma
    (config)#line console 0       # serial console 
    (config-line)#password 123
    (config-line)#login
    (config-line)#exit
    (config)#line vty 0 15        # telnet
    (config-line)#password 123
    (config-line)#login
    (config-line)#exit
    (config)#exit
    #show running-config
    
.. with default seetings, telnet users are rejected
    
SSH

    #configure terminal
    (config)#line vty 0 15
    (config-line)#login local                  # local users, no AAA
    (config-line)#transport input telnet ssh   # to improve security, leave out telnet
    (config-line)#exit
    (config)#username foo password 123
    (config)#ip domain-name example.com
    (config)#crypto key generate rsa
    (config)#^Z
    #show crypto key mypubkey rsa
    
Password encryption

* `service password-encryption` - all existing and future passwords encrypted (uses type 7 algorithm)

Banners

* MOTD (`banner`) -- shown before login
* Login (`banner login #`) -- shown before login
* Exec (`banner exec Z`) -- shown after login

Logging and timeout

.. normally logs are emitted anytime, including right in the middle of a command - to improve this

    logging synchronous
    
.. timeout (0 0 never times out)

    exec-timeout <minutes> <seconds>

### Switch configuration and operation

default (factory) switch configuration

* all interfaces enabled (`no shutdown`)
* autonegotation for ports with multiple speeds and duplex setting (`speed auto`, `duplex auto`)

IP Address

* needed only for mamangement
* VLAN 1 -- default VLAN used on all ports
* configure VLAN 1 interface to access the switch

.. Static IP address

    (config)#interface vlan 1
    (config-ig)#ip address 192.168.1.200 255.255.255.0
    (config-ig)#no shutdown
    (config-ig)#exit
    (config)#ip default-gateway 192.168.1.1

.. DHCP

    (config)#interface vlan 1
    (config-ig)#ip address dhcp
    (config-ig)#no shutdown
    (config-ig)#^Z
    #show dhcp lease

Interfaces

    (config)#interface FastEthernet 0/1
    (config-if)#duplex full
    (config-if)#speed 100
    (config-if)#description Server1 connects here

Port security

.. if you know what devices are to be connected to particular interfaces

    switchport mode access
    switchport port security
    switchport port-security maximum <number>  # defaults to 1
    switchport port-security violation { protect | restrict | shutdown }  # default is shutdown
    
    switchport port-security mac-address <mac-address>  # use multiple times to define more than one
        or
    switchport port-security mac-address sticky  # dynamically learn MAC addresses

... actions on security violation

* protect -- discard offeding traffic
* restrict -- protect + send log and SNMP message
* shutdown -- restrict + disable the interface (discard all traffic)

.. diagnostics

    show running-config
    show port-security interface fastEthernet 0/1

VLANs => vlans.md
	
Securing unused interfaces

.. Cisco interfaces are by default "plug and play" interfaces -- enabled (`no shutdown`), automatically negotiate speed and duplex, assigned to VLAN 1, use VLAN trunking and VTP

.. security recommendations (only the first is really required):

* administratively disable the interface (`shutdown`)
* disable VLAN trunking and VTP (`switchport mode access`)
* assign the port to an unused VLAN (`switchport access vlan <number>`)

## Switch troubleshooting

Sample CCNA exam questions - http://www.cisco.com/web/learning/wwtraining/certprog/training/cert_exam_tutorial.html

Organized (formalized) troubleshooting:

1. analyze/predict normal operation (documentation, `show`, `debug`)
2. isolate problem (`show`, `debug`)
3. root cause analysis -- find the cause of the problems

### CDP

.. proprietary protocol to learn about network topology -- uses multicast frames (when supported) or sends CDP updates to all data-link addresses

.. commands:

* `show cdp neighbors [type number]`
* `show cdp neighbors detail`
* `show cdp entry <name>` -- same as "detail", but only for named neighbors

.. shown info:

* device identifier -- hostname
* address list -- L2 and L3 addresses
* local interface -- interface from which `show cdp` was issued
* port identifier -- text
* capabilities list -- router, switch
* platform -- model, OS level

.. Cisco recommends to disable CDP where no needed:

* per interface -- `no cdp enable`
* globally -- `no cdp run`

### L1 and L2

`show interfaces`, `show interfaces description`

        .--------------------------------------------------------------------------------------------------------.
        |                                LAN switch interface status codes                                       |
        +-----------------------+----------------------+------------------+--------------------------------------+
        | Line status (L1)      | Protocol status (L2) | Interface status | Typical root cause                   |
        +-----------------------+----------------------+------------------+--------------------------------------+
        | Administratively down | Down                 | disabled         | shutdown command                     |
        | Down                  | Down                 | notconnect       | cable problems, other device down    |
        | up                    | Down                 | notconnect       | up/down state not expected on switch |
        | Down                  | down (err-disabled)  | err-disabled     | port security disabled the interface |
        | Up                    | Up                   | connected        | interface working                    |
        '-----------------------+----------------------+------------------+--------------------------------------'
<!-- Original table data:
Line status (L1);Protocol status (L2);Interface status;Typical root cause
Administratively down;Down;disabled;shutdown command
Down;Down;notconnect;cable problems, other device down
up;Down;notconnect;up/down state not expected on switch
Down;down (err-disabled);err-disabled;port security disabled the interface
Up;Up;connected;interface working
-->

### L1

`show interfaces gi0/1 status`:

* `a`-half, `a`-100 -- means autonegotiated, not set manually with `speed {10|100|1000}`, `duplex {half|full}`

`show interfaces fa0/13` (Indicator column):

        .---------------------------------------------------------------------------------------------------------.
        | Problem         | Indicator                         | Root cause                                        |
        +-----------------+-----------------------------------+---------------------------------------------------+
        | Excessive noise | Many input errors, few collisions | Cable problem (category, damaged, EMI)            |
        | Collisions      | Collisions > .1% of all frames    | Duplex mismatch, jabber, DOS                      |
        | Late collisions | Increasing late collisions        | Collision domain, too long cable, duplex mismatch |
        '-----------------+-----------------------------------+---------------------------------------------------'
<!-- Original table data:
Problem;Indicator;Root cause
Excessive noise;Many input errors, few collisions;Cable problem (category, damaged, EMI)
Collisions;Collisions > .1% of all frames;Duplex mismatch, jabber, DOS
Late collisions;Increasing late collisions;Collision domain, too long cable, duplex mismatch
-->

### L2

.. comands

* `show mac address-table`
* `show mac address-table dynamic`

.. switch forwarding logic

1. determine VLAN
2. look for destination MAC address, but only in the VLAN
 1. found (unicast) -- forward frame out of the matching interface
 2. not found (unicast) -- flood the frame within the VLAN
 3. broadcast or multicast -- flood the frame within the VLAN

.. port security filtering

* shutdown -- `show inteface` or `show inteface status`
* protect and restrict -- not so obvious, `show port-security interface`

## WLANs

<IMG SRC="https://raw.github.com/jreisinger/blog/master/files/ccna/wlan.jpg" ALT="Sample WLAN" WIDTH=200 HEIGHT=200>

### Concepts

* IEEE 802.3 -- Ethernet LAN
* IEEE 802.11 -- Ethernet WLAN

WLAN

* must use half-duplex (HDX)
* CSMA/CA


Standards:

        .---------------------------------------------------------------------.
        |                             | 802.11a | 802.11b | 802.11g | 802.11n |
        +-----------------------------+---------+---------+---------+---------+
        | Ratified                    | '99     | '99     | '03     | '09     |
        | Max. speed with DSSS (Mbps) | -       |      11 |      11 |      11 |
        | Max. speed with OFDM (Mbps) |      54 | -       |      54 |     150 |
        | Frequency band (GHz)        |       5 |     2.4 |     2.4 | Both    |
        | Non-overlapping channels    |      23 |       3 |       3 |       9 |
        '-----------------------------+---------+---------+---------+---------'
<!-- Original table data:
;802.11a;802.11b;802.11g;802.11n
Ratified;'99;'99;'03;'09
Max. speed with DSSS (Mbps);-;11;11;11
Max. speed with OFDM (Mbps);54;-;54;150
Frequency band (GHz);5;2.4;2.4;Both
Non-overlapping channels;23;3;3;9
-->

Modes

* ad hoc mode
* infrastructure mode -- cannot send frames directly to each other must go through an AP

Service sets

* BSS -- uses a single AP
* ESS -- uses more than one AP, allows for roaming

#### L1

* sending/receiving of radio waves
* frequency band -- range of consecutive frequencies

FCC (US) oversees the frequency ranges:

        .----------------------------------------------------------------------.
        |                      FCC unlicensed freq. bands                      |
        +-------------+-------+------------------------------------------------+
        | Freq. range | Name  | Sample devices                                 |
        +-------------+-------+------------------------------------------------+
        | 900 Mhz     | ISM   | Older cordless phones                          |
        | 2.4 Ghz     | ISM   | Newer cordless phones and 802.[11,11b,11g,11n] |
        | 5 Ghz       | U-NII | Newer cordless phones and 802.[11a,11n]        |
        '-------------+-------+------------------------------------------------'
<!-- Original table data:
Freq. range;Name;Sample devices
900 Mhz;ISM;Older cordless phones
2.4 Ghz;ISM;Newer cordless phones and 802.(11,11b,11g,11n)
5 Ghz;U-NII;Newer cordless phones and 802.(11a,11n)
-->

Encodings:

1. FHSS (802.11a) -- uses all frequencies in the band, hopping to different ones hoping to avoid intereference
2. DSSS (802.11[b, g]) -- uses one of the 11 overlapping channels (or frequencies); has a bandwidth of 82 MHz (2.402 - 2.483 Ghz); 3 (1, 6, 11) out of 11 channels are non-overlapping, i.e. can be used in the same space for WLAN communication and they won't interfere (important when designing ESS)
3. OFDM (802.11[a, g, n]) -- like DSSS, WLANs using OFDM can use multiple non-overlapping channels

802.11n -- uses mutliple antennas (MIMO)

Wireless interference

* passing through matter absorbs/reflects the radio signal
* SNR -- WLAN signal compared to undesired signals (noise)

Coverage -- transmit power of the AP cannot exceed the FCC limits

Speeed

* weaker signals can pass data at lower speeds (multiple speeds)
* generally: higher freq. => faster data transf. => smaller coverage area (exception - 802.11n)

Capacity -- non-overlapping channels multiply the WLAN capacity, as three devices can communicate with three APs at the same time

#### L2

* collisions can always occur (two or more devices sending at the same time, using overlapping frequencies)
* CSMA/CA used to minimize the chance of collision -- random wait time + aknowledgement of every frame

### Deployment

1) Verify the existing wired network

* DHCP working
* VLANs (ESS - all switch ports connecting the APs have to be in the same VLAN)
* Internet connectivity

2) Install the AP and configure the wired IP details

3) Configure the wireless details

* IEEE standard (a, b, g, or multiple)
* wireless channel
* SSID (must be the same within ESS)
* transmit power

4) Install and configure one wireless client

* M$ autoconfig tool: ACM (older WZC)

5) Verify the WLAN works from the client

### Security

Issues:

* War drivers -- gain Internet access for free
* Hackers -- find information or deny service
* Employees -- install the AP in his office with default configuration
* Rogue AP -- steel passwords

Counter-measures:

* Manual authentication (war drivers, hackers gaining access, rogue AP) -- password (called key) on client and server
* Encryption (hackers stealing info in a WLAN) -- secret key + math to scramble the contents of the WLAN frame
* IDS, IPS (employee AP installation, rogue AP) -- Cisco SWAN architecture

#### Security techniques

WEP

* 1997, IEEE, should not be used today

.. problems:

* Static preshared keys (PSK) -- configured manually on each AP and client (many people didn't bother to regularly change it)
* Easily cracked keys -- short keys (64 bits, only 40 were actual unique key)

SSID cloaking, MAC filtering

* not a real security
* cloaking -- AP doesn't send a periodic Beacon frame but the client sends the Probe message
* filtering -- attacker can sniff the MAC addresses and change their own

WPA

* dynamic key exchange via TKIP (Temporal key integrity protocol)
* user authentication via 802.1X or simple device authentication using preshared keys

WPA2 (802.11i) - AES

        .----------------------------------------------------------------------------------------------.
        |                                    WLAN security features                                    |
        +----------------+------------------+-----------------------+---------------------+------------+
        | Standard       | Key distribution | Device authentication | User authentication | Ecryption  |
        +----------------+------------------+-----------------------+---------------------+------------+
        | WEP            | Static           | Yes (weak)            | None                | Yes (weak) |
        | Cisco          | Dynamic          | Yes                   | Yes (802.1x)        | Yes (TKIP) |
        | WPA            | Both             | Yes                   | Yes (802.1x)        | Yes (TKIP) |
        | 802.11i (WPA2) | Both             | Yes                   | Yes (802.1x)        | Yes (AES)  |
        '----------------+------------------+-----------------------+---------------------+------------'
<!-- Original table data:
Standard;Key distribution;Device authentication;User authentication;Ecryption
WEP;Static;Yes (weak);None;Yes (weak)
Cisco;Dynamic;Yes;Yes (802.1x);Yes (TKIP)
WPA;Both;Yes;Yes (802.1x);Yes (TKIP)
802.11i (WPA2);Both;Yes;Yes (802.1x);Yes (AES)
-->

---

Source:

* W. Odom: CCENT/CCNA ICDN1 (2012)
