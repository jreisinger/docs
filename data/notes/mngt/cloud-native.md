A [cloud native](https://github.com/cncf/toc/blob/main/DEFINITION.md)
application incorporate all we've learned about running networked applications
at scale over the past 60 years.

* 1950s - maiframe computers + dumb terminals
* 1980s - inexpensive network connected PCs
* 1990s - WWW and dot-com gold rush
* 2006 - AWS launched

Scalability

* application doesn't need to be refactored after an increase in demand
* it's often difficult to refactor a service for scalability, so building it
    with scalability in mind can save time and money longterm
* an application that doesn't maintain state, or it distributes its state
    between carefully designed replicas, will be relatively easy to scale
* vertical scaling - adding memory or CPU (limited by available computing resources)
* horizontal scaling - adding (or removing) service instances
