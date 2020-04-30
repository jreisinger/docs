Find out whether RBAC is enabled on a cluster (one line for each control node):

```
> kubectl describe pod -n kube-system -l component=kube-apiserver | grep authorization
      --authorization-mode=Node,RBAC
      --authorization-mode=Node,RBAC
      --authorization-mode=Node,RBAC
```

Source:

* https://learning.oreilly.com/library/view/cloud-native-devops/9781492040750/ch11.html
