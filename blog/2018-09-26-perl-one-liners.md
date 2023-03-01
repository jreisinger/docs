# Perl one-liners

I think Perl one-liners are still super useful. They are small [Perl](https://www.perl.org/) programs that are run directly from command line. Like this one from the Kubernetes [job](https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/#running-an-example-job) documentation:

```
perl -Mbignum=bpi -wle 'print bpi(2000)' # calculate PI to 2000 digits
```

`perl` is the Perl language interpreter. `-M` and `-wle` are command line switches (or options) that modify the `perl`'s behaviour. See below for explanation of what they mean. The string within double quotes is the Perl code that gets executed. In this case it uses the `bpi` subroutine from the [bignum](https://perldoc.perl.org/bignum.html) module to calculate the PI with accuracy of 2000 digits. The command will take a while to finish.

## Switches

These are some of the most used command line switches:

* `-e '<code>'` -- **e**xecute `<code>`
* `-E '<code>'` -- **E**xecute `<code>` enabling new [feature](http://perldoc.perl.org/feature.html)s for your version of Perl
* `-w` -- enable **w**arnings (generally advisable)
* `-p` -- loop through lines, reading and **p**rinting them (in-script equivalent: `while (<>) { [<code>] print }`)
* `-n` -- loop through lines, reading but **n**ot printing them
* `-l` -- print a new**l**ine (`$/` actually) after each line of output and chomp input newlines if used with `-n` or `-p`
* `-i[<.ext>]` (**i**ntrepid) -- create backup file (with `<.ext>` extension if defined)
* `-a` (**a**utosplit) -- split the `$_` default variable into `@F` array (space is the default separator, change it with `-F`, ex. `-F:`)
* `-M<module>[=<subroutine>,...]` -- load subroutine(s) from a **m**odule

See [perlrun](http://perldoc.perl.org/perlrun.html) for more.

## Cut 

Cut out 2nd and 1st space separated field (column):

```
$ cat birthdays.txt
03/30/45 Eric Clapton
11/27/42 Jimi Hendrix
06/24/44 Jeff Beck

$ perl -wlane 'print $F[1]' birthdays.txt
Eric
Jimi
Jeff

$ perl -wlane 'print join " ", @F[1,0]' birthdays.txt
Eric 03/30/45
Jimi 11/27/42
Jeff 06/24/44
```

The field numbering starts at 0. We use [join](https://perldoc.perl.org/functions/join) to put a space between cut out fields.

## Search

Find lines in logs that contain error or warning:

```
perl -wne '/error|warning/i && print' /var/log/*.log
```

The thing between slashes is a [regular expression](https://perldoc.perl.org/perlre.html). It means match string `error` or string `warning` anywhere in the log line. `i` says to Perl to ignore the case. So it will match ERROR, error, Warning etc. If the regex finds a match (i.e. evaluates to true) the `&&` logical operator runs the `print` statement that will print the line containing the match.

Get IP addresses from logs:

```
journalctl --since "00:00" | perl -wlne '/((?:\d{1,3}\.){3}\d{1,3})/ && print $1' | \
sort | uniq > /tmp/ips.txt
```

The IP address regex explained:

```
(               # capturing parenthesis to be retrieved via $1
    (?:         # non capturing parenthesis, only for grouping
        \d{1,3} # one to three decimal numbers
        \.      # literal dot
    ){3}        # three times all within innermost parenthesis
        \d{1,3} # one to three decimal numbers
)               
```

For a more serious program where you want to cover possible edge cases you should use a well tested module [Regexp::Common](https://metacpan.org/pod/Regexp::Common) as suggested by [PerlMonks](https://perlmonks.org/?node_id=11127622).

## Replace

Replace `/bin/sh` with `/bin/bash` and emit the transformed passwd file to STDOUT:

```
perl -wpe 's#/bin/sh$#/bin/bash#' /etc/passwd
```

We used `#` instead of `/` as delimeter for better readibility since the strings themselves contain slashes. `$` means end of the string.

Replace `colour` with `color` in all text files. The original files will be kept with `.bak` suffix:

```
perl -i.bak -wpe 's/colour/color/g' *.txt
```

`g` (global) means replace all occurences in a string not just the first one.

Convert between DOS and Unix newline:

```
perl -i -wpe 's/\r//'  <file1> <file2> ... # dos-to-unix
perl -i -wpe 's/$/\r/' <file1> <file2> ... # unix-to-dos
```

## Calculate

Calculate the total size of log files older than 30 days:

```
find /opt/splunk/syslog/ -iname "*log*" -type f -mtime +30 | \
perl -wlne '$sum += (stat)[7]}{print $sum'
```

The [stat](https://perldoc.perl.org/functions/stat.html) function returns a 13-element list of status info about a file. We take the 8th element (with index `7`) which is the size of a file. We loop over the found files and add the size of each into the `$sum` variable. The handy [Eskimo Greeting Operator](http://www.catonmat.net/blog/secret-perl-operators/#eskimo) is for printing the `$sum` when the loop is over (suggested by [PerlMonks](http://www.perlmonks.org/?node_id=1172707)).

## More

[perl1line.txt](https://catonmat.net/ftp/perl1line.txt)

For a deeper dive see [Famous Perl One-Liners Explained](http://www.catonmat.net/blog/perl-one-liners-explained-part-one/). If you want a book have a look at [Minimal Perl for UNIX and Linux People](http://www.amazon.com/Minimal-Perl-UNIX-Linux-People/dp/1932394508/ref=sr_1_1?ie=UTF8&qid=1358096838&sr=8-1&keywords=minimal+perl+for+unix).
