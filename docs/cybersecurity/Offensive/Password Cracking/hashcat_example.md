#### run through list of hashes (`example0.hash`) and run brute force attack (`-a 3`) for 8-letter lower case password
```bash
hashcat -a 3 example0.hash ?l?l?l?l?l?l?l?l 
```

#### run through list of hashes (`example0.hash`) and run brute force attack (`-a 3`) for 6-letter lower case with 2 digits at the end 
```bash
hashcat -a 3 example0.hash ?l?l?l?l?l?l?d?d 
```

#### Crack MD5 password (`-m 0`), dictionary attack (`-a 0`) and output to `cracked.txt`, `target_hashes.txt` is the input file and `rockyou.txt` is the wordlist for dictionary attack 
```bash
hashcat -m 0 -a 0 -o cracked.txt target_hashes.txt /usr/share/wordlists/rockyou.txt
```

[Complete list of hash types](https://hashcat.net/wiki/doku.php?id=example_hashes)