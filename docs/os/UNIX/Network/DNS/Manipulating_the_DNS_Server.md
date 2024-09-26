# Manipulating DNS Server

For instance, enter `dig hackers-arise.com` and add the `ns` option (short for nameserver ). The nameserver for hackers-arise.com is displayed in the ANSWER SECTION of Listing 3-3 .

```bash
dig hackers-arise.com ns 

#-- snip -- ;; 

# QUESTION SECTION: ;

# hackers-arise.com. IN NS ;; 

# ANSWER SECTION: 
# hackers-arise.com. 5 IN NS ns7.wixdns.net.
# hackers-arise.com. 5 IN NS ns6.wixdns.net.

# ;; 
# ADDITIONAL SECTION: 
# ns6.wixdns.net. 5 IN A 216.239.32.100 # -- snip --
```

`ADDITIONAL SECTION` includes the DNS server serving host.

Get information about SMTP servers, can use

```bash
dig hackers-arise.com mx
```
