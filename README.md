# N-Puzzle
Project N-puzzle @ 42Lyon

# Goal

For this project, we had to implement different searching algorithms in order to solve a [n-puzzle](https://en.wikipedia.org/wiki/15_puzzle).
The mandatory part was about implementing an A* ("A star") algorithm with different heuristics.
The bonus part was about add a greedy search to compare with A* (complexity in time, size and the number of moves)

# Launch

```shell
make <build the project>
make re <rebuild the project>
make usage <launch the usage of the executable>
make test <build & launch project with the file $(TEST_FILE) written in the makefile>
```

# Output

If the renseigned n-puzzle is resolvable, you will see something like this running a simple `make test`:
```shell
No heuristics are selected, so calcs will be done with all heurisitcs
Solution found!
===============
[6 8 1]            |
[2 7 4]            | => the n-puzzle at time "0"
[3 5 0]            |
0 || 38            // Here 0 is the turn, 38 is the heuristic cost
=======
[6 8 1]
[2 7 0]
[3 5 4]
1 || 40
...
...
...
[1 2 3]
[0 8 4]
[7 6 5]
25 || 5
=======
[1 2 3]
[8 0 4]
[7 6 5]
26 || 0
=======

Complexity in time = 4903  // Total number of states ever selected in the "opened" set
Complexity in size = 2928  // Maximum number of states ever represented in memory at the same time during the search 
Total number of moves = 26
```

# Datastructure
To have the maximum note, we had to use a priority queue to sort all states in the opened set of A*.
I used an heapsort
