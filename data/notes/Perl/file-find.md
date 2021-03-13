# File::Find

Sometimes you need to do something to many/all files within certain directory.

    use File::Find;
    
    ## use a subroutine reference (coderef)
    sub process_file {
        # do something to each file found
    }
    find(\&process_file, @DIRLIST);
    
    ## or use anonymous subroutine
    find(
        sub {
            # do something to each file found
        },
        @DIRLIST
    }

.. `find` function from [File::Find](http://perldoc.perl.org/File/Find.html) scans directories in `@DIRLIST` recursively and for each file or directory calls the referenced function `process_file`

.. before calling your function `find` by default changes to the directory being scanned and sets the following (global) variables:

* `$File::Find::dir` -- visited directory path relative to the program's starting directory
* `$File::Find::name` -- full path of the file being visited relative to the program's starting directory
* `$_` -- basename of the file being visited

We are passing both data (the list of directories to search) and behaviour as parameters to the `find` routine.

We are using the subroutine reference as a *callback*.

## Find the largest file

NOTE: when dereferenced, a subroutine can see all visible lexical variables when reference to the subroutine is taken.

    #!/usr/bin/perl
    use strict;
    use warnings;
    use File::Find;

    @ARGV = qw(.) unless @ARGV;
    my ($biggest_size, $biggest_name) = (-1, "");

    sub biggest {
        return unless -f && -s _ > $biggest_size;
        $biggest_size = -s _;
        $biggest_name = $File::Find::name;
    }

    find(\&biggest, @ARGV);
    print "Biggest file $biggest_name in @ARGV is $biggest_size bytes long.\n";

.. `_` -- virtual filehandle that uses info from the last file lookup (`stat` - an expensive operation) that a file test operator (`-f`) performed

## Find the most recently changed file

    my ($age, $name);
    sub youngest {
        return if defined $age && $age > (stat($_))[9];
        $age = (stat($_))[9];
        $name = $File::Find::name;
    }

## File::Find::Rule

[File::Find::Rule](https://metacpan.org/module/File::Find::Rule) CPAN module is an alternative to File::Find:

    use File::Find::Rule;

    my @txt_files  = File::Find::Rule
        ->file
        ->name('*.txt')
        ->in('./');

    foreach (@txt_files) {
        print "$_\n";
    }

## More

* Intermediate Perl, Ch. 7
* Perl Cookbook, Recipe 9.7
* [Beginners guide to File::Find](http://www.perlmonks.org/?node_id=217166)
* [File::Find and permission error handling](http://perlmonks.org/?node_id=1023278) -- using `preprocess`
