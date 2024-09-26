# How to Find Process listening on port

```bash
PORT=4445
lsof -n -i :$PORT
```
