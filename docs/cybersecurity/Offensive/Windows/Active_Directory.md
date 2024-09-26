---
slug: active-directory
title: Active Directory
description: Active Directory
authors: [kbbgl]
tags: [cybersecurity,offensive,windows,ad,active_directory]
---

Authentication using Kerberos. Non-Windows devices can also authenticate using RADIUS/LDAP.


## Physical Components
### Domain Controller

A server with an AD DS role installed.

* Hosts a copy of the AD DS directory store.
* Provides authentication and authorization services.
* Replicate updates to other domain controllers.
* Allow administrative access to manage users/networks.

### AD DS Data Store

- Constists of the `Ntds.dit` file which holds all users, passwords for the domain.

- Stored by default in the `%SystemRoot%\NTDS` folder in all domain controllers.

- Accessible only through the domain controller processes and protocols.


## Logical Components

### AD DS Schema

- Defines every type of object that can be created and stored in the directory.
- Enforces rules regarding object creation and configuration.

Object types:
* Class Object - User, Computer
* Attribute Object - Display name

### Domains

Used to group and manage objects in an organization.

