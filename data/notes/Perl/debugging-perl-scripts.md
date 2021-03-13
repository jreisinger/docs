# Perl Debugging

There are several ways how to debug a Perl program.

## Print out the variables from inside the program

### print()

The simplest way is to print the contents of the variables, ex.:

    #!/usr/bin/perl
    use strict;
    use warnings;

    # A number between 1 and 100 (both included)
    my $secret = int( 1 + rand 100 );

    # If DEBUG environment variable set, print out the secret number
    print "The secret number is $secret\n"
      if $ENV{DEBUG};

Usage:

    $ DEBUG=1 perl script.pl
    The secret number is 3

### Data::Dumper, YAML

In case you're dealing with a more complicated data structure (like the `$server` variable below), you can use [Data::Dumper](http://perldoc.perl.org/Data/Dumper.html) (outputs valid Perl code) or [YAML](https://metacpan.org/module/YAML) (more compact output):

    #!/usr/bin/perl
    use strict;
    use warnings;
    use Data::Dumper;
    use YAML;

    my $name = 'foo';

    my $server = {
        'server' => {
            'imageRef'  => '8a3a9f96-b997-46fd-b7a8-a9e740796ffd',
            'flavorRef' => '2',
            'name'      => $name,
            'metadata'  => { 'My Server Name' => 'Ubuntu 12.10 (Quantal Quetzal)' },
        }
    };

    print "== Data::Dumper ==\n";
    print Dumper $server;

    print "== YAML ==\n";
    print Dump $server;

## Perl debugger

Perl debugger (`perl -d <program.pl>`) shows each line of code before it executes it. The most common commands:

 * `h` -- help
 * `s` -- single-step the program
 * `x` -- dump a variable value (for complex data use `x \%href` or even `x sort keys %hash`)
 * `n` -- step over a subroutine
 * `q` -- quit the debugger

To have the command line history, you need:

 * Debian/Ubuntu packages: `libncurses-dev libreadline-dev`
 * Perl module: `Term::ReadLine::Gnu`

## More

* [Blog post](http://jreisinger.blogspot.sk/2013/12/debugging-perl-scripts.html) based on this text.
* [LEARNING THE PERL DEBUGGER](http://techblog.net-a-porter.com/2014/03/learning-the-perl-debugger-introduction/) -- an interactive tutorial on GitHub
* [Basic debugging checklist](http://perlmonks.org/?node=Basic%20debugging%20checklist)
* [Debugging Perl with hdb](http://perlmaven.com/debugging-perl-with-hdb)
