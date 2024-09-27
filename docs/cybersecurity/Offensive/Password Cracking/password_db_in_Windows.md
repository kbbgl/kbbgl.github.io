# Password Database in Windows

The Security Account Manager (SAM) is a database file in Windows XP, Windows Vista, Windows 7, 8.1 and 10 that stores users' passwords.

It can be used to authenticate local and remote users.

Beginning with Windows 2000 SP4, Active Directory authenticates remote users.

SAM uses cryptographic measures to prevent unauthenticated users accessing the system.

## Location

`%SystemRoot%/system32/config/SAM` and `HKLM/SAM` in registry.

can use `meterpreter` command `hashdump` after reverse shell access to collect the
Then we can use [`johnthereaper`](https://www.openwall.com/john/)  or [`hashcat`](https://hashcat.net/hashcat/) to crack them.

### Hash structure

```
Jason:502:aad3c435b514a4eeaad3b935b51304fe:c46b9e588fa0d112de6f59fd6d58eae3:::
```

`Jason` is the user name

`502` is the relative identifier (`500` is an administrator, `502` here is a kerberos account.) - [adsecurity](https://adsecurity.org/?p=483)

`aad3c435b514a4eeaad3b935b51304f` is the LM hash

`c46b9e588fa0d112de6f59fd6d58eae3` is the NT hash

Valuable info on LM/NT hashes can be found [here](http://www.adshotgyan.com/2012/02/lm-hash-and-nt-hash.html)

Use [crackmapexec](https://github.com/byt3bl33d3r/CrackMapExec) or [psexec](https://learn.microsoft.com/en-us/sysinternals/downloads/psexec)
