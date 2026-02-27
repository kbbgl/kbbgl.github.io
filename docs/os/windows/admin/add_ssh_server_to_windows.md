---
slug: os-windows-admin-add-ssh-server-to-windows
title: "Add SSH Server"
authors: [kbbgl]
tags: [os, windows, admin, add_ssh_server_to_windows]
---

# Add SSH Server

## Check

```powershell
Get-WindowsCapability -Online

Name  : OpenSSH.Client~~~~0.0.1.0
State : Installed

Name  : OpenSSH.Server~~~~0.0.1.0
State : NotPresent
```

## Enable

```powershell
Add-WindowsCapability -Online -Name OpenSSH.Server~~~~0.0.1.0
```

start the `sshd` service:

```powershell
Start-Service sshd
Get-Service sshd
```

set to start automatically:

```powershell
Set-Service -Name sshd -StartupType 'Automatic'
```

[Optional] - add firewall rule on port 22:
