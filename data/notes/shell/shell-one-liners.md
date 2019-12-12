`su` to a user without shell

    su - www-data -s /bin/bash -c 'ulimit -Sn'

Count number of lines if Perl files

    for f in $(ack -f --perl); do wc -l $f; done | \
    sort -n | \
    perl -nE '$n += (split)[0]; print; END {say "-" x 40, "\n", $n}'

Sort running processes by memory use

    ps -eo pmem,pcpu,rss,vsize,args | sort -k 1 -r | less

Find files you might want to backup

    find / -type f ! -path '*/.*' -mtime -30 2>/dev/null | \
    grep -E -v \
    -e '^/proc' \
    -e '^/sys' \
    -e '^/var/(lib|cache|log|spool|mail|backups)' \
    -e '^/tmp' \
    -e '^/run' \
    -e '^/usr' \
    -e '^/lib' \
    -e '^/boot' \
    -e '^/dev'

Check which machines you can login into - [canssh](https://github.com/jreisinger/dotfiles/blob/master/bin/canssh)

Compute strings' SHA1 message digests

    for name in Aragorn Bilbo Gandalf; do echo -n $name | sha1sum; done

Find and rename multiple files (`*.log` => `*.LOG`)

    find . -type f -name '*.log ' | grep -v .do-not-touch | while read fname
    do
        echo mv $fname ${fname/.log/.LOG/}
    done

Search MS Word files for "robot" string:

    find /data -type f -iname "*.doc" -print0 | \
    xargs -0 strings -f | \
    grep -i 'robot' > find.out

* to get just the filenames

        cat find.out | cut -d: -f1 | sort | uniq

See also 

* http://www.commandlinefu.com
* http://www.catonmat.net/download/perl1line.txt
* http://wiki.reisinge.net/PerlOneLiners
