# Read Data from Sockets

## Create a Socket Server

```bash
ncat -k -l -p 6566
```

`ncat` is not the same as `nc`, need to install it `apt install ncat`. It provides the `-k` flag which keeps the connection alive.

Once we have a server, we can connect using a client

## Send Data to Server

From another terminal window:

```bash
> nc localhost 6566

hello
this is a test message
```

We will see in the server the messages sent:

```bash
> ncat -k -l -p 6566
hello
this is a test message
```
