*2014-01-06*

<img src="https://git-scm.com/images/logos/2color-lightbg@2x.png" style="max-width:100%;height:auto;float:right">

Although I'm more of a sysadmin than a developer I often write scripts (in Perl or Bash). I don't work within a big group of developers, so I try to keep things simple. I tend to use [Git](http://git-scm.com/) for tracking my programs. Git is a free & open source, distributed version (revision) control system created by Linus Torvalds. Every Git repository contains complete history of revisions and is not dependent on a central server or network access. Git uses an intelligent compression system to reduce the cost of storing the entire history. Branching and merging are fast and (relatively :-) easy to do.

# Configuration

## First-time setup

Introduce yourself to git with your name and public email address before doing any operation:

```bash
git config --global user.name "Jeffrey Lebowski"
git config --global user.email "jlebowski@dude.org"
git config --global github.login=jlebowski
git config --global github.email=jlebowski@gmail.com
```

You only need to do this *once*. You might also like to add some aliases and [configuration](https://github.com/jreisinger/dotfiles/blob/master/.gitconfig).

## Everyday configuration

You can see your configuration like this:

```bash
$ git config --list    # output depends on whether you're in a git repo directory or not
```

You can also manage you configuration via [~/.gitconfig](https://github.com/jreisinger/dotfiles/blob/master/.gitconfig) file.

If you want to manage the current repository configuration just leave out the `--global` option or use the local config file `.git/config`.

# Gir repository

To start using git, you can either get a project from the Internet or create a new one.

## Getting a git repository

```bash
## the Linux kernel (approx. 150MB download):
$ git clone git://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux-2.6.git
## any server with ssh and git
$ git clone ssh://[user@]server.xy/path/to/repo.git/
```

## Starting a git repository

```bash
$ mkdir project
$ cd project
$ git init
```

You've now initialized the working directory. You may notice a new directory created, named `.git`. Git stores all of its stuff in this hidden directory. If you want Git to stop tracking your files, just remove `.git`.

Next, tell git to take a snapshot [in svk: momentka] of the contents of all files under the current directory (note the `.`), with git-add:

```bash
$ git add .     # add files in working directory to a temporary storage - index
```

This snapshot is now stored in a *temporary staging area* which git calls the "staging index" or just "index". You can permanently store the contents of the index in the repository with `git-commit`:

```bash
$ git commit    # add files in index to a permanent storage - repository (.git directory)
```

This will prompt you for a commit message. You've now stored the first version of your project in git.

# Making changes

When dealing with git, it's best to work in small bits. Rule of thumb: *if you can't summarise it in a sentence, you've gone too long without committing*.

The typical working cycle is:

1. Work on your project.
2. Check which files you have changed:

        $ git status

3. Check what the actual changes were:

        $ git diff

4. Add any files/folders mentioned in step 2:

        $ git add file1 newfile2 newfolder3

5. You are now ready to commit. You can see what is about to be committed using `git-diff` with the `--cached` option:

        $ git diff --cached

    (Without `--cached`, `git-diff` will show you any changes that you've made but not yet added to the index.) If you need to make any further adjustments, do so now, and then add any newly modified content to the index. Finally, "commit" your changes with:

        $ git commit

    This will again prompt you for a message describing the change, and then record a new version of the project.

Alternatively, instead of running `git-add` beforehand, you can use

    $ git commit -am "commit message"

which will automatically notice any modified (but not new) files, add them to the index, and commit, all in one step.

A note on commit messages: Though not required, it's a good idea to begin the commit message with a single short (less than 50 character) line summarizing the change, followed by a blank line and then a more thorough description. Tools that turn commits into email, for example, use the first line on the `Subject:` line and the rest of the commit in the body.

# Excluding some files

To set certain files or patterns to be ignored by Git, you must either modify the `$GIT_DIR/info/exclude` file or create a file called `.gitignore` in your project’s root directory. The former is not shared between repositories (so you can set ignores for your specific environment), while the `.gitignore` is usually checked into Git and distributed along with everything else.

For example if you don't want to track dot-files, setup `.gitignore` like this:

```bash
.*
!/.gitignore
```

See more [here](http://github.com/guides/ignore-for-git).

If some of your files is already beeing tracked by git, you can untrack it like this:

```bash
git rm --cached <filename>
```

# Undo/Redo

Go back and forget about every change past a certain point

```sh
git reset --hard SHA1_HASH  # all newer commits get erased forever!!!
```

Travel back in time

```sh
git checkout SHA1_HASH      # newer commits are preserved
```

... if you now edit and commit, you will be in an alternate reality (called a branch)

# Branching and merging

Branch is a separate line of development.

To stop Git complaining, always commit or reset your changes before running checkout.

Create and switch to a new branch

```sh
git checkout -b experimental_idea
```

Compare two branches

```sh
# from within e.g. experimental_idea branch
git diff ..master
git diff master..
git log ..master
git log master..
```

Push a new local branch to remote

```
git push -u origin experimental_idea
```

Merge two branches

```sh
git checkout master
git merge experimental_idea
```

Delete a branch

```sh
git branch -d experimental_idea # you might need -D if not merged

# delete a remote branch
git push origin --delete experimental_idea
```

## My workflow

```
if [[ $CURRENT_BRANCH == "master" ]]; then
    run_cmd 'git pull'
    run_cmd 'git push'
else
    #run_cmd 'git pull'
    run_cmd 'git checkout master'
    run_cmd 'git pull'
    run_cmd "git checkout $CURRENT_BRANCH"
    run_cmd 'git merge master'
    run_cmd 'git push'
fi
```

See [git-sync](https://github.com/jreisinger/dotfiles/blob/master/bin/git-sync) for my current setup.

# Tips and Tricks

To to list tracked files:

```bash
# currently tracked files under master branch
git ls-tree -r master --name-only

# files that ever existed (i.e. including deleted files)
git log --pretty=format: --name-only --diff-filter=A | sort - | sed '/^$/d'
```

To remove files not tracked by Git (like log files, zipped files, compiled files)

```bash
git clean -n  # dry-run ...
git clean -f  # files removed!
```

Apply changes generated via `git diff [--binary]`:

```bash
git apply --ignore-space-change --ignore-whitespace
```

While in Git-tracked directory, print the filename followed by the author of the last commit:

```bash
for f in `find -type f`; do
        git log -1 --date=iso -- $f |
        grep ^Author |
        perl -wnla -s -F: -e 'print "$file --" . $F[1]' -- -file=$f
done
```

Ignore changes of the files' mode (for current repo):

```bash
git config core.fileMode false
```

# More

* [Pro Git - Online book](https://git-scm.com/book)
* See my [wiki](http://wiki.reisinge.net/Git) for more
* [My blog post derived from this text](http://jreisinger.blogspot.sk/2014/01/simple-source-code-management-with-git.html)
