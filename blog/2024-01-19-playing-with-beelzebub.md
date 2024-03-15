![image](https://github.com/jreisinger/docs/assets/1047259/18b6f3e2-05d5-4612-8ba7-41b96f4deecc)

While looking for a new project to hone my skills I came across the beelzebub. Wikipedia says Beelzebub was a Philistine god and later a major demon for some Abrahamic religions. In this case it's a honeypot written in Go :-).

My plan was something like:

1. Create a Kubernetes cluster on AWS using the EKS service
1. Deploy the honeypot into the cluster
1. Setup logs collection to see what's going on
1. Expose the honeypot to a dangerous network, like the Internet, and wait

## Create a Kubernetes cluster

Once I have [set up](https://docs.aws.amazon.com/eks/latest/userguide/setting-up.html) my access to AWS and installed all the necessary tools, the easiest way to create a Kubernetes cluster seemed to be this:

```
eksctl create cluster --name beelzebub-cluster --region eu-central-1
```

It took about 15 minutes but went smoothly.

## Deploy the honeypot into the cluster

Next, I just cloned the [beelzebub](https://github.com/mariocandela/beelzebub/) repo and created the Kubernetes resources from within the repo:

```
helm install beelzebub ./beelzebub-chart
```

## Setup logs collection

Now, a Kubernetes cluster provides logs from several components:

- control plane
- nodes
- applications

I was most interested in the application (or container) logs. For this I used the CloudWatch observability addon for EKS. To make it work I needed to attach some new policies to the worker nodes role and then create the addon.

```
aws iam attach-role-policy \
    --role-name <my-worker-node-role> \
    --policy-arn arn:aws:iam::aws:policy/CloudWatchAgentServerPolicy  \ 
    --policy-arn arn:aws:iam::aws:policy/AWSXrayWriteOnlyAccess
aws eks create-addon --addon-name amazon-cloudwatch-observability --cluster-name beelzebub-cluster
```

## Expose the honeypot to the Internet

Then, I created a Kubernetes service of type LoadBalancer:

```
kubernetes expose deployment beelzebub-beelzebub-chart --name beelzebub-public --type LoadBalancer --port 22 --target-port 2222
```

I had to wait for a bit so the load balancer is set up. I opened the CloudWatch Logs Insights, selected the `/aws/containerinsights/beelzebub-cluster/application` log group and entered the following query:

```
filter kubernetes.pod_name="beelzebub-beelzebub-chart-b86c7dff8-59ldz"
| fields @timestamp, log_processed.event.Msg, log_processed.event.User, log_processed.event.Password, log_processed.event.Command, log_processed.event.CommandOutput
| sort @timestamp desc
| limit 20
```

To make sure everything is working, I logged into the honeypot and observed the logs (it takes a while until the logs get to the CloudWatch):

```
ssh root@<some-string>.elb.eu-central-1.amazonaws.com
# obviously, the default password is root :-)
```

## Clean up

Once I was done, I removed the cluster (and all of its workloads and related AWS services):

```
eksctl delete cluster --name beelzebub-cluster --region eu-central-1
```

You might also want to delete the related CloudWatch Log groups.
