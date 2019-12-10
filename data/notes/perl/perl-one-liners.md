*2018-09-26*

I think Perl one liners are still super useful (even Kubernetes people like them to do their [jobs](https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/#running-an-example-job) :-). They are small Perl programs that are run directly from a command line (ex. on Unix/Linux, Cygwin). For a deeper dive see [Famous Perl One-Liners Explained](http://www.catonmat.net/blog/perl-one-liners-explained-part-one/). If you want a book have a look at [Minimal Perl for UNIX and Linux People](http://www.amazon.com/Minimal-Perl-UNIX-Linux-People/dp/1932394508/ref=sr_1_1?ie=UTF8&qid=1358096838&sr=8-1&keywords=minimal+perl+for+unix).

## `perl` command line switches

* `-e '<code>'` (execute) -- execute `<code>`
* `-E '<code>'` (execute) -- execute `<code>` enabling [feature](http://perldoc.perl.org/feature.html) bundle (like `use 5.010`) for your version of Perl
* `-p` (printing) -- loop through lines, reading and printing them (in-script equivalent: `while (<>) { [<code>] print }`)
* `-w` (warnings) -- enable warnings (generally advisable)
* `-n` (nonautoprinting) -- loop through lines, reading but not printing them
* `-l` (line) -- print a newline (`$/` actually) after each line of output and chomp newline if used with `-n` or `-p`
* `-i[<.ext>]` (intrepid) -- create backup file (with `<.ext>` extension if defined)
* `-a` (autosplit mode) -- split the `$_` into `@F` (space is the default separator, change it with `-F`, ex. `-F:`)
* `-s` (switch) -- rudimentary parsing of command line switches (see "Git-tracked directory" multi-liner below)

See [perlrun](http://perldoc.perl.org/perlrun.html) for more.

## Unix tools replacements

### `grep` replacement

Find lines containing `<regex>`:

```bash
perl -lne 'print if /<regex>/' aFile
```

Find DNS resource records of type A:

```bash
find /etc/bind -type f | xargs perl -ne '/\s+A\s+/ and print "$ARGV: $_"'
```

### `sed` replacement

Emit the transformed passwd file to STDOUT:

```bash
perl -pe 's#/bin/sh$#/bin/bash#' /etc/passwd
```

In-place editing with backups:

```bash
perl -i.bak -pe 's/colour/color/g' *.txt
```

### `awk` replacement

Switch columns:

```bash
$ cat birthdays.txt
03/30/45 Eric Clapton
11/27/42 Jimi Hendrix
06/24/44 Jeff Beck
$ perl -lane 'print "@F[1,0]"' birthdays.txt
```

Leave out the first column:

```bash
history | perl -anE 'say join " ", @F[1 .. $#F]' | sort | uniq
```

* see http://www.perlmonks.org/?node_id=739305 for why you can't use -1 as array index here

### `dos2unix` replacement

Convert DOS files to Unix files:

```bash
perl -i -pe 's/\r//' <file1> <file2> ...   # dos-to-unix
perl -i -pe 's/$/\r/' <file1> <file2> ...  # unix-to-dos
```

## Various

Total size of found files (using the [Eskimo Greeting Operator](http://www.catonmat.net/blog/secret-perl-operators/#eskimo) as suggested by [PerlMonks](http://www.perlmonks.org/?node_id=1172707)):

```bash
find /opt/splunk/syslog/ -iname "*log*" -type f -mtime +30 | \
perl -lne '$sum += (stat)[7]}{print $sum'
```

Remove comments and compress all consecutive blank lines into one ([more](http://www.catonmat.net/blog/perl-one-liners-explained-part-one/)):

```bash
cat /etc/ssh/sshd_config | perl -lne '!/^#/ and print' | perl -00 -pe ''
```

Create HTML anchor element:

```bash
perl -le 'print "<a href=\"$ARGV[1]\">$ARGV[0]</a>"' 'perldoc' http://perldoc.perl.org/
```

## Fun

Find big palindromes:

```bash
perl -lne 'print if $_ eq reverse and length >= 5' /usr/share/dict/words
```

Print a file system tree on UNIX like systems ([source](http://www.perlmonks.org/?node_id=1050343)):

```bash
ls -R | perl -ne'if(s/:$//){s{[^/]*/}{--}g;s/^-/\t|/;print}'
```

Greet user (stolen from [Utilitarian](http://perlmonks.org/?node_id=681898)) (-:

```bash
perl -E 'say "Good ".qw(night morning afternoon evening)[(localtime)[2]/6].", $ENV{USER}"'
```
