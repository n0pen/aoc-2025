## Day 10

### Challenge

The boss of the Advent. It took a lot of effort to solve, but I knew where to start. I remembered from my engineering
studies that there are methods for searching optimal solutions for this kind of problem, but I was terrible at algebra
and I hadn't done any matrix math in over 10 years. My way ahead was SIMPLEX

### Problem

Part 1 was a combinatory problem that didn't have any complexities

Part 2 in the other hand is a full-on optimization problem: Find the minimum amount of presses that gave summed the
final amounts

### Solution

Part 1 was easily solved by brute-force searching the combination of presses that yielded the correct solution

For part 2 I created my own implementation of a Simplex Solver

#### 1: Linear Programming and simplex

Simplex is quite a straightforward algorithm if you know it, but given I hadn't studied it in a long time, I forgot most
of its intricacies.

I had to read a full simplex step-by-step tutorial and program all that to get to a optimal solution. It took a very
long time but in the end, and after a lot of trial and error I ended up with a working simplex solver.

There were many challenges to create the solver, mainly because I was not very familiar with Simplex at this point, but
there was an unexpected challenge: Floating point arithmetic impression. I had to carefully consider my results when
iterating because a 0 value and a 0.000001 may look similar in the console but are not equal (!=)

#### 2: Linear is not always optimal

After finishing the simplex solver I found out the hard way that the linear optimal solution was not what I needed.
Given that the presses of a buton are discrete operations, I needed an integral (integer) solution and approximating the
values was not a correct solution

After some back and forth I ended up finding the method I wanted to use: Integral Branching.

For this to work I needed to find the linear optimal solution, and from there create new problems, adding new
constraints
that limited the values to integers. For this, each fractional value created 2 branches, one for the ceiling value and
one for the floor value, and branching one time for each solution, recursively and selecting the most optimal return
values, I could get the optimal integral solution.

### Score

- Fun: 10/10
- Challenge: 9/10