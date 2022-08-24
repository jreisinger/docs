Common commands (prefix all these with `:`)

```
set tabstop=4 shiftwidth=4 expandtab
retab # apply tabs -> spaces

syntax off

# enable spell checking in local buffer only
setlocal spell spelllang=en_us
```

Buffers

* list: `:ls[!]`
* add: `:n <filename>`
* delete: `:bd[<num>]`
* switch: `:b<num>`
* switch to previous: `<C-^>`

[vim-go](https://github.com/fatih/vim-go/) plugin

* install/update: `git clone https://github.com/fatih/vim-go.git ~/.vim/pack/plugins/start/vim-go`
* completion: `<C-x-o>`
* docs: `<K>`
