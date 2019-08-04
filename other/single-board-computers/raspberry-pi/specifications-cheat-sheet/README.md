# SPECIFICATIONS CHEAT SHEET

`specifications` _of Raspberry Pi models SBCs._

* [Raspi 1 Model B+ / Raspi 2 Model B](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/raspberry-pi/specifications-cheat-sheet#raspi-1-model-b--raspi-2-model-b)
* [Raspi 3 Model B / Raspi 3 Model B+ / Raspi 4 Model B](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/raspberry-pi/specifications-cheat-sheet#raspi-3-model-b--raspi-3-model-b--raspi-4-model-b)

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## Raspi 1 Model B+ / Raspi 2 Model B

|                  | Rasp Pi 1 Model B+ |  Rasp Pi 2 Model B |
|:-----------------|-------------------:|-------------------:|
| Year             |               2014 |               2015 |
| SoC              |            BCM2835 |            BCM2836 |
| - CPU / CORE     |              ARM11 |      ARM Cortex A7 |
| - bus            |             32-bit |             32-bit |
| - Hz             |             700MHz |             900MHz |
| - Cores          |                  1 |                  4 |
| - Instruct Set   |     ARMv6 (32-bit) |   ARMv7-A (32-bit) |
| - GPU            | 250MHz VideoCoreIV | 250MHz VideoCoreIV |
| - - OpenGL ES    |  1.1/2.0 24 GFLOPS |  1.1/2.0 24 GFLOPS |
| - - MPEG-2       |       Need License |       Need License |  
| - - H.264 encode |            1080p30 |            1080p30 |
| - - H.264 decode |            1080p30 |            1080p30 |
| - - H.265 encode |                  - |                 -  |
| - - H.265 decode |                  - |                 -  |
| - - VRAM         |             Shared |             Shared |  
| - RAM            |        512MB SDRAM |          1GB SDRAM |
| OS               |   Rasbian (32-bit) |   Rasbian (32-bit) |
| Storage          |            microSD |            microSD |
| Ethernet         |             10/100 |             10/100 |
| Wireless         |                  - |                  - |
| Bluetooth        |                  - |                  - |
| USB              |        (4) USB 2.0 |        (4) USB 2.0 |
| Video Out        |  HDMI(rev1.3)/Comp |  HDMI(rev1.3)/Comp |
| - resolution     |            1080p30 |            1080p30 |  
| Audio Out        |    HDMI/Headphone  |     HDMI/Headphone |
| GPIO             |                 40 |                 40 |  
| POWER SOURCE     |           microUSB |           microUSB |
| - idle           |         200mA (1W) |       220mA (1.1W) |
| - max            |      350mA (1.75W) |       820mA (4.1W) |
| PoE              |                  - |                  - |

## Raspi 3 Model B / Raspi 3 Model B+ / Raspi 4 Model B

|                  |  Rasp Pi 3 Model B | Rasp Pi 3 Model B+ |  Rasp Pi 4 Model B |
|:-----------------|-------------------:|-------------------:|-------------------:|
| Year             |               2016 |               2018 |               2019 |
| SoC              |            BCM2837 |          BCM2837B0 |            BCM2711 |
| - CPU / CORE     |     ARM Cortex A53 |     ARM Cortex A53 |     ARM Cortex A72 |
| - bus            |          64/32-bit |          64/32-bit |          64/32-bit |
| - Hz             |             1.2GHz |             1.4GHz |             1.5GHz |
| - Cores          |                  4 |                  4 |                  4 |
| - Instruct Set   | ARMv8-A(64/32-bit) | ARMv8-A(64/32-bit) | ARMv8-A(64/32-bit) |
| - GPU            | 400MHz VideoCoreIV | 400MHz VideoCoreIV |    ??? VideoCoreVI |
| - - OpenGL ES    |1.1/2.0 28.8 GFLOPS |1.1/2.0 28.8 GFLOPS |  2.0/3.0 ?? GFLOPS |
| - - MPEG-2       |       Need License |       Need License |         (Disabled) |
| - - H.264 encode |            1080p30 |            1080p30 |            1080p30 |
| - - H.264 decode |            1080p30 |            1080p30 |            1080p60 |
| - - H.265 encode |                 -  |                  - |                  - |
| - - H.265 decode |                 -  |                  - |       (HEVC) 4kp60 |
| - - VRAM         |             Shared |             Shared |             Shared |
| - RAM            |          1GB SDRAM |          1GB SDRAM |  1GB/2GB/4GB SDRAM |
| OS               |   Rasbian (32-bit) |   Rasbian (32-bit) |   Rasbian (32-bit) |
| Storage          |            microSD |            microSD |            microSD |
| Ethernet         |             10/100 |        10/100/1000 |        10/100/1000 |
| Wireless         |      b/g/n(2.4GHz) | b/g/n/ac(2.4/5GHz) | b/g/n/ac(2.4/5GHz) |
| Bluetooth        |           Blue 4.1 |           Blue 4.2 |           Blue 5.0 |
| USB              |        (4) USB 2.0 |        (4) USB 2.0 |    (2) 2.0 (2) 3.0 |
| Video Out        |  HDMI(rev1.3)/Comp |  HDMI(rev1.3)/Comp |   Dual u-HDMI(2.0) |
| - resolution     |            1080p60 |            1080p60 |  (1)4kp60 (2)4kp30 |
| Audio Out        |     HDMI/Headphone |     HDMI/Headphone |     HDMI/Headphone |
| GPIO             |                 40 |                 40 |                 40 |
| POWER SOURCE     |           microUSB |           microUSB |              USB-C | 
| - idle           |       300mA (1.5W) |       459mA (2.3W) |           ?mA (?W) | 
| - max            |       1.34A (6.7W) |      1.13A (5.66W) |            ?A (?W) |
| PoE              |                  - |  need add-on board |  need add-on board |

BCM chips are from Broadcom. Refer to my cheat sheet on
[ARM architectures](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/hardware-architectures/arm-cheat-sheet).

Note, you can use software decoders, rather than the hw ones.
