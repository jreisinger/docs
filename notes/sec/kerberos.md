* deals with authentication (like PAM); doesn't provide any additional security or encryption beyond that
* it's an authentication method (unlike PAM which is an authentication framework)
* PAM and Kerberos are generally used together
* a trusted server performs authentication for en entire network
* if you authenticate successfully the Kerberos service issues cryptographic credentials (tickets) which you present to other services as evidence of you identity

Improvements upon traditional password security

* unencrypted passwords are never transmitted over network
* provides a Single sign-on

Source: ULSAH, 4th

More: http://web.mit.edu/kerberos/www/dialogue.html
