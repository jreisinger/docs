# Design principles

This is from the [security](https://docs.aws.amazon.com/wellarchitected/latest/security-pillar/security.html) pillar of the AWS Well-Architected framework.

Implement a strong identity foundation 🔑

* implement principle of least privilege
* enforce separation of duties with authorization for each interaction with AWS resources
* centralize identity management
* aim to eliminate long-term static credentials

Enable traceability 🔍

* monitor, alert and audit actions in your environment in real time
* integrate log and metric collection with systems to automatically investigate and take action

Apply security at all layers 🏰

* apply a defense in depth approach 
* apply security controls to all layers (e.g., edge of network, VPC, load balancing, every instance and compute service, operating system, application, and code)

Automate security best practices 🤖

* use automated software-based security mechanisms
* create secure architectures
* define and manage security controls as code in version-controlled templates

Protect data in transit and at rest 🔒

* classify your data into sensitivity levels
* use encryption, tokenization, and access control where appropriate

Keep people away from data ⛔

* use mechanisms and tools to reduce direct access or manual processing of data
* this reduces the risk of mishandling of sensitive data and human error 

Prepare for security events 👮

* have incident management and investigation process suitable for your org. requirements
* run incident response simulations
* use tools and automation to increase speed for detection, investigation, and recovery

# Areas

Security foundations
IAM
Detection
Infrastructure protection
Data protection
Incident response
Application security

# Steps

1. IAM
2. Use automation
3. Enable detection
4. Prepare for an incident

More: https://youtu.be/u6BCVkXkPnM
