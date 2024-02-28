## JSON

JavaScript Object Notation (JSON) is a text based format for exchanging data 
(through APIs) between programs written in various languages.

Networked programs need to exchange information via messages. TCP and UDP
provide a transport mechanism to do this. However, at transport level the
messages are just sequences of bytes with no structure.

A program will typically build a complex data structure to hold the current
program state. To transfer this data outside of the program's own address space
(e.g. to another application over the network) it needs to be serialized. This
process is also called marshalling or encoding.

A JSON value MUST be an

- object `{}`
- array `[]`
- number
- string

or one of the literal names: 

- `true`
- `false`
- `null`

The literal names MUST be lowercase. No other literal names are allowed. [RFC
7159](https://www.rfc-editor.org/rfc/rfc7159.html#section-3)

## jq

JSON command line processor 

Sample JSON file:

```
$ cat animals.json
[
  {
    "animal": "Camel, aka Dromedary",
    "cover_src": "https://covers.oreilly.com/images/9780596004927/cat.gif",
    "link": "https://shop.oreilly.com/product/9780596004927.do",
    "title": "Programming Perl"
  },
  {
    "animal": "Botta's Pocket Gopher",
    "cover_src": "https://covers.oreilly.com/images/0636920046516/cat.gif",
    "link": "https://shop.oreilly.com/product/0636920046516.do",
    "title": "Introducing Go"
  }
]
```

Basic filters:

```
'.'      # pretty print everything
'.[]'    # all elems of an array
'.foo'   # value at key foo
```

Emit raw (**unquoted**) strings with `-r`:

```
$ jq -r '.[].animal' < animals.json
Camel, aka Dromedary
Botta's Pocket Gopher
```

**Join filters** using `|` and get values of **multiple keys**:

```
$ jq -r '.[] | "\(.title) => \(.animal)"' < animals.json
Programming Perl => Camel, aka Dromedary
Introducing Go => Botta's Pocket Gopher
```

Find objects with a **string match** against the title:

```
$ jq -r '.[] | select(.title=="Introducing Go")' < animals.json
{
  "animal": "Botta's Pocket Gopher",
  "cover_src": "https://covers.oreilly.com/images/0636920046516/cat.gif",
  "link": "https://shop.oreilly.com/product/0636920046516.do",
  "title": "Introducing Go"
}
```

Find an animal with a **regex match** against the title:

```
$ jq -r '.[] | select(.title|test("[Pp]erl")) | .animal' < animals.json
Camel, aka Dromedary
```

More

* [jq docs](https://stedolan.github.io/jq/manual/)
* [jq playground](https://jqplay.org/)
