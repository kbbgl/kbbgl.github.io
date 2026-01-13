---
title: Azure Authentication using IMDS
slug: azure-auth-imdsv2
app2or: kgal-akl
tags: [devops, azure, aks, az, imds, metadata]
---

```powershell
Invoke-RestMethod -Headers @{"Metadata"="true"} -Method GET -Uri "http://169.254.169.254/metadata/instance?api-version=2025-04-07" | ConvertTo-Json -Depth 64
```

