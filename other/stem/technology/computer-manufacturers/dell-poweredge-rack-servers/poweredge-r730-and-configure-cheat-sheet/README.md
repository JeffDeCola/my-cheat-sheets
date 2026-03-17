# POWEREDGE R730 AND CONFIGURE CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_A dell rack server._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/poweredge-r730-cheat-sheet#overview)
* [CONFIGURE RAID DRIVES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/poweredge-r730-cheat-sheet#configure-raid-drives)

## OVERVIEW

![IMAGE - dell-poweredge-r730-2016 - IMAGE](../../../../../../docs/pics/other/dell-poweredge-r730-2016.svg)

## CONFIGURE RAID DRIVES

To create virtual disk raid drives.

In iDRAC

```text
Storage → Controllers
* Click on your H730 Mini and you should see
  * Battery Backup Unit (BBU) — shows Present / Ready / Charging
  * Cache Memory — shows the amount (usually 1GB on H730 Mini)
```

If BBU field is ready we used Write Back to create virtual disk

```text
Storage → Virtual Disk -> Click Create Virtual Disk
* Virtual Disk Name: "SSD-Fast-1.6TB"
* Controller: PERC H730 Mini (embedded)
* Layout: RAID 10
* Media Type: SSD
* Stripe Element Size: 64KB (default is fine)
* Capacity: N/A
* Read Policy: Adaptive Read Ahead
* Write Policy: Write Back (if battery/cache present)
* Disk Cache Policy: Enabled
* T10 Protection Information Capacity: Disabled
* Span Count: N/A
```

Assign the Hot Spare Disk

```text
Storage → Virtual Disks -> Manage
  * Right-click the 5th WGP72 (your spare)
  * Select Assign Dedicated Hot Spare
  * Assign it to the SSD-Fast-1.6TB virtual disk
```

Check that it's a hot spare in physical disks.

If you want to do another one, do the same process as above but
may need to remove and old RAID configuration in
Storage -> Controllers.
