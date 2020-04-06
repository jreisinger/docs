Three algorithms/functions - **G**enerator, **S**igner, **V**erifier:

* `G() -> (pk, sk)` - returns public key (`pk`) and secret (private) key (`sk`)
* `S(sk, x) -> t` - returns tag `t` (string) for input `x` (string)
* `V(pk, x, t) -> accept|reject` - checks validity of tag `t` for given input `x` 

Correctness property

* `V(pk, x, S(sk, x)) = accept` should alway be true

Security property

* `V(pk, x, t) = accept` should almost never be true when `x` and `t` are chosen by an attacker
