A module is a collection of related Go packages that are versioned together as a single unit. Modules record precise dependency requirements and create reproducible builds.

Modules must be semantically versioned according to [semver](https://semver.org), usually in the form `v(major).(minor).(patch)`.

Relationship between repositories, modules and packages:

* A *repository* contains one (most often) or more Go modules
* Each *module* contains one or more Go packages
* Each *package* consists of one or more Go source files in a single directory

More

* https://github.com/golang/go/wiki/Modules
