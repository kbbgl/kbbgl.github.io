# Match Character Classes

## Vowels

```bash
echo abcdefghijk | grep -e "[aeiou]"
# will find a,e,i
```

## Range

```bash
echo abcdefghijkAC04 | grep -e "[a-zA-D0-9]"
```

## Negation character ^ - don't match

```bash
echo abcdefghijkAC04 | grep -e "[^a-zA-D0-9]"
# AC04
```

## select with or (alternation)

```bash
echo "cat dog" | grep -e "[cat|fish]" # selects cat
```

## Using metacharacters

### select only characters (including _)

```bash
echo "P45SW0%$#_#2#4RD" | grep -e "\w" # selects P45SW0RD
```

### select only special characters

```bash
echo "P45SW0%$#_#2#4RD" | grep -e "\W" # selects %$#_##
```

### select all digits (-P flag enables perl regex)

```bash
echo "P45SW0%$#_#2#4RD" | grep -P "\d" 
# matches
```

### select all except digits (-P flag enables perl regex)

```bash
echo "P45SW0%$#_#2#4RD" | grep -P "\D" # matches 
```

### select tabs, newlines

```bash
echo "some text with tab" | grep -P "\t" 
# matches tab  

echo "some text with tab" | grep -P "\n" 
# matches newline

echo "some text with tab" | grep -P "\s" 
# matches spaces

echo "some text with tab" | grep -P "\S" 
# matches sometextwiththab

# match square brackets
echo "[]" | grep -e "\[\]"
```
