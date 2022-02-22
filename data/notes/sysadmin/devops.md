*2016-11-20*

<img src="https://user-images.githubusercontent.com/1047259/152169337-a7181dcb-9c81-41c9-bbe2-68e153132ecd.png" title="DevOps">

Sysadmins have been traditionally separate from application developers and
project managers. The theory was that developers push products forward with new
features meanwhile change-resistant ops team would provide 24x7 management of
production environments. This arrangement often creates tremendous internal
conflict and fails to meet the needs of the business and its clients.

DevOps is a set of techniques to solve the chronic conflict between Development
an IT Operations. DevOps is an emerging field in operations. The practice of
DevOps typically appears in web applications and cloud environments, but its
influence is spreading to all parts of all industries.

The principles behind DevOps work patterns are the same principles that
transformed manufacturing. DevOps has been formed by:

* [Lean](https://www.amazon.com/Lean-Startup-Entrepreneurs-Continuous-Innovation/dp/0307887898/ref=sr_1_3?ie=UTF8&qid=1539069703&sr=8-3&keywords=lean+startup&dpID=51T-sMqSMiL&preST=_SY291_BO1,204,203,200_QL40_&dpSrc=srch) principles
* [Agile](http://agilemanifesto.org/) Community
* Infrastructure as Code
* CI, CD
* Innovation Culture
* Toyota Kata

# The Three Ways

"The Three ways" mentioned in The Phoenix Project book are the underpinning
principles that all the DevOps patterns can be derived from.

## The First Way - systems thinking

To maximaze left-to-right flow of work (from Dev to Ops to the customer) we need

* emphasize the performance of the entire system, as opposed to the performance of a specific silo of work or department
* each step is done in a repeatable way
* small batch sizes and intervals of work
* never passing defects to downstream work centers

Necessary practices:

* continuous build, CI, CD
* environments on demand
* limiting WIP
* building safe to change systems and organizations

## The Second Way - feedback

To prevent problems from happening again or enable faster detection and recovery we need:

* constant flow of fast feedback from right-to-left at all stages of the value stream

Practices:

* "stopping the production line" when builds/tests go wrong
* constantly elevating the improvements of daily work over daily work
* fast automated test suites
* shared goals and shared pain between Devs and Ops
* pervasive telemetry

## The Third Way - experimentation & repetition

To have a healthy work environment we need a culture that fosters:

* *continual* experimentation (risk taking and learning from success and failure)
* understanding that repetition and practice is the prerequisite to mastery (kata)

Practices:

* rewarding innovation and risk taking
* rewarding improvements
* at least 20% of Dev and Ops cycles allocated towards nonfunctional requirements

# CLAMS

> C[L]AMS: culture, [lean,] automation, measurement, and sharing. -- D. Edwards, J. Willis (2010)

Culture

* people work together and focus on the overall picture
* both devs and ops have on-call responsibilities
* automated testing is a must
* monitoring is a must
* all prod environments are mirrored by identical dev environments
* regular code and infrastructure reviews
* regular stand-up meetings

Lean

* use real time tools to communicate
* focus on solving component problems one at a time

Automation

* the most universally recognized aspect of DevOps
* if you need to perform a task more than twice, it should be automated
* don't automate what you don't understand
* allows brainpower to be used to solve new and more difficult challenges
* reduces the risk of human error
* you might need one, some, or all of tools like: Ansible, Salt, Jenkins, Packer, Terraform
* don't pick a tool and then look for the questions it answers
* reevaluate your tool set every year or two

Measurement

* tools like Icinga, Prometheus, Grafana, ELK
* having measurement data and doing something useful with it are two things
* metrics must be visible and evangelized to all interested parties

Sharing

* incentivize staff to share their work internally (presentations, how-to
  wiki articles) and externally (Meetups, blogs, conferences).


# DevOps without Devs

It's possible to apply these DevOps principles to any complex process:

1) Turn chaotic processes into repeatable, measurable ones.

* document process until consistent => automate => self-service
* to scale don't do IT tasks but maintain the system that does the IT tasks

3) Talk about problems and issues within and between teams.

* don't obscure or suffer through problems
* have channels for communicating problems and issues

4) Try new things, measure the results, keep the successes, and learn from the failures.

* create a culture of learning and experimentation
* use management practices that are blameless and reward learning

5) Do work in small batches so we can learn and pivot along the way.

* it's better to deliver some results each day than to hold back and deliver the entire result at the end
* you'll get feedback and fix problems sooner and avoid too much waisted effort

6) Drive configuration and infrastructure from machine readable sources kept under a source control system.

* IaC is flexible, testable and can benefit from software engineering techniques

7) Always be improving.

* always be identifying the next big bottleneck and experiment to fix it
* don't wait for a major disaster

8) Pull, don't push.

* build processes to meet demand, not accommodate supply
* determine the desired weekly output, allocate resources, and pull the work through the system to completion

9) Build community.

* you, your team, your company and world's IT community are interdependent

# Sources

* [The Phoenix Project](https://itrevolution.com/book/the-phoenix-project/)
* [TPoSaNA](http://the-sysadmin-book.com/)
* [TPoCSA](http://the-cloud-book.com/)
* ULSAH 5th

## More info

* https://www.oreilly.com/ideas/what-is-devops-yet-again
* http://sysadvent.blogspot.sk/2016/12/day-13-injecting-modern-concepts-into.html
