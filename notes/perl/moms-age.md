# Mathematical Riddle - Mom's Age

Question: Mom's 38, her two sons are 7 and 9. When will the sum of sons' ages be equal to mom's age?

Answer 1:

    (7 + x) + (9 + x) = 38 + x
                    x = 22

Answer 2:

    #!/usr/bin/perl
    use strict;
    use warnings;

    my $ma   = 38;
    my $son1 = 7;
    my $son2 = 9;

    until ( $son1 + $son2 == $ma ) {
        if ( $son1 + $son2 > $ma ) { # the age equaled in the past
            $ma--;
            $son1--;
            $son2--;
        } else {
            $ma++;
            $son1++;
            $son2++;
        }
    }

    print $ma . "\n";
