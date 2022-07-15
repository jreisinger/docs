* uses different key for encryption and decryption
* key pair solves the key distribution problem

<img width="433" alt="image" src="https://user-images.githubusercontent.com/1047259/179235563-06b3e4ec-fb34-4967-8487-a71639cd10ce.png">

```
# generate private key
openssl genpkey -algorithm RSA -out private_key.pem \
-pkeyopt rsa_keygen_bits:3072 # RSA key should be at least 2048 bits

# extract public key from private key
openssl rsa -pubout -in private_key.pem -out public_key.pem

# set permissions
chmod 600 private_key.pem
chmod 644 public_key.pem
```

* digital signature provides for nonrepudiation

<img width="428" alt="image" src="https://user-images.githubusercontent.com/1047259/179238518-81fd0115-17dc-4ae8-a7db-a40f409595a7.png">

<img width="390" alt="image" src="https://user-images.githubusercontent.com/1047259/179238799-4df6fc39-1c72-4ba9-80ac-2d2d5a7d1f1f.png">

* elliptic-curve key pairs are substantially more efficient at signing data and verifying signatures

Source: Full Stack Python Security (2021)
