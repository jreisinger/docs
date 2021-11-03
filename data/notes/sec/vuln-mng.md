After locking down access control, probably the best investment to improve security.

There is overlap among vulnerability, patch and configuration management.

You should use cloud API for assets inventory to avoid missing new systems as they come online.

Cloud, CI/CD and microservice architectures have changed how vulnerability management is done:

1. Automatically pull security updates (for libraries, OS components) as part of normal development.
2. Test the updates as part of the normal application tests flow.
3. Deploy new version, automatically creating new environment that includes security updates (or security configuration changes).
4. Discover additional vulnerabilities in test or prod environments and add them as bugs in the development backlog.

Data access

* e.g. leaving resoruces open to the public, using poor credentials, not removing access for individuals who no longer need it

Application

* no matter how good your team is, your code is most probably going to have bugs and some of them will impact security
* e.g. buffer overflow (nowadays many languages make this bug difficult to make), injection attacks, XML external entity attacks, cross-site scripting attacks, deserialization attacks
* a WAF can act as a safety net but detecting and fixing vulns is the first and most important line of defense
* frameworks can help avoid vulnerabilities but can be source of security bugs

Middleware

* e.g. databases, application servers, message queues
* vulns here are attractive to attackers - they can use the same exploit against many apps often without understanding the app at all
* if provided as a service you still need to worry about the configuration issues

Operating system

* just as with the middleware/platform layer, perform configuration benchmarking (health checking, baselining) when deploying and then regularly afterwards
* turn off or remove any component that's not needed (hardening)
* many containers contain the userspace part (not the kernel)
* hypervisors - special purpose OSs designed to hold other OSs - also require patching and secure configuration (they are usually hardened)

Network

1. patch and config mng of network components (routers, switches, firewalls)
2. managing which communications are allowed

Virtualized and physical infrastructure

* responsibility of cloud provider (except if you are running a private cloud)
* e.g. missing BIOS/microcode updates, poor configuration of the baseboard controller allowing remote mngt

When finding and fixing vulns pick the most important area for your org, and get **value** from it before moving on to other areas. A common pitfall is having five different tools and processes none of which are actually providing a lot of value in finding and fixing vulns.

You want to plug tooling and findings leaks so you don't have a lot of unknown risk:

![image](https://user-images.githubusercontent.com/1047259/138861332-f43d5650-276a-4eb0-8333-defd920c7e7c.png)

The size of the pipes is determined by a number of the problems you (expect to) find in an area and how critical to the business those problems might be.

Network vulnerability scanners (Nexpose)

* don't look at SW components, simply make network requests and check for vulnerable versions or configurations
* should be allowed to scan every component, even if it means weakening the perimeter network controls
* you should incorporate vuln scan of test environment into the deployment process
* have a good process for masking false positives, or teams might start ignoring scan results

Agentless/agent-based scanners and config management (Nexpose)

* use credentials to get into the systems being tested
* can find vulns invisible to the network scanners
* agent-based have to be deployed and kept up to date and vuln in them can put entire infra at risk
* agentless must have inbound network access and privileged credentials which makes them attractive target for attackers

Container scanners (Anchore, AF X-ray)

* traditional scanners work well for VMs but often not for containers
* your either scan images and/or the runnning containers (using agent on each container host)

Dynamic application scanners (ZAP)

* DAST = Dynamic Application Security Testing
* network scanners run against network addresses, dynamic web application vulnerability scanners against specific URLs or REST APIs
* can find issues like CSS or SQLi
* often require credentials

Static application scanners (SonarQube)

* SAST = Static Application Security Testing
* look directly at the code you've written
* good candidate for running as part of deployment pipeline
* can spot hard to find security errors like memory leaks or off-by-one errors
* tend to have a high false positive rate, which can leadn to "security fatigue" in developers

Software composition analysis scanners

* SCA = Software Composition Analysis
* extension of static code scanners
* look primarily at the open source dependencies

Interactive application scanners

* IAST == Interactive Application Security Testing
* a bit of both static scanning and dynamic scanning
* IAST agent is deployed alongside application code

Runtime application self-protection scanners

* RASP = Runtime Application Self-Protection
* similar to IAST but designed to block attacks rather than just detect vulnerabilities
* just as IAST can degrade performance because more code is running in the production environment
* offer some of the same protection as a distributed WAF

Manual code reviews

* can be effective but expensive and time-consuming
* may be used only for sections of code with special significance to security (access control, encryption)

Penetration tests

* on staff pentesters, external suppliers

User reports

* you should consider what to do when you get a security report from your users

Risk mng process

* you need a system to prioritize vulnerabilities and decide what to do with them
* in many cases you might accept the risk as a cost of doing business
* a simple register with an agreed-upon process for assigning severity to the risks goes a long way

Vulnerability mng metrics

* metrics are useful but dangerous (some teams will naturally have a harder job to keep up with vulns)
* tool coverage
* mean time to remediate
* systems/apps with open vulnerabilities
* percentage of false positives/negatives
* vulnerability recurrence rate 

A sample microservice application:

![image](https://user-images.githubusercontent.com/1047259/138866404-c95fe79a-1eab-42e3-b92f-138dbb699cff.png)

Source: Practical Cloud Security (2019)
