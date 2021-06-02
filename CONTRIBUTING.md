# How to contribute

The project mainly resides in the `parser/parser.go` file.

The idea is the following.

- Collect all `.rpy` files.
- Read and analyse each line.
  - If the line contains `renpy-graphviz`, search for tags (`tags.go` file)
  - Update the context thanks to regexes (`context.update` function)
  - Thanks to the given context, perform some actions on the graph (`Graph` function, the main one)
  - Perform the actions on the graph according to parameters (`graph.go` file)
- Before outputting anything, clean lonely nodes if asked
