Security problems at the Internet level started in 1988 with "Morris worm".

A systems administrator should:

* make sure systems are secured
* make sure systems are monitored
* make sure users are properly educated
* keep in touch with current security technology
* hire professionals to help with problems that exceed his knowledge

No system is secure, you can only make it more security. Security is a process
though.

You can't have both security and convenience.

Common areas of security compromise:

* social engineering
* software vulnerabilities

        #!/usr/bin/perl
        # Example of user input validation error

        open HTMLFILE, "/var/www/html/$ARGV[0] or die "trying\n";
        while (<HTMLFILE>) { print }
        close HTMLFILE;

    What if a user enters `../../../etc/passwd`?
* configuration errors (not-so-securely is often a default software configuration)

Attack phases

* Reconnaissance
* Scanning
* Gaining access
* Maintaining access
* Clearing tracks

Sources:

* ULSAH, Ch. 22 - Security
