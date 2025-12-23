## Day 9

### Challenge

Visualization for this input was the way that led me to a solution.
I also found out that the solution I had to create didn't have to
account for inexistent edge cases: The input is static.

### Problem

The problem for this one was finding the maximum area square created by 2 corner tiles
Part 1 was very easy. I didn't require any special considerations, part 2 on the other hand constrained the problem
quite a lot:
Your square could only be created inside the boundaries of a polygon created by a set of points

### Solution

Part 1 was simply found by iterating over all points and getting the maximum area

Part 2 in the other hand required a lot of thinking from myself.

#### Creating the criteria

After thinking a lot the first day, I went to sleep and thought about creating a visualization for the input.
I knew the polygon created by the points was quite big, so I ended up creating and SVG file with a polygon tag
containing all points, and used that to visualize my results.

The first criterion I deviced was checking the edges my square and the polygon's for intersections.
I knew that the polygons edges were axis-aligned and so where the square's edges so checking for intersection was easy.

The second criterion actually gave me the correct solution, and it was searching for aligned edges that broke my square
edge.
This criterion didn't give the correct solution for the test case, but it worked the real input.

I tried implementing a better solution, that considered other cases, but I leaved at that in the end

### Score

- Fun: 6/10
- Challenge: 8/10