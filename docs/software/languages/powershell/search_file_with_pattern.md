# Find in File by Regex

```powershell
Get-Content reporting-service.log | Select-String -Pattern '\"error\":[1-9]'
```

```powershell
Get-Content galaxy.log | Select-String -Pattern "Exporting" | Select-String -Pattern "WARN|ERROR"
 ```

## Exclude Pattern from File

```powershell
Get-Content "C:\programdata\app\prismweb\logs\prismwebserver.log" -Wait -Tail 10 | Where-Object { $_ -Match "ERR*" -and $_ -NotMatch "ERR-[01]0"}
```
