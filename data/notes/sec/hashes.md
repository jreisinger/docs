*2018-08-10*

There are several classes of hashes (hash functions) used for different purposes:

* hash datatype in Perl (maps a lookup string to a memory location; fast and short)
* networking protocols checks like CRC or Adler (rarely go over 32-bits in length)
* cryptographic hashes (this article is about them)

Cryptographic hashes

* provide an encryption using an algorithm and *no key*
* a variable length plaintext is "hashed" into a fixed-length hash value (also called a *message digest* or a *hash*)
* primarily used to provide integrity --> if the hash changed, the plaintext must have changed
* have an "avalanche effect" --> changing just one bit of input creates a completely different output

Collisions

* hashes are not unique - number of possible plaintexts is far larger than the number of possible hashes
* searching for collision to match a specific text should not be possible accomplish in a reasonable amount of time

Types

* MD5 - 128-bit, considered weak since late 90s
* SHA1 - 160-bit, developed by the NSA, [broken](https://www.schneier.com/blog/archives/2005/02/sha1_broken.html) in 2005, [completely broken](https://sha-mbles.github.io/) in 2020
* SHA2 - group of algorithms (224 to 512-bits), kind of secure
* SHA3 - Keccak won a [competition](https://csrc.nist.gov/projects/hash-functions/sha-3-project) announced in 2003, **recommended**

Perl

``` sh
$ perl -MDigest::SHA3=sha3_512_hex -E 'say sha3_512_hex( "plaintext" )'
```

Go

``` go
package main

import (
	"fmt"

	"golang.org/x/crypto/sha3"
)

func main() {
	hash := sha3.New512()
	hash.Write([]byte("plaintext"))
	fmt.Printf("%x\n", hash.Sum(nil))
}
```

Sources:

* http://www.wumpus-cave.net/2014/03/27/perl-encryption-primer-hashes
* http://wiki.reisinge.net/CISSP/03_Cryptography/HashFunctions
