# Assume role via aws CLI tool

1. Create policy "s3-create-bucket" that allows creating s3 buckets:

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "CreateBucket",
            "Effect": "Allow",
            "Action": [
                "s3:PutBucketPublicAccessBlock",
                "s3:ListAllMyBuckets",
                "s3:PutBucketOwnershipControls",
                "s3:CreateBucket"
            ],
            "Resource": "*"
        }
    ]
}
```

2. Create role "s3-create-bucket" that allows creating s3 buckets (i.e. refers
   the above policy) for the same account (e.g. 123456789012). It's trust
   policy will look like:

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "AWS": "arn:aws:iam::123456789012:root"
            },
            "Action": "sts:AssumeRole",
            "Condition": {}
        }
    ]
}
```

3. Create `assume-s3-create-bucket-role` policy:

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "AssumeRolePolicy",
            "Effect": "Allow",
            "Action": "sts:AssumeRole",
            "Resource": "arn:aws:iam::123456789012:role/s3-create-bucket"
        }
    ]
}
```

4. Attach `assume-s3-create-bucket-role` policy to your user.

5. Add to your `~/.aws/config`:

```
[profile assume-role]
role_arn = arn:aws:iam::123456789012:role/s3-create-bucket
source_profile = default
```

6. Create an s3 bucket using the `assume-role` profile defined above:

```
aws --profile assume-role s3api create-bucket --bucket GLOBALLY_UNIQUE_NAME
```
