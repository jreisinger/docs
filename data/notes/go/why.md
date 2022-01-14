2021-11-22

# Modern applications

Should be fully modern. Adding modern components alongside of legacy systems introduces significantly higher (operational and security) risk and complexity. Instead of simplifying you triple the complexity => multiple systems, languages and integrations. Rewrite in Go has fewer risks from long term perspective.

Must adapt to business needs => service oriented APIs that are

* isolated, modular and composable (no tight coupling which hinders progress and increases maintenance costs)
* easy to learn and maintain (constant turnover of engineers)
* easy to deploy
* scalable (up and down) and efficient
* cloud native architecture (automation, monitoring, repair)
* event driven (serverless) component

# Why should I be using Golang

## Simplicity 

> Everything should be made as simple as possible, but not simpler. -- Albert Einstein

* Go puts focus on simplicity: http://bradfitz.com/2020/01/30/joining-tailscale, https://talks.golang.org/2015/simplicity-is-complicated.
* You want simplicity because there's already enough (organizational, technological) chaos.
* You need to simplify your applications, not only containerize them.

## Learnability and usage

* Relatively easy to learn and adopt.
* Top 10 language.
* Productive programming.
* High readability => good for dynamic teams, (designed for) low maintenance.

## Large and mature ecosystem

* `go` tool - docs, tests, build, install
* standard library
* 3rd party libraries

## Security

Realatively new, built with safety and security in mind (this is not true of languages that are 20+ year old - it made sense in the pre Internet era).

[70%](https://raw.githubusercontent.com/microsoft/MSRC-Security-Research/master/presentations/2019_02_BlueHatIL/2019_02%20-%20BlueHatIL%20-%20Trends%2C%20challenge%2C%20and%20shifts%20in%20software%20vulnerability%20mitigation.pdf) of serious security bugs related to memory safety. Go has automating memory management (GC with low latency). Memory-safe by default. Type-safe polymorphism.

Right (latest) encryptions and protocols.

Minimal dependecies. "Efforts to target popular code registries like Node Package Manager (NPM) JavaScript registry, PyPI, and RubyGems have become commonplace and a new frontier for an array of attacks." -- https://thehackernews.com/2021/11/11-malicious-pypi-python-libraries.html

Tamper-evident dependency chain.

Single binary.

Supports distroless containers.

## Performance

* Fast.
* Concurrency, parallelism (usage of multiple CPUs).
* Built for Google needs.

## Cloud native language

* first class support on all cloud providers
* designed for containers, small footprint, fast tests and builds
* 75% of projects in CNCF are written in Go (including Docker and Kubernetes)
* low resource use, instant startup times => ideal for event driven/serverless

You can build performant and secure APIs using just Go standard library (including production ready web server).

Go's application structure mirrors architecture of services-oriented systems. One pattern at every level.

Composition over inheritance.

## Sources/more

* [Go Day 2021 - Modern Enterprise Applications](https://www.youtube.com/watch?v=5fgG1qZaV4w)
