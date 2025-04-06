# MESOS CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Mesos gives you the ability to run both containerized,
and non-containerized workloads in a distributed manner._

Table of Contents

* [OVERVIEW (MESOS, MESOSPHERE & MESOS DC/OS)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/orchestration/cluster-managers-resource-management-scheduling/mesos-cheat-sheet#overview-mesos-mesosphere--mesos-dc-os)
* [INSTALL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/orchestration/cluster-managers-resource-management-scheduling/mesos-cheat-sheet#install)
* [FRAMEWORK - MARATHON](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/orchestration/cluster-managers-resource-management-scheduling/mesos-cheat-sheet#framework---marathon)

Documentation and Reference

* [mesos DC/OS documentation](https://dcos.io/)
* [mesosphere documentation](https://mesosphere.com/)
* [mesos documentation](http://mesos.apache.org/)
* [marathon cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/orchestration/cluster-managers-resource-management-scheduling/marathon-cheat-sheet)
* [hello-go-deploy-marathon](https://github.com/JeffDeCola/hello-go-deploy-marathon)

## OVERVIEW (MESOS, MESOSPHERE & MESOS DC/OS)

`Mesos` is a project by Apache that gives you the
ability to run both containerized, and non-containerized
workloads in a distributed manner.
Mesos is complicated and hard to manage, so mesosphere
tries to make mesos into something understandable.

`Mesosphere` has a marathon “plugin” to mesos, which is an
easy way to manage container orchestration over mesos.

`Mesos DC/OS` (Data Center Operating System) simplifies mesos
even further and allows you to deploy your own mesos cluster,
with marathon, in a matter of minutes.

## INSTALL

This is out of the scope of this cheat sheet.
I used a vagrant file.

## FRAMEWORK - MARATHON

Mesos provides resource management and scheduling
while a framework such as marathon provides
scheduling and execution.

There are many types of frameworks that a designed to do different
things such as data storage, or processing.
Marathon is design to keep an app running forever.
