Q: How to exit a Python script?

A: Use `sys.exit()` in scripts, or `raise SystemExit()` if you prefer. -- https://stackoverflow.com/a/19747557

There are some predifined constants that can be used with [os._exit()](https://docs.python.org/3/library/os.html#os._exit). To get their numeric values:

```
for i in dir(os):
    if i.startswith('EX_'):
        print(f"{i} = {os.__getattribute__(i)}")
```
