# CONSUL CHEAT SHEET

`Consul` _is a distributed tool for service (VM, Instance, container)
discovery, configuration and orchestration._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## OVERVIEW

* SERVICE DISCOVERY

The location of services (VM/instance, etc..) frequently changes
by environment or when a new service version is deployed.
Services need to be able to find each other in this dynamic
environment without hardcoding configurations. `Services register`
themselves with Consul `so they can be discovered` by the
rest of the system.

* KV STORE

Can be used for dynamic configuration.
Configuration updates that take minutes or hours lead
to out-of-date service configurations and potential downtime.
`Consul updates configurations` for hundreds of thousands of
services across multiple regions in the sub-second timeframe
to maintain infrastructure resilience.

* HEALTH CHECKING

Centralized health monitoring is difficult to scale past
a few thousand services. `Consul scales health monitoring
for microservices` by performing health checks at the edges
rather than the center. The real-time health data is used
to immediately `remove unhealthy services from receiving traffic`.

* LOAD BALANCING

Multiple instances of a service need to balance traffic
between themselves. Consul `load balances traffic between
service instances` and `removes unhealthy instances` from
receiving traffic.

## INSTAL CONSUL ON UBUNTU

```bash
wget https://releases.hashicorp.com/consul/0.8.1/consul_0.8.1_linux_amd64.zip
sudo unzip consul_0.8.1_linux_amd64.zip -d /usr/bin
sudo chmod 755 /usr/bin/consul
rm consul_0.8.1_linux_amd64.zip
```

## AFTER INSTALL, VERSION

Check the version of consul,

```bash
consul -v
```

## CONSUL AGENT SERVER - NODES

When you install consul, you must start the Consul Agent Server.
Please see the documentation on how to do this.

The Agent maintains membership information, registers services,
runs checks, responds to queries, and more.

The Consul Server Agent must run on every node that is part
of a Consul Cluster. Ususlly 5 of these for production.

Query all the Nodes in this cluster via HTTP,

```bash
curl 10.240.1.114:8500/v1/catalog/nodes | jq
curl localhost:8500/v1/catalog/nodes | jq
```

Or query a Node via DNS,

```bash
dig @10.240.1.114 p-consul-1.node.consul
```

## RUN CONSUL AGENT CLIENT

A quick way to test Consul in development mode

```bash
/usr/bin/consul agent -dev
```

Then run in another window to test,

```bash
consul members
curl localhost:8500/v1/catalog/nodes
dig @127.0.0.1 -p 8600 p-jeff-test-hello-grp-20170419-2200-ql2l.node.consul
```

## METHOD 1 - REGISTER AND QUERY A SERVICE VIA SERVICE DEFINITION

Must have consul already installed.

`STEP 1` - Create directories,

```bash
sudo mkdir /etc/consul.d
sudo mkdir /var/consul
```

`STEP 2` - Create a Consul Service Definition File
/etc/consul.d/jeff-test-consul.json

```json
{
  "service": {
    "name": "p-jeff-test-hello-grp-20170419-2200-ql2l",
    "tags": ["jeff-test"],
    "address": "",
    "port": 8000,
    "enableTagOverride": false,
    "checks": [
      {
        "name": "jeffpingcheck",
        "DeregisterCriticalServiceAfter": "3m",
        "script": "/opt/pexeso/jeff-test-hello/healthcheck.sh",
        "interval": "10s"
      }
    ]
  }
}
```

Where /opt/pexeso/jeff-test-hello/healthcheck.sh could be

```bash
#!/bin/bash
myip=ip route get 8.8.8.8 | awk '{print $NF; exit}'
ping -c1 $myip >/dev/null || (exit 42)
```

`STEP 3` - Create a Consul configuration file /etc/consul.conf

Note, If you wanted to be a Agent Server, you would put it here.

```json
{
    "log_level": "debug",
    "datacenter": "us-central",
    "data_dir": "/var/consul",
    "server": false,
    "bind_addr": "{{ GetInterfaceIP \"eth0\" }}",
    "retry_join_gce": {
      "tag_value": "consulserver"
    }
}
```

Note: In GCE, the IP that your instances communicate through is usually set
up on eth0

`STEP 4` - Restart Consul Agent Client,

