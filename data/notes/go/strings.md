## What is a string

String is a read-only **slice of bytes**. A string can hold arbitrary bytes not just UTF-8 text or any other predefined format. Here is a string literal that uses the `\xNN` notation (hex values of a byte range from `00` to `FF`):

```go
const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
```

## Printing strings

```go
// Print the string directly.
fmt.Print(sample)                   // ��=� ⌘

// Get individual bytes by looping over the string.
for i := 0; i < len(sample); i++ {  // bd b2 3d bc 20 e2 8c 98
    fmt.Printf("%x ", sample[i])
}

// Print bytes in hex (same output as the byte-by-byte loop above).
fmt.Printf("% x\n", sample)         // bd b2 3d bc 20 e2 8c 98

// Escape any non-printable byte sequences ..
fmt.Printf("% q\n", sample)         // "\xbd\xb2=\xbc ⌘"

// .. escape also non-ASCII bytes while intepreting UTF-8.
fmt.Printf("%+q\n", sample)         // "\xbd\xb2=\xbc \u2318"
```

## UTF-8 and string literals

We can also create a "raw string" that can contain only literal text (regular string - created with double quotes - can contain escape sequences as shown above):

```go
const placeOfInterest = `⌘`

fmt.Printf("plain string: ")
fmt.Printf("%s\n", placeOfInterest)

fmt.Printf("quoted string: ")
fmt.Printf("%+q\n", placeOfInterest)

fmt.Printf("hex bytes: ")
fmt.Printf("% x\n", placeOfInterest)

// plain string: ⌘
// quoted string: "\u2318"
// hex bytes: e2 8c 98
```

This means that the Unicode character value U+2318, the "Place of Interest" symbol [⌘](http://unicode.org/cldr/utility/character.jsp?a=2318), is represented by the bytes e2 8c 98, and that those bytes are the UTF-8 encoding of the hexadecimal value 2318.

Source code in Go is *defined* to be UTF-8 text. That implies that the text editor places the UTF-8 encoding of the symbol ⌘ into the source code file (`0a` is line feed [control character](https://en.wikipedia.org/wiki/Control_character)):

```
$ cat a.go
⌘
$ hexdump a.go
0000000 e2 8c 98 0a
0000004
```

## Range loops

`for range` loop on a string treats it specially. It decodes one UTF-8-encoded rune (code point) on each iteration:

```go
const nihongo = "日本語" // Japanese
for index, runeValue := range nihongo {
    // %#U shows the code point's Unicode value and its printed representation.
    fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
}
// U+65E5 '日' starts at byte position 0
// U+672C '本' starts at byte position 3
// U+8A9E '語' starts at byte position 6
```

## More

* https://blog.golang.org/strings
