When transferring text streams over a network, use `base64` encoding. This is because it is the most widely-used charset so it would usually be transferred uncorruputed across the wire.

```bash
# use -n to not print newline
echo -n 'tomcat:tomcat' | base64
```