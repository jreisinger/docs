Some easy ways to create and destroy an [EKS cluster](https://docs.aws.amazon.com/eks/latest/userguide/what-is-eks.html).

Using [eksctl](https://docs.aws.amazon.com/eks/latest/userguide/getting-started-eksctl.html):

```
eksctl create cluster --name my-cluster --region region-code

eksctl delete cluster --name my-cluster --region region-code
```

Using [terraform](https://learn.hashicorp.com/tutorials/terraform/eks):

```
git clone https://github.com/hashicorp/learn-terraform-provision-eks-cluster
cd learn-terraform-provision-eks-cluster/
vi variables.tf
terraform init
terraform apply
rm ~/.kube/config # otherwise you'll get error: 'NoneType' object is not iterable
aws eks --region $(terraform output -raw region) update-kubeconfig --name $(terraform output -raw cluster_name)

terraform destroy
```
