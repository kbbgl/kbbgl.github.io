## DNS

DNS has been provided as CoreDNS by default as of v1.13. The use of CoreDNS allows for a great amount of flexibility. Once the container starts, it will run a Server for the zones it has been configured to serve. Then, each server can load one or more plugin chains to provide other functionality. As with other microservices, clients would it access using a service, kube-dns.

### Verifying DNS Registration

To make sure that your DNS setup works well and that services get registered, the easiest way to do it is to run a pod with a shell and network tools in the cluster, create a service to connect to the pod, then exec in it to do a DNS lookup.

Troubleshooting of DNS uses typical tools such as `nslookup`, `dig`, `nc`, `wireshark` and more. The difference is that we leverage a service to access the DNS server, so we need to check labels and selectors in addition to standard network concerns.

Other steps, similar to any DNS troubleshooting, would be to check the `/etc/resolv.conf` file of the container, as well as Network Policies and firewalls. We will cover more on Network Policies in the Security chapter.