Fast, modern, secure VPN tunnel.

- securely encapsulates IP packets over UDP
- you add wg interface, configure it with private key, peers' public keys and send packets accross it
- all issues of key distribution and pushed configuration are out of scope (unlike in OpenVPN)

```sh
# add interface
ip link add dev wg0 type wireguard

# add IP address
ip addr add dev wg0 192.168.2.1/24
# or if there are only two peers
ip addr add dev wg0 192.168.2.1 peer 192.168.2.2

# generate keys
umask 077
wg genkey > privkey
wg pubkey < privkey > pubkey

# configure interface
wg set wg0 listen-port 51820 private-key $PRIVKEY peer $PUBKEY allowed-ips 192.168.2.0/24 endpoint $PEER_IPADDR:51820
wg [show[conf]]

# activate interface
ip link set up dev wg0

# cleanup
ip link del dev wg0
```

More: https://www.wireguard.com/quickstart/