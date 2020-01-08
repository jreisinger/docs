There are several [steps](https://calleerlandsson.com/the-four-stages-of-compiling-a-c-program/) when compiling a C program:

1. Preprocessing - lines starting with `#` (include files, macros, conditionals) are interpreted by the preprocessor (`cpp` - but you rarely need to run it by itself)
2. Compilation - preprocessed code is translated to assembly instructions specific to the target CPU architecture
3. Assembly - assembly instructions are translated to machine code, or *object code*
4. Linking - object code can be understood by the CPU but some pieces of the program are out of order or missing

You need libraries to build complete programs. C library consists of *precompiled* functions that can be built into your program.

The two standard directories on Linux containing libraries are `/lib` and
`/usr/lib`. `/lib` should not contain static libraries.

Static libraries
----------------

When you link a program against a static library, the linker copies machine
code from the library file into your executable.

Advantages:
* the final executable does not need the original libraries to run
* the executable's behavior never changes

Disadvantages:
* waste of disk space and memory
* you need to recompile the executable if a library is found inadequate or
    insecure

Shared libraries
----------------

When you run a program linked against a shared library, the system loads the
library's code into the process memory space only when necessary.

Advantages:
* many processes can share the same library code in memory
* if you need to (slightly) modify the library code, no executable 
    recompilation is (usually) needed 

Disadvantages:
* library management
* missing libraries

Executables know just the names of shared libraries (performance and
flexibility reasons). A small program `ld.so` (runtime dynamic linked/loader)
finds and loads shared libraries for a program at runtime. To list the needed
shared libraries:

    $ ldd /bin/bash
        linux-vdso.so.1 =>  (0x00007ffe221ce000)
        libtinfo.so.5 => /lib64/libtinfo.so.5 (0x00007f2fb72df000)
        libdl.so.2 => /lib64/libdl.so.2 (0x00007f2fb70db000)
        libc.so.6 => /lib64/libc.so.6 (0x00007f2fb6d46000)
        /lib64/ld-linux-x86-64.so.2 (0x00007f2fb7506000)

* `<library name> => <where to find it>`
* the last line is the location of `ld.so`
* `so` - shared object

Header (include) files and directories
--------------------------------------

C header files - additional *source code* files that usually contain type and function declarations (ex. `stdio.h`).

Most paths that contain header files have "include" somewher in the name. The default include directory in Unix is `/usr/include`. If you want the compiler to look into different directory:

    $ cc -c -I/usr/junk/include badinclude.c
    
Double quotes (`#include "myheader.h"`) instead of angle brackets (`#include <stdio.h>`) mean that the header file is not in the system include directory - it's often in the same directory as the source file.

More
* http://wiki.reisinge.net/Linux1/InstalaciaLinuxuSpravaBalikov/ZdielaneKniznice
* How Linux Works, 2nd: Ch. 15
