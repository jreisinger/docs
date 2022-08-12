# ARN (Amazon Resource Name)

* globally unique identifier
* `arn:<partition>:<service>:[<region>]:<account-id>:<resource-id>`, e.g.:
  * partition: aws, aws-cn, aws-us-gov
  * service: ec2, s3, iam
  * region: us-east-1, eu-west-1, ap-south-1
  * account-id: 123456789012
  * resource-id: User/Chad, instance/i-XXXXXX, volume/vol-XXXXXX

# AWS (Amazon Web Services)

* world's most comprehensive and broadly adopted cloud platform
* over 200 services from globally distributed data centers (Lego)
* allows for lower costs and more agility

# AWS console

* the web interface
* orange buttons are important
* only some services are Region scoped

# AWS Region

* multiple, physically separate AZs
* traffic between AZs is encrypted
* service API endpoints are hosted here

<img width="553" alt="image" src="https://user-images.githubusercontent.com/1047259/184126933-29fb8020-fd6e-425f-996a-cb353689dd4e.png">

# AZ (Availability Zone)

* one or more collocated (within walking distance) data centers
* HA building block
* atomic unit of resource scope
* there are logical and physical AZ IDs, e.g. eu-central-1a (euc1-az2)

# EC2 (Elastic Compute Cloud)

* core part of AWS
* launched in 2006
* rent computing resources by the hour in the form of VMs aka instances
* idea from 1960s
* main advantage - cost savings due to flexibility, ex. instances can be launched/terminated based on the traffic levels
* AMI (Amazon Machine Image) - an image of a configured instance

# Edge locations

* separate infrastructure from Regions
* connected to Regions networks
* scope for global services (e.g. DNS, content caching)
* cca 10 times as many edge locations as Regions

# Storage

1. instance storage - attached to the physical host that runs your instance; recommended for fast temporary storage
2. EBS (Elastic Block Store) - attached over the network; recommended for most of use cases
3. SSDs - since 2012; similar to instance storage but massively higher IO speeds

* it's possible to attach multiple volumes (of either type) and build a SW RAID on them

# Networking

* every instance has a private IP (can only be used within EC2 network) and a public IP
* behind the scenes AWS will map a public/private IPs and create two DNS entries
* split-view DNS system - different responses depending on from where do you ask => always reference the public hostname of the instance (rather than the public IP) to save [costs](https://aws.amazon.com/ec2/pricing/#DataTransfer)