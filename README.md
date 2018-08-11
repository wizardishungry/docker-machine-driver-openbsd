# docker-machine-driver-openbsd
🐡 docker-machine driver plugin for vmm (native OpenBSD hypervisor) 

Work in progress.

## Caveats

* Doesn't work on my machine on 6.3 after installing latest syspatch
* Container host is insecure on multiuser system

## Hacking

Edit `/etc/pf.conf`
```
# NAT for the VMs
match out on egress from 100.64.0.0/10 to any nat-to (egress)
# Pass VM DNS to Cloudfront
pass in proto udp from 100.64.0.0/10 to any port domain rdr-to 1.1.1.1 port domain
```
Enable `vmd`
```bash
doas rcctl enable vmd
doas pfctl -f /etc/pf.conf
```
Build the driver
```bash
go get -u github.com/WIZARDISHUNGRY/docker-machine-driver-openbsd
```
Try to build a docker-machine
```bash
docker-machine create -d openbsd default
```