Why is security hard

Security is the capability of a system to achieve some goal in the presence of
an adversary. If your system is connected to the Internet it has adversaries.

High-level thinking about security

1. Policy - the goal you want to achive + C.I.A.
2. Threat model - assumptions about what adversaries could do
3. Mechanism - sw/hw/sys (e.g. accounts, passwords, permissions, encryption)

Why is security hard? Negative goal.

- Need to guarantee policy, assuming the threat model.
- Difficult to think of all possible ways that attacker might break in.
- Contrast: easy to check whether a positive goal is upheld, e.g. Alice can actually read file F.
- Weakest link matters.
- Simple mistake in mechanism can have serious security implications.
- Realistic threat modesl are open-ended (almost negative models).
- Iterative process: design, update threat model as necessary, etc.