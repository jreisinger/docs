Some of the [quotes](https://quotes.reisinge.net/) I like best ATM :-)

> The programmer, like the poet, works only sligtly removed from the pure thought-stuff. Yet the program construct, unlike the poet's words, is real in the sense that it moves and works, producing visible outputs separate from the construct itself. -- Frederick P. Brooks

> Money is like gas in the car - you need to pay attention or you'll end up on the side of the road - but a successful business or a well-lived life is not a tour of gas stations. -- Tim O'Reilly, WTF

---

# Why is it useful

I'm using shell (bash specifically) on daily basis. From time to time a need arised to run multiple commands in parallel. For example my [.bashrc](https://github.com/jreisinger/dotfiles/blob/master/.bashrc) runs the following commands to download or clone vim plugins I use:

```
curl -L -o $HOME/.git-completion.bash https://raw.githubusercontent.com/git/git/master/contrib/completion/git-completion.bash
#curl -L -o $HOME/.vim/autoload/pathogen.vim https://raw.github.com/tpope/vim-pathogen/master/autoload/pathogen.vim
rm -rf $HOME/.vim/pack/plugins/start/nerdtree && git clone https://github.com/scrooloose/nerdtree.git $HOME/.vim/pack/plugins/start/nerdtree
rm -rf $HOME/.vim/pack/plugins/start/vim-nerdtree-tabs && git clone https://github.com/jistr/vim-nerdtree-tabs.git $HOME/.vim/pack/plugins/start/vim-nerdtree-tabs
rm -rf $HOME/.vim/pack/plugins/start/vim-markdown && git clone https://github.com/plasticboy/vim-markdown.git $HOME/.vim/pack/plugins/start/vim-markdown
rm -rf $HOME/.vim/pack/plugins/start/bufexplorer && git clone https://github.com/jlanzarotta/bufexplorer.git $HOME/.vim/pack/plugins/start/bufexplorer
rm -rf $HOME/.vim/pack/plugins/start/ansible-vim && git clone https://github.com/pearofducks/ansible-vim.git $HOME/.vim/pack/plugins/start/ansible-vim
rm -rf $HOME/.vim/pack/plugins/start/vim-go && git clone https://github.com/fatih/vim-go.git $HOME/.vim/pack/plugins/start/vim-go
rm -rf $HOME/.vim/pack/plugins/start/grep.vim && git clone https://github.com/yegappan/grep.git $HOME/.vim/pack/plugins/start/grep.vim
# https://tpaschalis.github.io/vim-go-setup/
rm -rf $HOME/.vim/pack/plugins/start/vim-airline && git clone https://github.com/vim-airline/vim-airline $HOME/.vim/pack/plugins/start/vim-airline
```

The problem is that these commmands run sequentially and it takes a while until they are done. I was thinking of a way how to speed them up. So to scratch my itch a came up with [runp](https://github.com/jreisinger/runp). Now I can the those commands (I store them in `install-my-stuff.txt`) in parallel:

```
$ runp install-my-stuff.txt 
--> OK (0.30s): /bin/sh -c "curl -L -o /home/reisinge/.git-completion.bash https://raw.githubusercontent.com/git/git/master/contrib/completion/git-completion.bash"
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 70071  100 70071    0     0   247k      0 --:--:-- --:--:-- --:--:--  247k
--> OK (1.76s): /bin/sh -c "rm -rf /home/reisinge/.vim/pack/plugins/start/ansible-vim && git clone https://github.com/pearofducks/ansible-vim.git /home/reisinge/.vim/pack/plugins/start/ansible-vim"
Cloning into '/home/reisinge/.vim/pack/plugins/start/ansible-vim'...
--> OK (1.76s): /bin/sh -c "rm -rf /home/reisinge/.vim/pack/plugins/start/grep.vim && git clone https://github.com/yegappan/grep.git /home/reisinge/.vim/pack/plugins/start/grep.vim"
Cloning into '/home/reisinge/.vim/pack/plugins/start/grep.vim'...
--> OK (1.86s): /bin/sh -c "rm -rf /home/reisinge/.vim/pack/plugins/start/bufexplorer && git clone https://github.com/jlanzarotta/bufexplorer.git /home/reisinge/.vim/pack/plugins/start/bufexplorer"
Cloning into '/home/reisinge/.vim/pack/plugins/start/bufexplorer'...
--> OK (1.87s): /bin/sh -c "rm -rf /home/reisinge/.vim/pack/plugins/start/vim-nerdtree-tabs && git clone https://github.com/jistr/vim-nerdtree-tabs.git /home/reisinge/.vim/pack/plugins/start/vim-nerdtree-tabs"
Cloning into '/home/reisinge/.vim/pack/plugins/start/vim-nerdtree-tabs'...
--> OK (2.23s): /bin/sh -c "rm -rf /home/reisinge/.vim/pack/plugins/start/vim-markdown && git clone https://github.com/plasticboy/vim-markdown.git /home/reisinge/.vim/pack/plugins/start/vim-markdown"
Cloning into '/home/reisinge/.vim/pack/plugins/start/vim-markdown'...
--> OK (2.99s): /bin/sh -c "rm -rf /home/reisinge/.vim/pack/plugins/start/nerdtree && git clone https://github.com/scrooloose/nerdtree.git /home/reisinge/.vim/pack/plugins/start/nerdtree"
Cloning into '/home/reisinge/.vim/pack/plugins/start/nerdtree'...
--> OK (3.08s): /bin/sh -c "rm -rf /home/reisinge/.vim/pack/plugins/start/vim-airline && git clone https://github.com/vim-airline/vim-airline /home/reisinge/.vim/pack/plugins/start/vim-airline"
Cloning into '/home/reisinge/.vim/pack/plugins/start/vim-airline'...
--> OK (3.82s): /bin/sh -c "rm -rf /home/reisinge/.vim/pack/plugins/start/vim-go && git clone https://github.com/fatih/vim-go.git /home/reisinge/.vim/pack/plugins/start/vim-go"
Cloning into '/home/reisinge/.vim/pack/plugins/start/vim-go'...
```

Let's see how much time did I save:

```
$ time bash install-my-stuff.txt
<...>
real	0m15.690s
user	0m3.440s
sys	0m0.902s

$ time runp install-my-stuff.txt
<...>
real	0m3.678s
user	0m3.904s
sys	0m0.880s
```

Hmm, aroud 12 seconds. Not bad I think :-).

# A picture like [this](https://kapow.readthedocs.io/en/latest/)?

# Simple example or two

# How to install it

# More complex examples
