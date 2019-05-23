# GoDC - Datacenter and Cluster Management System
At the moment it's a simple CLI to gather VM info from host machines in datacenter, but in the future will manage and administer multiple datacenters.

## Quickstart
Download the godc binary:
```bash
wget https://github.com/girikuncoro/godc/releases/download/v0.3.1/godc-darwin-amd64
mv godc-darwin-amd64 /usr/local/bin/godc
chmod +x /usr/local/bin/godc
```

Create config file that contains list of host, DNS, and DHCP:
```bash
cat << EOF > datacenter.yaml
hosts:
  - <libvirt-host-1-ip>:<tcp-listen-port>
  - <libvirt-host-2-ip>:<tcp-listen-port>
  - <libvirt-host-3-ip>:<tcp-listen-port>
dhcpServer: <dhcp-ip-address>
dnsServer: <dns-ip-address>
EOF
```

Install the configuration:
```bash
godc install -f ./datacenter.yaml
```

List out VMs (make sure all hosts are reachable):
```bash
godc get vm
```

## Dev Dependencies
TODO (giri): This should be automated
```
$ brew install pkg-config
$ brew install libvirt
```
