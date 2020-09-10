Share data via SSH protocol.

Linux server

```
mkdir /data/sshfs
```

Linux client

```
sudo apt install sshfs
mkdir $HOME/sshfs
sshfs <user>@<server.net>:/data/sshfs $HOME/sshfs
```
