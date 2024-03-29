![aws-vpc drawio](https://user-images.githubusercontent.com/1047259/184867484-9f11dfbf-e883-4c88-af03-5035ecba78c8.png)

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
    * must have an IGW attached to VPC
    * must have a route (0.0.0.0/0) pointing to outside world via the IGW
    <img width="941" alt="image" src="https://user-images.githubusercontent.com/1047259/196145423-e659d894-8224-469d-804a-d3c800a9a204.png">
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
* this resource costs money! 

Network ACL (subnet level)

* associate with 1+ subnets (but a subnet can have only 1 NACL)
* state*less* firewall resource -> you need to consider both directions (inbound/outbound rules)
* rules order is important
* suggestion: 1 NACL for public subnets and individual NACLs for private subnets

<img width="846" alt="image" src="https://user-images.githubusercontent.com/1047259/196148897-689f3a8d-1b61-43a3-9a6e-8b2d16959951.png">

Security groups (instance level)

* associate with 1+ network interface
* state*full* firewall resource
* only allow rules (what is not allowed is denied)
* rules evaluated as a whole
* suggestion: 1 security group per application per tier (public, private, db)

<img width="1185" alt="image" src="https://user-images.githubusercontent.com/1047259/196147838-7cf7516d-0a51-4b05-a1b2-b55c19bdcaff.png">

---

Gateway endpoint (Endpoints)

* attach to VPC
* s3 and DynamoDB access
* same-region resources only
* Route table entry required
* suggestion: use it instead of NAT gateway to optimize cost

Interface endpoint (Endpoints)

* similar to Gateway endpoint (same icon)
* associate with 1 subnet
* associate with 1 Security group
* associate with 1 AWS service
* creates ENI in subnet
* overrides DNS
* suggestion: use to connect to private endpoints in VPCs in other accounts
