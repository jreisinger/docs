#!/usr/bin/perl
# Scope of Perl variables. For more see:
# - Mastering Perl, ch. 7 Symbol Tables and Typeglobs
# - Programming Perl, ch. 2
use v5.14;
use warnings;

our $global = q(I'm the global version);

show_me('at start');
lexical();
localized();
show_me('at end');

sub show_me {
    # show_me() is outside of the lexical scope of any other subroutine
    
    my $tag = shift;

    print "show_me() $tag: \$global --> $global\n";
    print '-' x 79, "\n";
}

sub lexical {
    # "my" defines lexical scope for variables which depends solely on the
    # inspection of the code
    my $global = q(I'm the lexical version);

    {
        our $global;
        print "In the naked block, our \$global --> $global\n";
    }

    print "In lexical(), my \$global --> $global\n";
    print "The package version, \$main::global --> $main::global\n";
    show_me('from lexical()');
}

sub localized {
    # "local" defines dynamic scope which depends on the state of the program
    # during runtime
    local $global = q(I'm the localized version);

    print "In localized(): local \$global --> $global\n";
    print "The package version, \$main::global --> $main::global\n";
    show_me('from localized()');
}
