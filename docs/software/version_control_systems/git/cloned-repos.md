# Cloned Repositories

Cloned repos do not have the `.git/objects/$HASH` files as in locally-created projects.

It has 2 files in `.git/objects/pack`:\

- `pack-$HASH.pack`
- includes compressed `git` objects.
- `pack-$HASH.idx`
- includes indexes of `git` objects.

To unpack:

```bash
# it will unpack all objects to `.git/objects/$HASH`
cat pack-$HASH.pack | git unpack-objects
```
