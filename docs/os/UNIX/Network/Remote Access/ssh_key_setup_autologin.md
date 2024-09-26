# SSH Autologin Setup

## Generate key

```bash
ssh-keygen
```

## Copy public key and permissions to remote server

```bash
ssh-copy-id -i /Users/me/.ssh/lnx12.pub me@lnx1
```
