# Security Design Principles

**Implement a strong identity foundation**: Implement the principle of least privilege and enforce separation of duties with appropriate authorization for each interaction with your AWS resources. Centralize identity management, and aim to eliminate reliance on long-term static credentials.

**Enable traceability**: Monitor, alert, and audit actions and changes to your environment in real time. Integrate log and metric collection with systems to automatically investigate and take action.

**Apply security at all layers**: Apply a defense in depth approach with multiple security controls. Apply to all layers (for example, edge of network, VPC, load balancing, every instance and compute service, operating system, application, and code).

**Automate security best practices**: Automated software-based security mechanisms improve your ability to securely scale more rapidly and cost-effectively. Create secure architectures, including the implementation of controls that are defined and managed as code in version-controlled templates.

**Protect data in transit and at rest**: Classify your data into sensitivity levels and use mechanisms, such as encryption, tokenization, and access control where appropriate.

**Keep people away from data**: Use mechanisms and tools to reduce or eliminate the need for direct access or manual processing of data. This reduces the risk of mishandling or modification and human error when handling sensitive data.

**Prepare for security events**: Prepare for an incident by having incident management and investigation policy and processes that align to your organizational requirements. Run incident response simulations and use tools with automation to increase your speed for detection, investigation, and recovery.

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
