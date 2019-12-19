JavaScript Object Notation is a popular text based format for exchanging data through APIs. It consists of these data types:

* objects (aka dictionary or hash) `{}`
* arrays `[]`
* numbers
* strings
* booleans
* "null"

# jq

Command line options:

`-r` (`--raw-output`) -- emit raw (unquoted) strings as output

[Basic filters](https://stedolan.github.io/jq/manual/#Basicfilters):

```
'.'         # pretty print everything

'.foo'      # value at key foo
'.foo.bar'

'.[]'       # all elems of an array
```

You can join filters using `|`:

```
jq '.data.result | .[] | .values | .[] | .[]'
```
