# Check an IP address

Sometimes I come across an unknown IP address. This happens, for example, when I'm reviewing logs and I see that someone or (most probably) something was trying to SSH into the system. Or it was enumerating the URL paths of a web application.

In such scenario I want to have a quick and easy way to check the IP address. I created a command line tool called [checkip](https://github.com/jreisinger/checkip) that does just that. For example, the following IP address definitely looks suspicious:

<img src="/static/checkip.png" style="max-width:100%;width:640px">

Of course, I can mix and match [checkip](https://github.com/jreisinger/checkip) with the standard shell tools. First let me get some IP addresses to check from a Linux box:

```
$ journalctl --since "00:00" |  perl -lne '/((?:\d{1,3}\.){3}\d{1,3})/ && print $1' | sort | uniq > /tmp/ips-all.txt
```

Now I check all of them and get only those considered suspicious by more than 3 checks (`checkip` exit code is the number of checks that say the IP address is not OK):

```
$ cat /tmp/ips-all.txt | xargs -I {} bash -c 'checkip {} > /dev/null; [[ $? -gt 3 ]] && echo {}'
101.32.22.56
101.33.116.189
101.89.219.59
^C
```

Or I can find out from where are my services being used:

```
$ cat /tmp/ips-all.txt | xargs -I {} checkip -check geo {} | sort | uniq -c | sort -n | tail -3
     17 Geolocation city unknown | France | FR
     40 Geolocation city unknown | United States | US
     41 Geolocation city unknown | China | CN
```

