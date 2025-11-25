# Cloud Development Kit

```plain
[CDK app]          like a city, definition of your infrastructure using code
 └─[stack]         like a building, CDK Stack == CloudFormation stack, deployment boundary
    ├─[construct]  like a building block, represents cloud component abstraction, e.g. Lambda function
    └─...
```

## Constructs 

- can be composed together into higher level abstractions via scopes. 
- there's [AWS construct Library](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-construct-library.html) organized into modules
- constructs are TS/JS classes with the `(scope, id, props)` signature
- they are like Go packages

bin/cdk-workshop.ts - holds the `CDK app`:

```ts
#!/usr/bin/env node
import * as cdk from 'aws-cdk-lib';
import { CdkWorkshopStack } from '../lib/cdk-workshop-stack';

const app = new cdk.App();
new CdkWorkshopStack(app, 'CdkWorkshopStack');
```

lib/cdk-workshop-stack.ts - holds the the CdkWorkshopStack `stack` and the Function `construct`:

```ts
import { Stack, StackProps } from "aws-cdk-lib";
import { Code, Function, Runtime } from "aws-cdk-lib/aws-lambda";
import { Construct } from "constructs";

export class CdkWorkshopStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    // defines an AWS Lambda resource
    const hello = new Function(this, "HelloHandler", {
      runtime: Runtime.NODEJS_22_X,     // execution environment
      code: Code.fromAsset("lambda"),   // code loaded from "lambda" directory
      handler: "hello.handler",         // file is "hello", function is "handler"
    });
  }
}
```

---

More: [CDK workshop](https://catalog.us-east-1.prod.workshops.aws/workshops/10141411-0192-4021-afa8-2436f3c66bd8/en-US)