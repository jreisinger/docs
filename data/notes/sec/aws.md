# Security Design Principles

Implement strong identity foundation

* implement least privilege and separation of duties
* centralize IAM
* aim to eliminate reliance on long-term static credentials

Enable traceability

* monitor, alert and audit actions and changes to you environment in real time
* integrate log and metric collection with systems

Apply security at all layers

* implement defense in depth with multiple security controls
* e.g. at network edge, VPC, load balancing, every instance, OS, application and code level

Automate security best practices

* create secure architectures
* define and manage security controls as code in version-controlled templates

Protect data in transit and at rest

* classify data into sensitivity levels
* use encryption, tokenization, and access control where appropriate

Keep people away from data

* reduce or eliminate need for direct access or manual processing of data

Prepare for security events

* have incident management and investigation policy and process that align to your organizational requirements
* use tools and automation to increase your speed for detection, investigation, and recovery
* run incident response simulations

Source: https://docs.aws.amazon.com/wellarchitected/latest/security-pillar/security.html

# Security Steps

1. IAM
2. Use automation
3. Enable detection
4. Prepare for an incident

Source: https://youtu.be/u6BCVkXkPnM

# IAM

Credential types

* Email + password: master account (root!) access
* Username + password: AWS web console
* Access key + secret key: CLI, SDK
* Access/secret key + session token: role-based access

IAM Best practices

* protect master (root) account credentials (email address + password) at all costs
* don't use root account for day-to-day operations
* delete any existing access/secret keys for root account
* enable (and enforce) MFA
* follow principle of least privilege
* rotate long-term credentials (access/secret keys, passwords)

Source: Amazon Web Services AWS LiveLessons 2nd Edition by Richard Jones (2019)
