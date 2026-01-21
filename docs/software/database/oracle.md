---
slug: oracledb-cheatsheet
title: Oracle Database Cheatsheet
description: A cheatsheet for Oracle database.
authors: [kgal-akl]
tags: [cheatsheet, oracle, cli, database]
---

## Login

Log in with `system` user as `sysdba`:
```bash
sqlplus system/password@localhost:1521/FREE as sysdba
```

## List Services
```sql
SQL> SELECT name FROM v$services;

NAME
----------------------------------------------------------------
freeXDB
SYS$BACKGROUND
SYS$USERS
freepdb1
free
```

## List Databases

```
SQL> SELECT name FROM v$database; 

NAME 
--------- 
FREE
```

## Create Pluggable DB
```
CREATE PLUGGABLE DATABASE testdb ADMIN USER ADMIN_USER2 IDENTIFIED BY "password" FILE_NAME_CONVERT=('/opt/oracle/oradata/FREE/pdbseed/', '/opt/oracle/oradata/FREE/testdb/');

ALTER PLUGGABLE DATABASE testdb OPEN;

CONNECT ADMIN_USER2/password@localhost:1521/testdb;
```


## List Users
```sql
SQL> SELECT USERNAME FROM ALL_USERS;
```

