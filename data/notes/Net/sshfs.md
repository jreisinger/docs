Share data via SSH protocol.

# Linux server

```
mkdir /data/sshfs
```

# Linux client

Installation

```
sudo apt install sshfs

mkdir sshfs
```

Manual mount

```
sshfs <user>@<server.net>:/data/sshfs sshfs
```

Auto mount

```
# /etc/fstab - permanent mount
sshfs#<user>@<server.net>:/data/sshfs <absolute-path>/sshfs fuse user,defaults,allow_other 0 0

# /etc/fuse.conf - Allow non-root users to specify the allow_other or allow_root mount options.
user_allow_other
```

# macOS client

Installation

```
brew cask install osxfuse
brew install sshfs

mkdir sshfs
```

Manual mount

```
sshfs <user>@<server.net>:/data/sshfs sshfs
```
