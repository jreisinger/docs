# Virtual LANs

Broadcast sent by one host in a VLAN => received and processed by all hosts

Devices in a VLAN need to be in the same subnet

* at least one router needs to be involved
* in real campus it's usually a multilayer (L3) switch

A switch interface can be:

* access interface -- sends/receives frames only in a single VLAN
* trunking interface -- sends/receives frames in multiple VLANs

## Trunking

* trunking -- allows passing frames from multiple VLANs over a single physical connection
* when there are multiple interconnected switches - VLAN trunking must be between switches
* Cisco supported trunking protocols: ISL, IEEE 802.1Q

### ISL

* proprietary
* only between two Cisco switches that support ISL
* encapsulates the original frame in another Ethernet header and trailer (adds VLAN number)

### IEEE 802.1Q

* more popular
* inserts an extra 4-byte VLAN header into the original frame's Ethernet header
* must recalculate FCS (frame check sequence)
* defines one VLAN (VLAN 1 by default) as a native VLAN (no header is added to frames in this VLAN)

## VLAN Trunking Protocol (VTP)

* Cisco proprietary
* exchange of VLAN configuration info
* when a switch changes its VLAN config => VTP causes that all switches synchronize VLAN IDs and names
* three VTP modes: server, client, transparent

Requirements for VTP to work between two switches:

* the links must be operating as a VLAN trunk (ISL or 802.1Q)
* case-sensitive domain names must match
* if configured, case-sensitive VTP passwords must match

### VTP Pruning

By default switch flood broadcasts to all active VLANs out all trunks.

Usually VLANs don't exist on all switches => VTP can determine which switches do not need a broadcast and then prune those VLANs from the trunks.

## Configuration

Stored in `vlan.dat` in flash memmory (good in case all switches lost power)

Switches in transparent mode - in both the running-config file as well as the vland.dat file in flash

To remove VLAN configuration: `delete flash:vlan.dat`

Adding a new VLAN (by default there is VLAN 1, to which all interfaces are assigned):

    (config)#vlan 2
    (config-vlan)#name My-vlan  # defaults to VLAN0002
    (config)#interface range fastethernet 0/13 - 14
    (config-if)#switchport accesss vlan 2
    (config-if)#switchport mode access  # optional step to disable trunking
    
    #show vlan brief
