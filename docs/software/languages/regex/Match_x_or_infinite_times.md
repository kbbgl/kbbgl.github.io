# Match X+ Characters

## use {x,} to match string with x+ characters

```bash
echo "hello world reallylongword" | grep -P "\w{6,}" -o 
#will match reallylongword
```
