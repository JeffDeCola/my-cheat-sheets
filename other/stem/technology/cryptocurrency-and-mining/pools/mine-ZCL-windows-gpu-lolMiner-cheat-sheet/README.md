# MINE ZCL (Zclassic) WINDOWS GPU USING LOLMINER CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_How to mine for Zclassic (ZCL)
on your Windows 10 GPU
using `lolMiner` mining software
connected to a pool._

My other mining cheat sheets

* **FULL NODE**
  * [mine-ZEC-macOS-cpu-zcashd](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/cryptocurrency-and-mining/full-node/mine-ZEC-macOS-cpu-zcashd-cheat-sheet)
* **MULTIPLE POOLS**
  * [mine-MULTICOINS-windows-gpu-awesome-miner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/cryptocurrency-and-mining/multiple-pools/mine-MULTICOINS-windows-gpu-awesome-miner-cheat-sheet)
  * [mine-MULTICOINS-windows-gpu-multipoolminer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/cryptocurrency-and-mining/multiple-pools/mine-MULTICOINS-windows-gpu-multipoolminer-cheat-sheet)
* **POOLS**
  * [mine-BEAM-windows-gpu-lolMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/cryptocurrency-and-mining/pools/mine-BEAM-windows-gpu-lolMiner-cheat-sheet)
  * [mine-BTG-windows-gpu-gminer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/cryptocurrency-and-mining/pools/mine-BTG-windows-gpu-gminer-cheat-sheet)
  * [mine-RVN-windows-gpu-t-rex](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/cryptocurrency-and-mining/pools/mine-RVN-windows-gpu-t-rex-cheat-sheet)
  * [mine-ZCL-windows-gpu-lolMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/cryptocurrency-and-mining/pools/mine-ZCL-windows-gpu-lolMiner-cheat-sheet)
    **<-YOU ARE HERE**
  * [mine-ZEC-windows-gpu-funakoshiMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/cryptocurrency-and-mining/pools/mine-ZEC-windows-gpu-funakoshiMiner-cheat-sheet)
  * [mine-ZEL-windows-gpu-gminer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/cryptocurrency-and-mining/pools/mine-ZEL-windows-gpu-gminer-cheat-sheet)

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/cryptocurrency-and-mining/pools/mine-ZCL-windows-gpu-lolMiner-cheat-sheet#overview)
* [MINER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/cryptocurrency-and-mining/pools/mine-ZCL-windows-gpu-lolMiner-cheat-sheet#miner)
* [POOL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/cryptocurrency-and-mining/pools/mine-ZCL-windows-gpu-lolMiner-cheat-sheet#pool)
* [RUN](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/cryptocurrency-and-mining/pools/mine-ZCL-windows-gpu-lolMiner-cheat-sheet#run)
* [MONITOR](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/cryptocurrency-and-mining/pools/mine-ZCL-windows-gpu-lolMiner-cheat-sheet#monitor)

Documentation and Reference

* [cryptocurrency](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/cryptocurrency-and-mining/overview/cryptocurrency/cryptocurrency-cheat-sheet)

## OVERVIEW

Here is an illustration of what we're going to do,

![IMAGE - mine-XYZ-windows-gpu-XYZminer - IMAGE](../../../../../../docs/pics/other/mine-XYZ-windows-gpu-XYZminer.svg)

## MINER

`lolMiner` is a AUTO192_7 CUDA Miner.

Grab the latest Windows binary from
[github.com/Lolliedieb/lolMiner-releases/releases/](https://github.com/Lolliedieb/lolMiner-releases/releases/)
and place the folder anywhere on your rig.

## POOL

You can use any pool. For this example we're going to use
[miningpoolhub.com](https://miningpoolhub.com).

## RUN

To run on `miningpoolhub.com` create a `start.bat` file,

```bash
setx GPU_FORCE_64BIT_PTR 1
setx GPU_MAX_HEAP_SIZE 100
setx GPU_USE_SYNC_OBJECTS 1
setx GPU_MAX_ALLOC_PERCENT 100
setx GPU_SINGLE_ALLOC_PERCENT 100

lolMiner.exe
--coin AUTO192_7
--pool us-east.equihash-hub.miningpoolhub.com
--port 20574
--user <USER>.<WORKER>
--apiport 10051
```

## MONITOR

You can locally monitor your mining software,

[localhost:10051/summary](http://localhost:10051/summary)

Thats it, you're mining and wasting electricity.
