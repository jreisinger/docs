# Keyed hashing

Hash functions are useful for ensuring data integrity - has the data changed?. But they won't help with data (or message) authentication - who authored this data?. For this you need:

* a key
* a keyed hash function

<img width="400" alt="image" src="https://user-images.githubusercontent.com/1047259/176613239-0a08bb23-e0a9-4cbd-9e17-28ed9939ad85.png">

For keys generation in Python use:

```python
>>> os.urandom(16)
b'\x08\xab\x0e\xe1\xc8\xeb\xefF`\x97\x8e\x91\x84\x98\x17\x89'

>>> secrets.token_bytes(16)
b">\xf9\xdf|\xec\x7f\xc1\xde\xc7'\xca\x96X_9\xe2"
>>> secrets.token_hex(16)
'a223e557df83dd609817ce2685f35129'
>>> secrets.token_urlsafe(16)
'4JPsW6AdzieK7DZWEIRcvA'
```

For keyed hasing in Python use:

```python
>>> from hashlib import blake2b
>>> m = b'message'
>>> x = b'key x'
>>> y = b'key y'
>>> blake2b(m, key=x).digest() == blake2b(m, key=x).digest()
True
>>> blake2b(m, key=x).digest() == blake2b(m, key=y).digest()
False
```

See [keyed_hashing.py](https://gist.github.com/jreisinger/e27f1f8c8c111225b40e9732ba6a704a) for a complete example.

# HMAC (Hash-based Message Authentication Code)

HMAC functions are a generic way to use any orginary hash function as though it were a keyed hash function.

<img width="389" alt="image" src="https://user-images.githubusercontent.com/1047259/176614275-c1ed5474-418e-4615-b486-fbc15b3b2de9.png">

```
>>> import hashlib, hmac
>>> hmac_sha256 = hmac.new(b'key', msg=b'message', digestmod=hashlib.sha256)
```

## Data authentication protocol

<img width="583" alt="image" src="https://user-images.githubusercontent.com/1047259/176614944-20f0977a-acb2-469c-b874-373a3f1fa1fa.png">

## Timing attacks

* a specific kind of side channel attacks
* a side channel attack derives unauthorized information by measuring a physical side channel (time, sound, power consumption, EM radiation, radio waves, heat)

Secure systems compare hash values in length-constant time, sacrificing a bit of performance to prevent timing attack vulnerabilities. In Python use `hmac.compare_digest()` instead of `==`.

# Sources

* Full Stack Python Security, ch. 3 (2021)
