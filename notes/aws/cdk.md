# Cloud Development Kit

```plain
[CDK App]
 `-[Stack] - like a bulding, == CloudFormation stack, deployment boundary
    `-[Construct] - like a bulding block, e.g. S3 bucket
    `-[Construct] - like a bulding block, e.g. Lambda function
```

*Constructs* represent cloud component abstractions which can be composed together into higher level abstractions via scopes.
