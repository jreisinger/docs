## Top-Down Design

* start with the general problem and try to express a solution in terms of smaller problems
* attack each of the smaller problems using the same technique
* eventually the problems get so small that they are trivial to solve

```
#!/usr/bin/env python3
# racquetball.py

from random import random

## Top-Level Design

def main():
    """ ALGORITHM:
        Get the inputs: probA, probB, n
        Simulate n games of racquetball using probA and probB
        Print a report on the wins for playerA and playerB
    """
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

## Prototyping and Spiral development

* not an alternative to top-down design but a complementary approach
* start with a simple version of a program or a program component
* try to gradually add features until it meets full specification
* useful when dealing with new or unfamiliar features or technologies (everything may seem new to a novice programmer!)
* useful when you get stuck at a step
