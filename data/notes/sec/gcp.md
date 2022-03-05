Like all major cloud providers, Google Cloud Platform (GCP) implents the shared responsibility model. GCP is secures its infra, users their cloud resources, workloads and data.

# GCP security

## Secure Service Deployment

*Crypto authn and authz* applied at the app level to all inter-service communication. This feature provides granular access control.

*Service account identity* associated with any service that runs on Google cloud infra. The service mus use it's crypto credentials to receive or make RPC call to other services, or identify itself to clients.

*Segmentation and firewall* - ingress and egress filtering at important network junctions, to prevent IP spoofing.

## Safeguards From Privileged Access Attacks

## Data Disposal Features

## Encryption of Data

## Secure Internet Communication

## Operational Security

# Users security (via GCP tools)

## Google Cloud KMS

## Google Cloud IAM

## Google Cloud Identity

## Stackdriver Logging

## Google Access Transparency

## Google Cloud Security Scanner

## Google Cloud Resource Manager

## Google Cloud Compliance

Source: https://www.aquasec.com/cloud-native-academy/cspm/google-cloud-security/