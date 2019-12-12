Basic concepts
--------------

EC2 (Elastic Compute Cloud)

* core part of AWS
* launched in 2006
* rent computing resources by the hour in the form of VMs aka instances
* idea from 1960s
* main advantage - cost savings due to flexibility, ex. instances can be launched/terminated based on the traffic levels
* AMI (Amazon Machine Image) - an image of a configured instance

Processing power

* "One EC2 Compute Unit provides the equivalent CPU capacity of a 1.0-1.2 GHz 2007 Opteron or 2007 Xeon processor."
* you need to benchmark it for your needs since AWS is built by incrementally adding commodity HW

Storage

1. instance storage - attached to the physical host that runs your instance; recommended for fast temporary storage
2. EBS (Elastic Block Store) - attached over the network; recommended for most of use cases
3. SSDs - since 2012; similar to instance storage but massively higher IO speeds

* it's possible to attach multiple volumes (of either type) and build a SW RAID on them

Networking

* every instance has a private IP (can only be used within EC2 network) and a public IP
* behind the scenes AWS will map a public/private IPs and create two DNS entries
* split-view DNS system - different responses depending on from where do you ask => always reference the public hostname of the instance (rather than the public IP) to save [costs](https://aws.amazon.com/ec2/pricing/#DataTransfer)

Resources

* AWS System Administration
