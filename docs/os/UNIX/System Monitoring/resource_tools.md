# System Resource Monitoring Tools

```bash
dstat

--total-cpu-usage-- -dsk/total- -net/total- ---paging-- ---system--
usr sys idl wai stl| read  writ| recv  send|  in   out | int   csw
  5   8  87   0   0|  41k 8151k|   0     0 |   0     0 |  13k   46k
  5  18  76   0   0|  16k 9848k|3252k 1737k|   0     0 |  43k   57k
  7  18  75   0   0|   0  4620k|3648k 2540k|   0     0 |  39k   49k
```

## Network traffic viewer

`iftop` – network traffic viewer. Ex:

```bash
sudo iftop -i eth0
```

## storage I/O statistics

```bash
iostat -h

Linux 4.15.0-106-generic (node1)  08/25/2020  _x86_64_ (16 CPU)

avg-cpu:  %user   %nice %system %iowait  %steal   %idle
           3.6%    0.0%   13.8%    5.4%    0.0%   77.2%

Device             tps    kB_read/s    kB_wrtn/s    kB_read    kB_wrtn
loop0
                  0.00         0.0k         0.0k       1.0M       0.0k
loop1
                  0.34         0.4k         0.0k      10.5M       0.0k
loop2
                  0.00         0.0k         0.0k     330.0k       0.0k
loop3
                  0.00         0.0k         0.0k       8.0k       0.0k
sda
                354.60        13.5M         1.5M     374.3G      41.3G
sdb
                  1.12        21.1k        45.2k     583.7M       1.2G
dm-0
                497.93        13.5M         1.8M     374.3G      48.7G
dm-1
                  0.01         0.2k         0.0k       4.2M       0.0k
```

run ever 1 seconds with MB output

```bash
iostat -h -m 1
```

focus on particular device:

```bash
iostat -p sda
```

more information:

```bash
iostat -m -p sda -x 
Linux 4.18.0-193.1.2.el8_2.x86_64 (rhel.test)     06/17/2020     _x86_64_    (4 CPU)
    
avg-cpu:  %user   %nice %system %iowait  %steal   %idle
           1.06    0.06    0.40    0.05    0.00   98.43
    
Device            r/s     w/s     rMB/s     wMB/s   rrqm/s   wrqm/s  %rrqm  %wrqm r_await w_await aqu-sz rareq-sz wareq-sz  svctm  %util
sda             12.20    2.83      0.54      0.14     0.02     0.92   0.16  24.64    0.55    0.50   0.00    45.58    52.37   0.46   0.69
sda2            12.10    2.54      0.54      0.14     0.02     0.92   0.16  26.64    0.55    0.47   0.00    45.60    57.88   0.47   0.68
sda1             0.08    0.01      0.00      0.00     0.00     0.00   0.00  23.53    0.44    1.00   0.00    43.00   161.08   0.57   0.00
```

`avgqu-sz` - average queue length of a request issued to the device
`await` - average time for I/O requests issued to the device to be served (milliseconds)
`r_await` - average time for read requests to be served (milliseconds)
`w_await` - average time for write requests to be served (milliseconds)

![tools](https://i.redd.it/lqqdjpphc4w61.png)
