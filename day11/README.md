## Day 11

### Challenge

At this I was expecting an impossible challenge. It ended up being a regular graph-theory problem

### Problem

Part 1 was simple: Find all the ways to connect 2 nodes. Given that this was a directed graph there was a clear path
forward

Part 2 was the same, but this time there were some stops to make

### Solution

Part 1 straightforward: Walk the graph recursively and return 1 at the end. I was surprised that with no optimization it
ran so fast

For part 2 I got ambushed by a 200 trillion ways to walk from start to end, which needed some optimization to actually
calculate.
At this point I knew attempt a naive solution trying to get the correct answer for my input, and I was right

My second attempt, after failing to brute force it (with memoization) was to subdivide the paths and sum it all at the
end. This didn't consider that some early paths could've contained the other required stop, but it ended up being the
solution anyway

### Score

- Fun: 6/10
- Challenge: 6/10