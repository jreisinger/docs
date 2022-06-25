Created: 2018-08-10

There are several classes of hashes (hash functions) used for different purposes

* hash datatype in Perl or dictionary (`hash` function) in Python maps a lookup string to a memory location (fast and short)
* networking protocols checksums like CRC or Adler (rarely go over 32-bits in length, used for error detection)
* cryptographic hashes (this article is about them)

Hash properties

* deterministic behavior - a given input always produces the same output
* fixed-length hash values (also called message digests) for arbitrary long messages
* avalanche effect - small diffs between messages produce large diffs between hash values

Cryptographic hash - additional properties

* one way function - easy to invoke, infeasable [1] to reverse engineer
* collision resistance - hash values for different messages must almost [1] never have the same value

[1] Reverting a hash value or searching for a collision should not be practically possible. This is a moving target and cryptographic strength weakens with time.

Cryptographic hash usage

* primarily used to provide data (message) integrity
* if the hash value changed, the plaintext must have changed

Safe algorithms

* SHA2 - family of algorithms (224 to 512-bits), use SHA-256 for general-purpose cryptographic hashing
* SHA3 - Keccak won a [competition](https://csrc.nist.gov/projects/hash-functions/sha-3-project) announced in 2003, use SHA3-256 in high-security environments
* BLAKE2 - not as popular but leverages modern CPU architecture to hash at extreme speeds -> good for large messages

Unsafe algorithms (should never be used for security purposes when creating a new system)

* MD5 - 128-bit, considered weak since late 90s
* SHA1 - 160-bit, developed by the NSA, [broken](https://www.schneier.com/blog/archives/2005/02/sha1_broken.html) in 2005, [completely broken](https://sha-mbles.github.io/) in 2020

Examples

```
$ python -c 'import hashlib; print(hashlib.sha256(b"hello").hexdigest())'
2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
$ ruby -e 'require "digest"; puts Digest::SHA256.hexdigest "hello"'
2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
$ perl -MDigest::SHA=sha256_hex -E 'say sha256_hex( "hello" )'
2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
```

* [Python](https://gist.github.com/jreisinger/833d02c2f544439e481b2e5ab5171baa)
* [Go](https://go.dev/play/p/tiT5N29hc4o)

Sources

* Full Stack Python Security (2021)
* http://www.wumpus-cave.net/2014/03/27/perl-encryption-primer-hashes
