# Health

Hard disks failures typically stem from either platter surface defects (bad blocks) or mechanical failures.

```
lsblk
sudo umount /dev/sda1
sudo fsck -y /dev/sda1 # check and repair
sudo fsck -c /dev/sda1 # check for bad blocks
sudo mount /dev/sda1

# If your disk supports SMART
sudo smartctl -a /dev/sda1 | grep 'SMART support'
sudo smartctl -H /dev/sda1
```

# I/O Speed

`fio` is a CLI tool for measuring disk read and write speeds

```
cd /tmp
fio --name=${HOSTNAME} --filename=test.fio --randrepeat=1 --ioengine=libaio --direct=1 \
--gtod_reduce=1 --bs=4k --iodepth=64 --size=4G --readwrite=randrw --rwmixread=75
rm test.fio
```
