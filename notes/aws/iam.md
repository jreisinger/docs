# Identity and Access Management

* AWS services are 100% API driven
* these service APIs require permissions for all requests
* IAM is one of the AWS services that provides authn (identity + verification) and authz (access level or type)

# AWS account and root user

* has access to unique tasks (e.g. account settings, support plan, billing)
* each AWS account is associated with one and only one root user
* root user uses generic login URL
* email becomes the username
* you should use email alias so you can have multiple accounts

# IAM resources

![156759604-c7e6dc08-6ddd-474a-836d-8fd06a2e5208](https://user-images.githubusercontent.com/1047259/184610339-fa8c4a0c-b853-4dcc-a05f-2d8f9b6960a9.png)

Basic steps

1. Create policies (JSON) that define access to services.
2. Create a group and assign (permission) policies to it.
3. Create a user and assign them to (user) groups.
4. Create roles (from the predifined ones). Roles are like users but for services (machine accounts).

You can use [terraform](https://github.com/vallard/EKS-Training/blob/master/segment02-iam/iam.tf) to create these ^.

## Roles

* IAM identity
* Associated with permission policies (inline, managed)
* assumed by other principals
* similar to `sudo`
* similar to user but in addition has role trust policy

## Permission policies and boundaries

Permission policies (grant permissions to perform actions)

* identity-based policies
  * (AWS- or customer-)managed: standalone resource, version controlled, can be associated with 1+ IAM users, groups or roles
  * inline: embedded with (parameter of) IAM user, group or role
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
            "Sid": "PartnerUploadOnlyWithACL",
            "Effect": "Allow",
            "Principal": {"AWS": "123456789012"},
            "Action": "s3:PutObject",
            "Resource": "arn:aws:s3:::brightkey-data-bucket/*",
            "Condition": {
                "StringEquals": {"s3:x-amz-acl": "bucket-onwer-full-control"}
            }
        }
    ]
}
```

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "AccessToSpecificBucketsOnly",
            "Effect": "Allow",
            "Principal": "*",
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

# Federation

* you can use an outside identity provider and manage only permissions via IAM
* SAML (e.g. external AD, Google Gsuite, Okta), OpenID (e.g. Amazon, Facebook, Google)
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
