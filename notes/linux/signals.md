*2015-10-20*

Signal -- a message from the kernel (or a process) to a process. Used for:

* errors (kernel saying: "You can't touch that area of memory!")
* events (death of a child, interrupt with Ctrl-C)

To ask the kernel to a send a signal:

``` shell
kill [-SIGNAL] PID  # default signal is TERM
```

Selected signal types:

* TERM (15) - terminate a process (polite request to die, i.e. can and should be caught)
* KILL (9) - terminate a process at the kernel level and remove it forcibly from memory (cannot be caught, blocked or ignored)
* INT (2) - interrupt, sent by the terminal driver on `Ctrl-C`. Simple programs usually just die, more important ones (ex. shells, editors) stop long-running operations.
* STOP - freeze the process (stays in memory ready to continue where it left
    off)
* CONT - continue running the STOPed process
* QUIT - similar to TERM but generates a core dump if not caugh (`Ctrl-\`)
* CHLD - one of the child processes stopped running - or, more likely, exited
* [HUP](http://world.std.com/~swmcd/steven/tech/daemon.html)
* TSTP - pressing `Ctrl-Z` sends TSTP to a process

Each process has a default disposition (what to do) for each possible signal. You may install your own handler or otherwise change the disposition of most signals. Only SIGKILL and SIGSTOP cannot be changed. The rest you can:

* ignore
* block (blocked signal is pending until it is later unblocked, i.e. removed from the signal mask)
* catch (trap)

This is how you can catch signals in Perl:

``` perl
#!/usr/bin/env perl
#
# signal-catcher -- send me a signal, e.g.:
#
# $ kill -2 <my-pid>
#
use 5.014;    # includes strict
use warnings;
use autodie;

our $shucks = 0;            # shuck - škrupina, šok?

sub catch_zap {             # zap - šleha?
    my $signame = shift();
    $shucks++;
    die "Somebody sent me a SIG$signame!";
}

$SIG{INT}   = \&catch_zap;
$SIG{QUIT}  = \&catch_zap;  # catch another signal, too

$|++;
print "Going to sleep ";
while (1) {
    print ".";
    sleep 1;
}
```

Catching a signal in Python:

``` python
#!/usr/bin/env python

import sys, signal, time

def now(): return time.ctime(time.time())       # current time string

def onSignal(signum, stackframe):               # python signal handler
    print('Got signal', signum, 'at', now())

signum = int(sys.argv[1])                       # from the cmd line

signal.signal(signum, onSignal)                 # install signal handler
while True: signal.pause()                      # wait for signals
```

Handling [signals in Go](https://gobyexample.com/signals).

Source:

* How Linux Works, 2nd
* Perl Cookbook, 2nd
* The Linux Programming Interface
* `man 7 signal` - overview of signals
* Programming Python
