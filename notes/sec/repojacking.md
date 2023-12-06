- repository hijacking

The problem

- repository usernames get changed or deleted
- a bad actor can create a repo with the same name and a pre-existing username that was changed or deleted
- this way a bad actor can mount an open-source software supply chain attack

Prevention

- organizations that undergo name changes should still own their previous name as placeholders to prevent such abuse
- GitHub's [popular repository namespace retirement](https://github.blog/2018-04-18-new-tools-for-open-source-maintainers/#popular-repository-namespace-retirement)
- developers should be aware of the modules they use, and the state of the repository that the modules originated from

Go modules

- particularly susceptible since (unlike e.g. npm or PyPI) they are decentralized - published via version control platforms (e.g. GitHub, BitBucket)
- an attacker can register the newly unused username, duplicate the module repository, and publish a new module to proxy.golang.org and go.pkg.dev
- since Go modules are cached by the [module mirror](https://proxy.golang.org/) there could be popular Go module repositories that have been cloned less than 100 times, thus bypassing the GitHub's popular repository namespace retirement prevention

More: https://vulncheck.com/blog/go-repojacking
