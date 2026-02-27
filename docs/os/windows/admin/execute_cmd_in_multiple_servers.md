---
slug: os-windows-admin-execute-cmd-in-multiple-servers
title: "Execute Command (with Arguments) on Multiple Windows Servers using PowerShell"
authors: [kbbgl]
tags: [os, windows, admin, execute_cmd_in_multiple_servers]
---

# Execute Command (with Arguments) on Multiple Windows Servers using PowerShell

This PowerShell script is useful when you have access to multiple Windows servers and need to run the same command on all of them.

For example,  I needed to run an executable that can be found on both nodes (`AppNode1`, `AppNode2`) in path `C:/scriptsmyapp.exe`  of a multi-node cluster:


```powershell
####### INPUT PARAMETERS #####

# Insert username
$username = "corpjohn.smith"

# Insert password
$password = "SOM3P45SW0RD"

# The executable's path which we want to run remotely
$exe = "C:/scriptsmyapp.exe"

# List of servers to run command on
$servers = "AppNode1", "AppNode2"

# Generate a PowerShell session secure credential
$passwordSecure = ConvertTo-SecureString -String $password -AsPlainText -Force
$cred = [pscredential]::new($username,$passwordSecure)

# Create object to hold command parameters
$parameters = @{

    ComputerName = $servers
    ScriptBlock = { & $exe -SOMEFLAG "SOME_FLAG_VALUE"}
    Credential = $cred

}

# Execute command
$commandOutput = Invoke-Command @parameters 

# Print command output to console
Write-Host $commandOutput
```