After locking down access control, probably the best investment to improve security.

There is overlap among vulnerability, patch and configuration management.

You should use cloud API for assets inventory to avoid missing new systems as they come online.

Cloud, CI/CD and microservice architectures have changed how vulnerability management used to be done:

1. Automatically pull security updates (for libraries, OS components) as part of normal development.
2. Test the updates as part of the normal application tests flow.
3. Deploy new version, automatically creating new environment that includes security updates (or security configuration changes).
4. Discover additional vulnerabilities in test or prod environments and add them as bugs in the development backlog.

Data access

* e.g. leaving resoruces open to the public, using poor credentials, not removing access for individuals who no longer need it

Application

* no matter how good your team is, your code is most probably going to have bugs and some of them will impact security
* e.g. buffer oveflow (nowadays many languages make this bug difficult to make), injection attacks, XML external entity attacks, cross-site scripting attacks, deserialization attacks
* a WAF can act as a safety net but detecting and finxing vulns is the first and most important line of defense
* frameworks can help avoid vulnerabilities but can be source of security bugs

Middleware

* e.g. databases, applications servers, message queues
* vulns here are attractive to attackers - they can use the same exploit against many apps often without understanding the app at all
* if provided as a service you still need to worry about the configuration issues

Source: Practical Cloud Security (2019)
