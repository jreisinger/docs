- tool for reconciling desired state with an external system
- interacts with external systems via *providers*: https://registry.terraform.io/
- provider describes the schema of resources and implements CRUD operations
- providers can manage almost anything: local files, databases, cloud infrastructure, messaging services, ...

```sh
❯ cat main.tf 
resource "local_file" "name" {
    content = "foo!"
    filename = "${path.module}/foo.txt"
    file_permission = 0644
}
❯ terraform init && terraform apply
❯ tree
.
├── foo.txt           # managed resource(s)
├── main.tf           # desired state of resource
└── terraform.tfstate # current state of resource
```
