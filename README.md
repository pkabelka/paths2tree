# paths2tree

`paths2tree` is a simple alternative to `tree` which does not use a backing
tree data structure and allocates very little memory.

## Known issues

### Prints only the last subdirectory name before its items

Expected:

```
a
|   b
|   |   c
|   |   |   d
|   |   |   |   e
|   |   |   |   |   foo
```

Outputs:

```
|   |   |   |   e
|   |   |   |   |   foo
```
