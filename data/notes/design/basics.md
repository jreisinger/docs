# Software design

Design is a combination of art and science. The essence of design is describing a system in terms of magical black boxes (functions or objects) and their interfaces. As usual, the key to success is practice.

## Top-Down Design

* invaluable tool for developing complex algorithms
* start with the general problem and try to express a solution (create an algorithm) in terms of smaller problems
* connect the top level problem and the lower level problems via an interface (function name, parameters and return values)
* attack each of the smaller problems using the same technique
* eventually the problems get so small that they are trivial to solve

```
#!/usr/bin/env python3
# rball.py

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

## Object-oriented design

* data-centered complement to top-down design

Design tips

1. Look for object candidates. Nouns are usually objects, verbs are methods. Primitive data types (numbers or strings) are probably not candidates for objects, grouping of related data items are.

2. Design iteratively. No one designs a program top to bottom in a linear, systematic fashion.

3. Try out alternatives. Good design involves a lot of trial and error. When you look at someone's code, it's finished work, not the process they went through to get there. "Plan to throw one away." -- Fred Brooks

4. Keep it simple. Don't design in extra complexity until it is *really* needed.

```
#!/usr/bin/env python3
# oo-rball.py -- simulation of a racquet ball game.
#                Illustrates design with objects.

from random import random

class Player:
    """ Player keeps track of service probability and score.
    """
    def __init__(self, prob):
        self.prob = prob
        self.score = 0
    def winsServe(self):
        return random() <= self.prob
    def incScore(self):
        self.score = self.score + 1
    def getScore(self):
        return self.score

class RBallGame:
    """ RBallGame represents a game in progress.  A game has two players and
        keeps track of which one is currently serving.
    """
    def __init__(self, probA, probB):
        # Create a new game having players with the given probs.
        self.playerA = Player(probA)
        self.playerB = Player(probB)
        self.server = self.playerA # playerA always serves first
    def play(self):
        # Play the game to completion.
        while not self.isOver():
            if self.server.winsServe():
                self.server.incScore()
            else:
                self.changeServer()
    def isOver(self):
        a, b = self.getScores()
        return a == 15 or b == 15 or \
            (a == 7 and b == 0) or (b == 7 and a == 0)
    def changeServer(self):
        if self.server == self.playerA:
            self.server = self.playerB
        else:
            self.server = self.playerA
    def getScores(self):
        return self.playerA.getScore(), self.playerB.getScore()

class SimStats:
    """ SimStats handles accumulation of statistics across multiple (completed)
        games. It tracks the wins and shutouts for each player.
    """
    def __init__(self):
        self.winsA = 0
        self.winsB = 0
        self.shutsA = 0
        self.shutsB = 0
    def update(self, aGame):
        # Determine the outcome of aGame and update statistics.
        a, b = aGame.getScores()
        if a > b:
            self.winsA = self.winsA + 1
            if b == 0:
                self.shutsA = self.shutsA + 1
        else:
            self.winsB = self.winsB + 1
            if a == 0:
                self.shutsB = self.shutsB + 1
    def printReport(self):
        # Print a nicely formatted report
        n = self.winsA + self.winsB
        print("Summary of", n , "games:\n")
        print("          wins (% total)   shutouts (% wins)  ")
        print("--------------------------------------------")
        self.printLine("A", self.winsA, self.shutsA, n)
        self.printLine("B", self.winsB, self.shutsB, n)
    def printLine(self, label, wins, shuts, n):
        template = "Player {0}:{1:5}  ({2:5.1%}) {3:11}   ({4})"
        if wins == 0:        # Avoid division by zero!
            shutStr = "-----"
        else:
            shutStr = "{0:4.1%}".format(float(shuts)/wins)
        print(template.format(label, wins, float(wins)/n, shuts, shutStr)) 

def getInputs():
    a = float(input("What is the prob. player A wins a serve? "))
    b = float(input("What is the prob. player B wins a serve? "))
    n = int(input("How many games to simulate? "))
    return a, b, n

def main():
    probA, probB, n = getInputs()
    stats = SimStats()
    for i in range(n):
        game = RBallGame(probA, probB)
        game.play()
        stats.update(game)
    stats.printReport()

main()
```

# Resources

* Python Programming: An Introduction to Computer Science, 2010
