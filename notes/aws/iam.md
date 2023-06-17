# Identity and Access Management

* AWS services are 100% API driven
* these service APIs require permissions for all requests
* IAM is one of the AWS services that provides authn (identity + verification) and authz (access level or type)

# AWS account and root user

```plain
Organization
    Account1
        root
        User1
        User2
        ...
    Account2
        root
        User1
        User2
        ...
    ...
```

* each AWS account is associated with one and only one root user
* root has access to unique tasks (e.g. account settings, support plan, billing)
* root uses generic login URL
* email becomes the username
* you should use email alias so you can have multiple accounts
* AWS accounts
  * master payer - contains billing data
  * linked - doesn't show billing data

# Organizations

to handle multiple accounts

terminology
* organization - hierarchical structure (family tree) of accounts
* root - container at the top of the tree
* OU - container to categorize/group accounts
* account - to configure and provision resources
* master account - account (with star) managing the organization; shouldn't be used to provision resources 
* SCP - what services and operations are allowed at most (applied recursively to root, OU or account)

# Service control policy (SCP)

* permission boundary (guardrail)
* doesn't grant access
* doesn't affect resource-based policies, only principals managed by your accounts
* doesn't affect actions performed by the master account

# IAM resources

![156759604-c7e6dc08-6ddd-474a-836d-8fd06a2e5208](https://user-images.githubusercontent.com/1047259/184610339-fa8c4a0c-b853-4dcc-a05f-2d8f9b6960a9.png)

Basic steps

1. Create policies (JSON) that define access to services.
2. Create a group and assign (permission) policies to it.
3. Create a user and assign them to (user) groups.
4. Create roles (from the predifined ones).

You can use [terraform](https://github.com/vallard/EKS-Training/blob/master/segment02-iam/iam.tf) to create these ^.

## Roles

* assumed by other principals (users, AWS services, applications)
* similar to `sudo`
* associated with permission policies (inline, managed)
* dynamically generated temporary credentials per session (no long-term credentials associated)

* trust relationship - who can assume the role, defined in the role
* assume policy - whether the principal can assume the role, defined in the principal

## Permission policies and boundaries

Permission policies (grant permissions to perform actions)

* identity-based policies
  * standalone resource, version controlled
  * can be associated with 1+ IAM users, groups or roles
  * AWS- or customer-managed
  * inline: embedded within IAM user, group or role
* resource-based policies
  * new element: Principal
  * Principal values vary by resource
  * many resource policies are optional (e.g. S3), some are required (e.g. KMS)
* session policies (there's no on-prem analogy)
  * parameter passed during creation of temporary session
  * used with IAM Role and federated users
 
Permission boundaries (define the maximum permissions for a resource)

* IAM permissions boundaries
* AWS organizations SCPs

[Policy evaluation logic](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html#policy-eval-denyallow).

[Policy format (JSON)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_grammar.html), e.g.:

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "AccessToSpecificBucketsOnly",
            "Effect": "Allow",
            // "Principal": "*",
            "Principal": {"AWS": "123456789012"},
            "Action": [ "s3:GetObject", "s3:PutObject" ],
            "Resource": [
                "arn:aws:s3:::brightkey-data1",
                "arn:aws:s3:::brightkey-data1/*",
                "arn:aws:s3:::brightkey-config1",
                "arn:aws:s3:::brightkey-config1/*"
            ]
        }
    ]
}
```

# AWS IAM Identity Center

* susccessor to AWS Single Sign-On
* centrally manage workforce access to multiple AWS ccounts and applications

# Federation

* you can use an outside identity provider and manage only permissions via IAM
* SAML (e.g. external AD, [jumpcloud](https://support.jumpcloud.com/s/article/Single-Sign-On-SSO-With-AWS-SSO), Google Gsuite, Okta), OpenID (e.g. Amazon, Facebook, Google)
* account scope or multiple accounts with AWS SSO (all accounts must be part of the same organization)
* mobile app federation with AWS Cognito

# Credential types

* Sing-in credentials: username + password [+ MFA], for AWS console (web browser) access
* Access (API) keys: access key ID + secret key, for CLI and SDK access
* Temporary credentials: access keys + session token, for role-based access
* CodeCommit credentials
* Keyspaces credentials

# Best security practices

* protect root account credentials (email address + password) at all costs
* don't use root account for day-to-day operations
* delete any existing access/secret (API) keys for root account
* enable (and enforce) MFA
* rotate long-term credentials (access/secret keys, passwords)
* follow principle of least privilege
