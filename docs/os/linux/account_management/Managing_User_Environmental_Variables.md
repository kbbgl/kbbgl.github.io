---
slug: os-linux-account-management-managing-user-environmental-variables
title: "Environmental Variable Management"
authors: [kbbgl]
tags: [os, linux, account_management, managing_user_environmental_variables]
---

# Environmental Variable Management

types of variables:

* **Environment** variables are system-wide variables built into your system and interface that control the way your system looks, acts, and “feels” to the user, and they are inherited by any child shells or processes.
* **Shell** variables are typically listed in lowercase and are only valid in the shell they are set in.

To view environmental variables:

```bash
env
```

To view all variables:

```bash
set
```

---

## Change variable for session

```bash
# HISTSIZE is a variable that controls the number of commands to remember
HISTSIZE=0
```

## Change variable permanently

Backup before modifying:

```bash
set > env.bkp
```

To make the change permanent, use the `export` command:

```bash
HISTSIZE=0
export HISTSIZE
```

## Create user-defined variable

```bash
NEW_VAR="some value"

echo $NEW_VAR
# some value
```

To delete a variable:

```bash
unset NEW_VAR
```
