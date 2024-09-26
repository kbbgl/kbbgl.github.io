# Output to File and STDOUT

```bash
mongo --eval "var cleanup=true" /path/to/sanity/tools/scripts/scripts_loader.js 2>&1 | tee sanity.log
```
