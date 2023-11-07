To fight the imposter syndrome ...

2023

- did and tested Datadog-Pagerduty integration to call us when there's an issue
- unified log attributes in Datadog
- wrote `sectools/scripts/httpver.sh` to check HTTP versions of our endpoints
- delivered company-wide presentation about data breach and security
- obtained "AWS Certified Cloud Practitioner" certification
- built `sectools/dd-awsinfo`
- found and reported leaked git secrets
    - documented and explained to engineer how to prevent it
    - engineers with my help reviewed commited secrets
- found and reported public S3 buckets
    - built `sectools/aws-s3pub`
- lowered AWS and Datadog costs
    - buitl `sectools/aws-dbinfo`

2022

- understood and documented *Kubernetes ingress WAF*
    - operated Kubernetes ingress WAF: whitelisting and deploying rules
    - built `kubectl-modsec` to extract ModSecurity WAF information from Kubernetes
- built (with P's help) a dummy gRPC application `demologger-grpc` to demonstrate and test logging and tracing
- interviewed and evaluated cloud team candidates - mostly developers
- set up and documented *security monitoring*: logs, detection rules, alerting
    - learned Datadog
    - delivered company-wide presentation about security monitoring
    - [ongoing] operating security monitoring and handling alerts
