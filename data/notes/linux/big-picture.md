To study complex (software) systems we use *abstraction* - a fancy way to say that we
don't care about details of the components (subsystems, modules, packages,
...).

Levels of abstraction in a Linux system
=======================================

We need to organize the components somehow, however. In this case we create
groupings going from user to hardware:

User processes
* GUI
* servers
* shell

Kernel
* memory mngt.
* process mngt.
* syscalls
* device drivers

Hardware
* CPU
* RAM
* disks
* network ports

Difference between running kernel and user processes:
* code running in *user mode* has access only to a subset of memory and safe
    CPU operations
* code running in *kernel mode* has unrestricted access
* user and kernel modes are the processor's states (see "Main memory" below)

Main memory (RAM)
-----------------

Just a big storage area for a bunch of 0s and 1s (bits).

The running kernel and all user processes reside here. Also all I/O from peripheral devices flows through RAM. They are all just big collections of bits.

A CPU is just an operator on memory:

    CPU <-- reads instructions and data --- RAM
    CPU ---         writes data         --> RAM

State
* a particular arrangements of bits in RAM (three different states of 4 bits: 0001, 0101, 1000)
* as it can consists of millions of bits we use abstract terms to describe it, ex. process is waiting for input, and we use the term *image* for a particular physical arrangements of bits

Kernel
------

Nearly all kernel's tasks revolve around main memory:
* split memory into subdivisions
* keep state info about subdivisions
* make sure processes use only their subdivisions

Memory mngt.
* it is a complex task for the kernel - modern CPUs come with a help => memory
    management unit (MMU) using the virtual memory
* MMU interfaces processes' access to physical memory via memory address map
* kernel must still maintain the memory address map
* *page table* - name for the implementation of a memory address map

Process mngt.
* which processes are allowed to use CPU
* *multitasking* - (the appearance of) multiple processes running at the same time
* *context switch* - the act of one process giving up control of CPU to another process
 * the kernel is responsible for context switching
 * kernel runs between processes' time slices during a context switch
 * in *multi-CPU* systems kernel doesn't need (but usually does so anyway) to relinquish control of its current
    CPU in order to allow a different process to run on a different CPU

System calls a.k.a. *syscalls* - `man 2`
* the kernel's API
* feature of kernel allowing user processes to request specific actions, ex.
    opening, reading and writing files, creating new processes
* to execute a system call the kernel must temporarily switch to kernel mode,
    verify syscall's arguments and transfer data between user and kernel
    memory

(C library functions a.k.a. library calls - `man 3`)
* functions provided by standard C library (`glibc` on Linux)
* some library functions employ system calls, other perform tasks entirely within user space (ex. the string manipulation libraries)
* often designed to provide a more user-friendly interface than the underlying
    syscall, ex: `printf()` function provides output formatting and data
    buffering, whereas the `write()` syscall just outputs a block of bytes

All user processes (except for init) start as a result of `fork()` usually
   followed by `exec()`, ex. - running `ls` command in shell:

    shell ---> fork() ---> shell
                       |
                       +-> copy of shell ---> exec(ls) ---> ls
                           
* `exec()` is actually an entire family of syscalls for similar tasks

*psesudodevice* looks like a device but it's another kernel feature
  (implemented purely in software) - ex. `/dev/random`

Userspace and users
-------------------

* userspace - the main memory allocated by the kernel to user processes
* user - object for supporting permissions and boundaries
* group - a set of users used mainly for sharing files access

Devices
=======

Device files
------------

Block device (`b`)

* fixed total size
* split in fixed chunks
* easy to index
* processes have random access to any block in the device

Character device (`c`)

* no size
* works with data streams

Pipe device (`p`)

* like character device, with another process at the other end of the I/O stream instead of a kernel driver

Socket device (`s`)

* for interprocess communication
* often found outside of the `/dev` directory

`Major` and `minor` device numbers help the kernel to identify the device. Similar devices usually have the same major number.

Not all devices are represented as device files, e.g. network interfaces.

/dev and /sys/devices
---------------------

/dev
* traditional Unix way of representing devices as files
* convenient for user processes to reference and interface with devices supported by the kernel
* little information about devices
* the kernel assigns devices in the order in which they are found => may have different names between reboots

/sys/devices
* part of sysfs (a virtual FS provided by the kernel that exports information about kernel subsystems)
* system of files and directories (with symlinks)
* uniform view for attached devices based on their HW attributes
* to find the path and other attributes: `udevadm info --query=all --name=/dev/sda`

Resources
=========

* Brian Ward: How Linux Works, 2nd  Edition; No Starch Press 2014
* Michael Kerrisk: The Linux Programming Interface, No Starch Press 2010
