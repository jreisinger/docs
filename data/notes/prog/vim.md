## Commands

NOTE - prefix all these with `:`

Common settings:

```
set tabstop=4 shiftwidth=4 expandtab
retab # apply tabs -> spaces

set nofoldenable

set textwidth=0 colorcolumn=0

syntax off

# enable spell checking in local buffer only
setlocal spell spelllang=en_us
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
