# MINE BEAM (BEAMassic) WINDOWS GPU LOLMINER CHEAT SHEET

`mine-BEAM-windows-gpu-lolMiner` _will show you
how to mine for BEAM (BEAM)
on your Windows 10 GPU
using `lolMiner` mining software
connected to a pool._

My other mining cheat sheets,

* **FULL NODE**
  * [mine-ZEC-macOS-cpu-zcashd](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/full-node/mine-ZEC-macOS-cpu-zcashd-cheat-sheet)
* **POOLS**
  * [mine-BEAM-windows-gpu-lolMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-BEAM-windows-gpu-lolMiner-cheat-sheet)
  **YOU ARE HERE**
  * [mine-MULTI-windows-gpu-multipoolminer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-MULTI-windows-gpu-multipoolminer-cheat-sheet)
  * [mine-ZCL-windows-gpu-lolMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-ZCL-windows-gpu-lolMiner-cheat-sheet)
  * [mine-ZEC-windows-gpu-funakoshiMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-ZEC-windows-gpu-funakoshiMiner-cheat-sheet)

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-BEAM-windows-gpu-lolMiner-cheat-sheet#overview)
* [MINER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-BEAM-windows-gpu-lolMiner-cheat-sheet#miner)
* [RUN](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-BEAM-windows-gpu-lolMiner-cheat-sheet#run)

Check out my cheat sheet on all the popular
[cryptocurrency](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/cryptocurrency/cryptocurrency-cheat-sheet).

## OVERVIEW

Here is an illustration of what we're going to do,

![IMAGE - mine-XYZ-windows-gpu-XYZminer - IMAGE](../../../../docs/pics/mine-XYZ-windows-gpu-XYZminer.jpg)

## MINER

`lolMiner` is also a BEAM CUDA Miner.

Grab the latest Windows binary from
[github.com/Lolliedieb/lolMiner-releases/releases/](https://github.com/Lolliedieb/lolMiner-releases/releases/)
and place the folder anywhere on your rig.

## RUN

To run on `2miners.com` SOLO pool create a `start.bat` file,

```bash
setx GPU_FORCE_64BIT_PTR 1
setx GPU_MAX_HEAP_SIZE 100
setx GPU_USE_SYNC_OBJECTS 1
setx GPU_MAX_ALLOC_PERCENT 100
setx GPU_SINGLE_ALLOC_PERCENT 100

lolMiner.exe
--coin BEAM
--pool us-solo-beam.2miners.com
--port 5454
--user <WALLET>.<WORKER>
```

To run on `2miners.com` normal pool create a `start.bat` file,

```bat
lolMiner.exe
--coin BEAM
--pool us-beam.2miners.com
--port 5252
--user <WALLET>.<WORKER>
```

Thats it, you're mining and wasting electricity.
