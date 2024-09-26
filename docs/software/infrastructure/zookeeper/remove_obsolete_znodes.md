---
slug: zk-rm-obsolete-znodes
title: Remove Obsolete ZNodes
authors: [kbbgl]
tags: [zk, query, k8s, windows]
---

```bash
zkCli.sh -server localhost:2181 ls /my_app | awk -F'[][]' '{print $2}' | tr ',' '\n' | sed -e 's/^[[:space:]]*//' | grep "_" > /tmp/my_app.txt
```

add connector path to each line

```bash
sed -i -e 's/^/rmr \/my_app\//' my_app.txt
```

```bash
./bin/zkCli.sh -server localhost:2181 <<EOF
rmr /my_app/query_35554281-a497-4f8a-b828-7681fc7a74b0
rmr /my_app/query_6a3726d1-a0a6-45c5-a3b6-4edd2db2dd6f
rmr /my_app/query_6f7542a1-523a-46db-9374-938757df41ae
rmr /my_app/query_9bcea156-f255-4b3a-abd9-70f44c690824
...
quit
EOF
```

Full script

```bash
echo "./bin/zkCli.sh -server localhost:2181 <<EOF" >> clean_zk_my_app.sh
zkCli.sh -server localhost:2181 ls /my_app | awk -F'[][]' '{print $2}' | tr ',' '\n' | sed -e 's/^[[:space:]]*//' | grep "_" | sed -e 's/^/rmr \/my_app\//' >> clean_zk_my_app.sh
echo quit >> clean_zk_my_app.sh
echo EOF >> clean_zk_my_app.sh
chmod u+x clean_zk_my_app.sh
. clean_zk_my_app.sh
```
