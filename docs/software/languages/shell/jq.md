---
title: jq Cheatsheet
slug: jq-cheatsheet
app2or: kgal-akl
tags: [cheatsheet, utils, tools, cli, jq, json, javascript]
---

## Select

The format is:

```bash
... | jq '<filter> | select(<condition>)' | ...
```

### Select objects where a specific field equals a value:

```bash
jq '.[]	| select(.id == 101)'
```

### Select objects based on multiple conditions

```bash
jq '.[]	| select(.status == "active" and .count > 5)'
```

test husky https://corp:8000

### Select items that match a value in a nested array

```bash
jq '.[]	| select(.tags[] == "urgent")'
```

### Select based on the presence of a key

```bash
`jq '.[] | select(has("optional_field"))'
```


### Select specific data from filtered objects

```bash
jq -r '.files[] | select(.fileName=="FOO")'
```

### Select with contains

```bash
jq -r '.files[] | select(.<field> | contains("<value>")'
```