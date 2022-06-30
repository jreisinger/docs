NOTE: gpg key pairs are associated with identities not machines (like SSH keys). So you should have one personal and one work gpg key.

```
# list imported keys
gpg --list-keys --keyid-format=long

# generate key pair
gpg --full-generate-key

# export public key (e.g. for pasting into GitHub, GitLab)
gpg --armor --export KEY

# transfer the keys to new machine (Linux, Mac)
mv ~/.gnupg ~/.gnupg.old          # on the new machine
scp -rp OTHERMACHINE:~/.gnupg ~/  # on the new machine
```

```
# 1) import a public key
gpg --import john_doe_gpg.asc

# 2) verify and sign the public key (the sender should supply a fingerprint to compare)
gpg --fingerprint john.doe@example.org
gpg --sign-key john.doe@example.org

# 3) encrypt a file with the public key
gpg --encrypt --sign --armor --recipient john.doe@example.org my-secret.conf
```

```
# decrypt a file (with your private key)
gpg --decrypt their-secret.conf.asc > their-secret.conf
```
