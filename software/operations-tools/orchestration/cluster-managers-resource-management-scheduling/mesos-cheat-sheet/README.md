# MESOS CHEAT SHEET

`mesos` _gives you the ability to run both containerized,
and non-containerized workloads in a distributed manner._

Documentation and reference,

* [Mesos DC/OS Documentation](https://dcos.io/)
* [Mesosphere Documentation](https://mesosphere.com/)
* [Mesos Documentation](http://mesos.apache.org/)
* My 
  [marathon cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/orchestration/cluster-managers-resource-management-scheduling/marathon-cheat-sheet)

My Repo example is
[hello-go-deploy-marathon](https://github.com/JeffDeCola/hello-go-deploy-marathon).

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

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
