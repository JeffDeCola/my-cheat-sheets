# MINE MULTIPLE COINS WINDOWS GPU USING AWESOME MINER CHEAT SHEET

`mine-MULTI-windows-gpu-awesome-miner` _will show you
how to mine for multiple coins
on your Windows 10 GPU
using `awesome-miner` mining software
connected to the `zergpool.com` pool._

My other mining cheatsheets,

* **FULL NODE**
  * [mine-ZEC-macOS-cpu-zcashd](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/full-node/mine-ZEC-macOS-cpu-zcashd-cheat-sheet)
* **MULTIPLE POOLS**
  * [mine-MULTICOINS-windows-gpu-multipoolminer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/multiple-pools/mine-MULTICOINS-windows-gpu-multipoolminer-cheat-sheet)
  * [mine-MULTICOINS-windows-gpu-awesome-miner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/multiple-pools/mine-MULTICOINS-windows-gpu-awesome-miner-cheat-sheet)
    **YOU ARE HERE**
* **POOLS**
  * [mine-BEAM-windows-gpu-lolMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-BEAM-windows-gpu-lolMiner-cheat-sheet)
  * [mine-BTG-windows-gpu-gminer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-BTG-windows-gpu-gminer-cheat-sheet)
  * [mine-RVN-windows-gpu-t-rex](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-RVN-windows-gpu-t-rex-cheat-sheet)
  * [mine-ZCL-windows-gpu-lolMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-ZCL-windows-gpu-lolMiner-cheat-sheet)
  * [mine-ZEC-windows-gpu-funakoshiMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-ZEC-windows-gpu-funakoshiMiner-cheat-sheet)
  * [mine-ZEL-windows-gpu-gminer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-ZEL-windows-gpu-gminer-cheat-sheet)

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/multiple-pools/mine-MULTI-windows-gpu-awesome-miner-cheat-sheet#overview)
* [MINER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/multiple-pools/mine-MULTI-windows-gpu-awesome-miner-cheat-sheet#miner)
* [POOL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/multiple-pools/mine-MULTI-windows-gpu-awesome-miner-cheat-sheet#pool)
* [RUN](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/multiple-pools/mine-MULTI-windows-gpu-awesome-miner-cheat-sheet#run)
* [MONITOR](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/multiple-pools/mine-MULTI-windows-gpu-awesome-miner-cheat-sheet#monitor)

Check out my cheat sheet on all the popular
[cryptocurrency](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/cryptocurrency/cryptocurrency-cheat-sheet).

## OVERVIEW

Instead of using a one miner and one pool, you will be
using multiple miners and multiple pools.

A "mining coordinator" will be used to determine what pool to mine to
based on your rig in order to maximize profits.
It will pick the best miner to use at the time and switch
between miners and pools as coin prices change, etc...

Since you are mining to multiple pools and will have
little bits of coins everywhere, the pool will automatically
exchange all those coins into one wallet.
I like to put everything into litecoin since its fast.

They like to call these `Multialgo (Multi algorithm), Multicoin, Auto-Exchange Pools`.

Here is an illustration of what we're going to do,

![IMAGE - mine-MULTICOINS-windows-gpu-XYZmulti-miner - IMAGE](../../../../docs/pics/mine-MULTICOINS-windows-gpu-XYZmulti-miner.jpg)

## MINER

Grab the latest Windows binary from
[awesomeminer.com](https://www.awesomeminer.com/download)

Place the folder anywhere on your rig.

## POOL

First, lets pick your pool. I picked
[zergpool.com](https://zergpool.com).

No registration is required, they do payouts in the currency
of you wallet address.  Simple.

## RUN

Its installed in your system so just kick it off.

I suggest doing a benchmark first.

Normally, you can just run because the software is all setup for
your pool under `options -> profit switching`.  You just add your wallet.
The default is BTC.

If you want to payout to a particular wallet other than BTC like litecoin,
add `c=LTC` in each of the pools under `options -> online services`.

You can even get more selective if you want to create your own pools and pool groups.
For example, I would like to mine FOLM (FLM) coin and put my coins into my FLM wallet
using a SOLO pool.

My switches for zergpool would be placed in `options -> pools`,

```txt
Pool URL: stratum+tcp://phi.mine.zergpool.com:8333
Worker Name: Fpz2px36Ptaa7PKBK4HLA7DprDrqGNRsXN
Worker password: c=FLM,mc=FLM,m=solo
```

## MONITOR

You can play around with the interface.
[Awesomeminer.com](https://www.awesomeminer.com/download)
is the best source of information.

Thats it, you're mining and wasting electricity.
