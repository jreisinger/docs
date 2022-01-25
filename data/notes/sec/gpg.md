```
# list imported keys
gpg --list-keys --keyid-format=long

# generate key pair
gpg --full-generate-key

# export public key (e.g. for pasting into GitHub, GitLab)
gpg --armor --export KEY

# transfer the keys to new machine (Linux, Mac)
mv ~/.gnupg ~/.gnupg.old
scp -rp OTHERMACHINE:~/.gnupg ~/

# verify and sign a key (the sender should supply a fingerprint to compare)
gpg --fingerprint john.doe@example.org
gpg --sign-key john.doe@example.org

# encrypt a file
gpg --encrypt --sign --armor -r john.doe@example.org secret.conf

# decrypt a file
gpg --decrypt secret.conf.asc > secret.conf
```
