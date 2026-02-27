---
slug: os-linux-system-monitoring-check-volume-groups-size-free
title: "Check Volume Groups Size"
authors: [kbbgl]
tags: [os, linux, system_monitoring, check_volume_groups_size_free]
---

# Check Volume Groups Size

## Physical

```bash
ubuntu@ip-10-11-58-68:~$ kubectl exec glusterfs-dp4th -- pvs
  PV           VG                                  Fmt  Attr PSize   PFree
  /dev/nvme1n1 vg_e827dd98ecf1ea496850de5810917757 lvm2 a--  999.87g 30.06g
```

## Virtual

```bash
ubuntu@ip-10-11-58-68:~$ kubectl exec glusterfs-dp4th -- vgs
  VG                                  #PV #LV #SN Attr   VSize   VFree
  vg_e827dd98ecf1ea496850de5810917757   1  16   0 wz--n- 999.87g 30.06g
  ```
