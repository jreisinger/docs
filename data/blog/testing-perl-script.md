---
title: "Testing Perl script"
date: 2015-12-22
categories: [prog]
tags: [perl, testing]
---

At [$work](https://www.eset.com/), I was to upgrade several Debians from Squeezy through Wheezy to
Jessie (6 to 8). I wanted to be sure that after the upgrade (mostly) the same
processes are running as before. I whipped up a
[script](https://github.com/jreisinger/checkprocs/blob/master/checkprocs), that
simply stores the list of running processes before the upgrade. When run
subsequently it reports missing processes (if any). 

To make the script reliable and easy to maintain I wanted to test it somehow.
To do that I turned the script into a
[modulino](http://www.perlmonks.org/index.pl?node_id=396759) following brian d
foy's advice in chapter 17 of Mastering Perl. The trick was to put all the code
into subroutines that can be tested and using the
[caller()](http://perldoc.perl.org/functions/caller.html) function to decide
whether the script is used as a script or as a module. The script looks something 
like this now:

    #!/usr/bin/env perl
    use strict;
    use warnings;
    use 5.010;
    use autodie;
    use Getopt::Long;
    use Pod::Usage;
    use Storable qw(freeze thaw);
    
    GetOptions(
        "h|?|help"  => \my $help,
        "l|print"   => \my $print,
        "v|verbose" => \my $verbose,
        "n|net"     => \my $net,
    ) or pod2usage(1);
    pod2usage( -exitval => 0, -verbose => 2, -noperldoc => 1 ) if $help;
    
    run() unless caller();
    
    sub run {
        # code
    }
    
    sub missing_procs {
        # code
    }

    sub get_procs {
        # code
    }

After this modification I created a symlink

    ln -s checkprocs checkprocs.pm 

and wrote a couple of tests in `checkprocs.t`

    use strict;
    use warnings;
    use Test::More tests => 3;
    
    use_ok('checkprocs');
    
    #<<<
    my $old = [(
        'proc1',
        '/path/to/proc2',
        'proc3',
        'proc4 --with-arg',
        '/path/to/proc5 -w',
    )];
    my $new = [(
        'proc1',
        'proc3',
        '/path/to/proc5'
    )];
    #>>>
    
    {
        my @missing_procs = main::missing_procs( $old, $new );
        is(
            "@missing_procs",
            '/path/to/proc2 proc4',
            'Found missing process w/o args'
        );
    }
    
    {
        my @missing_procs = main::missing_procs( $old, $new, { verbose => 1 } );
        is(
            "@missing_procs",
            '/path/to/proc2 proc4 --with-arg',
            'Found missing process w/ args'
        );
    }

Since I need to run the script under different Perl versions (Squeeze had
5.10.1, Wheezy 5.15.2 and Jessie 5.20.2) I used [perlbrew](http://perlbrew.pl)
to test it:

    $ perlbrew exec prove checkprocs.t
    perl-5.10.1
    ==========
    checkprocs.t .. ok
    All tests successful.
    Files=1, Tests=3,  0 wallclock secs ( 0.01 usr  0.00 sys +  0.04 cusr  0.00 csys =  0.05 CPU)
    Result: PASS
    
    perl-5.14.2
    ==========
    checkprocs.t .. ok
    All tests successful.
    Files=1, Tests=3,  0 wallclock secs ( 0.01 usr  0.00 sys +  0.04 cusr  0.00 csys =  0.05 CPU)
    Result: PASS
    
    perl-5.20.2
    ==========
    checkprocs.t .. ok
    All tests successful.
    Files=1, Tests=3,  0 wallclock secs ( 0.01 usr  0.00 sys +  0.03 cusr  0.00 csys =  0.04 CPU)
    Result: PASS

