> Hope is not a strategy. -- SRE saying

Two approaches to run (operate) a complex computing system: sysadmin, SRE.

# Sysadmin approach

* traditional way
* assemble and deploy SW components and then run them
* running means sysadmins are responding to events and updates as they occur

Advantages

* relatively easy to implement for companies (familiar industry paradigm, lots
  of examples to emulate, wide talent pool, existing tools and consultants)

Disadvantages

* team size scales linearly with the load generated by the system
* you end up with two different teams (product developers and product
  operators) that often live in conflict

# SRE approach

* invented by Google (planet-scale products)
* hire software engineers to run products i.e. to do operations
* they create SW systems that do the work otherwise done by sysadmins

Advantages

* lower costs (less people needed than with sysadmin approach)
* no dev/ops split dysfunctionality
* rapid innovation

Disadvantages

* it's hard to find SREs - you are competing with product development teams,
  high hiring bar requirements (coding skills, Unix internals, L1 - L3
  networking), relatively new and unique discipline
* SRE team's unorthodox approach to service management requires strong
  management support

# DevOps or SRE?

* DevOps term is older - coined in 2008
* DevOps is a generalization of several SRE core principles to a wider range of ogranizations, management structures and personnel
* SRE is a specific [implementation of DevOps](https://cloud.google.com/blog/products/gcp/sre-vs-devops-competing-standards-or-close-friends) with some idyosyncratic extensions 

# Source

* Site Reliability Engineering (2016)
