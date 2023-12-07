repojacking == repository hijacking

The problem

- repository usernames get changed or deleted
- a bad actor can create an account using this pre-existing username
- and a repository with a pre-existing name that is used as dependency/library
- this way a bad actor can mount an open-source software supply chain attack

Prevention

- organizations that undergo name changes should still own their previous name as placeholders to prevent such abuse
- GitHub's [popular repository namespace retirement](https://github.blog/2018-04-18-new-tools-for-open-source-maintainers/#popular-repository-namespace-retirement)
- developers should be aware of the modules they use, and the state of the repository that the modules originated from

Go modules

- particularly susceptible since (unlike e.g. npm or PyPI) they are decentralized - published via version control platforms (e.g. GitHub, BitBucket)
- an attacker can register the newly unused username, duplicate the module repository, and publish a new module to proxy.golang.org and go.pkg.dev
- since Go modules are cached by the [module mirror](https://proxy.golang.org/) there could be popular Go module repositories that have been cloned less than 100 times, thus bypassing the GitHub's popular repository namespace retirement prevention

Find Go dependencies (modules) susceptible to repojacking

```
export GHUSER=someuser
mkdir /tmp/$GHUSER && cd /tmp/$GHUSER
gh repo list $GHUSER --language Go --limit 1000 --json name |\
 jq -r '.[] | .name' |\
 runp -p 'gh repo clone'
gorepojack
```

- [gh](https://cli.github.com/)
- [runp](https://github.com/jreisinger/runp)
- gorepojack -> [tools](https://github.com/jreisinger/tools)

More

- https://github.com/jreisinger/tools/tree/main/cmd/gorepojack
- https://vulncheck.com/blog/go-repojacking
- https://blog.aquasec.com/github-dataset-research-reveals-millions-potentially-vulnerable-to-repojacking
