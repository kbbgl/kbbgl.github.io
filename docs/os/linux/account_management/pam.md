---
slug: os-linux-account-management-pam
title: "Pluggable Authentication Modules"
authors: [kbbgl]
tags: [os, linux, account_management, pam]
---

# Pluggable Authentication Modules

Provide mechanism to ensure that users/applications are properly identified and authenticated.

Various applications have used `libpam` to be able to configure a uniform authentication method. Its modules provide flexibility in authenticating, password and session management.

Each PAM-aware application or service can be configured in a configuration file in:

```text
/etc/pam.d/
```

Each file in `/etc/pam.d` corresponds to a service and each line in the file specifies a rule. The rule is formatted as:

```text
type control module-path module-arguments
```

## PAM Rules

- `type`: specifies management group is associated with. Possible values are:
  - `auth`: Instructs app to prompt for identification.
  - `account`: Checks user's account such as password aging, ACL.
  - `password`: Update user authentication token.
  - `session`: Used to provide functions before and after the session is established (e.g., setting up environment, logging).

- `control`: controls success or failure of module on the rest of the flow:
  - `required`: Must return success.
  - `requesite`: Same as `required` except failure in any module termninates stack and sends application return status.
  - `optional`: module is not required.
  - `sufficient`: if module succeeds, no subsequent modules are executed. If it fails and is the only module in the stack, it will fail.

- `module-path`: gives file name of the library that can be found in `/lib*/security`.
- `module-arguments`: given to modify the PAM module's behavior.

### Steps Involved in Authentication

1) User invokes PAM-aware application, e.g. `login`, `ssh`.
2) The application calls `libpam`.
3) The library checks for configuration files in `/etc/pam.d/`.
4) Each referenced module is executed according to the configuration.

### LDAP

LDAP can be integrated using PAM.
DAP uses PAM and `system-config-authentication` or `authconfig-tui`. One has to specify the server, search base DN (domain name) and TLS (Transport Layer Security). Also required is `openldap-clients`, `pam ldap` and `nss-pam-ldapd`.

When you configure a system for LDAP authentication, five files are changed:

```bash
/etc/openldap/ldap.conf
/etc/pam_ldap.conf
/etc/nslcd.conf
/etc/sssd/sssd.conf
/etc/nsswitch.conf
```

You can edit these files manually or use one of the utility programs available (`system-config-authentication` or `authconfig-tui`).

We can use LDAP search to check for successful client authentication:

```bash
ldapsearch -H ldaps://ldap.kbbgl.github.io:636/ -b "DC=kbbgl,DC=io" -D "CN=kbbgl.gh.io,OU=svcacct,DC=kbbgl,DC=io" -w "$PASSWORD" -ZZ -x -v -d 1 -o nettimeout=10 -o tls_cacert=./ca-chain.pem tls ca cert "cn=*" dn
```