I think Perl one-liners are still super useful. They are small Perl programs that are run directly from a command line (ex. on Unix/Linux, Cygwin). Like this one from the Kubernetes [job](https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/#running-an-example-job) documentation:

```
perl -Mbignum=bpi -wle "print bpi(2000)"
```

`perl` is the Perl language interpreter. `-M` and `-wle` are command line switches (or options) that modify the `perl`'s behaviour. See below for explanation of what they mean. The string within doubles quotes is the Perl code that gets executed. In this case it calculates the PI with accuracy of 2000 digits. The command will take a while to finish.

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

## Search and replace

Find lines in logs that contain error or warning:

```
perl -wne '/error|warning/i && print' /var/log/*.log
```

The thing between slashes is a regular expression. It means match string `error` or string `warning` anywhere in the log line. `i` says to Perl to ignore the case. So it will match ERROR, error, Warning etc. If the regex finds a match (i.e. evaluates to true) the `&&` logical operator runs the `print` statement that will print the line containing the match.

Replace `/bin/sh` with `/bin/bash` and emit the transformed passwd file to STDOUT:

```
perl -pe 's#/bin/sh$#/bin/bash#' /etc/passwd
```

We used `#` instead of `/` as delimeters for better readibility since to strings themselves contain slashes.

Replace `colour` with `color` in all text files. The original files will be kept with `.bak` suffix:

```
perl -i.bak -pe 's/colour/color/g' *.txt
```

`g` (global) means replace all occurences (in a line) not just the first one.

## Columns

Print 2nd and 1st column:

```
$ cat birthdays.txt
03/30/45 Eric Clapton
11/27/42 Jimi Hendrix
06/24/44 Jeff Beck
$ perl -lane 'print "@F[1,0]"' birthdays.txt
Eric 03/30/45
Jimi 11/27/42
Jeff 06/24/44
```

Convert DOS files to Unix files:

```
perl -i -pe 's/\r//'  <file1> <file2> ... # dos-to-unix
perl -i -pe 's/$/\r/' <file1> <file2> ... # unix-to-dos
```

Calculate the total size of found log files:

```
find /opt/splunk/syslog/ -iname "*log*" -type f -mtime +30 | \
perl -lne '$sum += (stat)[7]}{print $sum'
```

We are using here the so called [Eskimo Greeting Operator](http://www.catonmat.net/blog/secret-perl-operators/#eskimo) as suggested by [PerlMonks](http://www.perlmonks.org/?node_id=1172707).

Remove comments and compress all consecutive blank lines into one:

```
cat /etc/ssh/sshd_config | perl -lne '!/^#/ && print' | perl -00 -pe ''
```

Find big palindromes:

```
perl -lne 'print if $_ eq reverse and length >= 5' /usr/share/dict/words

Greet user (stolen from [Utilitarian](http://perlmonks.org/?node_id=681898)) (-:

```
perl -E 'say "Good ".qw(night morning afternoon evening)[(localtime)[2]/6].", $ENV{USER}"'
```

For a deeper dive see [Famous Perl One-Liners Explained](http://www.catonmat.net/blog/perl-one-liners-explained-part-one/). If you want a book have a look at [Minimal Perl for UNIX and Linux People](http://www.amazon.com/Minimal-Perl-UNIX-Linux-People/dp/1932394508/ref=sr_1_1?ie=UTF8&qid=1358096838&sr=8-1&keywords=minimal+perl+for+unix).
