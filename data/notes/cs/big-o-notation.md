NOTE: Logarithms (log) mentioned here are base 2.

# Big-O Notation

* a mathematical way of describing *scaling* - e.g. growth rate of algorithm's run(ning) time
* it's a gross generalization that focuses on the trend as its input grows, not the specific run-time
* used to classify (compare) systems based on how they respond to changes in input size
* *O* is used because the growth rate of an algorithm's run time is known as its *order*
* order of magnitude is related but different - how many digits are in a number describing size (1000 is an order of magnitude larger than 100)
* you can't really tell how fast an algorithm will be without benchmarking (CPU, disk, [RAM complexities](http://queue.acm.org/detail.cfm?id=1814327)) but O() notation can serve as a general guide

Comparison of simple and binary search algorithms:

| n     | Simple search (n) | Binary search (log n) |
|:------|:------------------|:----------------------|
| 10    | 10 steps          | 4 steps               |
| 100   | 100 steps         | 7 steps               |
| 1 000 | 1 000 steps       | 10 steps              |

The expression in the parenthesis in the 2nd and 3rd column defines the number of steps (guesses, operations) for the *worst-case* scenario. For example, if you binary search through 100 elements (n = 100) you will do 7 (log n) steps at most. If you use simple search you will do 100 (n) steps at most. Also notice that the growth rate of the binary algorithm is much smaller.

## Sub-linear scaling (growth is less than linear)

* O(1) - **constant** - no matter the scale of the input, performance of the system does not change; ex. hash-table lookup in RAM; such algorithms are rare
* O(log n) - logarithmic - grows slower as the size of the corpus being searched grows; ex. binary search

## Linear scaling

* O(n) - **linear** - ex. simple search; twice as much data requires twice as much processing time

## Super-linear scaling

* O(n^m) - exponential - as input size grows the system slows down disproportionately
* O(n^2) - quadratic - but everybody says **exponential** when they mean quadratic :-)

# Resources

* TPoCSA, Appendix C
* Grokking Algorithms
* http://bigocheatsheet.com/
