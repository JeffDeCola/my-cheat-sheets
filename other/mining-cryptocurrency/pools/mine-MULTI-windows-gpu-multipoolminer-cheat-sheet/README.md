# MINE MULTIPLE COINS WINDOWS GPU CHEAT SHEET

`mine-MULTI-windows-gpu-multipoolminer` _will show you
how to mine for multiple coins
on your Windows 10 GPU
using `multipoolminer` mining software
connected to the `miningpoolhub.com` pool._

My other mining cheatsheets,

* **FULL NODE**
  * [mine-ZEC-macOS-cpu-zcashd](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/full-node/mine-ZEC-macOS-cpu-zcashd-cheat-sheet)
* **POOLS**
  * [mine-BEAM-windows-gpu-lolMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-BEAM-windows-gpu-lolMiner-cheat-sheet)
  * [mine-MULTI-windows-gpu-multipoolminer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-MULTI-windows-gpu-multipoolminer-cheat-sheet)
  **YOU ARE HERE**
  * [mine-RVN-windows-gpu-t-rex](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-RVN-windows-gpu-t-rex-cheat-sheet)
  * [mine-ZCL-windows-gpu-lolMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-ZCL-windows-gpu-lolMiner-cheat-sheet)
  * [mine-ZEC-windows-gpu-funakoshiMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-ZEC-windows-gpu-funakoshiMiner-cheat-sheet)
  * [mine-ZEL-windows-gpu-gminer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-ZEL-windows-gpu-gminer-cheat-sheet)

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-MULTI-windows-gpu-multipoolminer-cheat-sheet#overview)
* [POOL (miningpoolhub.com)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-MULTI-windows-gpu-multipoolminer-cheat-sheet#pool-miningpoolhubcom)
  * [CONFIGURE AUTOEXCHANGE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-MULTI-windows-gpu-multipoolminer-cheat-sheet#configure-autoexchange)
  * [ADD A HUB WORKER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-MULTI-windows-gpu-multipoolminer-cheat-sheet#add-a-hub-worker)
* [MINER (multipoolminer)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-MULTI-windows-gpu-multipoolminer-cheat-sheet#miner-multipoolminer)
  * [INSTALL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-MULTI-windows-gpu-multipoolminer-cheat-sheet#install)
* [RUN](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-MULTI-windows-gpu-multipoolminer-cheat-sheet#run)

Check out my cheat sheet on all the popular
[cryptocurrency](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/cryptocurrency/cryptocurrency-cheat-sheet).

## OVERVIEW

Here is an illustration of what we're going to do,

![IMAGE - mine-MULTI-windows-gpu-multipoolminer - IMAGE](../../../../docs/pics/mine-MULTI-windows-gpu-multipoolminer.jpg)

## POOL (miningpoolhub.com)

First, lets pick you pool. I registered at
[miningpoolhub.com](https://miningpoolhub.com).

There are other pools, but I choose this one.

### CONFIGURE AUTOEXCHANGE

This is a fancy way to say your base wallet.
I choose want LTC to be my base currency.

Add your LTC Address and a threshold of when to send.

### ADD A HUB WORKER

Add a Worker name and a password.
I used `WORKER1` as my worker.

## MINER (multipoolminer)

This is just compute software that does all the hash computations
it is told to do from the various multi pools.  It will determine which is
best for your machine.

This software Supports multiple platforms,

* AMD
* NVIDIA
* CPU

And Supports the following pools,

* MiningPoolHub (I'm using this one)
* Zpool
* HashRefinery
* Nicehash
* Ahashpool
* BlockMunch
* ItalYiiMP
* YiiMP.eu

It makes it money by taking out a little bit of coin per day.

There are other programs like `awesome miner` and `minerstat`,
but I choose this one.

### INSTALL

Now lets get multipoolminer.exe for your rig.
Grab the latest Windows binary from
[https://github.com/MultiPoolMiner/MultiPoolMiner/releases](https://github.com/MultiPoolMiner/MultiPoolMiner/releases)

Place the folder anywhere on your rig.

## RUN

Its actually very simple. Open `Start-MiningPoolHub.bat` and simply edit
the following line to something like,

```bash
set "command=& .\multipoolminer.ps1
-API_Key <YOUR_API_KEY>
-APIPort 3999
-UserName <YOUR_USERNAME>
-WorkerName <YOUR_WORKER_NAME>
-Region US
-Currency usd
-DeviceName nvidia
-PoolName miningpoolhub-Algo
-Donate 10
-Watchdog
-MinerStatusURL https://multipoolminer.io/monitor/miner.php
-MinerStatusKey <YOUR KEY>
-SwitchingPrevention 2"
```

If you want to only use a particular miner (algorithm) use,

```txt
-Algorithm Ethash
```

If you want to mine for a particular coin like `Feathercoin` use,

```txt
-CoinName Feathercoin
```

I wanted to get Ethereum coin to work but can't seem to figure it out.

The API key allowed you to use this,

 [http://localhost:3999]( http://localhost:3999)

Thats it, you're mining and wasting electricity.
