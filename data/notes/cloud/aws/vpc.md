![aws-vpc drawio](https://user-images.githubusercontent.com/1047259/184865788-37330a4e-5af1-4793-9ddf-6779a0006fa6.png)

VPC = Virtual Private Cloud!

* virtual data center network
* Region scope
* use RFC 1918 IPv4 CIDR range or bring your own
* up to 5 CIDR ranges, largest /16, smallest /28 (16 - 5 = 11 IP addresses)
* AWS creates a default VPC (and corresponding subnets) in every Region for you
  with CIDR 172.31.0.0/16

VPC subnets

* contiguous range of IP addresses in a VPC
* AZ scope (or Local Zone scope)
* associate with Route table and Network ACL
* types
  * public: bidirectional Internet access via IGW
  * private: outbound Internet access via NAT GW
  * VPC/VPN only subnet: no Internet access, or only via VPN/DX
* 5 reserved IP addresses:
  * 0: network
  * 1: VPC router
  * 2: DNS (if base VPC CIDR)
  * 3: Reserved for future use
  * last: broadcast address (not used)

Route table

* associate with 1+ subnets
* evaluated at subnet boundary for outbound traffic (what's the next hop)
* suggestion: 1 route table per subnet to minimize blast radius of changes

Internet Gateway (IGW)

* attach to VPC
* AWS public network access (AWS service APIs)
* Internet access
* requires a subnet route table entry and public IP address for clients to use

NAT Gateway

* outbound access to Internet
* outbound access to remote networks
* deploy into subnet
* AZ scope (not as fault tolerant as Region scope)
* suggestion: 1 NAT GW per AZ for higher resilience
 
