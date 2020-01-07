## Networking (Neutron)

[Allowed Address Pairs](https://docs.openstack.org/dragonflow/latest/specs/allowed_address_pairs.html) - feature that allows adding additional IP/MAC address pairs on a port to allow traffic that matches those specified values

[Floating IP address (FIP)](https://docs.openstack.org/ocata/user-guide/cli-manage-ip-addresses.html) - Each instance has a private, fixed IP address and can also have a public, or floating IP address. Private IP addresses are used for communication between instances, and public addresses are used for communication with networks outside the cloud, including the Internet.

[Virtual IP address (VIP)](https://medium.com/jexia/virtual-ip-with-openstack-neutron-dd9378a48bdf) - is an IP address that is shared among two or more instances (VMs). Can be implemented using the Allowed Address Pair ^.

Useful commands:

```
openstack network list              # all nets
openstack network show <id>         # details
openstack ip availability show <id> # IP addresses
```

## Orchestration (Heat)

```
heat stack-list
heat stack-show <id>
heat resource-list <id>                         # list stack resources
heat resource-show <stack-id> <resource-name>   # resource details
```

## Tips and tricks

[Releases](https://en.wikipedia.org/wiki/OpenStack#Release_history) (Ocata, Pike, ...)

Initiliaze env. vars:

```sh
source ~/.openrc
```

Monitor stack progress when creating/deleting/updating:

```sh
watch 'openstack stack event list <stack-name> | tail -n 30'
```
