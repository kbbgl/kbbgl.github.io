# Get Occurrences of String in File

```powershell
(Get-Content .\galaxy.log | Select-String "error running jaql:  Timeout Connection 30000 ms" |
 Measure-Object -line).Lines
```
