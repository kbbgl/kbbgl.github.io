# Read all git Object Contents"

```bash
find .git/objects/ -type f | while read object; \
  do; \
  hash=$(grep -E "[0-9a-f]{2}/\\b[0-9a-f]{5,40}\\b" --only-matching | tr -d "/"); \
  echo $hash | xargs -I {} git cat-file -p {}; \
  done;
  ```
