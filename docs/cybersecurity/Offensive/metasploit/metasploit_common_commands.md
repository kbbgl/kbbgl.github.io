```bash
msf5> use post/multi/recon/local_exploit_suggester

msf5> post(multi/recon/local_exploit_suggester) options 

msf5> set session 1

msf5> run
```

Get back to active meterpreter session:
```bash
sessions $SESSION_NUMBER

#ex:
msf5> post(multi/recon/local_exploit_suggester) sessions 1
```

Put session in background:
```bash
meterpreter> background
```

search for exploits
```bash
msf5> post(multi/recon/local_exploit_suggester) search ms16-032
```