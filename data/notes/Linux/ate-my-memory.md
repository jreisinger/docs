Linux ate my memory!
====================

Linux uses free memory for *disk* caching and buffering but this memory can be
allocated to applications immediately if needed.

    $ free -m
                 total       used       free     shared    buffers     cached
    Mem:          3965       2821       1144          0         96       1038
    -/+ buffers/cache:       1686       2279
    Swap:         1903        454       1449

So the machine above has `1144` MB of memory free from the kernel point of view
since Linux considers memory used for caching/buffering as used. However it has `2279` MB
of memory free from the applications point of view because programs don't
consider memory used by caches and buffers.

To show this, let's run this memory eating application (`munch.c`):

    #include <stdlib.h>
    #include <stdio.h>
    #include <string.h>
    
    int main(int argc, char** argv) {
        int max = -1;
        int mb = 0;
        char* buffer;
    
        if(argc > 1)
            max = atoi(argv[1]);
    
        while((buffer=malloc(100*1024*1024)) != NULL && mb != max) {
            memset(buffer, 0, 100*1024*1024);
            mb++;
            printf("Allocated %d MB\n", 100*mb);
            sleep(1);
        }
    
        return 0;
    }

To see it in action:

    sudo swapoff -a                 # turn off swapping; can take a while
    watch free -m                   # in a different shell
    dd if=/dev/urandom of=/tmp/bla  # cache some data
    ./munch                         # run until SIGINT or killed by kernel
    
    # Don't forget to cleanup ... :-)
    rm /tmp/bla
    sudo swapon -a

See [mem-muncher](https://github.com/jreisinger/mem-muncher) for a dockerized version of the `munch` program.

Resources
=========

* http://www.linuxatemyram.com/

