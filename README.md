# parity
Unordered diffs between files

This application finds unique line count differences between files and then outputs a sorted list of those differences.

Example before.txt file
```
line a
line b
line c
line a
```

Example after.txt file:
```
line d
line c
line a
line a
line c
line c
```

Parity usage:
```
parity before.txt after.txt
    -1 line b
     2 line c
     1 line d
```
