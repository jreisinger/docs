Common commands (prefix all these with `:`)

```
set tabstop=4 shiftwidth=4 expandtab
retab # apply tabs -> spaces

set nofoldenable

set textwidth=0 colorcolumn=0

syntax off

# enable spell checking in local buffer only
setlocal spell spelllang=en_us
```

Buffers

* list: `:ls`
* switch: `:b<num>` 
* switch to previous: `<C-^>` or `<C-6>`
* simplified: `:nnoremap <F5> :buffers<CR>:buffer<Space>`
