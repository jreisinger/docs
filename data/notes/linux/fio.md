`fio` is a CLI tool for measuring disk read and write speeds

```
cd /tmp
fio --name=${HOSTNAME} --filename=test.fio --randrepeat=1 --ioengine=libaio --direct=1 \
--gtod_reduce=1 --bs=4k --iodepth=64 --size=4G --readwrite=randrw --rwmixread=75
rm test.fio
```
