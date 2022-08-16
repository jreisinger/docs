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

# Region

* multiple, physically separate AZs
* traffic between AZs is encrypted
* service API endpoints are hosted here
 
![aws-region drawio](https://user-images.githubusercontent.com/1047259/184850533-2ec4ecc2-0d07-4774-8e77-777415a2e08f.png)

# AZ (Availability Zone)

* one or more collocated (within walking distance) data centers
* HA building block
* atomic unit of resource scope
* there are logical and physical AZ IDs, e.g. eu-central-1a (euc1-az2)

# EC2 (Elastic Compute Cloud)

* core part of AWS, launched in 2006
* rent computing resources by the hour in the form of VMs aka instances
* idea from 1960s
* AMI (Amazon Machine Image) - an image of a configured instance

# Edge locations

* separate infrastructure from Regions
* connected to Regions networks
* scope for global services (e.g. DNS, content caching)
* cca 10 times as many edge locations as Regions
