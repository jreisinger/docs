(Up-to-date [source](https://github.com/jreisinger/blog/blob/master/posts/big-o-notation.md) of this post.)

* a mathematical way of describing *scaling* - e.g. growth rate of algorithms' run(ning) times
* it's a gross generalization that focuses on the trend as its input grows, not the specific run-time
* used to classify (compare) systems based on how they respond to changes in input size
* *O* is used because the growth rate of an algorithm's run time is known as its
    *order*
* order of magnitude is related but different - how many digits are in a number describing size (1000 is an order of magnitude larrger than 100)
* you can't really tell how fast an algorithm will be without benchmarking (CPU, disk, [RAM complexities](http://queue.acm.org/detail.cfm?id=1814327)) but O() notation can serve as a general guide

The expression in the parenthesis defines the number of operations for the *worst-case* scenario. For example, if you binary search through 100 elements (n = 100) you will do 7 (log n) guesses at most. If you use simple search you will do 100 (n) guesses at most. Also notice that the growth rate of the binary algorithm is much smaller:

| # elements (n) | Simple search (n)     | Binary search (log n) |
|----------------|-----------------------|-----------------------|
| 100            | 100 operations        | 7 operations          |
| 10 000         | 10 000 operations     | 14 operations         |
| 1 000 000 000  | 10 000 000 operations | 32 operations         |
    
Sub-linear scaling

* O(1) - **constant** - no matter the scale of the input, performance of the system
    does not change; ex. hash-table lookup in RAM; such algorithms are rare
* O(log n) - logarithmic - ex. [binary search](https://github.com/jreisinger/algorithms-with-perl/blob/master/binary-search) grows slower as the size of the corpus being searched grows; it's growth is less than linear

Linear scaling

* O(n) - **linear** - ex. simple search; twice as much data requires twice as much processing time

Super-linear scaling

* O(n^m) - exponential - as input size grows the system slows down
    disproportionately
* O(n^2) - quadratic - but everybody says **exponential** when they mean quadratic :-)

## See also

* http://bigocheatsheet.com/

## Resources

* TPoCSA, Appendix C
* Grokking Algorithms
