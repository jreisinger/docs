Common commands (prefix all these with `:`)

```
# whitespace
set list            # show whitespace characters
set tabstop=4       # how many spaces a tab is worth
set shiftwidth=4    # how many spaces an indentation is worth
set expandtab       # use spaces
set noexpandtab     # use tabs
retab               # change tabs to spaces
retab!              # change spaces to tabs

# syntax
syntax off                      # don't highlight
setlocal spell spelllang=en_us  # spell checking in local buffer only

# buffers
ls[!]           # list
n <filename>    # add
bn              # next
bp              # previous
b<num>          # switch
bd[<num>]       # delete
```

[vim-go](https://github.com/fatih/vim-go/) plugin

* install/update: `git clone https://github.com/fatih/vim-go.git ~/.vim/pack/plugins/start/vim-go`
* completion: `<C-x-o>`
* docs: `<K>`
