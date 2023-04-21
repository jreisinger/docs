![](https://user-images.githubusercontent.com/1047259/233588004-d2feae9d-5962-462a-bc6f-6d426ab8f026.png)

I keep most of my code and prose on GitHub in public repos. I do it because I get reliable storage for free that I can access from any computer. And it might be helpful to others. Also it engages my [hubris](https://thethreevirtues.com/) :-). It's working nicely but sometimes I get a bad feeling when I push stuff. I'm worried that I might leak some sensitive information like passwords, API keys or tokens. The obvious solution is to think twice before committing and pushing data.

## New commits

But there are also some helpful tools, like gitleaks. It basically finds and reports secrets in the files you are about to commit. I want to run it whenever I commit something in *any* of my repos. These are the steps to make that happen:

1. Install [gitleaks](https://github.com/gitleaks/gitleaks).
2. Add this to your `~/.gitconfig`:

```
[core]
    hooksPath = ~/.git-global-hooks
```

3. Create `~/.git-global-hooks/pre-commit`:

```sh
#!/bin/bash
# Detect secrets in a git repos using https://github.com/zricethezav/gitleaks

if [[ $SKIP == "gitleaks" ]]; then
    echo "skipping gitleaks checks ..."
    exit 0
fi

set -xe

# Check uncommitted changes (parsing output of 'git diff') that had been 'git add'ed.
gitleaks protect --no-banner --staged
```

## Existing commits

The steps above will prevent you from committing secrets from now on. But you should also check existing commits because you might have committed a secret in the past. You can either do it on each commit by adding these lines to `~/.git-global-hooks/pre-commit`:

```
# Check existing commits (parsing output of 'git log -p').
gitleaks detect --no-banner
```

But on bigger repos this might take several seconds every time you commit. To avoid this you can check all your historical commits in all your repos once. I used [gh](https://cli.github.com/) and [runp](https://github.com/jreisinger/runp) to do it:

```sh
export GHORG=jreisinger # CHANGE ME
mkdir /tmp/$GHORG && cd /tmp/$GHORG
# clone all repos in parallel
gh repo list $GHORG --limit 1000 | cut -f 1 | runp -n -p 'gh repo clone'
# check existing commits in all repos in parallel
ls | runp -p 'gitleaks detect --no-banner -s'
```

If `runp` exits with 0, all is good. Otherwise scroll up to review the output. To check a repo for committed leaks:

```sh
cd $REPO && gitleaks detect --no-banner -v
```

After you are done, you might want to remove the temporary repos:

```sh
rm -rf /tmp/$GHORG
```
