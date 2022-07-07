* encryption ensures confidentiality, the guarantee of privacy
* the encryption/decryption algorithm is called a cipher

![image](https://user-images.githubusercontent.com/1047259/177801568-f230e8b8-7858-482e-8291-a9b05a25c0f9.png)

* in Python, use the [cryptography](https://cryptography.io) package

```
# Python package management (alternative to pip and venv)
sudo apt install pipenv
cd /tmp/myproj
pipenv install
pipenv shell
pipenv install cryptography
```

* `Fernet` is a safe and easy way to symmetrically encrypt and authenticate data

```
>>> from cryptography.fernet import Fernet
>>> key = Fernet.generate_key()
>>> fernet = Fernet(key)
>>> token = fernet.encrypt(b'plaintext')
>>> fernet.decrypt(token)
b'plaintext'
```

* `Fernet` doesn't just encrypt the plaintext; it hashes the ciphertext as well

![image](https://user-images.githubusercontent.com/1047259/177803130-404c4f02-7297-4a0a-aa74-57b782befffd.png)

* a `Fernet` object can decrypt any `Fernet` token created by a `Fernet` object with the same key
* you can throw away an instance of `Fernet`, but the key must be saved and protected

# Block ciphers

* encrypt plaintext as a series of fixed-length blocks
* each block of plaintext is encrypted to a block of ciphertext
* the block size depends on the encryption algorithms
* larger block sizes are generally more secure

Triple DES (3DES)

* don't use it
* 64-bit block size, key size of 56, 112, or 168 bits
* slow (runs DES 3 times) and deprecated by NIST and OpenSSL

Blowfish

* don't use it
* invented by Bruce Schneier in early 1990s
* 64-bit block size, variable key size of 32 to 448 bits
* found vulnerable to SWEET32 attack in 2016

Twofish

* don't use it (not that popular)
* developed in late 1990s as successor to blowfish
* 128-bit block size, key size of 128, 192, or 256 bits

**AES**

* Rijndael standardize byt NIST in 2001
* 128-bit block size, key size of 128, 192, or 256 bits
* used by `Fernet`, HTTPS, compressions, filesystems, hashing, VPNs

# Stream ciphers

* plaintext is processed as a stream of individual bytes
* good at encrypting continuous or uknown amounts of data
* ofter used by networking protocols

RC4

* don't use it
* half dozen vulnerabilities

**ChaCha**

* used by TLS

# Encryption modes

Application developers tipically discuss which encryption mode to run AES in.

ECB

* don't use it in production
* exceptionally week
* reveals patterns within plaintext and between plaintexts

**CBC**

* produces different ciphertexts when encrypting identical plaintexts with the same key
* this is done by individualizing plaintext with initialization vector (IV)
* IV is not secret and is typically kept alongside the ciphertext (need also for decryption)
* `Fernet` encrypts data with AES in CBC mode

# Source

* Full Stack Python Security (2021)
