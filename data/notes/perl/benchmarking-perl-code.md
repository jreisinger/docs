(Up-to-date
[source](https://github.com/jreisinger/blog/blob/master/posts/benchmarking-perl-code.md)
of this post.)

Sometimes my code takes a really long time to run and I'd like to know which of
the alternatives runs faster.

In this example I compare three sorting subroutines; a "naive" approach, "the
Schwartzian Transform" and "the Orcish Manoeuvre". The first subroutine just
compares all files' sizes to each other. The second first precomputes the size
of each file and then does the comparisons. The third only computes the size if
not already cached in a hash.

    #!/usr/bin/perl
    use strict;
    use warnings;
    use 5.010;
    use Benchmark qw(timethese);

    # the minimum number of CPU seconds to tun
    my $CPU_SECS = shift // 2;

    chdir;    # change to my home directory
    my @files = grep -f, glob '* .*';
    print "found ", scalar @files, " files\n";

    timethese(
        "-$CPU_SECS",
        {
            naive => sub {
                my @sorted =
                    sort { -s $a <=> -s $b }
                @files;
            },
            schwartzian => sub {
                my @sorted =
                    map  { $_->[0]                }
                    sort { $a->[1] <=> $b->[1]    }
                    map  { [ $_, -s $_ ]          }
                @files;
            },
            orcish => sub {
                my %size;  # cache of files' sizes
                my @sorted =
                    sort { ( $size{$a} //= -s $a ) <=> ( $size{$b} //= -s $b ) }
                @files;
            },
        }
    );


The program's output:

    found 25 files
    Benchmark: running naive, orcish, schwartzian for at least 2 CPU seconds...
         naive:  2 wallclock secs ( 1.00 usr +  1.03 sys =  2.03 CPU) @ 7182.27/s (n=14580)
        orcish:  2 wallclock secs ( 1.55 usr +  0.51 sys =  2.06 CPU) @ 22422.33/s (n=46190)
    schwartzian:  2 wallclock secs ( 1.58 usr +  0.50 sys =  2.08 CPU) @ 22015.38/s (n=45792)

The output says that the Schwartzian Transform and the Orcish Manoeuvre are much
faster, i.e. the function ran more times in 2 seconds. The reason is that we don't
ask for the file size (a relatively expensive operation) each time we want to 
compare two files sizes; we ask just once for each file and we cache the result.
This way we run the expensive function N times instead of N.log(N) times (N is 
the number of files, N.log(N) is the number of comparisons).

If you want to see a comparison chart of the runtimes use `cmpthese()` instead of `timethese()`.

## See Also

* http://perldoc.perl.org/Benchmark.html
* Intermediate Perl, 2nd, p. 144
* http://www.perlmonks.com/?node_id=393128
