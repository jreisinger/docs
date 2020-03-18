Design is a combination of art and science. As usual, the key to success is practice.

## Top-Down Design

* invaluable tool for developing complex algorithms
* start with the general problem and try to express a solution (create an algorithm) in terms of smaller problems
* connect the top level problem and the lower level problems via an interface (function name, parameters and return values)
* attack each of the smaller problems using the same technique
* eventually the problems get so small that they are trivial to solve

```
#!/usr/bin/env python3
# racquetball.py

from random import random

## Top-Level Design

def main():
    """ PROBLEM:
        Simulate a game of racquetball to find out whether a slightly
        better player can win significant number of games.
        SOLUTION algorithm:
        Get the inputs: probA, probB, n
        Simulate n games of racquetball using probA and probB
        Print a report on the wins for playerA and playerB
    """
    
    # Algorithm written in terms of functions that don't yet exist.
    probA, probB, n = getInputs()
    winsA, winsB = simNGames(n, probA, probB)
    printSummary(winsA, winsB)

## Second-Level Design

def getInputs():
    a = float(input("What is the prob. player A wins a serve? "))
    b = float(input("What is the prob. player B wins a serve? "))
    n = int(input("How many games to simulate? "))
    return a, b, n

def simNGames(n, probA, probB):
    """ Simulate n games and return winsA and winsB.
    """
    winsA = winsB = 0
    for i in range(n):
        scoreA, scoreB = simOneGame(probA, probB)
        if scoreA > scoreB:
            winsA = winsA + 1
        else:
            winsB = winsB + 1
    return winsA, winsB

def printSummary(winsA, winsB):
    n = winsA + winsB
    print("\nGames simulated:", n)
    print("Wins for A: {0} ({1:0.1%})".format(winsA, winsA/n))
    print("Wins for B: {0} ({1:0.1%})".format(winsB, winsB/n))

## Third-Level Design

def simOneGame(probA, probB):
    """ Simulate one game and return scoreA and scoreB.
    """
    scoreA = scoreB = 0
    serving = "A"
    while not gameOver(scoreA, scoreB):
        if serving == "A":
            if random() < probA:
                scoreA = scoreA + 1
            else:
                serving = "B"
        else:
            if random() < probB:
                scoreB = scoreB + 1
            else:
                serving = "A"
    return scoreA, scoreB

def gameOver(a, b):
    return a==15 or b==15

if __name__ == '__main__': main()
```

## Bottom-Up Implementation

A good way to approach the implementation of a modest size program is to start at the lowest levels of the structure chart and work your way up, testing each component (*unit testing*) as (or before) you complete it. 

## Prototyping and Spiral development

* not an alternative to top-down design but a complementary approach
* useful when dealing with new or unfamiliar features or technologies 
* (everything may seem new to a novice programmer!)
* useful when you get stuck at a step
* start with a simplified version of a program or a program component (prototype)
* try to gradually add features until it meets full specification

```
#!/usr/bin/env python3
# prototype.py

from random import random

def simOneGame():
    # Play just 30 rallies. Each player has
    # a 50-50 chance of winning a given point.
    scoreA = scoreB = 0
    serving = "A"
    for i in range(30):
        if serving == "A":
            if random() < .5:
                scoreA = scoreA + 1
            else:
                serving = "B"
        else:
            if random() < .5:
                scoreB = scoreB + 1
            else:
                serving = "A"
    print(scoreA, scoreB)

if __name__ == '__main__': simOneGame()
```

# Resources

* Python Programming: An Introduction to Computer Science, 2010
