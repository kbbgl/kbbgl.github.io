---
slug: os-windows-admin-enable-running-remote-commands
title: "Enable Running Remote Commands"
authors: [kbbgl]
tags: [os, windows, admin, enable_running_remote_commands]
---

# Enable Running Remote Commands

```powershell
# Enable remoting on server where RPC will run in
Enable-PSRemoting -Force

# Check if WinRM is running on remote server
Test-WsMan QUERY_NODE1_HOSTNAME

$username = "corp\myuser"
$password = ConvertTo-SecureString -String "mypswd" -AsPlainText -Force

$cred = [pscredential]::new($username,$password)

$args = "-t", "list_only"
$parameters = @{

    ComputerName = "s1", "s2"
    ScriptBlock = { & "C:\scripts\elasticube_data_cleaner_11Feb2019 Update\elasticube_data_cleaner.exe" -t "list_only" }
    Credential = $cred

}

Invoke-Command @parameters 
```
