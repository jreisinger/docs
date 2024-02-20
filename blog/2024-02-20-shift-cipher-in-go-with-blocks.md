In the previous post we implemented a shift cipher that accepts key of variable length. Now, the cryptographic algorithms can be classified into two groups. Stream ciphers and block ciphers. Stream ciphers process data a bit at a time and are good for continuous or unknown amounts of data like in networking. Block ciphers operate with fixed-length blocks of data and are suitable for handling variable sized data. What group does our current implementation belong to? Well, it works on a byte at a time so it sounds like a stream cipher. On the other hand, it could be regarded as a block cipher where the block size is one byte. In practice though it always enciphers or deciphers the whole message so the block size equals the message size.

And this is a problem. How come? Well, if the message is not large, it's fine. But if it's bigger, say 10GB, we might run out of memory because we read all of it at once

```
message, err := io.ReadAll(os.Stdin)
```

and we pass all the bytes we read to Encipher or Decipher function. But these functions only need to work on one byte at a time!

In order to turn our existing code into a practical block cipher, we don't need to change the cipher scheme, as such. We just need to make it work with chunks, or blocks, of data. For this, there's a special interface in the standard library's package `crypto/cipher`:

```
type Block interface {
    BlockSize() int
    Encrypt(dst, src []byte)
    Decrypt(dst, src []byte)
}
```

Standard library interfaces (other famous ones are io.Reader and io.Writer) define standardized "connectors" that allow plugging together different parts of code. So let's implement it, i.e. let's create a type with the methods defined in the interface:


```
// Cipher implements crypto/cipher.Block interface.
type Cipher struct {
    key [BlockSize]byte
}

func NewCipher(key []byte) (cipher.Block, error) {
    if len(key) != BlockSize {
        return nil, fmt.Errorf("%w %d (must be %d)", ErrKeySize, len(key), BlockSize)
    }
    return &Cipher{
        key: [BlockSize]byte(key),
    }, nil
}

func (c *Cipher) Encrypt(dst, src []byte) {
    for i, b := range src {
        dst[i] = b + c.key[i]
    }
}

func (c *Cipher) Decrypt(dst, src []byte) {
    for i, b := range src {
        dst[i] = b - c.key[i]
    }
}

func (c *Cipher) BlockSize() int {
    return BlockSize
}
```

Fine but if you look at the signatures of the Encrypt and Decrypt functions they still take all the data as input. Maybe we have a little look at the documentation:

```
$ go doc crypto/cipher Block
<...>
    It provides the capability to encrypt or decrypt individual blocks. The mode
    implementations extend that capability to streams of blocks.
```

Aha, we need some other code, called mode, that will chop data for the Encrypt and Decrypt functions into chunks:

```
type BlockMode interface {                                                                                                                                  
    BlockSize() int                                                                                                                                     
    CryptBlocks(dst, src []byte)                                                                                                                        
}                                                                                                                                                           
```

Let's implement also this interface (I show here only the code for encrypting):

```
type Encrypter struct {
    cipher    cipher.Block
    blockSize int
}

func NewEncrypter(c cipher.Block) Encrypter {
    return Encrypter{
        cipher:    c,
        blockSize: c.BlockSize(),
    }
}

func (e Encrypter) CryptBlocks(dst, src []byte) {
    if len(src)%e.blockSize != 0 {
        panic("encrypter: input not full blocks")
    }
    if len(dst) < len(src) {
        panic("encrypter: output smaller than input")
    }
    // Keep chopping block-sized pieces off the plaintext
    // and enciphering them until there are no more pieces.
    for len(src) > 0 {
        e.cipher.Encrypt(dst[:e.blockSize], src[:e.blockSize])
        dst = dst[e.blockSize:]
        src = src[e.blockSize:]
    }

}

func (e Encrypter) BlockSize() int {
    return e.blockSize
}
```

This is awesome and there's one major problem. The CryptBlocks function will panic if the plaintext is not aligned with the blockSize (32 bytes in our case). It means we can work only with messages whose length in bytes is a multiple of 32. Interesting but a bit limiting. We improve on this situation by padding all messages to be multiples of 32. We define the padding scheme like this. Both the number and the value of padded bytes is equal to the difference from the nearest multiple of block size. If the message size is aligned with the block size, the number and the value of padded bytes is equal to the block size. And here's the code:

```
func Pad(data []byte, blockSize int) []byte {
    n := blockSize - len(data)%blockSize
    padding := bytes.Repeat([]byte{byte(n)}, n)
    return append(data, padding...)
}

func Unpad(data []byte, blockSize int) []byte {
    n := int(data[len(data)-1])
    return data[:len(data)-n]
}
```

Finally we have all the necessary code to create commands that can encrypt and decrypt arbitrary data:

```
$ export KEY=0101010101010101010101010101010101010101010101010101010101010101
$ go run ./cmd/encipher/ -key $KEY < ../shift/testdata/tiger.txt | go run ./cmd/decipher/ -key $KEY 
The tiger appears at its own pleasure. When we become very silent at that
place, with no expectation of the tiger, that is when he chooses to appear...
When we stand at the edge of the river waiting for the tiger, it seems that the
silence takes on a quality of its own. The mind comes to a stop. In the Indian
tradition that is the moment when the teacher says, “You are that. You are that
silence. You are that.”
--Francis Lucille, “The Perfume of Silence”
```

See <https://github.com/jreisinger/pocs/tree/main/crypto/shift-block> for the full code.
