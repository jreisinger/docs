* [ZFS](#zfs)
  * [Pools](#pools)
  * [Filesystems](#filesystems)
  * [Arrays](#arrays)
  * [Encrypted backups with snapshots (on external HDD)](#encrypted-backups-with-snapshots-on-external-hdd)

# ZFS

* refererred to as a filesystem but it's a comprehensive storage management (LVM, RAID)
* could not be included into the Linux kernel due to licence terms (although it's open source)
* Ubuntu 16.04 included it in the form of a loadable kernel module
* ZFS secretly writes a GPT-style partition table and allocates all disks' space to its first partition
* organized around *copy-on-write* principle

![ZFS architecture](https://www.safaribooksonline.com/library/view/unix-and-linux/9780134278308/image/ZFSArchitecture.png)

## Pools

* ~ volume group
* composed of virtual devices - raw storage (disks, partitions, SAN), mirror groups, RAID arrays

Create a pool (add a disk):

```
zpool create mypool sdb
```

* disk was labeled
* `mypool` pool was created
* filesystem root inside `mypool` was created
* filesystem was mounted as `/mypool` (will be remounted automatically on (re)boot)

See what have we created:

```
zpool list
zpool status
```

## Filesystems

* all filesystems living in a pool can draw from pool's available space
* unlike traditional filesystems which are independent of each other, hierarchically dependent (property inheritance)
* automounted as soon as created

Create a filesystem

```
zfs create mypool/myfs
zfs list -r mypool    # -r -- recurse through child filesystems
```

Change default mount point (a property) of the root filesystem

    zfs set mountpoint=/opt/mypool mypool
    zfs get all mypool/myfs    # filesystem properties

Snapshots

```
touch /opt/mypool/myfs/file
zfs snapshot mypool/myfs@friday
rm /opt/mypool/myfs/file
ls /opt/mypool/myfs/.zfs/snapshot/friday
zfs rollback mypool/myfs@friday  # can only revert FS to its most recent snapshot
```

* copy-on-write brought to the user level (just as in LVM)
* per-filesystem not per-volume
* comptele identifier: `<filesystem>@<snapshot>`
* read-only
* not true filesystems however can be turned into one:

```
zfs clone mypool/myfs@friday mypool/myfs_clone
```

## Arrays

* ZFS's RAID-Z ~ RAID 5

Adding (five) disks:

```
zpool destroy mypool
zpool create mybigpool raidz1 sdb sdc sdd    # raidz<parity>
zpool add -f mybigpool mirror sde sdf
zpool status mybigpool
```

See [ULSAH](https://www.safaribooksonline.com/library/view/unix-and-linux/9780134278308/Storage.xhtml) for more.

## Encrypted backups with snapshots (on external HDD)

Setup external disk (once)

```
dd if=/dev/zero of=/dev/sdc bs=1M count=10
zpool create extusb /dev/sdc
zfs create extusb/backup
encfs /extusb/backup/.encrypted /extusb/backup/decrypted
```

Mount the disk

```
#sudo zpool import [-f] extusb # not needed
sudo /etc/init.d/zfs-fuse restart
sudo zpool list # you should see the 'extusb' pool
sudo encfs /extusb/backup/.encrypted /extusb/backup/decrypted
sudo ls -l /extusb/backup/decrypted/
```

Backup data

```bash
#!/bin/bash
#
# zfs-backup.sh

# Remote host
RHOST=host.org
RPORT=22
RUSER=$USER

# Local user ssh key
SSHKEY=$HOME/.ssh/id_rsa

# Make sure decrypted backups are mounted.
mount | grep /extusb/backup/decrypted > /dev/null
EV=$?
if [[ $EV -ne 0 ]]; then
        echo "Backups not running, because encrypted FS is not mounted. Run:"
        echo
        echo "    encfs /extusb/backup/.encrypted /extusb/backup/decrypted"
        exit 1
fi

# Rsync data.
rsync --quiet --delete -az \
        --exclude 'public' \
		--rsync-path="sudo rsync" \
		--rsh "ssh -i $SSHKEY -p $RPORT -l $RUSER" \
        $RHOST:/data \
        /extusb/backup/decrypted/alarm/

# Create snapshot with a timestamp.
zfs snapshot extusb/backup@`date +%F_%T`
```

```
sudo ./zfs-backup.sh
```

Check backups

```
sudo zfs list -t snapshot
```

Unmount the disk

```
sudo fusermount -u /extusb/backup/decrypted  # encfs
sudo umount /extusb/backup                   # zfs
sudo umount /extusb                          # zfs root
sudo /etc/init.d/zfs-fuse stop
```

Restore data (once)

```
zfs clone extusb/backup@2015-03-13 extusb/2015-03-13
encfs /extusb/2015-03-13/.encrypted /extusb/2015-03-13/decrypted/
#### take the files you need from /extusb/2015-03-13/decrypted/
fusermount -u /extusb/2015-03-13/decrypted
zfs destroy extusb/2015-03-13
```

Cleanup (once)

```bash
# list snapshots
zfs list -t snapshot

# remove 2017 backups
zfs list -t snapshot -o name | grep backup@2017 | tac                                # check
zfs list -t snapshot -o name | grep backup@2017 | tac | xargs -n 1 zfs destroy -r    # remove      
```
