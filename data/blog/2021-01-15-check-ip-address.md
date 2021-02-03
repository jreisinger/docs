# Check an IP address

Sometimes I come across an unknown IP address. This happens, for example, when I'm reviewing logs and I see that someone or (most probably) something was trying to SSH into the system. Or it was enumerating the URL paths of a web application.

In such scenario I want to have a quick and easy way to check the IP address. I created a command line tool called [checkip](https://github.com/jreisinger/checkip) that does just that. For example, the following IP address definitely looks suspicious:

<img src="/static/checkip.png" style="max-width:100%;width:640px">

Of course, I can mix and match [checkip](https://github.com/jreisinger/checkip) with the standard shell tools. First let me get some IP addresses to check from a Linux box:

```
$ journalctl --since "00:00" |  perl -lne '/((?:\d{1,3}\.){3}\d{1,3})/ && print $1' | sort | uniq > /tmp/ips-all.txt
```

Now I check all of them and get only those suspicious (`checkip` exits non-zero if at least one check says the IP address is not OK):

```
$ cat /tmp/ips-all.txt | xargs -I {} bash -c 'checkip -check ipsum {} > /dev/null || echo {}' > /tmp/ips-suspicious.txt
```

Almost one third of the IP addresses is suspicious (well, the Internet is a weird and terrible thing):

```
$ wc -l /tmp/ips-*
     318 /tmp/ips-all.txt
      98 /tmp/ips-suspicious.txt
```

Or I can find out from where are people (or programs) engaging with my services:

```
$ cat /tmp/ips-all.txt | xargs -I {} checkip -check geo {} | sort | uniq -c | sort -n | tail -3
     17 Geolocation city unknown | France | FR
     40 Geolocation city unknown | United States | US
     41 Geolocation city unknown | China | CN
```

