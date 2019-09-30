# MINE ZCASH MACOS CPU USING ZCASHD CHEAT SHEET

`mine-ZEC-macOS-cpu-zcashd` _will show you
how to mine for Zcash (ZEC)
on your macOS CPU
using `zcashd` node, wallet and mining software._

NOTE: cpu mining really not worth it even if you join a pool.
Also, you don't need a full node to mine solo anymore,
just use a solo pool.

My other mining cheatsheets,

* **FULL NODE**
  * [mine-ZEC-macOS-cpu-zcashd](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/full-node/mine-ZEC-macOS-cpu-zcashd-cheat-sheet)
    **YOU ARE HERE**
* **MULTIPLE POOLS**
  * [mine-MULTICOINS-windows-gpu-multipoolminer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/multiple-pools/mine-MULTICOINS-windows-gpu-multipoolminer-cheat-sheet)
  * [mine-MULTICOINS-windows-gpu-awesome-miner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/multiple-pools/mine-MULTICOINS-windows-gpu-awesome-miner-cheat-sheet)
* **POOLS**
  * [mine-BEAM-windows-gpu-lolMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-BEAM-windows-gpu-lolMiner-cheat-sheet)
  * [mine-BTG-windows-gpu-gminer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-BTG-windows-gpu-gminer-cheat-sheet)
  * [mine-MULTI-windows-gpu-multipoolminer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-MULTI-windows-gpu-multipoolminer-cheat-sheet)
  * [mine-RVN-windows-gpu-t-rex](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-RVN-windows-gpu-t-rex-cheat-sheet)
  * [mine-ZCL-windows-gpu-lolMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-ZCL-windows-gpu-lolMiner-cheat-sheet)
  * [mine-ZEC-windows-gpu-funakoshiMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-ZEC-windows-gpu-funakoshiMiner-cheat-sheet)
  * [mine-ZEL-windows-gpu-gminer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-ZEL-windows-gpu-gminer-cheat-sheet)

tl;dr,

```bash
# ADD PATH TO .bashrc
export PATH=/Users/jeffdecola/src/github.com/zcash/zcash/src:$PATH
# START
zcashd
zcashd -daemon
# STOP
zcash-cli stop
# ZCASH-CLI
    zcash-cli help
    # GET STATUS
    zcash-cli getinfo
    zcash-cli getmininginfo
    # LIST WALLETS
    zcash-cli getaddressesbyaccount ""
    zcash-cli z_listaddresses
    # WALLET BALANCE
    zcash-cli z_gettotalbalance
    # GET PRIVATE WALLET KEY
    zcash-cli dumpprivkey "t-addr"
# DATA
cd ~/Library/Application\ Support/Zcash/
# EDIT CONFIG
nano ~/Library/Application\ Support/Zcash/zcash.conf
```

Table of Contents,

* **FULL NODE**
  * [mine-ZEC-macOS-cpu-zcashd](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/full-node/mine-ZEC-macOS-cpu-zcashd-cheat-sheet)
    **YOU ARE HERE**
* **MULTIPLE POOLS**
  * [mine-MULTICOINS-windows-gpu-awesome-miner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/multiple-pools/mine-MULTICOINS-windows-gpu-awesome-miner-cheat-sheet)
  * [mine-MULTICOINS-windows-gpu-multipoolminer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/multiple-pools/mine-MULTICOINS-windows-gpu-multipoolminer-cheat-sheet)
* **POOLS**
  * [mine-BEAM-windows-gpu-lolMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-BEAM-windows-gpu-lolMiner-cheat-sheet)
  * [mine-BTG-windows-gpu-gminer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-BTG-windows-gpu-gminer-cheat-sheet)
  * [mine-RVN-windows-gpu-t-rex](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-RVN-windows-gpu-t-rex-cheat-sheet)
  * [mine-ZCL-windows-gpu-lolMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-ZCL-windows-gpu-lolMiner-cheat-sheet)
  * [mine-ZEC-windows-gpu-funakoshiMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-ZEC-windows-gpu-funakoshiMiner-cheat-sheet)
  * [mine-ZEL-windows-gpu-gminer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-ZEL-windows-gpu-gminer-cheat-sheet)

Check out my cheat sheet on all the popular
[cryptocurrency](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/cryptocurrency/cryptocurrency-cheat-sheet).

## OVERVIEW

Here is an illustration of what we're going to do,

![IMAGE - mine-ZEC-macOS-cpu-zcashd - IMAGE](../../../../docs/pics/mine-ZEC-macos-cpu-zcashd.jpg)

## MINER (zcashd)

There are many different hardware options but I am going to use
my macOS CPU. The easiest is the Zcash software provided by
Zcash company which is a complete package that will let you,

* Run a full Node
* Mine with your CPU
* Built in Wallet for sending and receiving Zcash
* Ability to hook up to pools (Optional)

The benefit of a full node is that you can mine solo,
rather than in a pool, but they have solo pools, so
you really don't need to run a full node anymore.

### INSTALL

Lets start the installation,

I got a lot of information from the official zcash mining guide
[here](https://zcash.readthedocs.io/en/latest/rtd_pages/zcash_mining_guide.html).

First make sure you have everything for the build,

```bash
xcode-select --install
brew install git pkgconfig automake autoconf wget libtool coreutils
sudo easy_install pip
sudo pip install pyblake2 pyzmq
open /Library/Developer/CommandLineTools/Packages/macOS_SDK_headers_for_macOS_10.14.pkg
```

Downloading Zcash source,

```bash
git clone https://github.com/zcash/zcash.git
```

Now place keys in `~/Library/Application Support/ZcashParams`
These keys are just under 1.7GB in size, so it
may take some time to download them.

```bash
cd zcash
git checkout v2.0.7-2
./zcutil/fetch-params.sh
```

### BUILD FOR NUMBER OF CORES

In the zcash directory,

```bash
./zcutil/build.sh -j$(nproc)
./zcutil/build.sh -j$(4)
```

`nproc` is number of cores.

Create a data directory,

```bash
cd ~/Library/Application\ Support
mkdir -p Zcash
```

### CONFIGURE FOR MINING

Create and edit your config file,

```txt
cd ~/Library/Application\ Support/Zcash
touch zcash.conf
nano ~/Library/Application\ Support/Zcash/zcash.conf
```

And add the following to it,

```txt
addnode=mainnet.z.cash
gen=1
genproclimit=-1
```

## RUN

Add path in `.bashrc`,

```txt
export PATH=/Users/jeffdecola/src/github.com/zcash/zcash/src:$PATH
```

Start, run in background or stop,

```bash
zcashd
zcashd -daemon
zcashd -damemon=0 -showmetrics=1
zcash-cli stop
```

get status,

```bash
zcash-cli getinfo
zcash-cli getmininginfo
```

## ZCASH-CLI COMMANDS

```bash
zcash-cli help
```

### WALLETS

Create, List and get private key for your t-addr wallet,

```bash
zcash-cli getnewaddress
zcash-cli getaddressesbyaccount ""
zcash-cli dumpprivkey "t-addr"
```

Create and List z-addr wallet,

```bash
zcash-cli z_getnewaddress
zcash-cli z_listaddresses
```

Wallet balance,

```bash
zcash-cli z_gettotalbalance
```

## POOLS (OPTIONAL)

At some point this may be a good idea to join a pool,

[List of ming pools](https://www.zcashcommunity.com/mining/mining-pools/).
