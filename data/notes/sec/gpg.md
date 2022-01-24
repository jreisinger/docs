```
# list imported keys
gpg --list-keys

# verify and sign a key (the sender should supply a fingerprint to compare)
gpg --fingerprint john.doe@example.org
gpg --sign-key john.doe@example.org

# encrypt a file
gpg --encrypt --sign --armor -r john.doe@example.org secret.conf

# decrypt a file
gpg --decrypt secret.conf.asc > secret.conf

# export a key (e.g. for pasting into GitHub, GitLab)
gpg --armor --export <KEY>
```
