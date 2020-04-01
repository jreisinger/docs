(Ubuntu) packages containing basic tools

```
apt-get update && apt-get install procps net-tools vim
```

common `date` formats

```
        Format/result         |       Command              
------------------------------+----------------------------
YY-MM-DD_hh:mm:ss             | date +%F_%T                
YYMMDD_hhmmss                 | date +%Y%m%d_%H%M%S        
YYMMDD_hhmmss (UTC version)   | date --utc +%Y%m%d_%H%M%SZ 
YYMMDD_hhmmss (with local TZ) | date +%Y%m%d_%H%M%S%Z      
YYMMSShhmmss                  | date +%Y%m%d%H%M%S         
YYMMSShhmmssnnnnnnnnn         | date +%Y%m%d%H%M%S%N       
Seconds since UNIX epoch:     | date +%s                   
Nanoseconds only:             | date +%N                   
Nanoseconds since UNIX epoch: | date +%s%N                 
ISO8601 UTC timestamp         | date --utc +%FT%TZ         
ISO8601 Local TZ timestamp    | date +%FT%T%Z              
```

Vim config options I often use

```
set tabstop=4 shiftwidth=4 expandtab
set nofoldenable
set textwidth=0
syntax off
```

find files within Vim

```
grep -iR what .
cw
```

convert from hex to decimal

```
perl -le 'print hex "0xAf"'
perl -le 'print hex   "aF"' # same
```
