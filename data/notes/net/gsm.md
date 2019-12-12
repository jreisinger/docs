(Up-to-date <a href="https://github.com/jreisinger/blog/blob/master/posts/gsm.md">source</a> of this post.)

## Cellular network

 * a radio network distributed over land areas called cells
 * each cell is served by at least one transceiver - **BTS** (Base Transceiver Station) = cell site
 * this enables a large number of portable transceivers (e.g. mobile phones) to communicate with each other
 * example of a cellular network: **the mobile phone network** or [PLMN](http://en.wikipedia.org/wiki/PLMN)

## GSM

 * World's most popular standard for mobile telephony systems (80% of mobile market uses the standard)
 * both signaling and speech channels are digital (1G was analog, ex. NMT)
 * second generation (2G) of mobile phone system
 * GSM release '97 - added packet data capabilities via GPRS
 * GSM release '99 - higher data transmission via EDGE
 * UMTS (Universal Mobile Telecommunications System) - 3G mobile cellular technology for networks based on GSM standards
 * LTE - 4G, standard for wireless communication of high-speed data for mobile phones and data terminals, based on the GSM/EDGE and UMTS/HSPA


<img src="https://raw.github.com/jreisinger/blog/master/files/mobile_technology_roadmap.png" alt="Mobile Technology Roadmap">

## Network Structure

GSM PLMN has two main logical domains:

1. access network - most used access networks in western Europe as of 2009 (can be deployed in parallel):
 * GERAN (GSM EDGE radio access network)
 * UTRAN (UMTS terrestrial radio access network) - HSPA can be implemented into UMTS to increase data transfer speed
2. core network
 * circuit switched domain
 * packet switched domain
 * IP multimedia subsystem (IMS)

GPRS/UMTS architecture with the main interfaces:

<img src="https://raw.github.com/jreisinger/blog/master/files/plmn.jpeg" alt="PLMN" height="400" width="700">

The network is structured into a number of discrete sections:

 * the base station subsystem (**BSS**) - handles traffic and signaling between a mobile phone and the NSS (access network)
 * the network and switching subsystem (**NSS**) - part of the network most similar to a fixed network (VOICE, circuit switched)
 * the **GPRS core network** - optional part for packet based Internet connections (NON-VOICE, packet switched)
 * operations support system (**OSS**) for maintenance

![GSM network](https://raw.github.com/jreisinger/blog/master/files/gsm_structure.png)

See [this picture](https://raw.github.com/jreisinger/blog/master/files/gsm_communication.jpg) for **GSM communication**.

**BSC** = Base Station Controller

 * intelligence behind the BTSs (allocation of radio channels, measurements from the mobile phones, handover control from BTS to BTS)
 * concentrator towards the mobile switching center (MSC)
 * the **most robust** element in the BSS
 * often based on a distributed computer architecture

**PCU** = Packet Control Unit

 * late addition to the GSM standard
 * processing tasks for packet data

**MSC** = Mobile Switching Centre

**HLR** = Home Location Register

 * database of subscribers
 * a central database that contains details of each mobile phone subscriber that is authorized to use the GSM and/or WCDMA core network of this PLMN

**VLR** = Visitor Location Register

 * register of roaming subscribers

**AUC**

 * database of authentication keys

**EIR**

 * stolen devices (phones) register

**SS7** = Signaling System #7

 * a set of telephone signaling protocols
 * main purpose: setup/tear down telephone calls
 * other uses: number portability, SMS, etc.

**SGSN** = Serving GPRS Support Node

 * delivery of data packets from and to mobile stations withing its geographical service area
 * packet routing and transfer, mobility management, logical link management, authentication and charging functions

**GGSN** = Gateway GPRS Support Node

 * **main component** of the GPRS network
 * inter-networking between the GPRS network and external packet switched networks
 * router to a sub-network

## AT commands

Huawei, Android

 * `at+cgmi` - manufacturer
 * `at+cgmm` -  model
 * `at+cimi` - IMSI
 * `at+cmgw="0914123456",145,"STO UNSENT"` - store message to memory
 * `at+cmgl="all"` - show stored messages
 * `at+cmss=3` - send message n. 3 from memory
 * `at+cmgd=2` - delete message n. 2 from memory

## Links

General

* [Mobile Internet Usage](http://lib.tkk.fi/Dipl/2009/urn100072.pdf) -- Thesis by a Finnish student
* [Mobile network](http://en.wikipedia.org/wiki/Mobile_network)
* P. Luptak: [Strucny prehlad do bezpecnosti GSM](http://www.nethemba.com/gsm-zranitelnosti.pdf) (in Slovak)

AT commands

* [Send SMS using AT commands](http://www.smssolutions.net/tutorials/gsm/sendsmsat/) - I was able to send an SMS following this guide
* [AT+C commands of GSM devices](http://gatling.ikk.sztaki.hu/~kissg/gsm/at+c.html)
* http://www.traud.de/gsm/
* [SMS Tutorial](http://www.developershome.com/sms/)

Hacking

 * [Osmocom OpenBSC](http://openbsc.osmocom.org/trac/) - functionality of BSC (Base Station Controller), MSC (Mobile Switching Center), HLR (Home Location Register), AuC (Authentication Center), VLR (Visitor Location Register), EIR (Equipment Identity Register)
 * [AirProbe](https://svn.berlin.ccc.de/projects/airprobe/) - GSM-Sniffer
 * [Kraken](http://reflextor.com/trac/a51) - cryptographic weaknesses found in today's cellular networks
 * [Nové trendy v GSM odpočúvaní](https://www.nethemba.com/sk/blog/-/blogs/nove-trendy-v-gsm-odpocuvani) (P. Luptak)
 * [GSM security map](http://gsmmap.org/)
 * [Decrypting GSM phone calls](http://srlabs.de/research/decrypting_gsm/) - tools
 * [28c3: Defending mobile phones](http://www.youtube.com/user/28c3#p/search/0/YWdHSJsEOck) (Video) - impersonating another MS
  * [28c3: Defending mobile phones](http://events.ccc.de/congress/2011/Fahrplan/attachments/1994_111217.SRLabs-28C3-Defending_mobile_phones.pdf) (PDF)
  * ["We will release tools to be used at the camp"](http://srlabs.de/events/gprs-intercept-wardriving-phone-networks-at-the-ccc-camp-finowfurt-august-10-2011/)
 * [27c3: Wideband GSM Sniffing](http://www.youtube.com/watch?v=lsIriAdbttc) (Video) -- Call/SMS interception and decrypting
  * [clarifications about 27C3 GSM Sniff Talk](http://lists.osmocom.org/pipermail/baseband-devel/2010-December/000912.html) -- you can't get the tools used for cracking A5/1 and traffic sniffing

PDUSpy

 * [WAP Push SI exploit](http://www.silentservices.de/adv03-2009.html)
 * [PDUSpy manual](http://www.nobbi.com/pduspy.html)
 * [PDUSpy download](http://www.nobbi.com/download.html#pduspy)

## Books

 * M. Grayson et al.: IP Design for Mobile Networks (Cisco Press, 2009)
 * A. Henry-Labordere, V. Jonack: SMS and MMS Interworking in Mobile Networks (Artech House, 2004)

