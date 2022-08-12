# Identity and Access Management

* AWS services are 100% API driven
* these service APIs require permissions for all requests
* IAM is one of the AWS services that provides authn (identity) and authz (permissions)

# AWS account and root user

* has access to unique tasks (e.g. account settings, support plan, billing)
* each AWS account is associated with one and only one root user
* root user uses generic login URL
* email becomes the username
* you should use email alias so you can have multiple accounts

# IAM resources

![image](https://user-images.githubusercontent.com/1047259/156759604-c7e6dc08-6ddd-474a-836d-8fd06a2e5208.png)

Basic steps

1. Create policies (JSON) that define access to services.
2. Create a group and assign (permission) policies to it.
3. Create a user and assign them to (user) groups.
4. Create roles (from the predifined ones). Roles are like users but for services (machine accounts).

You can use [terraform](https://github.com/vallard/EKS-Training/blob/master/segment02-iam/iam.tf) to create these ^.

# Permissions

* group
* inline
* managed

* boundary permissions

# Credential types

* Sing-in credentials: username + password [+ MFA], for AWS console (web browser) access
* Access (API) keys: access key ID + secret key, for CLI and SDK access
* Access keys + session token, for role-based access
* CodeCommit credentials
* Keyspaces credentials

# Best security practices

* protect root account credentials (email address + password) at all costs
* don't use root account for day-to-day operations
* delete any existing access/secret (API) keys for root account
* enable (and enforce) MFA
* rotate long-term credentials (access/secret keys, passwords)
* follow principle of least privilege