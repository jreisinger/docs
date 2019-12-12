When a Unix program finishes, it leaves an exit code (stored in `$?` special
shell variable) for its parent process. The exit code is a small integer and
it's also called *error code* or *exit value*.

Zero *always* means success, non zero exit status usually means error.

Exit code can be used in conditionals like this:

    ## 'x' is used to get around empty $1
    $ if [ x"$1" == x"hi" ]; then echo 1st arg was hi; else echo where\'s hi?; fi
    where's hi?

    $ ls / > /dev/null && echo ok || echo error
    ok
    
    $ ls-bad / > /dev/null 2&>1 && echo ok || echo error
    error

    $ grep root /etc/passwd > /dev/null && echo ok || echo error
    ok
    
    $ if grep -q bla /etc/passwd; then echo ok; else echo error; fi
    error

The shell runs the command after the `if` keywords. If the command exits with
0, the `then` clause commands are executed, else `else` clause commands are
run.

Note that some programs (like `grep` and `diff`) use non zero exit codes to
indicate normal conditions:

    ## grep returns 0 if it finds a match and 1 if it doesn't
    $ grep blabla /etc/passwd > /dev/null && echo ok || echo 'error (not really)'
    error (not really)

Search for "exit" or "diagnostics" in man pages if not sure.

Resources:

* How Linux Works, 2nd
