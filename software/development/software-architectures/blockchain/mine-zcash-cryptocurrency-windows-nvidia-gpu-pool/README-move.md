# MINE ??? USING MULTIPOOLMINER CHEAT SHEET

`mine-zcash-cryptocurrency-windows-evga-gtx1080` _will show you how to mine for
zcash(ZEC) on the Zcash mainnet on your Windows 10 using your GPU (EVGA GTX1080).
I will be using open source sw `funakoshiMiner` connected to `SlushPool`._

This is **pool based**, meaning the software will not work independently
since you would need a full node.  **The pool keeps the node**.

Table of Contents,

* tbd

## WINDOWS 10 NVIDIA GPU

There are many different hardware and software options but
I am going to use my gtx1080 on my windows rig.
I will be using the open source funakoshiMiner with slush pool.

`Funakoshi` is a Equihash CUDA Miner.  So it works with ZEC and ETH
but not bitcoin.

Here is an illustration of what we're going to do,

![IMAGE - mine-zcash-cryptocurrency-windows-evga-gtx1080 - IMAGE](../../../../../docs/pics/mine-zcash-cryptocurrency-windows-evga-gtx1080.jpg)

## INSTALL funakoshiMiner

Grab the latest Windows binary from
[github.com/funakoshi2718/funakoshi-miner](https://github.com/funakoshi2718/funakoshi-miner)

Place the folder anywhere on your rig.

## GET A SLUSH POOL ACCOUNT

You really can use any pool, but I choose
[slushpool.com](https://slushpool.com).

Provide them a wallet to put your mined coins.
And then created a worker for ZEC coins in your dashboard.

## CONFIGURE

Its actually very simple. Open Start.bat and simply edit,

```bash
funakoshiMiner.exe --server us-east.zec.slushpool.com --port 4444 --user <USER>.<WORKER>
```

Thats it, you are mining and now wasting electricity.
