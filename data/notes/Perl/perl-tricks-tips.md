Handle STDIN vs. @ARGV
----------------------

    # https://varlogrant.blogspot.sk/2017/05/count-is-not-uniq.html
    my %seen ;
    
    map { $seen{$_}++ } do {
        @ARGV ? @ARGV : map { chomp ; $_ } <>;
        } ;
    
    while ( my ( $k, $v ) = each %seen ) {
        say join "\t", $v, $k ;
        }


Uninstall a module
------------------

    cpanm --uninstall Module::Name

See also

* https://perltricks.com/article/3/2013/3/27/How-to-cleanly-uninstall-a-Perl-module/
* http://stackoverflow.com/questions/2626449/how-can-i-de-install-a-perl-module-installed-via-cpan

Reading a directory
-------------------

    opendir my $DIR, $dir or die "Cannot open $dir: $!\n";
    my @names = readdir $DIR;
    closedir $DIR;
    
    for my $name (@names) {
        next if $name eq '.' or $name eq '..';
        # do something with the $name
    }

Sending an email
----------------

Using the `mail` command (you might need to setup something like [ssmtp](http://jreisinger.blogspot.sk/2014/02/fixing-email-aliases-when-using-ssmtp.html)):

    sub send_mail {
        my $recipient = shift;
        my $subject   = shift;
        my $body      = shift;

        my $mailexe = '/usr/bin/mail';

        open my $MAIL, "|$mailexe -s '$subject' '$recipient'"
            or die "Cannot send mail: $!\n";

        print $MAIL $body;
    }
    
Other ways

* http://learn.perl.org/examples/email.html
* http://perldoc.perl.org/Net/SMTP.html
* https://rjbs.manxome.org/rubric/entry/2104
* see Perl Cookbook

See also
--------

http://wiki.reisinge.net/PerlTips
