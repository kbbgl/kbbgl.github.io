# Match Chunks

## match 0 or more times using * quantifier

```bash
echo "hello, how are you?" | grep -e "[a-z]*" 
# will match hello, how, are, you
```

## match all numbers in chunks with length of one or more with + quantifier

```bash
echo "47427 8381481 5813471" | grep -P "[0-9]+" 
echo "47427 8381481 5813471" | grep -P "\d+" 
```
