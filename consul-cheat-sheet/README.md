# CONSUL CHEAT SHEET

`consul` _is a distributed tool for service (VM, Instance, container)
discovery, configuration and orchestration._

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## OVERVIEW

* Uncertain service locations:  Register with consul.

The location of services (VM/instance, etc..) frequently changes
by environment or when a new service version is deployed.
Services need to be able to find each other in this dynamic
environment without hardcoding configurations. Services register
themselves with Consul so they can be discovered by the
rest of the system.

Stale service configurations: Update configurations.

Configuration updates that take minutes or hours lead
to out-of-date service configurations and potential downtime.
Consul updates configurations for hundreds of thousands of
services across multiple regions in the sub-second timeframe
to maintain infrastructure resilience.

Unscalable health monitoring: Health checks

Centralized health monitoring is difficult to scale past
a few thousand services. Consul scales health monitoring
for microservices by performing health checks at the edges
rather than the center. The real-time health data is used
to immediately remove unhealthy services from receiving traffic.

Load balancing

Multiple instances of a service need to balance traffic
between themselves. Consul load balances traffic between
service instances and removes unhealthy instances from
receiving traffic.

ADDING CONFIGURATION HOOKS

DYNAMIC CONFIGURATION

Flexible key/value store for dynamic configuration,
feature flagging, coordination, leader election and more.
Long poll for near-instant notification of configuration changes.

There are 3 main components of consul: consul itself, the consul agent
and the consul template daemon.

## 1. CONSUL

Runs as a cluster on you server.  You may interact with consul via,

* Web Browser
* RESTful HTTP API (see below)
* CLI (ssh into it)

### QUERY A SERVICE in a INSTANCE/VM via HTTP API

Once the agent is started and the service is synced,
we can query the service using either the DNS or HTTP API.

Using RESTful HTTP API to Query:

```bash
curl http://<ip>:8500/v1/catalog/services
```

or just for a particular service,

```bash
curl http://<ip>:8500/v1/catalog/service/<service names>
```

## 2. CONSUL AGENT

Sits in your image/instance/VM.

The Consul agent is the core process of Consul. 
The agent maintains membership information, registers services,
runs checks, responds to queries, and more.
The agent must run on every node that is part of a Consul cluster.

## 3. CONSUL TEMPLATE DAEMON

Consul template runs as a daemon and allows you to update
key/value pairs. Its useful if you want to interject
key:values into a running service or restart a service.

You update the key:value via consul.  This will 
take the `template.ctmpl` and write to the `config.yml` file.

```bash
$ consul-template \
    -template "/tmp/template.ctmpl:/tmp/config.yml"
```

Now lets create a config file for your service we want to
pass configuration info to via consul.

* Update key:value via consul (browser, http api or cli)
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

### TEMPLATE FILE

The template file key value may looks like:

my_test_number = {{keyOrDefault "staging/jeff-test-hello/number" "45"}}

### MATCH TEMPLATE FILE WITH CONSUL

You have the template file on your instance/VM.  But you now need to
provide those hooks to consul.

A. Via Browser

Just open consul browser, goto KEY/VALUE and match whats in your
template file.

B. Via RESTful HTTP API


curl http://10.240.1.114:8500/v1/catalog/service/jeff-test-hello


C. Via CLI

---TBD---

