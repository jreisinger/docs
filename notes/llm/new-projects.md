# Creating new software projects with LLM agents

## Categories of projects

1. Low importance projects where deep code understanding is unnecessary
   (prototypes, throw away code): can be "vibe-coded", i.e. committing agent code
   without even reviewing it.
2. High importance which I want to maintain: review and guide all code agent
   writes before it's committed (or shortly after, see below)

NOTE: to really learn something new, you have to work through it from scratch,
yourself, reading, designing, writing the code. Agents can still be used as
props for learning.

## Workflow phases

Design

- requires extensive design
- iterate on the design with the agent
- outcome is a committed MD file

Implementation

- start asking the agent to create small diffs in a logical order that make sense to you
- review, cleanup and refactor the diffs before committing
- if it's not possible the make the commit small, clean up and refactor the code with subsequent commits
- you might need to spend several hours instructing the agent to clean up and refactor
- if you find out you go in the wrong direction revert the whole sequence of commits

## Practical workflow

I find it difficult to let the agent send a PR and review that. I find the following method more reliable:

1. I'm running the agent locally in the repo and ask it to update the code there. 
2. In parallel, I review the diffs in VScode.
3. I make my own tweaks and code changes if needed. Sometimes I ask questions.
4. Once I'm pleased with the change, I manually create a commit.

NOTE: it's imperative to keep making progress in small chunks, that a human can
fully understand in a single review. It's very tempting to sprint ahead
submitting thousands lines of code every day, but this temptation has to be
avoided.

## Testing

Agents produce - by far - the best results when they have a solid test suite to
test their code against. You need to start with tests. Find some or build them
from scratch. It's dangerous to trust agents for both the tests and the
implementations tested against them.

## Language choice

Go is a fantastic language for agents to write, because it was designed to be
very readable by humans. So it's relatively easy to review the agents' code.
When using agents humans spent most of the time *reading* the code.

---

Adapted from: https://eli.thegreenplace.net/2026/thoughts-on-starting-new-projects-with-llm-agents
