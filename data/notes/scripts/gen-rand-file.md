*2014-03-20*

How to generate a file of a defined size (ex. 100MB) with random content

Binary file:

    dd if=/dev/urandom of=file.dat bs=1M count=100

or

    dd if=/dev/urandom of=file.dat bs=1024 count=`echo $((100*1024))`

* `bs` -- block size in bytes

Text file:

    base64 /dev/urandom | dd of=file.txt bs=1M count=100 iflag=fullblock

* `base64` represents (encodes) binary data using printable ASCII characters

Empty file:

    dd if=/dev/zero of=file0.dat bs=1M count=100
