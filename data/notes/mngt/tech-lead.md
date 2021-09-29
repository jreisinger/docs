# The Main Roles of a Tech Lead

_Source: The Manager's Path (2017)_

Let’s say you’re partnering with a product manager and a team of four other engineers on a big multiweek effort to launch a new initiative. The tech lead has a number of responsibilities in this scenario, depending on where you are in the project lifecycle. Sure, you’ll need to write some code and make some technical decisions. But that’s only one of the roles you’ll play, and it’s likely not even the most important one.

## Systems architect and business analyst

In the systems architect and business analyst roles, you identify the critical systems that need to change and the critical features that need to be built in order to deliver upcoming projects. The goal here is to provide some structure for basing estimates and ordering work. You need not perfectly identify every single element of a project, but there’s a lot of value in spending time thinking through the externalities and issues related to a project. This role requires you to have a good sense of the overall architecture of your systems and a solid understanding of how to design complex software. It probably also requires you to be able to understand business requirements and translate them into software.

## Project planner

Project planners break work down into rough deliverables. With this hat on, you’re learning to find efficient ways of breaking down the work so that the team can work quickly. Part of the challenge here is getting as much productive work done in parallel as possible. This can be tough because you are probably used to thinking about only your own work, instead of the work of groups of people. Finding places to apply agreed-upon abstractions to enable parallel work is key. For example, if you have a frontend that consumes JSON objects from an API, there should be no need for the API to be completely finished for the frontend development to begin. Instead, agree upon the JSON format and start to code to that format using dummy objects. If you are lucky, you’ve seen this happen before and are simply pattern-matching your previous work. At this stage, you will want to gather input from the experts on your team, and talk to the people who know the affected parts of the software deeply, so that they can help with the details here. You will also want to start identifying priorities as part of this process. Which pieces are critical, and which are optional? How can you work on the critical items early in the project?

## Software developer and team leader

Software developers and team leaders write code, communicate challenges, and delegate. As projects move forward, unexpected obstacles arise. Sometimes tech leads are tempted to go to heroics and push through these obstacles themselves, working excessive overtime to get it all done. In your position as tech lead, you should continue writing code, but not too much. Even if you are tempted to pull a rabbit out of the hat yourself, you must communicate this obstacle first. Your product manager should know as early as possible about any possible challenges. Enlist the help of your engineering manager as needed. In a healthy organization, there is no shame or harm in raising issues early. Teams often fail because they overworked themselves on a feature that their product manager would have been willing to compromise on. As a large project nears its delivery date, there will be compromises on functionality. Start looking for opportunities to delegate work, especially if there is part of the system you expected to build yourself that you have not had the time to tackle.

As you can see from these descriptions, in the process of being a tech lead, you have to act as a software developer, a systems architect, a business analyst, and a team leader who knows when to do something single-handedly, and when to delegate the work to others. Fortunately you don’t have to do all of these tasks at once. It may be uncomfortable at first, but you’ll find a balance with time and practice.

My final advice is to remember that you can switch tracks if you want. It is common for people to try out management at some point, realize they don’t enjoy it, and go back to the technical track [enginneer, individual contributor]. Nothing about this choice has to be permanent, but go in with your eyes wide open. Each role has benefits and drawbacks, and it’s up to you to feel out what you enjoy the most.

# Stop Writing Code and Engineering in the Critical Path

_Source: https://charity.wtf/2019/01/04/engineering-management-the-pendulum-or-the-ladder/_

There is an enormous demand for technical engineering leaders — far more demand than supply. The most common hackaround is to pair a people manager (who can speak the language and knows the concepts, but stopped engineering ages ago) with a tech lead, and make them collaborate to co-lead the team. This unwieldy setup often works pretty well.

But most of those people managers didn’t want or expect to end up sidelined in this way when they were told to stop engineering.

If you want to be a pure people manager and not do engineering work, and don’t want to climb the ladder or can’t find a ladder to climb, more power to you. I don’t know that I’ve met many of these people in my life. I have met a lot of people in this situation by accident, and they are always kinda angsty and unhappy about it. Don’t let yourself become this person by accident. Please.

Which brings me to my next point.

## You Will Be Advised to Stop Writing Code or Engineering

✨ **FUCK THAT.** ✨

Everybody’s favorite hobby is hassling new managers about whether or not they’ve stopped writing code yet, and not letting up until they say that they have. This is a terrible, horrible, no-good VERY bad idea that seems like it must originally have been a botched repeating of the correct advice, which is:

## Stop Writing Code and Engineering in the Critical Path

Can you spot the difference?  It’s very subtle. Let’s run a quick test:

* Authoring a feature?  ⛔️
* Covering on-call when someone needs a break?  ✅
* Diving on the biggest project after a post mortem?  ⛔️
* Code reviews?  ✅
* Picking up a p2 bug that’s annoying but never seems to become top priority?  ✅
* Insisting that all commits be gated on their approval?  ⛔️
* Cleaning up the monitoring checks and writing a library to generate coverage?  ✅

The more you can keep your hands warm, the more effective you will be as a coach and a leader. You’ll have a richer instinct for what people need and want from you and each other, which will help you keep a light touch. You will write better reviews and resolve technical disputes with more authority. You will also slow the erosion and geriatric creep of your own technical chops.

I firmly believe every line manager should either be in the on call rotation or pinch hit liberally and regularly, but that’s a different post.

