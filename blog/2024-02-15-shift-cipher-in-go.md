I started to read <https://bitfieldconsulting.com/books/crypto>. I'll also try to write the code and take some notes in the form of blog posts, like this one.

A simple way to encipher (or encrypt) some data is by using the shift cipher. We do this by going through the data byte by byte adding a key to each of the bytes. It's possible because in Go bytes are equivalent to 8-bit numbers ranging from 0 to 255 (`byte` data type is an alias for `uint8`).

```
func Encipher(plaintext []byte, key byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	for i, b := range plaintext {
		ciphertext[i] = b + key
	}
	return ciphertext
}
```

To decipher we need to do the same but in reverse, i.e. we detract the key from each byte of the enciphered data.

```
func Decipher(ciphertext []byte, key byte) []byte {
	return Encipher(ciphertext, -key)
}
```

This way Alice and Bob can exchange data in somehow secure way. If Eve wants to learn what are they talking about she needs to know the encryption algorithm and the key. Let's say she finds out they are using the Caesar cipher so she just needs to crack the key. The standard way to do this is called brute forcing, i.e. trying out all possibilities; in our case all possible keys. She also needs to know some bytes from the beginning of the "plaintext" data; this we call a crib. 

```
func Crack(ciphertext, crib []byte) (key byte, err error) {
	for guess := 0; guess < 256; guess++ {
		result := Decipher(ciphertext[:len(crib)], byte(guess))
		if bytes.Equal(result, crib) {
			return byte(guess), nil
		}
	}
	return 0, errors.New("no key found")
}
```

If we call these functions from within commands (`package main`) it looks like this:

```
$ echo hello world | go run ./cmd/encipher -key 10 | go run ./cmd/decipher -key 10                                                                          
hello world

$ echo hello world | go run ./cmd/encipher -key 10 | go run ./cmd/crack -crib hell                                                                          
hello world
```

See <https://github.com/jreisinger/pocs/tree/main/crypto/shift> for the whole code and more.
