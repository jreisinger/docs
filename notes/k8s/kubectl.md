Delete pod immediately

    export now="--force --grace-period=0"
    k delete pod nginx $now

Explain manifest fields

    k explain pod.spec.containers.ports [--recursive]

Run a command inside a cluster

    k run busybox --image=busybox --rm -it --restart=Never --command -- wget -qO- example.com --timeout 2

Run a shell inside a cluster

    k run alpine --image=alpine --rm -it --restart=Never --command -- /bin/sh
    / # apk --update add bind-tools curl

Copy files

    k cp <pod>:/path/to/remote/file /path/to/local/file # or vice versa

Forward a port from localhost to cluster

    k port-forward <pod> 8080:8080 # tunnel: localhost -> k8s master -> k8s worker node
