# ARM CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_ARM architecture is is a family of reduced instruction
set computing (RISC) architectures for computer processors._

Table of Contents

* [ARM HOLDINGS LICENSE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/hardware-architectures/arm-cheat-sheet#arm-holdings-license)
* [WHY USE SAME ARCHITECTURE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/hardware-architectures/arm-cheat-sheet#why-use-same-architecture)
* [ARM CORES USE ARM ARCHITECTURE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/hardware-architectures/arm-cheat-sheet#arm-cores-use-arm-architecture)

Documentation and Reference

* [Hummingboard Specs](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/single-board-computers/hummingboard/specifications-cheat-sheet)
* [Raspberry Pi Specs](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/single-board-computers/raspberry-pi/specifications-cheat-sheet)

## ARM HOLDINGS LICENSE

Arm Holdings develops CPU IP cores based on architecture
and licenses it to other companies, who design their own products.

## WHY USE SAME ARCHITECTURE

The benefit of using the same HW architecture is
everyone can use the same software library for an OS distribution.
For example, ubuntu on a Hummingboard can get software from
the ARM7 distro.

## ARM CORES USE ARM ARCHITECTURE

This highlights the main cores and the architecture used.

| ARM ARCHITECTURE |      ARM HOLDINGS IP CORE |       THIRD PARTY CHIP |  Single Board Computers |
|:-----------------|--------------------------:|-----------------------:|------------------------:|
| ARMv1            |                      ARM1 |                        |                         |
| ARMv2            |                      ARM2 |                        |                         |
| ARMv6            |                     ARM11 |       Broadcom BCM2835 |      Rasp Pi 1 Model B+ |
| ARMv7-A          |             ARM Cortex-A7 |       Broadcom BCM2836 |       Rasp Pi 2 Model B |
|                  |             ARM Cortex-A9 |        Freescale i.MX6 |            Hummingboard |
| ARMv8-A          |            ARM Cortex-A53 |       Broadcom BCM2837 |  Rasp Pi 3 Model B & B+ |
| ARMv8-A          |            ARM Cortex-A72 |       Broadcom BCM2711 |       Rasp Pi 4 Model B |
| ARMv8.2-A        |            ARM Cortex-A55 |                        |                         |
