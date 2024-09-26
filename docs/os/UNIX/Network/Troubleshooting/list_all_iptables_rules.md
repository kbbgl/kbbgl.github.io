# List All IP Tables Rules

```bash
sudo iptables -L -v

...
Chain cali-FORWARD (1 references)
 pkts bytes target     prot opt in     out     source               destination
1882K 1570M MARK       all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* cali:vjrMJCRpqwy5oRoX */ MARK and 0xfff1ffff
1882K 1570M cali-from-hep-forward  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* cali:A_sPAO0mcxbT9mOV */ mark match 0x0/0x10000
1881K 1570M cali-from-wl-dispatch  all  --  cali+  *       0.0.0.0/0            0.0.0.0/0            /* cali:8ZoYfO5HKXWbB3pk */
 3212  758K cali-to-wl-dispatch  all  --  *      cali+   0.0.0.0/0            0.0.0.0/0            /* cali:jdEuaPBe14V2hutn */
 1870  112K cali-to-hep-forward  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* cali:12bc6HljsMKsmfr- */
 1870  112K ACCEPT     all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* cali:MH9kMp5aNICL-Olv */ /* Policy explicitly accepted packet. */ mark match 0x10000/0x10000
...
```
