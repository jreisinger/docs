In the previous blog post we developed a simple crypto system. It's algorithm is based on shifting bytes by the number represented by a single byte we call a key. It means Eve has to do at maximum 256 (1 byte is 8 bits and that means 2^8 possibilities) guesses to find out the key. Let's try to improve the situation here by supporting longer keys.

To encipher a message we go through it byte by byte and we also go byte by byte through the key. Usually the key is much shorter than the message we want to encrypt. So we need to go through the key multiple times. The math trick to do this is called modulo. A mod B is the remainder that's left after dividing A by B as many times as you can. E.g. 5 mod 2 = 1. Modular arithmetic is sometimes called "clock arithmentic" because it wraps around like an analog clock; 12 hours later than 5 o'clock can't be 17 o'clock, it's 5 o'clock again. To put it another way, 17 mod 12 = 5.

To illustrate how modulo (`%` in Go) works let's write a short program:

```
func main() {
        B := 3
        for A := range 10 {
                fmt.Printf("%d mod %d = ", A, B)
                fmt.Println(A % B)
        }
}

```

The program will produce this output - notice the result is never greater than 2 which is handy for a slice (or array) index:

```
0 mod 3 = 0
1 mod 3 = 1
2 mod 3 = 2
3 mod 3 = 0
4 mod 3 = 1
5 mod 3 = 2
6 mod 3 = 0
7 mod 3 = 1
8 mod 3 = 2
9 mod 3 = 0
```

OK, let's use modulo operation to encrypt a message:


```
func Encipher(plaintext []byte, key []byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	for i, b := range plaintext {
		ciphertext[i] = b + key[i%len(key)]
	}
	return ciphertext
}
```

To decrypt a message encrypted by a multi-byte key we do the same in reverse:

```
func Decipher(ciphertext []byte, key []byte) []byte {
	plaintext := make([]byte, len(ciphertext))
	for i, b := range ciphertext {
		plaintext[i] = b - key[i%len(key)]
	}
	return plaintext
}
```

Now, how do we pass a multi-byte key as a command line argument? The key, in some sense, is just a single number, no matter how many bytes it takes to express it. For example, if we had a 32-byte (that is, 256-bit) key, we could express it as either a series of 32 integers (one for each byte), or as a single very large integer. But Go's `int64` can hold only 8 bytes (or 64 bits) worth of information... There's a neat and concise way to write large integers: as a string, using hexadecimal notation. For example the decimal number 3 735 928 559 can be represented as DEADBEEF (4 bytes) in hex, isn't that funny? :-) If fact, any given byte can be written as exactly two hex digits, which is convenient.

```
$ echo hello | go run ./cmd/encipher -key DEADBEEF
F*[M�
```

Also notice that unlike with the single-byte version, the same plaintext letter does not always produce the same ciphertext letter. The "ll" is enciphered as "[M". This makes the frequency analysis a lot harder for Eve.

But what troubles Eve even more is that her function for brute-forcing single key shift ciphers won't work anymore:

```
func Crack(ciphertext, crib []byte) (key byte, err error) {
	for guess := 0; guess < 256; guess++ {
		plaintext := Decipher(ciphertext[:len(crib)], byte(guess))
		if bytes.Equal(plaintext, crib) {
			return byte(guess), nil
		}
	}
	return 0, errors.New("no key found")
}
```

She has to solve couple of issues:

- Repeat the guessing of the key byte multiple times. The number of repetitions will be either the length of the encrypted message or some value defined by us (`MaxKeyLen`); whatever is smaller.
- In order to use the Decipher function she needs to create byte slice out of a byte to match the function's arguments type.
- She has to check the whole key is correct after each cracked key byte.

```
const MaxKeyLen = 32 // bytes

func Crack(ciphertext, crib []byte) (key []byte, err error) {
	for k := range min(MaxKeyLen, len(ciphertext)) {
		for guess := range 256 {
			plaintext := Decipher([]byte{ciphertext[k]}, []byte{byte(guess)})
			if plaintext[0] == crib[k] {
				key = append(key, byte(guess))
				break
			}
		}
		if bytes.Equal(Decipher(ciphertext[:len(crib)], key), crib) {
			return key, nil
		}
	}
	return nil, errors.New("no key found")
}
```

The longer key is harder to brute-force but it's still possible:

```
$ go run ./cmd/encipher -key DEADBEEF < ../shift/testdata/tiger.txt | go run ./cmd/crack -crib 'The tiger'
The tiger appears at its own pleasure. When we become very silent at that
place, with no expectation of the tiger, that is when he chooses to appear...
When we stand at the edge of the river waiting for the tiger, it seems that the
silence takes on a quality of its own. The mind comes to a stop. In the Indian
tradition that is the moment when the teacher says, “You are that. You are that
silence. You are that.”
--Francis Lucille, “The Perfume of Silence”
```

However, there is a limitation (or a bug): the crib must be at least as long as the key; it this case 4 bytes, i.e. 'The '.

See <https://github.com/jreisinger/pocs/tree/main/crypto/shift-multibytekey> for all the code including tests.
