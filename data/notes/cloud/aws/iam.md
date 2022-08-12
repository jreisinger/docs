# IAM (Identity and Access Management)

* AWS services are 100% API driven
* these service APIs require permissions for all requests
* each AWS account is associated with one an only one root user (email becomes the username)

![image](https://user-images.githubusercontent.com/1047259/156759604-c7e6dc08-6ddd-474a-836d-8fd06a2e5208.png)

1. Create policies (JSON) that define access to services.
2. Create a group and assign (permission) policies to it.
3. Create a user and assign them to (user) groups.
4. Create roles (from the predifined ones). Roles are like users but for services (machine accounts).

You can use [terraform](https://github.com/vallard/EKS-Training/blob/master/segment02-iam/iam.tf) to create these ^.

Credential types

* Email + password: master (root!) account
* Username + password: AWS console (web browser)
* Access key + secret key: CLI, SDK
* Access/secret key + session token: role-based access

IAM Best practices

* protect master (root!) account credentials (email address + password) at all costs
* don't use root account for day-to-day operations
* delete any existing access/secret (API) keys for root account
* enable (and enforce) MFA
* follow principle of least privilege
* rotate long-term credentials (access/secret keys, passwords)