---
title: "Data Sections"
---

# Data Sections

Sections in an `asm` file are a way to organize a program. Sections have the following syntax:

```
section SECTION_NAME [data, code, readable, writable, executable]
```

The `SECTION_NAME` could be anything but by conventions it's usually `.data`, `.bss` and `.text` and `.idata`.

- `.data` sections includes initialized data.
- `.bss` sections include uninitialized (future) data.
- `.text` includes the code.
- `.idata` connects us to external modules (imports).

## Portable Executable Format

When opening an executable with a hex viewer, we're able to see that the sections include information about different memory location references.

![image.png](https://boostnote.io/api/teams/OdksQwLn2/files/9798d0fb1c9a374457ab1212a117ef6f3a6b73808616e33a11938400e5a5fe34-image.png)
