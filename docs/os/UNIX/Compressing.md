# Compressing Files

## `tar`

### Create

```bash
tar -czvf filename.tar.gz /path/to/dir1
```

### List Files

```bash
tar -tvf file.tar.gz
```

The `dd` command makes a bit-by-bit copy of a file, a filesystem, or even an entire hard drive. This means that even deleted files are copied

```bash
dd if=/dev/sdb of=/tmp/sdb_copy
```

`if` is the input file, `of` is the output file. Other useful arguments are `bs` which is the block size of each copy (sector size is 4096 bytes) and `conv:noerror` which ignores errors.

## Concat Two Zips

```bash
zip -s 0 masked-logs-prod.zip --out full_logs.zip
```
