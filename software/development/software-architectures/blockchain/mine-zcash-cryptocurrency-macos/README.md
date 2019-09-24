# MINE ZCASH CRYPTOCURRENCY CHEAT SHEET

`mine-zcash-cryptocurrency-macOS` _will show you how to mine for zcash(ZEC) on the
Zcash mainnet on your macOS._

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
    # GET PRIVATE WALLET KEY
    zcash-cli dumpprivkey "t-addr"
# DATA
cd ~/Library/Application\ Support/Zcash/
# EDIT CONFIG
nano ~/Library/Application\ Support/Zcash/zcash.conf
```

Table of Contents,

* [OVERVIEW OF ZCASH](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/blockchain/mine-zcash-cryptocurrency-macos#overview-of-zcash)
* [macOS CPU](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/blockchain/mine-zcash-cryptocurrency-macos#macos-cpu)
* [INSTALL FROM SOURCE ON YOUR MACOS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/blockchain/mine-zcash-cryptocurrency-macos#install-from-source-on-your-macos)
* [RUN](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/blockchain/mine-zcash-cryptocurrency-macos#run)
* [ZCASH-CLI COMMANDS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/blockchain/mine-zcash-cryptocurrency-macos#zcash-cli-commands)

## OVERVIEW OF ZCASH

* Launch: 2016
* Official website:
  [z.cash](https://z.cash/)
* Can you mine: YES
* Hash: Equihash based
* Max Supply: 21,000,000
* Proof Type: Proof of Work
* Example of a Wallet:
  [explorer.zcha.in](https://explorer.zcha.in/accounts/t1h1xStMimJTxAo9DvLY7koDj9UkKDACtxb)

The Block Rewards will issue a total of 50 Zcash (ZEC) every 10 mins.
The block spacing is 2.5 mins, so each block produces 12.5 ZEC.
10% of all coins mined will go to the Founders Reward.

## macOS CPU

There are many different hardware options but I am going to use
my macOS CPU. The easiest is the Zcash software provided by
Zcash company which is a complete package that will let you,

* Run a full Node
* Mine with your CPU
* Built in Wallet for sending and receiving Zcash

Here is an illustration of what we are going to do,

![IMAGE - mine-zcash-cryptocurrency - IMAGE](../../../../../docs/pics/mine-zcash-cryptocurrency.jpg)

## INSTALL FROM SOURCE ON YOUR MACOS

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
zcash-cli stop
```

get status,

```bash
zcash-cli getinfo
zcash-cli getmininginfo
```

## ZCASH-CLI COMMANDS

````bash
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
