A [cloud native](https://github.com/cncf/toc/blob/main/DEFINITION.md) application incorporates all we've learned about running networked applications at scale over the past 60 years.

* 1950s - maiframe computers + dumb terminals
* 1980s - inexpensive network connected PCs
* 1990s - WWW and dot-com gold rush
* 2006 - AWS launched

Cloud native techniques and technologies exist for no other reasons than to make it possible to leverage the benefits of a “cloud” (quantity) while compensating for its downsides (lack of dependability).

Scalability

* application doesn't need to be refactored after an increase in demand
* it's often difficult to refactor a service for scalability, so building it with scalability in mind can save time and money longterm
* an application that doesn't maintain state, or it distributes its state between carefully designed replicas, will be relatively easy to scale
* vertical scaling - adding memory or CPU (limited by available computing resources)
* horizontal scaling - adding (or removing) service instances

Loose coupling

* a system's components have minimal knowledge of any other components
* changes to one component generally don't require changes to the other
* example: web servers and web browsers

Resilience

* how well a system withstands and recovers from errors and faults

Manageability

* it's possible to sufficiently alter its behavior without having to alter it's code

Observability

* metrics, logging and tracing
* but their existence is not enough - data is not information

IaaS vs PaaS vs SaaS

<img src="https://user-images.githubusercontent.com/1047259/136159426-797502ab-ef09-4739-9d31-2ae1a7bed71d.png" style="max-width:100%;height:auto;"> 

Source: Cloud Native Go (2021), Practical Cloud Security (2019)
