# HTML Parsing

I wanted to store the names of [Slovak MPs](http://www.nrsr.sk/web/?sid=poslanci/zoznam_abc) in a text file. Perl programming [documentation](http://perldoc.perl.org/perlfaq6.html#How-do-I-match-XML%2c-HTML%2c-or-other-nasty%2c-ugly-things-with-a-regex%3f) says to forget about regular expressions when messing with ugly things like HTML. So I used [HTML::TreeBuilder](https://metacpan.org/module/HTML::TreeBuilder) and read a few articles about it:

* [HTML::Tree(Builder) in 6 minutes](http://www.perlmonks.org/?node_id=280461)
* [Analyzing HTML with Perl](http://www.perl.com/pub/2006/01/19/analyzing_html.html)

I also needed to take care of the Unicode stuff:

* [Perl Unicode Cookbook](http://www.perl.com/pub/2012/04/perlunicook-standard-preamble.html)

And here is the result:

    #!/usr/bin/perl
    # Get the names of Slovak MPs
    use strict;                 # quote strings, declare variables
    use warnings;               # on by default
    use open qw(:std :utf8);    # undeclared streams in UTF-8
    use LWP::Simple;
    use HTML::Element;
    use HTML::TreeBuilder;

    my $url     = 'http://www.nrsr.sk/web/?sid=poslanci/zoznam_abc';
    my $content = get $url;
    my $tree    = HTML::TreeBuilder->new;                              # empty tree
    $tree->parse_content($content);

    my @elements = $tree->look_down( 'href', qr/PoslanecID/ );

    for my $element (@elements) {
        print $element->as_text, "\n";
    }

Other HTML parsing programs:

* [dlznik.pl](https://github.com/jreisinger/audit/blob/master/scripts/dlznik.pl)
* [foaf.pl](https://github.com/jreisinger/audit/blob/master/scripts/foaf.pl)
* [orsr](https://github.com/jreisinger/audit/tree/master/orsr)
