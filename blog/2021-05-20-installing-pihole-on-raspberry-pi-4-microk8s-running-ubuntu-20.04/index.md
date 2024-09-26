---
slug: install-pihole-rpi4-microk8s-ubuntu-2004
title: Installing PiHole On Raspberry Pi 4, MicroK8s running Ubuntu 20.04 (focal)
description: Fill me up!
authors: [kbbgl]
tags: [docker,how-to,k8s,kubernetes,linux,pihole,raspberry_pi,troubleshooting]
---

## PiHole, What’s That?

The [Wikipedia definition](https://en.wikipedia.org/wiki/Pi-hole) should be sufficient in explaining what the software does:

> Pi-hole or Pihole is a Linux network-level advertisement and Internet tracker blocking application which acts as a DNS sinkhole and optionally a DHCP server, intended for use on a private network

I wanted to deploy it for a few reasons:

- I have a spare Raspberry Pi 4 lying around.
- Because I’m working on getting my CKAD (Certified Kubernetes Application Developer) certification and thought it would be a great hands-on practice.
- I couldn’t find a good enough article that described how to install PiHole on Kubernetes. The majority did not go throught the whole procedure, were aimed for Docker/Swarm and Raspbian (Raspberry Pi flavored Linux distribution).
- I got tired of all the advertisements and popups on all the devices while surfing the web at home.

This post is here to explain how was able to deploy PiHole on Kubernetes and how I resolved some of the problems that occurred during the deployment process.

## Setting Up Kubernetes on Ubuntu

There are a few ways to do this. I was looking at the different options but decided to choose MicroK8s in the end (over Charmed Kubernetes or kubeadm) simply because the Canonical team maintains it (Canonical is the publisher of Ubuntu) so I thought it would be the wisest decision long term as any kernel/software upgrades on the OS level would likely be QA’d in the future in accordance with MicroK8s maintenance.
Since MicroK8s is bundled as a snap (an additional package manager for Ubuntu), it already includes all the binaries necessary to set up Kubernetes. So we can run the following command to install it:

```bash
sudo snap install microk8s --classic
```

We also need to ensure that we’re allowing the different Kubernetes components to communicate with each other. To modify the firewall settings, we run the following:

```bash
sudo ufw allow in on cni0 && sudo ufw allow out on cni0
sudo ufw default allow routed
```

We also need to enable a DNS for the Kubernetes deployment. To do this we run:

```bash
microk8s enable dns
```

We can then verify that the basic Kubernetes resources are up and running by:

```bash
microk8s kubectl get all --all-namespaces
```

We should see that the kube-system namespace has the CNI (by  default Calico) controllers and node, coredns are running.
You may have noticed that we need to prefix the `kubectl` commands with microk8s, which bothered me a bit because I was used to interacting with the Kubernetes API server using only `kubectl [ACTION] [RESOURCE]`. I decided to install `kubectl` from the `snap` store to prevent typing out an extra prefix:

```bash
snap install kubectl --classic
```

We now have a running bare Kubernetes single node and ready to create all the necessary resources!

## Creating Storage and Kubernetes Resources

One of the great advantages of Kubernetes is the ability to isolate applications into their own scope, called Namespaces. Since I did see myself using the cluster for other projects in the future, I thought it would be practical to separate the Pihole project from the future projects into its own Namespace. To create the new namespace, I ran the following:

```bash
kubectl create namespace pihole
```

Since all of the following commands will be run in this namespace (and we want to save typing `-n pihole` in every command), we can just set the context of the following commands to the newly-created namespace. To do this:

```bash
kubectl config set-content --current --namespace pihole
```

Next I tackled the subject of storage. Since Pihole requires some persistent storage for configuration files, logs and data for its SQLLite database, we need to create a place in the filesystem that will be used as the mount for the persistent storage resources we’ll set up for the Pod. So I created a directory in my home directory to hold it:

```bash
mkdir ~/pihole/data
```

Make sure that you have enough space in the directory you choose (you can use `df u /path/to/pihole/data` to verify).
Next, we need to create some resources to bind the host (Raspberry Pi) filesystem to the Kubernetes resources which will run the Pihole container. We need to create 3 things:

1. Default `StorageClass` – This is just a resource that we use to describe the different types of storage the Kubernetes deployment offers. In our case, the `StorageClass` will be a simple one that’s provisioned by the local machine.
1. 2 `PersistentVolume`s – These are abstractions of the volumes we’ll be using to store the data for Pihole. We need to specify 2 of these, one for the assets stored in the /etc filesystem of the container (such as DNS server lists, domain lists , pihole FTL configuration, etc) and one for the `dnsmasq` filesystem (includes the initial Pihole configuration).
1. 2 `PersistentVolumeClaim`s – These are the actual requests for storage from the PVs created above.

To create the `StorageClass`, I defined the following specification called `storageclass.yaml`:

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: manual
provisioner: manual
reclaimPolicy: Delete
volumeBindingMode: Immediate
```

The `volumeBindingMode: Immediate` ensures that upon specifying the `PersistentStorage` to the same provisioner type (manual), the `StorageClass` will immediately bind to it. The `reclaimPolicy` ensures that ones the `PersistentVolumeClaims` are discarded, so will the `StorageClass`.
We can create the resource by running:

```bash
kubectl apply -f storageclass.yaml
```

Now that we have a `StorageClass` set up, we can create the 2 PVs (`volume-etc.yaml` and `volume-dnsmasq.yaml`, respectively):

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: pihole-volume
  labels:
    type: local
spec:
  storageClassName: manual
  capacity:
    storage: 1Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: "/home/ubuntu/pihole/data/"
```

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: pihole-dnsmasq
  labels:
    type: local
spec:
  storageClassName: manual
  capacity:
    storage: 1Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: "/home/ubuntu/pihole/data/"
```

We’re defining 2 volumes with `1GB` of storage to bind with `StorageClass` named manual in the host filesystem path `~/pihole/data` (same path we created in the first step above). Keep in mind this path as we’ll come back to this later.
We can create both PVs by running:

The final step to set up the storage is to create the two `PersistentVolumeClaims`. Here are the two specifications for them (`claim-etc.yaml` and `claim-dnsmasq.yam`l, respectively):

And create the PVCs:

```bash
kubectl apply -f volume-etc.yaml
kubectl apply -f volume-dnsmasq.yaml
```

Great, let’s verify that the storage is running set up correctly. It should look something like this:

```bash
kubectl get sc,pv,pvc
NAME                                 PROVISIONER   RECLAIMPOLICY   VOLUMEBINDINGMODE   ALLOWVOLUMEEXPANSION   AGE
storageclass.storage.k8s.io/manual   manual        Delete          Immediate           false                  3m
NAME                              CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                            STORAGECLASS   REASON   AGE
persistentvolume/pihole-dnsmasq   1Gi        RWO            Retain           Bound    pihole/pihole-dnsmasq-pv-claim   manual                  2m
persistentvolume/pihole-volume    1Gi        RWO            Retain           Bound    pihole/pihole-etc-pv-claim       manual                  2m
NAME                                            STATUS   VOLUME           CAPACITY   ACCESS MODES   STORAGECLASS   AGE
persistentvolumeclaim/pihole-dnsmasq-pv-claim   Bound    pihole-dnsmasq   1Gi        RWO            manual         1m
persistentvolumeclaim/pihole-etc-pv-claim       Bound    pihole-volume    1Gi        RWO            manual         1m
```

We can see that the `PersistentVolumeClaims`s are bound to their respective `PersistentVolume` and the `PersistentVolume`s are bound to the `StorageClass`. Looks good!
Now that we have the storage set up, we need to create two more specifications:

1. A `Service` to specify the access to Pihole.
1. A `Deployment` which will pull the latest Pihole image from Dockerhub, create a container from this image, allow the Pod to use the storage we set up to hold its data and configuration files and bind to a Service so that we can access Pihole dashboard.

Let’s start with the service `svc.yaml`. We’ll expose the Pihole server externally so we can access it from within our network. To do that, we need to know our internal IP address. We can find it easily by running:

```bash
hostname -i | cut -d " " -f1
10.100.102.95
```

We can then use that value to finish off our `Service` specifications:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: pihole
spec:
  selector:
    app: pihole
  clusterIP: 10.152.183.2
  ports:
    - port: 80
      targetPort: 80
      name: pihole-admin
    - port: 53
      targetPort: 53
      protocol: TCP
      name: dns-tcp
    - port: 53
      targetPort: 53
      protocol: UDP
      name: dns-udp
  externalIPs:
  - 10.100.102.95
```

We then need to create the `Deployment` which binds the storage and network settings together. This is what it looks like (`deployment.yaml`):

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: pihole
  labels:
    app: pihole
spec:
  replicas: 1
  selector:
    matchLabels:
      app: pihole
  template:
    metadata:
      labels:
        app: pihole
        name: pihole
    spec:
      containers:
        - name: pihole
          image: pihole/pihole:latest
          env:
            - name: TZ
              value: 'Asia/Jerusalem'
            - name: WEBPASSWORD
              value: 'YOUR_PASSWORD'
            - name: TEMPERATUREUNIT
              value: c
          volumeMounts:
            - name: pihole-local-etc-volume
              mountPath: '/etc/pihole'
            - name: pihole-local-dnsmasq-volume
              mountPath: '/etc/dnsmasq.d'
      volumes:
        - name: pihole-local-etc-volume
          persistentVolumeClaim:
            claimName: pihole-etc-pv-claim
        - name: pihole-local-dnsmasq-volume
          persistentVolumeClaim:
            claimName: pihole-dnsmasq-pv-claim
```

Most of this is boilerplate and optional. The important sectors are the volumes where we specify the PVC bindings and the matchLabels which binds the `Service` to the `Deployment`. Also, you can set a password for the Pihole dashboard admin page by changing the value of `spec.template.spec.containers[0].WEBPASSWORD`.
I ran the following command to create the deployment:

```bash
kubectl apply -f deployment.yaml
```

Unfortunately, I noticed that the pihole `Pod` was in a `CrashLoopBackoff`!

## Troubleshooting PiHole Pod CrashLoopBackOff

The first step in troubleshooting a `CrashLoopBackoff` is to review the `Pod` logs. This is what I saw:

```bash
kubectl logs pihole-64678974cd-p7spj 
# ...
::: Preexisting ad list /etc/pihole/adlists.list detected ((exiting setup_blocklists early))
https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts
dnsmasq: bad option at line 1 of /etc/dnsmasq.d/adlists.list
::: Testing pihole-FTL DNS: [cont-init.d] 20-start.sh: exited 1.
[cont-finish.d] executing container finish scripts...
[cont-finish.d] done.
[s6-finish] waiting for services.
[s6-finish] sending all processes the TERM signal.
```

So it seemed that there was some sort of unexpected issue when reading a file named adlists.list. But since the `Pod` was in a `CrashLoopBackoff`, I could not have direct access to the Pod to check the file because it was constantly restarting.
Therefore, I went the hard route and decided to download the pihole image to review the source code and pinpoint the failure.

## Installing Docker to Troubleshoot Image Initialization Failure

To install Docker on the Raspberry Pi, I needed to figure out first what architecture the processor is running.
I ran the following command and found that the I was running `aarch64`:

```bash
uname -m
aarch64
```

But when reviewing the Docker documentation how to install the engine on Ubuntu, I saw that the only tabs available were `x86_64`/`amd64`, `armhf` or `arm64`.
So I did some research and found that the GNU triplet for the 64-bit ISA is `aarch64`. So essentially `aarch64` is `arm64`.
I ran the following commands to install Docker:

```bash
sudo apt install apt-transport-https ca-certificates curl gnupg lsb-release
echo \\n  "deb [arch=arm64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \\n  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt update
sudo apt install docker-ce docker-ce-cli containerd.io
```

I then created the pihole container:

```bash
sudo docker run -it pihole/pihole bash
```

On another ssh session, I ran the following command on the host machine to download the content of the pihole image into an archive `pihole.tar`:

```bash
# -l returns the latest container
# -q quiet output
container_id=`sudo docker ps -lq`
echo $container_id
4565bc8fe1e1
sudo docker export $container_id -o pihole.tar
```

Now that I had the contents of the image, I can terminate the container by simply exiting the session created by docker run -it.
I could then decompress the tar:

```bash
mkdir /tmp/pihole
tar xvf pihole.tar -C /tmp/pihole
```

and review the source code to check why startup was failing.
Since I knew that the failure occurs on the following line:

```plaintext
::: Testing pihole-FTL DNS: [cont-init.d] 20-start.sh: exited 1.
```

I could use the search for that particular line in the whole image filesystem I just extracted and see which file has the logic.

```bash
sudo grep -rnw "Testing pihole" ./* --exclude pihole.tar
./bash_functions.sh:260:    echo -n '::: Testing pihole-FTL DNS: '
```

So we can see that this line is called in the shell script `bash_functions.sh` in line 260. This is what the scope looks like:

```bash
test_configs() {
     set -e
     echo -n '::: Testing pihole-FTL DNS: '
     sudo -u ${DNSMASQ_USER:-root} pihole-FTL test || exit 1
     echo -n '::: Testing lighttpd config: '
     lighttpd -t -f /etc/lighttpd/lighttpd.conf || exit 1
     set +e
     echo "::: All config checks passed, cleared for startup ..."
}
```

So seems that we’re running the command `pihole-FTL test`  as `root` and send an exit code of 1 if the command fails (which it does in this case). The next step is to figure out what’s the command: `pihole-FTL test`.
We can find the binary by searching for it in the whole extracted image filesystem:

```bash
find /tmp/pihole -name "*pihole-FTL*"
/usr/bin/pihole-FTL
```

I decided to recreate the container so I could interact with this binary:

```bash
# Inside pihole container
root@371cad2a9105:/# pihole-FTL --help
pihole-FTL - The Pi-hole FTL engine
Usage:    sudo service pihole-FTL <action>
where '<action>' is one of start / stop / restart
Available arguments:
            debug           More verbose logging,
                            don't go into daemon mode
            test            Don't start pihole-FTL but
                            instead quit immediately
        -v, version         Return FTL version
        -vv                 Return more version information
        -t, tag             Return git tag
        -b, branch          Return git branch
        -f, no-daemon       Don't go into daemon mode
        -h, help            Display this help and exit
        dnsmasq-test        Test syntax of dnsmasq's
                            config files and exit
        regex-test str      Test str against all regular
                            expressions in the database
        regex-test str rgx  Test str against regular expression
                            given by rgx
        --lua, lua          FTL's lua interpreter
        --luac, luac        FTL's lua compiler
        dhcp-discover       Discover DHCP servers in the local
                            network
        sqlite3             FTL's SQLite3 shell
Online help: https://github.com/pi-hole/FTL
```

The strange is that running the same command, in a Docker container where PiHole loads successfully, also returns an exit code of 1:

```bash
# Inside pihole container
root@371cad2a9105: sudo -u pihole-FTL test;echo $?
1
```

This threw me off a bit. My main question was: How come the Docker container’s initialization script finishes successfully but the same container running in a Kubernetes Pod fails?
I had to take a step back and try to understand what’s the biggest difference between the Docker method that works and the Kubernetes method that fails. The only possible differences that I could think of were (1) no port forwarding, host network on Docker (2) no presistent storage set up in Docker.
Since I always seem to have a problem with using the `StorageClass`, `PersistentVolume` and `PersistentVolumeClaim` APIs in Kubernetes, my gut told me that I had mis-configured something there.
When navigating back to the mount path that I set in the PV YAML specification, I noticed that in both PVs (etc, `dnsmasq`) I set the path to be the same, namely `/home/ubuntu/pihole/data`. I decided I would create two different mount points, and retest it.
Below are the modifications I made (see comment):

```yaml
# volume-etc.yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: pihole-volume
  labels:
    type: local
spec:
  storageClassName: manual
  capacity:
    storage: 1Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: "/home/ubuntu/pihole/data/etc" # Changed from /home/ubuntu/data/
```

```yaml
# volume-dnsmasq.yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: pihole-dnsmasq
  labels:
    type: local
spec:
  storageClassName: manual
  capacity:
    storage: 1Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: "/home/ubuntu/pihole/data/dnsmasq"
```

I recreated all Kubernetes resources after the updates and found that the Pod ran successfully!

```bash
watch -t kubectl get pv,pvc,sc,deploy,svc,pod
NAME                              CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                            STORAGECLASS   REASON   AGE
persistentvolume/pihole-dnsmasq   2Gi        RWO            Retain           Bound    pihole/pihole-dnsmasq-pv-claim   manual                  13m  
persistentvolume/pihole-volume    2Gi        RWO            Retain           Bound    pihole/pihole-etc-pv-claim       manual                  13m  
NAME                                            STATUS   VOLUME           CAPACITY   ACCESS MODES   STORAGECLASS   AGE
persistentvolumeclaim/pihole-dnsmasq-pv-claim   Bound    pihole-dnsmasq   2Gi        RWO            manual         13m  
persistentvolumeclaim/pihole-etc-pv-claim       Bound    pihole-volume    2Gi        RWO            manual         12m  
NAME                                 PROVISIONER   RECLAIMPOLICY   VOLUMEBINDINGMODE   ALLOWVOLUMEEXPANSION   AGE
storageclass.storage.k8s.io/manual   manual        Delete          Immediate           false                  7h4m 
NAME                     READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/pihole   1/1     1            1           12m  
NAME             TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)                AGE
service/pihole   ClusterIP   10.152.183.2   10.100.102.95        80/TCP,53/TCP,53/UDP   12m  
NAME                          READY   STATUS    RESTARTS   AGE
pod/pihole-64678974cd-mxwcw   1/1     Running   0          12m
```

Nice! Now we have PiHole running locally on the server!
The last step was to ensure that we direct all DNS queries from our home network to the Pihole DNS server instead of our router.

## Changing DNS Settings on Devices

The first device I changed was a 2013 MacBook Air running the latest Ubuntu Desktop that I used to SSH into the Raspberry Pi to set up Pihole.
It was quite easy to change the DNS server. Just go to _WiFi > Choose currently connected WiFi network kegwheel > IPv4 tab > Disable Automatic DNS and set the IP of your Raspberry Pi IP address_ (Search up for `hostname -i` in this page to see the command again). After I restarted my laptop, I began seeing lots of queries getting blocked:

![rpi](https://tilsupport.files.wordpress.com/2021/05/image-1.png?w=1024)

## Setting Up Static IP for the Cluster

Unfortunately for me, I’m using my ISP’s router and not one I have full control of. This means that I do not have access to some of the router settings such as configuring the DNS server for the whole network and I can’t also control the DHCP settings. This is an important limitation because I cannot control how the router assigns IPs and I was anticipating a scheduled job I have configured in crontab to reboot the server at some point causing the router to assign a new IP within the specified range. I needed to somehow ensure that the Raspberry Pi requests a specific (static) IP address to be assigned to it after reboot.
To set up a static IP address, we need to use a tool called [`netplan`](https://netplan.io/), the default Ubuntu network configuration utility.
But first, we must confirm which interface we’re going to configure to request the static IP for. In my case, I was connected to the router using the WiFi interface `wlan0`. You can find yours by running `ip link` although it’s usually `en0` for Ethernet (wired) connection and `wlan0` for WiFi.
Next, we can start interacting with the `netplan` configuration.
First, let’s create a backup file (always good practice):

```bash
sudo cp /etc/netplan/50-cloud-init.yaml /etc/netplan/50-cloud-init.yaml.bak
```

Next, let’s open the `50-cloud-init.yaml` configuration file to add the necessary configuration. In my case, since the interface is wlan0, I will be modifying the `network.wifis.wlan0` object but it should be the same for `network.ethernets.eth0` in case you use Ethernet.
Let’s open the file for editing:

```bash
sudo vim /etc/netplan/50-cloud-init.yaml
```

And see the comments for the added fields:

```yaml
network:
  ethernets:
    eth0:
      dhcp4: true
      optional: true
  version: 2
  wifis:
    wlan0:
      optional: true
      addresses: # add section
        - 10.100.102.95/24 # add node IP
      gateway4: 10.100.102.1 # add router IP
      nameservers: # add section
        addresses: [10.100.102.95, 8.8.8.8] # add node IP and Google as alternate DNS
      access-points:
        "YOUR_AP_NAME":
          password: "YOUR_AP_PW"
      dhcp4: false # true -> false
```

To know what’s your router/gateway IP for `gateway4`, you can run the following command:

```bash
ip route | grep default | cut -d " " -f3 | head -n1
10.100.102.1
```

We can then apply the changes running:

```bash
sudo netplan apply
```

And we can confirm it’s set up by running:

```bash
ip addr show dev wlan0
3: wlan0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether dc:a6:32:be:b8:c4 brd ff:ff:ff:ff:ff:ff
    inet 10.100.102.95/24 brd 10.100.102.255 scope global dynamic wlan0
       valid_lft 3550sec preferred_lft 3550sec
```

We can then restart the Raspberry Pi for the changes to take effect. Now we’ll be able to run maintenance on the device that requires restarts without breaking the whole cluster because of an incorrect address assignment by the router.

## Maintenance, `kubectl drain/uncordon`

After about a week of running, I performed an `apt upgrade` which required me to reboot the server for the changes to apply. I needed to take the Pihole (and the whole Raspberry Pi) down.
To do this, we first need to ensure that we evict all running Kubernetes resources and that we stop scheduling of new Pods to the cluster.
We need to run the following sequence of commands to be able to ensure that our server restart runs smoothely.

```bash
# Get node name
node_name=`kubectl get node -o=jsonpath='{.items[0].metadata.labels}' | jq '."kubernetes.io/hostname"'`
 
# Drain node, will evict all resources
kubectl drain $node_name --ignore-daemonsets --delete-local-data --force
  
# Reboot the server
sudo reboot
```

After reboot, wait until `kubelet.service` is up and run:

```bash
node_name=`kubectl get node -o=jsonpath='{.items[0].metadata.labels}' | jq '."kubernetes.io/hostname"'`
 
kubectl uncordon $node_name
```

We should then see that the Node is schedulable and that the all Pihole resources are back up and running.
