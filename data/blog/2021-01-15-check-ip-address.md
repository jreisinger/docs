# Check an IP address

Sometimes I come across an unknown IP address. This happens, for example, when I'm reviewing logs and I see that someone or (most probably) something was trying to SSH into the system. Or it was enumerating the URLs of a web application.

In such scenario I want to have a quick and easy way to check the IP address. I created a command line tool called [checkip](https://github.com/jreisinger/checkip) that does just that.

For example the following IP address definitely looks suspicious:

<img src="/static/checkip.png" style="max-width:100%;width:640px">

Of course, you can mix and match `checkip` with the standard shell tools. For example, to find out from where are people trying to use my web services:

```
cloud ~ [] $ for ip in \
>     $( \
>         # get logs since midnight
>         journalctl --since "00:00" | \
>         # filter WAF logs
>         grep waf | \
>         # filter IP addresses
>         perl -wlne '/((?:\d{1,3}\.){3}\d{1,3})/ and print $1' | \
>         # sort
>         sort | \
>         # deduplicate
>         uniq \
>     )
>     do
>         # get only geolocation
>         echo -ne "$ip\t"
>         checkip $ip | grep Geolocation
>     done
101.0.34.228	Geolocation Ludhiana | India | IN
101.108.180.255	Geolocation Chatuchak | Thailand | TH
103.117.20.162	Geolocation Cyberjaya | Malaysia | MY
104.144.171.75	Geolocation Buffalo | United States | US
104.160.11.93	Geolocation Washington | United States | US
```
