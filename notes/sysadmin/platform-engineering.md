# Why Platform Engineering is becoming essential

## Definition of terms

Plaftorm
- self-service API, tools, services, knowledge and support arranged as compelling internal product
- autonomous appliaction teams can use it to deliver product features at higher pace with reduced coordination
- it requires you to be doing platform engineering
- wiki page - no, there's no engineering
- "the cloud" - no, overwhelming array of offerings, too large to be seen as a coherent platform

Platform engineering
- discipline of developing and operating platforms
- goal: manage overall system complexity in order to deliver leverage to business
- how: curated product approach to develop software-based abstractions used by broad base of application developers

Leverage
- work of few engineers on platform team reduces work of the greater organization
- application engineers are more productive
- duplicated work accross engineering teams is eleminated

Product
- it's essential to view a platform as a compelling product
- this means we must take customer-centric approach when deciding on the platform features

## Current status: the over-general swamp

- public cloud and OSS have driven lot of industry change in last 25 years
- they make applications easier to build but harder to maintain
- to more your system grows, the slower you get -> you feel like walking through swamp
- [most](https://emauton.org/2016/03/22/a-software-lifecycle-update/) of the software cost is related to its upkeep, support and maintenance
- rather than reducing maintenance overhead cloud + OSS have amplified this problem
- because they provide and ever growing layer of promitives: general-purpose [building blocks](https://www.aboutamazon.com/news/company-news/amazon-ceo-andy-jassy-2023-letter-to-shareholders) that provide broad capabilities but are not integrated with one another
- to function they need **glue**: integration code, one-off automation, configuration, and management tools
- glue holds the blocks together but also creates stickiness making future changes hard
- the over-general swamp forms as the glue spreads
- each application team makes independent choices accross the array of primitives
- they select those that allow them easily build their own application
- they create whatever custom glue needed to hold everything together
- as this repeats over time, the company ends up with

_The over-general swamp, held together by glue_

<img width="651" alt="image" src="https://github.com/user-attachments/assets/fd0a954f-060c-4f9a-b1bb-5c0b8e0997c1" />

What's the problem with the sticky mess above?
- hard to understand
- difficult to change
- that's important because
  - applications are constantly evolving (new features or operational requirements)
  - cloud primitives + every OSS also undergo regular changes

> With the glue smeared everywhere, seemingly trivial updates to primitives (say, a security patch) require extensive organization-wide engineering time for integration and then testing, creating a massive tax on organizational productivity.

The way out of the swamp is to reduce the amount of glue by having "more boxes and fewer lines". Platform allows us to do so
- by abstracing over a limited set of OSS and cloud primitives in an opinated manner, specific to your organizational needs
- by encapsulation the underlying complexity and exposing usefula and easy to use interfaces

_How platforms reduce the amount of glue_

<img width="657" alt="image" src="https://github.com/user-attachments/assets/ee20ab72-e6d3-4b54-8670-8c83e640c4ea" />

# How we got stuck in the over-general swamp

Quick walk through last 25 years ...

## Explosion of applications and infrastructure

- Internet generated incredible demand for new software
- software (even serverless :-) runs on hardware
- compnies start buying more servers and network gear
- application teams interacting with infrastructure team often end in conflict
- public cloud came along providing APIs to the infrastructure
- PaaS (like Hereku) struggles to support wide range of applications and to integrate with existing applications and infrastructure
- most companies doing in-house software development embrace IaaS accepting the added complexity in order to get flexibility of choice
- Kubernetes attemps to simplify IaaS ecosystem by forcing applications to be "cloud native" and thus needing less infrastructure glue
- Kubernetes standardizes infrastructure but is not a complexity win
- still too much detailed configuration is needed, we just traded Terraform glue for YAML glue
- Kubernetes is also an example of the OSS proliferation of choice
- teams usually find and OSS solution optimal for them but not for others within the company
- to solution let's them quickly ship but initial launch eventually turns into burden

## How is going to operate all these applications and infrastructure

At first there were two professions
- software delevoper - architecture, coding, testing the software delivered as a monolith and handed over to SA
- system administrator - production operation of software on the company's computers

As the Internet took off and in-house software became more important these roles started to mutate.

DevOps - model to integrate application development and operations activities - has two on the ground implementations:
- Split: keep the separation but have operations team do some amount of development, particularly around creating glue for pushing code to production and integrating with infrastructure
- Merged: everyone who works on a system is on the same team ("you build it, you run it")

SRE
- lot of excitement but not a widespread susccess
- relied too much on the specific cultural capital that came from Google being the world's search engine

## Drowning in the swamp

> So you’ve got more application teams, making more choices, over a more complex set of underlying OSS and cloud primitives. Application teams get into this situation because they want to deliver quickly, and using the best systems of the day that fit the problem (or the systems they know best) will help them do that. Plus, if they’ve gotta own all the operational responsibility for this system themselves, they might as well pick their own poison!

> Add to this that application engineers with new features are not the only ones wanting to ship as quickly as possible. The increasing surface of internet-accessible systems has led to an escalation of cyberattacks and vulnerability discoveries, which in turn means that infrastructure and OSS are changing faster to address these risks. We’ve seen upgrade cycles for systems and components move from years to months, and these changes mean work for application teams who must update their glue and retest or even migrate their software in response.

> The pressure for change has created a swampy mess of glue mixed with the long-term consequences of individual team decisions. Every new greenfield project adds more choices and glue to this bog of complexity, and over time your developers get stuck in the mire. It’s hard to navigate, slow to move through, and full of hungry operational alligators (or worse, crocs!). How do you extract yourself from this morass? It’s no surprise that we think the answer is platform engineering, and next we will cover the ways in which it helps you do just that.

---

Source: Platform Engineering book by Camille Fournier and Ian Nowland (2024)
