# Intro

* `systemd` provides a central logging for all kernel and userland processes
* the system that collects and manages logs is called the journal
* the journal is implemented with the `journald` daemon
* `journald` daemon collects data from all available sources and stores them in binary format
* `systemd` journal can either be used with an existing `syslog` implementation, or it can replace the `syslog` functionality
* journal data are accessed and manipulated using the `journalctl` tool

## System time

* the `systemd` suites comes with a tool `timedatectl` that controls the system time and date

```
timedatectl list-timezones
timedatectl set-timezone ZONE
timedatectl status
```

# Logs viewing

```
journalctl --utc # display timestamp in UTC
journalctl -b    # all entries since the most recent boot
journalctl --list-boots
journalctl -b -1 # entries from the previous boot
```

## Display kernel messages

```
journalctl -k # or --dmesg
journalctl -k -b -2
```

## Filter by time

Format for absolute time values:

```
YYYY-MM-DD HH:MM:SS # some parts can be left off
```

```
journalctl --since "2019-12-24" --until "2019-12-31 23:59"
```

Relative values:

```
yesterday, today, tomorrow, now, ago, ...
```

```
journalctl --since yesterday
journalctl --since 09:00 --until "1 hour ago"
```

## Filter by unit

```
systemctl list-unit-files --all
journalctl --field _SYSTEMD_UNIT

journalctl -u nginx.service
journalctl -u nginx.service -u docker.service --since today
```

## Filter by process, user or group

```
man systemd.journal-fields

journalctl _PID=1

id -u http         # get numeric id of user http
journalctl _UID=33 --since today

journalctl -F _GID # list group IDs journal has entries for
```

## Filter by priority

```
journalctl -p err -b # error level and above
```

The journal implements the standard `syslog` message levels:

```
0: emerg
1: alert
2: crit
3: err
4: warning
5: notice
6: info
7: debug
```

# Modifying output

```
journalctl --no-full  # no right arrow possibility
journalctl -a         # also unprintable characters

journalctl --no-pager # no less

journalctl -o json
journalctl -o json-pretty

journalctl -n 20      # last 20 entries
journalctl -f         # like tail -f
```

# Journal maintenance

```
journalctl --disk-usage

journalctl --vacuum-size=1G     # remove old entries to shrink journal size
journalctl --vacuum-size=1years # keep only entries from last year

vi /etc/systemd/journald.conf
```

Source

* https://www.digitalocean.com/community/tutorials/how-to-use-journalctl-to-view-and-manipulate-systemd-logs
