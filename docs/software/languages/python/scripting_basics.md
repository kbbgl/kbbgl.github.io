# Installing Packages

modules can be found in [PyPi](http://www.pypi.org/).

We can install modules directly using `pip`:

```bash
apt install python3-pip
pip3 install $PACKAGE_NAME
```

The package is automatically placed in `/usr/local/$python-version/dist-packages`.

To verify the location:

```bash
pip3 show $PACKAGE_NAME
```

To manually install a package:

```bash
wget $PACKAGE_URL
python3 setup.py install
```

## Running Python Scripts

Create script:

```bash
# test.py

#!/usr/bin/python3

...
```

Give execution permissions:

```bash
chmod u+x test.py
```

Run script:

```bash
./test.py
```

## Opening Connections

```bash
# banner_grab.py
#!/usr/bin/python3

import socket

s = socket.socket()
s.connect(("192.168.1.1", 22))
answer = s.recv(1024)

print(answer)

s.close()
```

## Listening for Connections

```bash
# tcp_server.py

#!/usr/bin/python3

import socket

TCP_IP = "192.168.181.191"
TCP_PORT = 6996
BUFFER_SIZE=100

s = socket.socket(socket.AF_NET, socket.SOCK_STREAM)
s.bind((TCP_IP, TCP_PORT))
s.listen(1)

conn, addr = s.accept() 

print ('Connection address: ', addr ) 

while 1: 
 data=conn.recv(BUFFER_SIZE) 
    if not data: break 
    print ("Received data: ", data)
     conn.send(data) #echo 

conn.close

```
