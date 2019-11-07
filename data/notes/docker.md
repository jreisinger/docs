# Docker

Install basic packages inside a container:

```
apt-get update && apt-get install procps net-tools vim
```

Clean up host running `dockerd`:

```bash
# you will be asked to confirm
docker system prune               # containers, images
docker images prune               # only images
```

[More](https://github.com/jreisinger/notes/blob/master/content/posts/docker.md).
