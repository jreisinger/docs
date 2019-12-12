(Up-to-date
[source](https://github.com/jreisinger/blog/blob/master/posts/module-build.md)
of this post.)

The development and (Github, CPAN) release cycle

    # Work on the project and test locally
    perl Build.PL && ./Build && ./Build test && ./Build install
    
    # Prepare the distro for CPAN
    vi MANIFEST.SKIP       # do once; #!include_default
    ./Build manifest       # only files listed in MANIFEST will go into the distibution archive
    vi lib/App/Monport.pm  # increase VERSION string - search BUILD.PL for 'version' or 'version_from'
    vi Changes
    podselect lib/App/Monport.pm > README.pod
    perl Build.PL && ./Build && ./Build test && ./Build install && ./Build disttest && ./Build dist
    
    # Commit to GitHub and tag it with the version from Changes
    git commit -am 'commit message from Changes'
    git tag v<version>  # <version> from Changes, ex. v1.01
    git push --tags -u origin master
    
    # Upload the distro using CPAN::Uploader
    cpan-upload App-Monport-<version>.tar.gz --user reisinge

# Creating a Module::Build Distribution

We show here how to create a Perl distribution using `Module::Build` build system with `Module::Starter`. The other Perl build system (we don't show here) is `ExtUtils::MakeMaker`. For sophisticated distribution creation see [Dist::Zilla](http://dzil.org/).

Create config file `~/.module-starter/config`:

    author: Foo Bar
    email: foo@bar.org
    builder: Module::Build
    verbose: 1
    # Allow adding new modules to existing distro.
    plugins: Module::Starter::AddModule
    
... or use `module-starter` (see below) with command line arguments like:

    --author="Foo Bar" \
    --email=foo@bar.org \
    --mb \
    --verbose \
    --plugin=Module::Starter::AddModule
    
Run basic commands

* install needed modules: `cpanm Module::Build Module::Starter Module::Starter::AddModule`
* create (a working skeleton of) module distribution: `module-starter --module=Animal` 
* change to the created distro directory: `cd Animal`
* create the `Build` script: `perl Build.PL`
* build the distro (modules from `lib` copied to `blib` staging area and embedded documenation translated into Unix manpage in `blib/libdoc`): `./Build`
* make sure the tests pass: `./Build test` (or run individual tests - see below)
* test the distro: `./Build disttest`
* create the distro: `./Build dist`

Add modules

* add new modules: `module-starter --module=Sheep,Cow,Horse --dist=Animal`
* add new modules (we are inside our distribution directory): `module-starter --module=Sheep,Cow,Horse --dist=.`

Run individual tests

* rebuild distro and run test including modules from `blib/lib`: `./Build && perl -Iblib/lib -T t/Cow.t`
* rebuild distro and run test including modules from `blib/lib`: `./Build && perl -Mblib -T t/Cow.t`

Measure test coverage

* run `testcover` target: `./Build testcover`
* turn the collected statistics into human-readable reports: `cover`

Generate LICENSE using `App::Software::License`

    software-license --holder 'Jozef Reisinger' --license Perl_5 --type notice --year $(date +"%Y") > LICENSE

For more see:

* Alpaca Book, Ch. 12
* [The Perl Toolchain: developing your module](http://blogs.perl.org/users/neilb/2016/04/the-perl-toolchain-developing-your-module.html)
* [How to upload a script to CPAN](http://perltricks.com/article/how-to-upload-a-script-to-cpan/)
* http://blogs.perl.org/users/neilb/2014/08/put-your-cpan-distributions-on-github.html
