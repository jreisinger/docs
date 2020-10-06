* aka generate and test, brute force, trial and error, exhaustive enumeration
* general problem solving technique and algorithm
* systematically enumerate all possible candidates for a solution and check whether each satisfies the problem's statement

Implementation of guess and check algorithm with approximation:

```
# Find cube root of a given number.

cube = 27

step = 0.01 # decreasing step -> slower program
diff = 0.1  # increasing diff -> less accurate answer

count = 0
guess = 0

# Look for close enough answer and make sure you don't
# accindentally skip the close enough boundary.
while abs(guess**3-cube) >= diff and guess <= cube:
    guess += step
    count += 1

print("numbers of guesses: {}".format(count))

if abs(guess**3-cube) < diff:
    print("{} is close to the cube root of {}".format(guess, cube))
else:
    print("failed to find cube root of {}".format(cube))
```

More

* [MIT Open courseware](https://www.youtube.com/watch?v=SE4P7IVCunE&list=PLUl4u3cNGP63WbdFxL8giv4yhgdMGaZNA&index=11) (video)
