* global service supporting all regions
* used for monitoring, auditing and security
* captures API calls (console, CLI, SDKs and other services)
* captured events are stored in S3 buckets
* can be sent also to CloudWatch for monitoring 

```
CloudTrail (API event) --- log ---> S3 (log file)
                            |
                            +-----> CloudWatch logs
```

Use cases

* security monitoring and analysis
* evidence for compliance
* debugging ops issues
* tracking changes to AWS environment (AWS Config can also be used for this)
