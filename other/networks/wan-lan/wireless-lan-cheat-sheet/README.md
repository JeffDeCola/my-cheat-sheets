# WIRELESS LAN CHEAT SHEET

`wireless lan` _also called Wi-Fi, is technology for radio wireless
local area networking of devices based on the IEEE 802.11 standards._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## WHY IS IT CALLED WI-FI

Who knows.  I would guess because it sounds cool.

## VERSIONS

These are pretty big approximations,

|  Ver |  ANTENNAS |  FREQ (GHz)|  DATA RATE (Mbps) |
|-----:|----------:|-----------:|------------------:|
|    b |         1 |        2.4 |                11 |
|    g |         1 |        2.4 |                54 |
|    n |         1 |      2.4/5 |               130 |
|      |         2 |      2.4/5 |               300 |
|      |         3 |      2.4/5 |               450 |
|   ac |         1 |          5 |               450 |
|      |         2 |          5 |               900 |
|      |         3 |          5 |              1300 |

## SECURITY LEVELS

In order of security (with encryption),

* WEP (RC4) - Never use, easily hacked
* WPA (TKIP) - Don't use
* WPA (AES) - Don't use
* WPA-PSK (TKIP) - Don't use
* WPA-PSK (AES) - Don't use
* WPA2-PSK (TKIP) - Don't use
* **WPA2-PSK (AES)** - Use this, simple and easy, also called personal or home
* WPA2 (TKIP) - Don't use
* **WPA2 (AES)** - This is the highest level and also called enterprise

Do not use mixed mode on your routers, just select WPA2-PSK (AES).

PSK is Pre-Shared Key. Also called personal.  Provide router with Plain english key.
