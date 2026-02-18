# Primitives of GitHub Actions (hosted CI/CD tool)

workflow
- defines when automation runs and what it does

job
- a unit inside a workflow that runs on a runner
- ordered list of steps
- jobs run in parallel by default; can be made dependent using `needs`

step
- runs shell script you write
- invokes an action

action
- a reusable task (packaged code) that a step can run