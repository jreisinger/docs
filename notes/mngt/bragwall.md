To fight the imposter syndrome ...

2023

- wrote github.com/jreisinger/mathpractice
- trying to build cybersecurity program based on NIST CSF 2.0 (`notes/sec/nist-csf.md`)
- reviewing API security (`notes/sec/api.md`)
- embedding vulnerability scanning into CI/CD (`.github/workflows/scan-vulns-go.yml`)
- reviewed Go repojacking (`notes/sec/repojacking.md`, `tools/cmd/gorepojack`)
- reviewing and improving email security (DKIM, SPF, DMARC)
- did and tested Datadog-Pagerduty integration to call us when there's an issue
- unified log attributes in Datadog
- wrote `sectools/scripts/httpver.sh` to check HTTP versions of our endpoints
- delivered company-wide presentation about data breach and security
- obtained "AWS Certified Cloud Practitioner" certification
- review what AWS accounts are integrated with DD and security monitored (`sectools/dd-awsinfo`)
- found and reported leaked git secrets
    - documented and explained to engineers how to prevent it
    - helped engineers to reviewed commited secrets
- found and reported public S3 buckets (`sectools/aws-s3pub`)
- lowered AWS and Datadog costs (`sectools/aws-dbinfo`)
- reviewed Kubernetes clusters security

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
