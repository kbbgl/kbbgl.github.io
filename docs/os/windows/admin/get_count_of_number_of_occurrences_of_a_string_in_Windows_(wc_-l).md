---
slug: os-windows-admin-get-count-of-number-of-occurrences-of-a-string-in-windows-wc-l
title: "Get Occurrences of String in File"
authors: [kbbgl]
tags: [os, windows, admin, get_count_of_number_of_occurrences_of_a_string_in_windows_wc_l]
---

# Get Occurrences of String in File

```powershell
(Get-Content .\galaxy.log | Select-String "error running jaql:  Timeout Connection 30000 ms" |
 Measure-Object -line).Lines
```
