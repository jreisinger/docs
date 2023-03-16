Created: 2015-06-01

GRUB 2 vs. GRUB (Legacy)

* simpler code structure
* modular model
* internationalization (support for non-ASCII character sets)
* GUI interface

`/boot/grub/grub.cfg` still exists but is not supposed to be edited.

Files and directories

* `/boot/grub/grub.cfg` - automatically created central config file
* `/etc/default/grub` - general options (no menu entries)
* `/etc/grub.d` - helper scripts; you can add scripts here to add menu entries

GRUB 2 generally automates lot of steps (helper scripts search the partitions
for known OSs and create the required entries).

When GRUB starts it launches all executable scripts in `/etc/grub.d/`. If you modify scripts in this directory run `update-grub`.

## Adding a Menu Entry

Custom startup script to boot the System Rescue CD from the hard disk:

    01 echo "Modified entry for launching a Rescue CD as an Image" >&2
    02 cat << EOF
    03
    04 menuentry "Starting System Rescue CD from hard disc" {
    05         set root=(hd0,10)
    06         linux   /sysrcd/rescuecd subdir=sysrcd setkmap=de
    07         initrd  /sysrcd/initram.igz
    08 }
    09 EOF

* 05 - the 10th partition of the first hard disk is the root filesystem
    (partition numbering starts from 1, not from 0 as in GRUB Legacy)
* 06 - defines the kernel path (`root` in GRUB Legacy) with two parameters: `subdir` and `setkmap`
* 07 - defines the initial RAM disk path

## Booting from grub-rescue>

If you see `grub rescue>` that means it couldn't find `normal.mod`, so it probably couldn't find any of your boot files.

How does this happen? The kernel might have changed drive assignments or you moved your hard drives, you changed some partitions, or installed a new operating system and moved things around. In these scenarios your boot files are still there, but GRUB can't find them. So you can look for your boot files at the GRUB prompt, set their locations, and then boot your system and fix your GRUB configuration.

```
grub rescue> ls
grub rescue> set prefix=(hd0,1)/boot/grub
grub rescue> set root=(hd0,1)
grub rescue> insmod normal
grub rescue> normal
```

## Making Permanent Repairs

When you have successfully booted your system, run these commands to fix GRUB permanently:

```
# update-grub
# grub-install /dev/sdc
```

When you run `grub-install` remember you're installing it to the boot sector of your hard drive and not to a partition, so do not use a partition number like `/dev/sdc1`.

## Source

* Linux Magazine 111, February 2010
* https://www.linux.com/learn/how-rescue-non-booting-grub-2-Linux
