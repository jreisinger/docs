# Cloud Development Kit

```plain
[CDK App]           definition of your infrastructure using code
 `-[Stack]          bulding, == CloudFormation stack, deployment boundary
    `-[Construct]   bulding block, e.g. S3 bucket
    `-[Construct]   bulding block, e.g. Lambda function
```

*Constructs* represent cloud component abstractions which can be composed together into higher level abstractions via scopes. They are like Go packages or Perl modules. There's AWS construct Library.
