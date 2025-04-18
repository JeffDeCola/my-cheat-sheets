# WIRELESS LAN CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Wireless lan, also called Wi-Fi, is technology for radio wireless
local area networking of devices based on the IEEE 802.11 standards._

Table of Contents

* [WHY IS IT CALLED WI-FI](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/networks/wan-lan/wireless-lan-cheat-sheet#why-is-it-called-wi-fi)
* [VERSIONS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/networks/wan-lan/wireless-lan-cheat-sheet#versions)
* [SECURITY LEVELS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/networks/wan-lan/wireless-lan-cheat-sheet#security-levels)

## WHY IS IT CALLED WI-FI

Who knows. I would guess because it sounds cool.

## VERSIONS

These are pretty big approximations,

|  Ver |    WIFI | DATE | ANTENNAS |  FREQ (GHz)|  DATA RATE (Mbps) | INDOOR (ft) | OUTDOOR (ft) |
|-----:|--------:|-----:|---------:|-----------:|------------------:|------------:|-------------:|
|    b |  Wi-Fi 1 | 1999 |        1 |        2.4 |                11 |         115 |          460 |
|    g |  Wi-Fi 3 | 2003 |        1 |        2.4 |                54 |         125 |          460 |
|    n |  Wi-Fi 4 | 2009 |        1 |      2.4/5 |               130 |         230 |          820 |
|      |          |      |        2 |            |               300 |             |              |
|      |          |      |        3 |            |               450 |             |              |
|   ac |  Wi-Fi 5 | 2013 |        1 |          5 |               450 |         150 |          300 |
|      |          |      |        2 |            |               900 |             |              |
|      |          |      |        4 |            |              3500 |             |              |
|   ax |  Wi-Fi 6 | 2019 |        1 |      2.4/5 |               540 |          98 |          390 |
|      |          |      |        4 |            |              4800 |             |              |
|      |          |      |        8 |            |              9600 |             |              |
|      | Wi-Fi 6E | 2020 |        1 |    2.4/5/6 |               540 |          98 |          390 |
|      |          |      |        4 |            |              4800 |             |              |
|      |          |      |        8 |            |              9600 |             |              |

## SECURITY LEVELS

In order of security (with encryption),

* GENERATION 1
  * WEP (RC4) - Never use, easily hacked
  * WPA (TKIP) - Don't use
  * WPA (AES) - Don't use
  * WPA-PSK (TKIP) - Don't use
  * WPA-PSK (AES) - Don't use
* GENERATION 2
  * WPA2-PSK (TKIP) - Don't use
  * **WPA2-PSK (AES)** - Use this, simple and easy, also called personal or home
  * WPA2 (TKIP) - Don't use
  * WPA2 ENT (AES) - Enterprise mode
* GENERATION 3
  * **WPA3-PSK (SAE)** - Personal - Simultaneous Authentication of Equals
  * WPA3 (Enhance Open Mode) - For open networks
  * WPA3 ENT (AES with GCMP) - Enterprise Mode

Do not use mixed mode, just select WPA2-PSK (AES).

PSK is Pre-Shared Key. Also called personal.  Provide router with Plain english key.
