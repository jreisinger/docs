## Commands

NOTE - Prefix all these with `:`

Common settings:

```
set tabstop=4 shiftwidth=4 expandtab
retab # apply tabs -> spaces

set nofoldenable

set textwidth=0
set colorcolumn=0

syntax off
setlocal spell spelllang=en_us # enable spell checking in local buffer only
```

How to search:

```
grep -iR what .
cw
```

```
vimgrep /PATTERN/ FILE(s)
cnext
cprex
cfnext
cfprev
```

Know your alphabet:

```
help normal-index
help insert-index
```

## Keyboard shortcuts

Scroll inside `:term`:

```
Ctrl-w N
```

Jumping to previously visited locations ([link](https://vim.fandom.com/wiki/Jumping_to_previously_visited_locations)):

```
Ctrl-O  # back
Ctrl-I  # forth
```
