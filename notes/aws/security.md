# Security principles

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
* apply security controls at multiple layers (e.g., edge of network, VPC, load balancing, every instance and compute service, operating system, application, and code)

Automate security best practices 🤖

* automated software-based security mechanisms allow you to scale more securely, rapidly and cost-effectively
* create secure architectures
* define and manage security controls as code in version-controlled templates

Protect data in transit and at rest 🔒

* classify your data into sensitivity levels
* use encryption, tokenization, and access control where appropriate

Keep people away from data ⛔

* use mechanisms and tools to reduce direct access or manual processing of data
* this reduces the risk of mishandling of sensitive data and human error 

Prepare for security events 👮

* have incident management and investigation process adequate to your organization requirements
* run incident response simulations
* use tools with automation to increase speed for detection, investigation, and recovery

# Security steps

1. IAM
2. Use automation
3. Enable detection
4. Prepare for an incident

Source: https://youtu.be/u6BCVkXkPnM
