Three algorithms/functions - **G**enerator, **S**igner, **V**erifier:

* `G() -> (pk, sk)` - returns public key (`pk`) and secret (private) key (`sk`)
* `S(sk, x) -> t` - returns tag `t` (string) for input `x` (string)
* `V(pk, t, x) -> accept|reject` - checks validity of tag `t` for given input `x` 
