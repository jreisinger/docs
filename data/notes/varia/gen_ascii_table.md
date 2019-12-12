# Generate ASCII Table

I sometimes need to insert a table into a [blog](http://openhouse.sk/blog) post (written in [MarkDown](http://daringfireball.net/projects/markdown/syntax)). Having source data in `data.csv`:

    Line status (L1) ; Protocol status (L2) ; Interface status ; Typical root cause
    Administratively down ; Down ; disabled ; shutdown command
    Down ; Down ; notconnect ; cable problems, other device down
    up ; Down ; notconnect ; up/down state not expected on switch
    Down ; down (err-disabled) ; err-disabled ; port security disabled the interface
    Up ; Up ; connected ; interface working

.. I run [gen_ascii_table.pl](https://github.com/jreisinger/varia/blob/master/gen_ascii_table.pl):

    perl gen_ascii_table.pl --title "Switch Interface Status Codes" --orig | perl -p -e "s/(^[\|+'.])/\t\1/"

Then I just copy/paste the output into the blog entry.
