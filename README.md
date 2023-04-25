# paths2tree

`paths2tree` is a simple alternative to `tree --fromfile .` which does not use
a backing tree data structure and allocates very little memory.

## Example

```
echo "a/foo\na/bar\nb/baz\nc/qux\nc/d/quux" | paths2tree
```

Output:

```
a
|   foo
|   bar
b
|   baz
c
|   qux
|   d
|   |   quux
```

## Known issues

### Prints only the last subdirectory name before its items

Input:

```
a/b/c/d/e/foo
```

Expected output:

```
a
|   b
|   |   c
|   |   |   d
|   |   |   |   e
|   |   |   |   |   foo
```

Actual output:

```
|   |   |   |   e
|   |   |   |   |   foo
```
