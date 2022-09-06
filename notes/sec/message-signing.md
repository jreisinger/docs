Three algorithms/functions for message signing

* `G() -> (pk, sk)` -- generator returns public key (`pk`) and secret or private key (`sk`)
* `S(sk, x) -> t` -- signer returns tag `t` (string) for input `x` (string)
* `V(pk, x, t) -> accept|reject` -- verifier checks validity of tag `t` for given input `x` 

Correctness property

* `V(pk, x, S(sk, x)) = accept` -- should alway be true

Security property

* `V(pk, x, t) = accept` -- should almost never be true when `x` and `t` are chosen by an attacker

Source: https://web.stanford.edu/class/cs253/
