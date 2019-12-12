(Up-to-date <a href="https://github.com/jreisinger/blog/blob/master/posts/osmocombb.md">source</a> of this post.)

[OsmocomBB](http://bb.osmocom.org/trac/) (Open source mobile communications BaseBand) is an GSM Baseband software implementation. It intends to completely replace the need for a proprietary GSM baseband software. By using OsmocomBB on a compatible phone, you are able to make and receive phone calls, send and receive SMS, etc. based on Free Software. You can [learn](https://raw.github.com/jreisinger/blog/master/files/gsm_communication.jpg), hack and audit mobile networks with this tool.

Follow notes on how I got OsmocomBB runnning on Motorola C118 (brought to me by Mate :-).

## Compile
 1. [get started](http://bb.osmocom.org/trac/wiki/GettingStarted)
 1. `cd ~/osmocom-bb/src/target/firmware/`
 1. uncomment `CFLAGS += -DCONFIG_TX_ENABLE` in `Makefile`
 1. read [this](http://baseband-devel.722152.n3.nabble.com/I-wanna-make-sure-why-LOCATION-UPDATE-REQUEST-is-always-faild-td2655847.html) and [this](http://bb.osmocom.org/trac/wiki/SIMReader)

## Run
load layer1 code into mobile phone RAM

 1. `cd ~/osmocom-bb/src/host/osmocon`
 1. `sudo -E ./osmocon -p /dev/ttyUSB0 -m c123xor ../../target/firmware/board/compal_e88/layer1.compalram.bin`
 1. shortly press _On/Off_ button

run `mobile` - application implementing a regular GSM mobile phone (and more)

 1. `cd ~/osmocom-bb/src/host/layer23/src/mobile`
 1. `sudo -E ./mobile -i 127.0.0.1`

start terminal connection to `mobile`

 1. `cd ~/osmocom-bb/src/host/osmocon`
 1. `telnet localhost 4247`
  * `enable`
  * `sim pin`
  * `show ms 1 <PIN>`
  * `show subscriber`

## Wireshark

To install and run follow [this](http://bb.osmocom.org/trac/wiki/WiresharkIntegration). Quick how-to run wireshark:

    nc -u -l 127.0.0.1 4729 > /dev/null &   ## to discard ICMP port unreachable messages
    sudo wireshark -k -i lo -f 'port 4729'  ## listen on loopback device, port 4729

System information type 4

 * This message is sent on the BCCH (Broadcast Control Channel) by the network to all mobile stations within the cell giving information of control of the RACH (Random Access Channel), of location area identification (LAI), of cell identity and various other information about the cell.
 * Source: [I-ETS 300 022-1](http://www.scribd.com/doc/58945903/46/System-information-type-4) (1998)
 * See also: [Signaling Channels](http://www.gsmfordummies.com/tdma/logical.shtml)
