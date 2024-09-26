---
slug: smb-client-cheatsheet
title: SMB Client Cheatsheet
description: SMB Client Cheatsheet
authors: [kbbgl]
tags: [cybersecurity,offensive,windows,smb,cheatsheet]
---


## List files
```bash
smbclient -L \\\\10.10.10.10\\path\\to\\smb PASSWORD
```

```
Server time is Sat Aug 10 15:58:27 1996
Timezone is UTC+10.0
Password: 
Domain=[WORKGROUP] OS=[Windows NT 3.51] Server=[NT LAN Manager 3.51]

Server=[ZIMMERMAN] User=[] Workgroup=[WORKGROUP] Domain=[]

        Sharename      Type      Comment
        ---------      ----      -------
        ADMIN$         Disk      Remote Admin
        public         Disk      Public 
        C$             Disk      Default share
        IPC$           IPC       Remote IPC
        OReilly        Printer   OReilly
        print$         Disk      Printer Drivers


This machine has a browse list:

        Server               Comment
        ---------            -------
        HOPPER               Samba 1.9.15p8
        KERNIGAN             Samba 1.9.15p8
        LOVELACE             Samba 1.9.15p8
        RITCHIE              Samba 1.9.15p8
        ZIMMERMAN  
```

When connection established:
```bash
smb: \> h
ls             dir            lcd            cd             pwd            
get            mget           put            mput           rename         
more           mask           del            rm             mkdir          
md             rmdir          rd             prompt         recurse        
translate      lowercase      print          printmode      queue          
cancel         stat           quit           q              exit           
newer          archive        tar            blocksize      tarmode        
setmode        help           ?              !              
smb: \> 
```