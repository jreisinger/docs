*No source* provides completely random numbers but some sources are more random
than others:

* intervals between keystrokes on keyboard
* rapidly changing areas of RAM
* network packets arrival times
* disk seek times (speed variations are caused by air turbulence)
* hardware random number generators (unstable free-running oscillators,
    radioactive decay, thermal noise from an amplified diode)

*Pseudorandom* numbers - if you know what number the computer just picked, you
can predict the next number.

Congruential generators
* lame but fast; so they are used often :-)
* algorithm:
    1. take the last random number (Nj)
    2. multiply it by something (A)
    3. add something (B)
    4. take the remainder when you divide by something else (C)

        Nj+1 = ANj + B(mod C)
* Every random number is derived by the previous random number except for the
    first, which comes from you. This first number is called the *seed* of the
    random number generator, and you provide it to Perl (and C) with `srand()`.

Source:
* Jon Orwant: Computer Science & Perl Programming (2002)
