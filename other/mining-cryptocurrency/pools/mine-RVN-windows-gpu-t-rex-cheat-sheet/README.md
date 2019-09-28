# MINE RVN (Ravencoin) WINDOWS GPU USING T-REX CHEAT SHEET

`mine-RVN-windows-gpu-t-rex-` _will show you
how to mine for RVN (Ravencoin)
on your Windows 10 GPU
using `t-rex` mining software
connected to a pool._

My other mining cheat sheets,

* **FULL NODE**
  * [mine-ZEC-macOS-cpu-zcashd](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/full-node/mine-ZEC-macOS-cpu-zcashd-cheat-sheet)
* **POOLS**
  * [mine-BEAM-windows-gpu-lolMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-BEAM-windows-gpu-lolMiner-cheat-sheet)
  * [mine-MULTI-windows-gpu-multipoolminer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-MULTI-windows-gpu-multipoolminer-cheat-sheet)
  * [mine-RVN-windows-gpu-t-rex](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-RVN-windows-gpu-t-rex-cheat-sheet)
  **YOU ARE HERE**
  * [mine-ZCL-windows-gpu-lolMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-ZCL-windows-gpu-lolMiner-cheat-sheet)
  * [mine-ZEC-windows-gpu-funakoshiMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-ZEC-windows-gpu-funakoshiMiner-cheat-sheet)

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-RVN-windows-gpu-t-rex-cheat-sheet#overview)
* [MINER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-RVN-windows-gpu-t-rex-cheat-sheet#miner)
* [RUN](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-RVN-windows-gpu-t-rex-cheat-sheet#run)
* [MONITOR](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-RVN-windows-gpu-t-rex-cheat-sheet#monitor)

Check out my cheat sheet on all the popular
[cryptocurrency](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/cryptocurrency/cryptocurrency-cheat-sheet).

## OVERVIEW

Here is an illustration of what we're going to do,

![IMAGE - mine-XYZ-windows-gpu-XYZminer - IMAGE](../../../../docs/pics/mine-XYZ-windows-gpu-XYZminer.jpg)

## MINER

`t-rex` can mine quit a bit of different coins.

Grab the latest Windows binary from
[github.com/trexminer/T-Rex/releases](https://github.com/trexminer/T-Rex/releases/)
and place the folder anywhere on your rig.

## RUN

### SOLO POOL

To run on `2miners.com` SOLO pool create a `start.bat` file,

```bash
t-rex.exe
-a x16r
--fork-at x16rv2=2019-10-01T16:00:00
-o stratum+tcp://solo-rvn.2miners.com:7070
-u <WALLET>.<WORKER>
-p x
```

### NORMAL POOL

To run on `2miners.com` normal pool just change the pool and port,

```bat
-o stratum+tcp://rvn.2miners.com:6060
```

## MONITOR

You can also monitor `t-rex` on your rig,

[http://10.141.141.1:4067/trex](http://10.141.141.1:4067/trex)

Thats it, you're mining and wasting electricity.
