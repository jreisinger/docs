Security problems at the Internet level started in 1988 with "Morris worm".

A systems administrator should:

* make sure systems are up to date
* make sure systems are backed up
* make sure systems are monitored
* make sure users are properly educated
* keep in touch with current security technology
* hire professionals to help with problems that exceed his knowledge

No system is secure, you can only make it more secure.

Security is a process, it's never done.

You can't have both security and convenience.

Common areas of security compromise:

* social engineering
* software vulnerabilities
```perl
#!/usr/bin/perl
# Example of user input validation error. 
# What if a user enters ../../../etc/passwd?

open HTMLFILE, "/var/www/html/$ARGV[0] or die "trying\n";
while (<HTMLFILE>) { print }
close HTMLFILE;
```
* configuration errors (not-so-securely is often a default software configuration)

Attack phases

* Reconnaissance
* Scanning
* Gaining access
* Maintaining access
* Clearing tracks

Sources:

* ULSAH (2013), Ch. 22 - Security
