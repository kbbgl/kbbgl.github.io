Script can be used when we need to sniff messages on a particular service.
Since every service has an address/port which assigned by calico, we can use the routing table to sniff the traffic of that particular service.

```bash
# retrieve 
ip=$(kubectl get endpoints galaxy -o=jsonpath='{.subsets[0].addresses[0].ip}')

interface=$(ip route | grep $ip | cut -d" " -f3)

sudo tcpdump -i $interface -s 0 -A
```