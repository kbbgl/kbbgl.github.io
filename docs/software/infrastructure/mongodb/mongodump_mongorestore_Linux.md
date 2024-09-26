---
slug: mongodb-dump-restore
title: Dump and Restore Database
authors: [kbbgl]
tags: [mongodb, dump, backup, restore]
---

## Linux

### Dump

```bash
DB_NAME=foo
mongodump --db=$DB_NAME --out=./
```

### Restore

```bash
mongorestore --drop --db=$DB_NAME mongo_dump/$DB_NAME/
```

## Windows

### Dump

```powershell
"c:\program files\mongodb\mongodump.exe" --host=localhost --port=27017 -u $USER -p $PASS /d $DB_NAME /authDB $ADMIN
```

### Restore

```powershell
"c:\program files\mongodb\mongorestore.exe" --host=localhost --port=27017 -u $USER -p $PASS /d $DB_NAME /authDB $ADMIN /dir {path/to/dump} /drop
```
