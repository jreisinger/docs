*2015-06-14*

Taking stock of hardware
========================

Sources of hardware information:

    lscpu
    /proc/cpuinfo       # one entry for each core seen by the OS
    
    free -m
    /proc/meminfo
    
    hpacucli ctrl all show config [detail]  # HP physical RAID
    lsblk [-o KNAME,TYPE,SIZE,MODEL]        # installed on RHEL/CentOS
    parted -l                               # supports both MBR (msdos) and GPT
    fdisk -l                                # only traditional MBR partition table
    cat /proc/diskstats
    dmsetup ls                              # LVM device mapper
    udevadm info --query=all --name=/dev/sda

Desktop Management Interface (DMI, aka SMBIOS):

    dmidecode -t <type>    # see "DMI TYPES" in manpage

Network:

    ifconfig -a
    ip a s

CPU
===

Overall utilization
-------------------

Is CPU the bottleneck?

    $ vmstat 5 5 -S M    # in MB
    procs -----------memory---------- ---swap-- -----io---- -system-- ----cpu----
     r  b   swpd   free   buff  cache   si   so    bi    bo   in   cs us sy id wa
     1  0      0    230    687  44366    0    0  2923  3037    1    0  4  3 85  7
     0  0      0    218    687  44380    0    0 76160    10 2814 4233  3  1 96  0
     0  0      0    224    687  44377    0    0 79462     0 3253 5979  3  2 95  0
     0  0      0    230    687  44374    0    0 82432    18 3069 5674  3  1 95  0
     1  0      0    233    687  44372    0    0 86400    18 3705 5215  3  2 95  0

* first line reports averages since system's boot (the entire uptime), subsequent lines are averages within the previous sample period (default is 5 seconds)
* `r`  - runnable processes
* `b`  - processes blocked for I/O
* `in` - interrupts
* `cs` - context switches (number of times the kernel switches into kernel code; i.e. changing which process is running)
* `us` - user time (the percentage of time the CPU is spending on user tasks)
* `sy` - system (kernel) time
* `id` - idle time 
* `wa` - waiting for I/O

On multiprocessor machines, most tools present an average of processor statistics across all processors.

High `us` numbers generally indicate computation, high `sy` numbers mean that processes are doing lot of syscalls or I/O. A rule of thumb for a general server is that the system should spend 50% in user space and 50% in system space; the overall idle time should be 0.

Extremely high `cs` or `in` values typically indicate a misbehaving or misconfigured hardware device.

Load average
------------

How many pieces is the CPU divided into?

Average number of runnable (ready to run) processes:

    $ uptime 
     13:03:23 up 8 days, 13:06,  2 users,  load average: 1.13, 1.31, 1.38

* 5, 10, and 15-minute averages 
* processes waiting for input (from keyboard, network) are not considered
 ready to run - only processes that are actually doing something 
 (including wating for disk I/O) contribute to load average
* on a single-core system - 1 means that the CPU is exactly at capacity, i.e. the CPU has
 just enough to do all the time
* on a multi-core system - if number of cores = load average all cores are exactly at capacity
 
If your load average is high and your system still responds well, don't panic. The system just has a lot of processes sharing the CPU.

The system load average is an excellent metric to track as part of a system baseline. If you know your system’s load average on a normal day and it is in that same range on a bad day, this is a hint that you should look elsewhere (such as the network) for performance problems. A load average above the expected norm suggests that you should look at the processes running on the system itself.

Search for "/proc/loadavg" in `man 5 proc`.

http://www.brendangregg.com/blog/2017-08-08/linux-load-averages.html

Per process consumption
-----------------------

Which processes are hogging resources?

Snapshot of current processes:

    $ ps auxw    # BSD style options (other styles: Unix, GNU)

* `x` - show all your running processes
* `ax` - all processes on the system, not just yours
* `u` - more detailed info
* `w` - show full command names
* `m` - show threads

Processes and other system information regularly updated:

    $ top

* `z`, `x`   - turn on colors and highlight sort column
* `Spacebar` - update display immediately
* `M`        - sort by current resident memory usage
* `T`        - sort by total (cumulative) CPU usage
* `H`        - toggle threads/processes display
* `u`        - display only one user's processes
* `f`        - select statistics to display

On a busy system, at least 70% of the CPU is often consumed by just one or two processes. Deferring the execution of the CPU hogs or reducing their priority makes the CPU more available to other processes.

How much CPU time a process uses:

    $ time ls    # or /usr/bin/time

* user time - time the CPU spent running the program's *own* code
* system time - time the CPU spent running kernel code doing the process's work (ex. reading files or directories)
* real/elapsed time - total time it took to run the process, including the time
 the CPU spent running other tasks

