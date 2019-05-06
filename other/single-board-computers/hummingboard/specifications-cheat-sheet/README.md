# SPECIFICATIONS CHEAT SHEET

`specifications` _of SolidRun's Hummingboard SBC._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## SOLIDRUN

A company that makes the hummingboard.

## HUMMINGBOARD PRODUCT LIST

The hummingboard is composed of two parts,

* Carrier Board
* SoM (System on a module).

SolidRun’s Hummingboard-i4 Single Board Computer (SBC) uses Freescale's SoC
Quad Core ARMv7 CPU (based on
[ARM](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/hardware-architectures/arm-cheat-sheet)
cortex A9).
The Hummingboard is a development board composed of two smaller boards:
a Carrier Board and a Micro-SOM (System on Modules) Board.
It requires 5V, 2A microUSB power.

## CARRIER BOARD OPTIONS & SPECS

|            |   mSATA II |        RTC | ANALOG AUD |
|:-----------|-----------:|-----------:|-----------:|
| BASE       |         No |         No |         No |
| PRO        |        yes |        Yes |        Yes |

### Carrier Board (Pro)

* Model          Hummingboard Pro
* OS             Arch Linux ARM or Debian (linux distro for embedded system)
* Linux Kernel   3.14.x
* HDMI           1.4
* USB            2 External ports. 2 Internal
* Boot Up        Boots from microSD (4GB)
* Power          microUSB
* Ethernet       10/100/1000
* mSATA II       512GB mSATA SSD
* RTC Chip       NXO PCF8523 (external battery (CR1620) using 2 pin header)
* GPIO           26 pin - Same configuration as
                 [Raspberry Pi](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/raspberry-pi/specifications-cheat-sheet)
* Audio          Analog Stereo and MIC in

## MICRO-SOM (SYSTEM ON MODULE) OPTIONS

|      | SoC i.MX6 |                     RAM |         3D GPU |  WiFi |
|:-----|----------:|------------------------:|---------------:|------:|
|   i1 |      Solo |  32-bit ,512MB, 800Mbps |  Vivante GC880 |    No |
|   i2 | Dual-Lite |    64-bit ,1GB, 800Mbps |  Vivante GC880 |    No |
| i2eX |      Dual |    64-bit ,1GB, 800Mbps | Vivante GC2000 |   Yes |
|   i4 |      Quad |   64-bit ,2GB, 1066Mbps | Vivante GC2000 |   Yes |
| i4x4 |      Quad |   64-bit ,4GB, 1066Mbps | Vivante GC2000 |   Yes |

### i4

* Model                i4
* SoC                  Freescale’s i.MX6 Quad
  * CORE               Quad Core (4) – ARM Cortex A9 (ARMv7 Instruction Set)
  * CPU                1 GHz ARMv7
  * 3D GPU             Vivante GC2000 (OpenGL ES profile v3.0)
  * 2D GPU             Vivante GC355 (OpenVG 1.1) and GC320 (BLIT)
  * Video En/Decode    1080p (1920x1080) 60 h.264, VP8, etc… (ON VPU)
* WiFi 802.11 a/b/g/n  Broadcom BCM 4330 Chip
* Bluetooth 4.0+HS     Broadcom BCM 4330 Chip (Not sure if supported)
* RAM                  2GB DDR3 (1066MHz)

### i2eX

* Model                i2eX
* SoC                  Freescale’s i.MX6 Dual
  * CORE               Dual Core (2) – ARM Cortex A9 (ARMv7 Instruction Set)
  * CPU                1 GHz ARMv7
  * 3D GPU             Vivante GC2000 (OpenGL ES profile v3.0)
  * 2D GPU             Vivante GC355 (OpenVG 1.1) and GC320 (BLIT)
  * Video En/Decode    1080p (1920x1080) 60 h.264, VP8, etc… (ON VPU)
* WiFi 802.11 a/b/g/n  Broadcom BCM 4330 Chip
* Bluetooth 4.0+HS     Broadcom BCM 4330 Chip (Not sure if supported)
* RAM                  1GB (800MHz)