```bash
sudo /usr/bin/consul agent -config-dir /etc/consul.d/ \
-config-file=/etc/consul.conf -join 104.154.212.93 -syslog
```

Should get seomthing like this back

```text
           Version: 'v0.8.1'
           Node ID: '9bafb1a7-2ee0-f0b2-cfd2-9769fc009fda'
         Node name: 'p-jeff-test-hello-grp-20170419-2200-ql2l'
        Datacenter: 'us-central'
            Server: false (bootstrap: false)
       Client Addr: 127.0.0.1 (HTTP: 8500, HTTPS: -1, DNS: 8600)
      Cluster Addr: ????????????????? (LAN: 8301, WAN: 8302)
    Gossip encrypt: false, RPC-TLS: false, TLS-Incoming: false
             Atlas: <disabled>
```

## METHOD 2 - REGISTER AND QUERY A SERVICE VIA HTTP

Now that you have a Cluster of Nodes running (Agents),
lets register a service.

$ sudo mkdir /etc/consul.d

VIA CLI

Create a service json file

Then register to your Consul Agent Server via http,

```bash
curl -X PUT --data-binary @http1_checkhttp.json http://<ip>:8500/v1/agent/service/register
```

You can also register your vm/instance to the consul agent
server so it can monitor health.

## METHOD 3 - REGISTER AND QUERY A SERVICE VIA GO

You may also use go to register your service.

## CHECK HEALTH STATUS

To check health status,

```bash
curl http://<ip>:8500/v1/health/checks/<service_name> | jq
curl http://10.240.1.114:8500/v1/health/checks/\
p-stack-to-graph-grp-20170417-1707-fzhj | jq
```

## QUERY SERVICE

You may interact with consul via,

* Web Browser
* RESTful HTTP API (see below)
* CLI

VIA HTTP API

Once the agent is started and the service is synced,
we can query the service using either the DNS or HTTP API.

Using RESTful HTTP API to Query all services:

```bash
curl http://<ip>:8500/v1/catalog/services | jq
curl http://10.240.1.114:8500/v1/catalog/services | jq
```

or just for a particular service,

```bash
curl http://<ip>:8500/v1/catalog/service/<service_name> | jq
curl http://10.240.1.114:8500/v1/catalog/service\
/p-stack-to-graph-grp-20170417-1707-fzhj | jq
```

Query Consul DNS Server Database for the service

```bash
dig @<ip> <service_name>.service.consul
dig @10.240.1.114 p-stack-to-graph-grp-20170417-1707-fzhj.service.consul
```

With the port number,

```bash
dig @<ip> <service_name>.service.consul SRV
dig @10.240.1.114 p-stack-to-graph-grp-20170417-1707-fzhj.service.consul SRV
```

## REGISTER

TBD

## CONSUL AGENT (/CHECKS /SERVICES)

CHECKS,

```bash
curl http://10.240.1.114:8500/v1/agent/checks | jq
```

SERVICES,

```bash
curl http://10.240.1.114:8500/v1/agent/services
```

/health/checks/:service

## CONSUL TEMPLATE DAEMON - PROVIDING A WAY TO UPDATE A SERVICE CONFIG FILE

Consul template runs as a daemon and allows you to update
KEY/VALUE pairs. Its useful if you want to interject
KEY:VALUE into a running service or restart a service.

You update the KEY:VALUE via Consul.  This will
take the `template.ctmpl` and write to a configuration file
(e.g. config.yml).

```bash
$ consul-template \
    -template "/tmp/template.ctmpl:/tmp/config.yml"
```

Create a config file for your service we want to
pass configuration info to via Consul.

* Update KEY:VALUE via Consul (Browser, HTTP API or CLI)
* Consul template updates config.yml
* Restarts your service

```bash
/usr/bin/consul-template
-consul <ip>:<port>
-template "/your/path/template.ctmpl:/your/path/config.yml:systemctl restart {MAIN}.service"
-retry 30s -log-level info
```

To make life simple, you can have `consul-template` always running
as a service itself at boot.

The template file key value may looks like:

my_test_number = {{keyOrDefault "staging/jeff-test-hello/number" "45"}}

You have the template file on your instance/VM.  But you now need to
provide those hooks to Consul.

Just open consul browser, goto KEY/VALUE and copy your template file.
