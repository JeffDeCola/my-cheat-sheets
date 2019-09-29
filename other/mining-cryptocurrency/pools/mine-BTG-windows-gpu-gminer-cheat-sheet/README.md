# MINE BTG (Bitcoin Gold) WINDOWS GPU USING GMINER CHEAT SHEET

`mine-BTG-windows-gpu-gminer` _will show you
how to mine for BTG (Bitcoin Gold)
on your Windows 10 GPU
using `gminer` mining software
connected to a pool._

My other mining cheat sheets,

* **FULL NODE**
  * [mine-ZEC-macOS-cpu-zcashd](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/full-node/mine-ZEC-macOS-cpu-zcashd-cheat-sheet)
* **MULTIPLE POOLS**
  * [mine-MULTICOINS-windows-gpu-multipoolminer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/multiple-pools/mine-MULTICOINS-windows-gpu-multipoolminer-cheat-sheet)
  * [mine-MULTICOINS-windows-gpu-sniffdogminer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/multiple-pools/mine-MULTICOINS-windows-gpu-sniffdogminer-cheat-sheet)
* **POOLS**
  * [mine-BEAM-windows-gpu-lolMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-BEAM-windows-gpu-lolMiner-cheat-sheet)
  * [mine-BTG-windows-gpu-gminer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-BTG-windows-gpu-gminer-cheat-sheet)
    **YOU ARE HERE**
  * [mine-RVN-windows-gpu-t-rex](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-RVN-windows-gpu-t-rex-cheat-sheet)
  * [mine-ZCL-windows-gpu-lolMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-ZCL-windows-gpu-lolMiner-cheat-sheet)
  * [mine-ZEC-windows-gpu-funakoshiMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-ZEC-windows-gpu-funakoshiMiner-cheat-sheet)
  * [mine-ZEL-windows-gpu-gminer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-ZEL-windows-gpu-gminer-cheat-sheet)

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-BTG-windows-gpu-gminer-cheat-sheet#overview)
* [MINER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-BTG-windows-gpu-gminer-cheat-sheet#miner)
* [POOL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-BTG-windows-gpu-gminer-cheat-sheet#pool)
* [RUN](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-BTG-windows-gpu-gminer-cheat-sheet#run)
* [MONITOR](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-BTG-windows-gpu-gminer-cheat-sheet#monitor)

Check out my cheat sheet on all the popular
[cryptocurrency](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/cryptocurrency/cryptocurrency-cheat-sheet).

## OVERVIEW

Here is an illustration of what we're going to do,

![IMAGE - mine-XYZ-windows-gpu-XYZminer - IMAGE](../../../../docs/pics/mine-XYZ-windows-gpu-XYZminer.jpg)

## MINER

`gminer` can mine quit a bit of different coins.

Grab the latest Windows binary from
[github.com/develsoftware/GMinerRelease/releases](https://github.com/develsoftware/GMinerRelease/releases/)
and place the folder anywhere on your rig.

## POOL

You can use any pool. For this example we're going to use
[2miners.com](https://2miners.com/)
which doesn't require an account.

The only thing you need is a personal wallet address
that you will place in the .bat file below.

## RUN

### SOLO POOL

To run on `2miners.com` SOLO pool create a `start.bat` file,

```bash
miner.exe
--algo 144_5
--pers BgoldPoW
--server us-solo-btg.2miners.com
--port 4040
--user <WALLET>.<WORKER>
--pass x
--api 10050
```

### NORMAL POOL

To run on `2miners.com` normal pool just change the pool and port,

```bat
--server us-btg.2miners.com
--port 4040
```

## MONITOR

You can locally monitor your mining software,

[localhost:10050](http://localhost:10050/)

Thats it, you're mining and wasting electricity.
