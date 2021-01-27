# Check an IP address

Sometimes I come across an unknown IP address. This happens, for example, when I'm reviewing logs and I see that someone or (most probably) something was trying to SSH into the system. Or it was enumerating the URLs of a web application.

In such scenario I want to have a quick and easy way to check the IP address. I created a command line tool called [checkip](https://github.com/jreisinger/checkip) that does just that. The following IP address definitely looks suspicious:

<img src="/static/checkip.png" style="max-width:100%;width:640px">

Of course, I can mix and match `checkip` with the standard shell tools.

To get get only malicious IP addresses from a list of IP addresses (checkip exits non-zero if at least one checker thinks the IP address is not OK):

```
$ cat suspicious-ips.txt | xargs -I {} bash -c 'checkip {} > /dev/null || echo {} is not OK'
212.114.52.42 is not OK
```

Or to find out from where are people (or programs) engaging with my services I run:

```
$ journalctl --since "00:00" |  perl -wlne '/((?:\d{1,3}\.){3}\d{1,3})/ and print $1' | sort | uniq | xargs -I {} checkip -check geo {} | sort | uniq -c | sort -n | tail
      8 Geolocation London | United Kingdom | GB
      8 Geolocation city unknown | India | IN
      8 Geolocation city unknown | Russia | RU
      9 Geolocation Amsterdam | Netherlands | NL
     11 Geolocation Toronto | Canada | CA
     13 Geolocation Frankfurt am Main | Germany | DE
     13 Geolocation city unknown | Singapore | SG
     17 Geolocation city unknown | France | FR
     40 Geolocation city unknown | United States | US
     41 Geolocation city unknown | China | CN
```

