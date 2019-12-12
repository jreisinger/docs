# RAID

RAID (Redundant Array of Independent/Inexpensive Drives/Disks) -- a technology allowing to use two or more "disks" with the aim of

* read/write performance improvement (by striping data across multiple drives)
* fault-tolerance improvement (by replicating data across multiple drives)

RAID is __not__ a backup replacement!

## Creating RAID 1 (mirror)

Assemble the array

    mdadm --create --verbose /dev/md0 --level=mirror --raid-devices=2 /dev/sda /dev/sdb
    
Check the results

1. `cat /proc/mdstat`
2. `mdadm --detail /dev/md0`

## Creating RAID 5

Assemble the array (using partitions instead of raw disks for consistency)

    mdadm --create --verbose /dev/md0 --level=5 --raid-devices=3 /dev/sda1 /dev/sdb1 /dev/sdc1

## Simulating RAID 10 failure

Simulate a failed disk

    mdadm /dev/md0 -f /dev/sdc1
    
Check the array status

    cat /proc/mdstat
    
Remove the failed disk

    mdadm /dev/md0 -r /dev/sdc1
    
(System shutdown and physical replacement of the disk should follow in real failure).

Add the disk back

    mdadm /dev/md0 -a /dev/sdc1
    
Whatch the array rebuilding - `watch cat /proc/mdstats`
    
    Every 2.0s: cat /proc/mdstat                                               Tue Jul  2 09:55:13 2013

    Personalities : [raid10]
    md0 : active raid10 sdc1[2] sda1[0] sdd1[3] sdb1[1]
          312574976 blocks super 1.2 512K chunks 2 near-copies [4/3] [UU_U]
          [>....................]  recovery =  0.6% (1024896/156287488) finish=68.1min speed=37959K/sec

    unused devices: <none>

## Configuration

Via command line arguments or file (recommended)

    echo DEVICE /dev/sda /dev/sdb > /etc/mdadm.conf
    mdadm --detail --scan >> /etc/mdadm.conf
        
Enable array (at startup)

    mdadm -As /dev/md0
    
Stop array

    mdadm -S /dev/md0
    
## Monitoring

Set `MAILADDR` in mdadm.conf to get notified

Arrange for automatic start of `mdadm --monitor` at boot

    ubuntu# update-rc.d mdadm enable
    redhat# chkconfig mdmonitor on
