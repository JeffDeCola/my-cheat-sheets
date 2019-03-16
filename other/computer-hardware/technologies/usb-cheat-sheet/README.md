# USB CHEAT SHEET

`usb` _universal serial bus specifications._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## OVERVIEW

Released in 1996, the USB standard is currently maintained by the USB
Implementers Forum (USB IF). There have been three generations of USB
specifications: USB 1.x, USB 2.0, USB 3.x.

## USB

|        VER| RELEASE | MAX DATA | STANDARD |   MINI |    MICRO | DUPLEX |
|----------:|--------:|---------:|---------:|-------:|---------:|-------:|
|       1.0 |    1996 |  12 Mb/s |      A,B |      - |        - |      - |
|       2.0 |    2001 | 480 Mb/s |      A,B | A,B,AB |   A,B,AB |      - |
|       3.0 |    2011 |   5 Gb/s |  A,B(ss) |      - | B,AB(ss) |      - |
| 3.1 Gen.1 |    2011 |   5 Gb/s |  A,B(ss) |      - | B,AB(ss) |      C |
| 3.1 Gen.2 |    2014 |  10 Gb/s |  A,B(ss) |      - | B,AB(ss) |      C |
|       3.2 |    2017 |  20 Gb/s |  A,B(ss) |      - | B,AB(ss) |      C |

ss is superspeed.

## APPLE'S THUNDERBOLT

|  VER      | RELEASE | MAX DATA |        PORT TYPE |
|----------:|--------:|---------:|-----------------:|
|         1 |    2011 |  10 Mb/s | Mini DisplayPort |
|         3 |    2013 |  20 Mb/s | Mini DisplayPort |
|         2 |    2015 |  40 Mb/s |                C |

USB 3.x is different then Thunderbolt 3 but can use the
same Type-C connector.