**jq** - command line JSON processor 

* JSON - JavaScript Object Notation
* JSON - text based format for exchanging data through APIs

JSON consists of these data types:

* objects (dictionaries) `{}`
* arrays `[]`
* numbers, strings, booleans, "null"

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

Command line options:

```
-r (--raw-output) -- emit raw (unquoted) strings as output
```

Basic filters:

```
'.'         # pretty print everything

'.foo'      # value at key foo
'.foo.bar'  # value at key foo.bar

'.[]'       # all elems of an array
```

You can **join filters** using `|`:

```
jq '.data.result | .[] | .values | .[] | .[]'
```

Get values of **multiple keys**:

```
jq '.rrsets[] | "\(.name) \(.type)"'
```

Find an animal with a **regex match** against the title:

```
jq '.[] | select(.title|test("Perl")) | .animal' < animals.json
```

# More

* [jq docs](https://stedolan.github.io/jq/manual/)
* [jq playground](https://jqplay.org/)