Threads
-------

Some processes can be divided into pieces called *threads*:

* very similar to processes: have TID, are scheduled and run by the kernel
* processes don't share system resources
* all threads inside a single process share system resources (I/O connections, 
    memory)

Many processes have only one thread - *single-threaded* processes (usually
called just processes).

All processes start out single-threaded. This starting thread is called *main
thread*. The main thread then starts new threads in similar fashion a process
calls `fork()` to start a new process.

Threads are useful when process has a lot to do because threads can run
simultaneously on multiple processors and *start faster* than processes and
*intercommunicate more efficiently* (via shared memory) than processes (via
network connection or pipe).

It's usually not a good idea to interact with individual threads as you would
with processes.

Memory
======

See also [posts/linux-ate-my-memory](https://github.com/jreisinger/blog/blob/master/posts/linux-ate-my-memory.md).

Amount of paging (swap) space that's currently used:

    # swapon -s
    Filename                Type        Size    Used    Priority
    /dev/sdb2               partition   7815616 0       -1

* in kilobytes

`vmstat` (see above) fields:

* `si` - swapped in (from the disk)
* `so` - swapped out (to the disk) => if your system has constant stream of page outs, buy more memory

Storage I/O
===========

    $ iostat 5 5 [-d sda]
    Linux 3.2.0-4-amd64 (backup2)   06/14/2015  _x86_64_    (16 CPU)
    
    avg-cpu:  %user   %nice %system %iowait  %steal   %idle
               3.80    0.34    3.17    7.49    0.00   85.20
    
    Device:            tps    kB_read/s    kB_wrtn/s    kB_read    kB_wrtn
    sdb              49.61      1852.45       349.64 1369392967  258461851
    sdc             301.74     21510.91     24545.93 15901546498 18145130448
    sdd              75.02      6184.17      6195.25 4571531985 4579724644
    sda             307.37     16906.94     17127.65 12498149921 12661307662
    dm-0            131.14      8082.58      9533.25 5974897325 7047285056
    dm-1            172.96     13428.25     15012.67 9926593437 11097845392
    dm-2            107.96      1612.16       347.05 1191762057  256547336

* the first report provides statistics since the system was booted, subsequent reports cover the time since the previous report
* `tps` - number of transfers per second (IOPS?)
* `kB_read/s` - average number of kilobytes read per second
* `kB_read` - total kilobytes read

Processes using file or directory on `/usr` filesystem (mount point):

    $ sudo fuser -cv /usr
                         USER        PID ACCESS COMMAND
    /usr:                root     kernel mount /
                         root          1 .rce. init
                         root          2 .rc.. kthreadd

.. ACCESS:
* `f`,`o`  - the process has a file open for reading or writing
* `c`      - the process's current directory is on the filesystem
* `e`, `t` - the process is currently executing a file
* `r`      - the process's root directory (set with `chroot`) in on the filesystem
* `m`, `s` - the process has mapped a file or shared library

List all open files:

    $ lsof    # pipe output to pager, grep or use options
    
Network I/O
===========

To see info on network connections:

    # netstat -tulpan
    
* `-t` - print TCP ports info
* `-u` - print UDP ports info
* `-l` - print listening ports
* `-p` - print name and PID of the program owning the socket
* `-a` - print all active ports
* `-n` - don't reverse-resolve IP addresses
* `Recv-Q` and `Send-Q` columns show the number of bytes in the sockets' read and write buffers, respectively.

To list all programs using or listening to ports (when run as regular user, only shows user's processes):

    # lsof -i -nP

* `-i` - list all Internet and x.25 (HP-UX) network files
* `-n` - don't reverse-resolve IP addresses
* `-P` - disable /etc/services port name lookups

To list Unix domain sockets (not to be confused with network sockets although similar) currently in use on your system:

    # lsof -U    # unnamed sockets have "socket" in NAME column

lsof network connections filtering
----------------------------------

by protocol, host and port:

    lsof -i[<protocol>@<host>]:<port>

.. ex.

    lsof -i:22
    lsof -iTCP:80

by connection status:

    lsof -iTCP -sTCP:LISTEN

Other tools
===========

iotop - per process I/O utilization and monitoring

    $ iotop

pidstat - per process monitoring

    $ pidstat -p <PID> <interval>
    
dstat - versatile replacement for vmstat, iostat and ifstat

    $ dstat
    
sar - record resource utilization over time

Resources
=========

* ULSAH, 4th, Ch. 29 
* How Linux Works, 2nd, Ch. 8
* Corresponding `man` pages

