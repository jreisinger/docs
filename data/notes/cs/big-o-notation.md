# Big O notation

NOTE: [Logarithms](https://www.rapidtables.com/calc/math/Log_Calculator.html) (log) mentioned here are base 2.

* a mathematical way of describing *scaling* - e.g. growth rate of algorithm's run(ning) time
* used to classify (compare) systems based on how they respond to changes in input size
* you can't really tell how fast an algorithm will be without benchmarking (CPU, disk, RAM [complexities](http://queue.acm.org/detail.cfm?id=1814327)) but *O()* notation can serve as a general guide
* *O* is used because the growth rate of an algorithm's run time is known as its *order*
* *order of magnitude* is related but different - how many digits are in a number describing size (1 000 is an order of magnitude larger than 100)

For example, *simple search* is O(n) algorithm and *binary search* is O(log n) algorithm. It means that if you have for example 100 (sorted) items you need to search through you need at most (worst-case scenario) 100 steps using simple search and 7 steps (guesses, operations) using binary search. Also notice that the growth rate of the binary search algorithm is much smaller.

| n     | Simple search O(n) | Binary search O(log n) |
|------:|-------------------:|-----------------------:|
| 10    | 10 steps           | 4 steps                |
| 100   | 100 steps          | 7 steps                |
| 1 000 | 1 000 steps        | 10 steps               |

All this means that binary search is faster than simple search and it gets a lot faster as the input size increases.

## Sub-linear scaling

* growth is less than linear

O(1) - constant

* no matter the scale of the input, performance of the system does not change
* ex. hash-table lookup in RAM
* such algorithms are rare

O(log n) - logarithmic

* grows slower as the size of the corpus being searched grows
* ex. binary search

## Linear scaling

O(n) - linear

* twice as much data requires twice as much processing time
* ex. simple search

## Super-linear scaling

O(n^m) - exponential

* as input size grows the system slows down disproportionately

O(n^2) - quadratic

* everybody says exponential when they really mean quadratic :-)
* ex. selection sort

# Resources

* TPoCSA, Appendix C
* Grokking Algorithms
