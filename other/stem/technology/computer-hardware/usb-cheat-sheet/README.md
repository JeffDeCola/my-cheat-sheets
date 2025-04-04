# USB CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Universal serial bus specifications._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-hardware/usb-cheat-sheet#overview)
* [USB](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-hardware/usb-cheat-sheet#usb)
* [APPLE THUNDERBOLT](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-hardware/usb-cheat-sheet#apple-thunderbolt)

## OVERVIEW

Released in 1996, the USB standard is currently maintained by the USB
Implementers Forum (USB IF). There have been three generations of USB
specifications: USB 1.x, USB 2.0, USB 3.x.

PORTS AND CONNECTOR TYPES

* STANDARD
  * USB-A (Most common one)
  * USB-B (Like printers use)
* MINI
  * USB MINI A (Tapered)
  * USB MINI B (Rectangle)
* MICRO
  * USB MICRO A (Rectangle)
  * USB MICRO B (Tapered)
* TYPE C
  * USB-C (Reversible)

## USB

|                  VER| RELEASE |           AKA | LANES | MAX DATA | STANDARD |   MINI |    MICRO | TYPE C |
|--------------------:|--------:|--------------:|------:|---------:|---------:|-------:|---------:|-------:|
|             USB 1.0 |    1996 |               |     - |  12 Mbps |      A,B |      - |        - |      - |
|             USB 1.1 |    1998 |               |     - |  12 Mbps |      A,B |      - |        - |      - |
|             USB 2.0 |    2000 |               |     - |  480 Mbs |      A,B |    A,B |      A,B |      C |
|             USB 3.0 |    2008 |               |     - |   5 Gbps |      A,B |      - |        B |      C |
|             USB 3.1 |    2014 |               |     - |  10 Gbps |          |      - |          |      C |
|       USB 3.1 Gen.1 |    2011 |       USB 3.0 |     - |   5 Gbps |        A |      - |          |      C |
|       USB 3.1 Gen.2 |    2014 |       USB 3.1 |     - |  10 Gbps |          |      - |          |      C |
|             USB 3.2 |    2017 |               |     2 |  20 Gbps |          |      - |          |      C |

The Big Reset,

|                  VER| RELEASE |           AKA | LANES | MAX DATA | STANDARD |   MINI |    MICRO | TYPE C |
|--------------------:|--------:|--------------:|------:|---------:|---------:|-------:|---------:|-------:|
|   **USB 3.2 Gen 1** |    2011 | USB 3.1 Gen.1 |     1 |   5 Gbps |        A |      - |        B |      C |
|                     |         |       USB 3.0 |       |          |          |        |          |        |
|   **USB 3.2 Gen 2** |    2014 | USB 3.1 Gen.2 |     1 |  10 Gbps |        A |      - |        B |      C |
|                     |         |       USB 3.1 |       |          |          |        |          |        |
| **USB 3.2 Gen 1x2** |    2017 |               |     2 |  10 Gbps |          |      - |          |      C |
| **USB 3.2 Gen 2x2** |    2017 |       USB 3.2 |     2 |  20 Gbps |          |      - |          |      C |
|        USB4 Gen 2x2 |    2019 |               |     2 |  20 Gbps |          |      - |          |      C |
|        USB4 Gen 3x2 |    2019 |               |     2 |  40 Gbps |          |      - |          |      C |

NOTE: USB 3.2 Gen 1x2 means 5 Gbps x 2 lanes, for a connection speed of 10 Gbps.

NOTE: USB4 Gen 3x2 means 20 Gbps x 2 lanes, for a connection speed of 40 Gbps.

## APPLE THUNDERBOLT

|  VER      | RELEASE | MAX DATA |        PORT TYPE |
|----------:|--------:|---------:|-----------------:|
|         1 |    2011 |  10 Mb/s | Mini DisplayPort |
|         2 |    2013 |  20 Mb/s | Mini DisplayPort |
|         3 |    2015 |  40 Mb/s |                C |
|         4 |    2020 |  40 Mb/s |                C |

USB 3.x is different then Thunderbolt 3 but can use the
same Type-C connector.
