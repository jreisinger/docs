I think Perl one-liners are still super useful. They are small Perl programs that are run directly from a command line (ex. on Unix/Linux, Cygwin). Like this one from the Kubernetes [job](https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/#running-an-example-job) documentation:

```
perl -Mbignum=bpi -wle "print bpi(2000)"
```

`perl` is the Perl language interpreter. `-M` and `-wle` are command line switches (or options) that modify the `perl`'s behaviour. See below for explanation of what they mean. The string within doubles quotes is the Perl code that gets executed. In this case it calculates the Pi with accuracy of 2000 digits. The command will take a while to finish.

# Switches

These are some of the most used command line switches:

* `-e '<code>'` -- **e**xecute `<code>`
* `-E '<code>'` -- **E**xecute `<code>` enabling [feature](http://perldoc.perl.org/feature.html) bundle (like `use 5.010`) for your version of Perl
* `-w` -- enable **w**arnings (generally advisable)
* `-p` -- loop through lines, reading and **p**rinting them (in-script equivalent: `while (<>) { [<code>] print }`)
* `-n` -- loop through lines, reading but **n**ot printing them
* `-l` -- print a new**l**ine (`$/` actually) after each line of output and chomp newline if used with `-n` or `-p`
* `-i[<.ext>]` (**i**ntrepid) -- create backup file (with `<.ext>` extension if defined)
* `-a` (**a**utosplit) -- split the `$_` into `@F` (space is the default separator, change it with `-F`, ex. `-F:`)
* `-s` -- rudimentary parsing of command line **s**witches (see "Git-tracked directory" multi-liner below)
* `-m<module>[=<subroutine>,...]` -- load subroutine(s) from a **m**odule

See [perlrun](http://perldoc.perl.org/perlrun.html) for more.

# Examples

Seach and replace:

```
# Find lines containing `<regex>`:
perl -lne 'print if /<regex>/' aFile

# Find DNS resource records of type A:
find /etc/bind -type f | xargs perl -ne '/\s+A\s+/ and print "$ARGV: $_"'

# Emit the transformed passwd file to STDOUT:
perl -pe 's#/bin/sh$#/bin/bash#' /etc/passwd

# In-place editing with backups:
perl -i.bak -pe 's/colour/color/g' *.txt
```

Switch columns:

```
$ cat birthdays.txt
03/30/45 Eric Clapton
11/27/42 Jimi Hendrix
06/24/44 Jeff Beck
$ perl -lane 'print "@F[1,0]"' birthdays.txt
```

Leave out the first column:

```
history | perl -anE 'say join " ", @F[1 .. $#F]' | sort | uniq
```

* see http://www.perlmonks.org/?node_id=739305 for why you can't use -1 as array index here

Convert DOS files to Unix files:

```
perl -i -pe 's/\r//' <file1> <file2> ...   # dos-to-unix
perl -i -pe 's/$/\r/' <file1> <file2> ...  # unix-to-dos
```

Total size of found files (using the [Eskimo Greeting Operator](http://www.catonmat.net/blog/secret-perl-operators/#eskimo) as suggested by [PerlMonks](http://www.perlmonks.org/?node_id=1172707)):

```bash
find /opt/splunk/syslog/ -iname "*log*" -type f -mtime +30 | \
perl -lne '$sum += (stat)[7]}{print $sum'
```

Remove comments and compress all consecutive blank lines into one ([more](http://www.catonmat.net/blog/perl-one-liners-explained-part-one/)):

```bash
cat /etc/ssh/sshd_config | perl -lne '!/^#/ and print' | perl -00 -pe ''
```

Fun

```
# Find big palindromes:
perl -lne 'print if $_ eq reverse and length >= 5' /usr/share/dict/words

# Greet user (stolen from [Utilitarian](http://perlmonks.org/?node_id=681898)) (-:
perl -E 'say "Good ".qw(night morning afternoon evening)[(localtime)[2]/6].", $ENV{USER}"'
```

For a deeper dive see [Famous Perl One-Liners Explained](http://www.catonmat.net/blog/perl-one-liners-explained-part-one/). If you want a book have a look at [Minimal Perl for UNIX and Linux People](http://www.amazon.com/Minimal-Perl-UNIX-Linux-People/dp/1932394508/ref=sr_1_1?ie=UTF8&qid=1358096838&sr=8-1&keywords=minimal+perl+for+unix).
