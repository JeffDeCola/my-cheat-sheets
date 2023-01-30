# WILLIAMS 1990 DINER PINBALL MACHINE CHEAT SHEET

_Parts, part numbers and settings._

Table on Contents,

* [DINER OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/williams-diner-pinball-machine-cheat-sheet#diner-overview)
* [FUSES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/williams-diner-pinball-machine-cheat-sheet#fuses)
* [LAMPS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/williams-diner-pinball-machine-cheat-sheet#lamps)
* [LEFT FLIPPER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/williams-diner-pinball-machine-cheat-sheet#left-flipper)
* [MAIN BOARDS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/williams-diner-pinball-machine-cheat-sheet#main-boards)
  * [SYSTEM 11C CPU BOARD (D-11883-571)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/williams-diner-pinball-machine-cheat-sheet#system-11c-cpu-board-d-11883-571)
  * [AUDIO BOARD (D-11581-571)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/williams-diner-pinball-machine-cheat-sheet#audio-board-d-11581-571)
  * [POWER SUPPLY BOARD (D-12246)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/williams-diner-pinball-machine-cheat-sheet#power-supply-board-d-12246)
  * [AUX POWER DRIVER BOARD (D-12247-566)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/williams-diner-pinball-machine-cheat-sheet#aux-power-driver-board-d-12247-566)
  * [BACKBOARD INTERCONNECT BOARD (D-12313-571)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/williams-diner-pinball-machine-cheat-sheet#backboard-interconnect-board-d-12313-571)
  * [MASTER DISPLAY BOARD (D-12232-1)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/williams-diner-pinball-machine-cheat-sheet#master-display-board-d-12232-1)
* [GAME ADJUSTMENTS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/williams-diner-pinball-machine-cheat-sheet#game-adjustments)
  * [GENERAL SETTINGS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/williams-diner-pinball-machine-cheat-sheet#general-settings)
  * [GAME DIFFICULTY SETTINGS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/williams-diner-pinball-machine-cheat-sheet#game-difficulty-settings)
  * [OTHER SETTINGS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/williams-diner-pinball-machine-cheat-sheet#other-settings)

Documentation and reference,

* [The Internet Pinball Machine Database](https://www.ipdb.org/machine.cgi?id=681)
* A great resource for system 11C systems from
  [pinwiki](https://www.pinwiki.com/wiki/index.php/Williams_System_9_-_11)
* A resource from troubleshooting
  [here](https://techniek.flipperwinkel.nl/wms11/index1.htm)
* [GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## DINER OVERVIEW

* Model: 571
* Serial Number: 571 420760
* System: 11C (Forth gen of system 11)
* Manufactured: Williams Electronic Games, Inc.
* Year: September, 1990
* Produced: 3,552
* Dimensions (WxDxH): 29" x 56" x 75"
* Power: 115V 60Hz, 8 AMPS
* Weight: 250 lbs
* CPU: 1 Mhz 8 bit
* RAM: 2Kx8 Memory with (3) 1.5V Alkaline AA batteries for power
* ROM (2) 32Kx8 (Game Program) Version L-3 (latest version is L-4)
* Players: 4
* Flippers: 2
* Ramps: 2
* Bumpers: 3
* Multi-ball: 2
* 3-Bank Drop Targets: 2
* Coins: 25 cents

![IMAGE - diner-pinball.jpg - IMAGE](../../../docs/pics/diner-pinball.jpg)

## FUSES

There are a lot of fuses.

![IMAGE - williams-diner-pinball-machine-fuses.jpg - IMAGE](../../../docs/pics/williams-diner-pinball-machine-fuses.jpg)

## LAMPS

* WEDGE BASE
  * Part# 24-8768 - Bulb #555 (6.3V 0.25A) - A lot of these
  * Part# 24-8802 - Bulb #906 (13V, 0.69A) - I think 4 of these
* BAYONET BASE
  * Part# xx - Bulb #44 (6.3V 0.25A) - Less than 10 of these
    * Can use Bulb #47 (6.3V .15A) but they are dimmer
  * Part# 24-8704 - Bulb #89 (13V, 0.58A) - Maybe 10

## LEFT FLIPPER

Since I needed to fix the left flipper, here are my notes.

* LEFT FLIPPER ASSEMBLY (C-13174-L)
  * Flipper Coil "Solenoid" (FL11630/50VDC)
  * Coil Stop (A-12111)
  * Lane change Switch Assembly: Not Used

## MAIN BOARDS

There are 6 main boards.

### SYSTEM 11C CPU BOARD (D-11883-571)

* System 11C CPU Board Part #D-11883-571 (571 is the game number)
  * 6802 Microprocessor (U15) Part #5400-09150-00
    * STMicroelectronics Part#EF6802P Date Code 8936 (I think 1989 Week 36)
    * 40 pin MPU
    * 8-bit
    * Address bus: 16 bits (64K memory RAM)
    * Data Bus: 8 bits
    * 72 instructions (172 opcodes)
    * 1 MHz
    * Transistors: 7200
    * Assembly language
  * RAM (U25) Part #5340-10139-00
    * Fujitsu Part #MB8416A-15L Date Code 8413 R45 (I think 1984 Week 13)
    * SRAM 6116
    * 2Kx8
    * 24 pin
    * (3) 1.5V Alkaline AA batteries for power
  * ROM (U26 & U27) Part #A-5343-571-2
    * EPROM 27256
    * 32Kx8
    * Version L-3 (Latest Version is L-4)
    * U26 : Checksum: C706 (for L-4)
    * U27 : Checksum: 10C3 (for L-4)
  * Peripheral Interface Adapter (U10, U38, U41, U42, U51, U54)
    * 6821
    * 40-pin

### AUDIO BOARD (D-11581-571)

For this system 11C, it contains all sounds on 3 ROM chips.

### POWER SUPPLY BOARD (D-12246)

Supplies power.

### AUX POWER DRIVER BOARD (D-12247-566)

Supplies Aux power.

### BACKBOARD INTERCONNECT BOARD (D-12313-571)

The 571 is for the game number.

### MASTER DISPLAY BOARD (D-12232-1)

The LED display.

## GAME ADJUSTMENTS

There are 70 adjustments. 31-48 are game specific (e.g. diner).
I will list the notable ones.

There are three buttons on the door,

* ADVANCE
* AUTO (UP)/MANUAL (DOWN)
* HIGH SCORE RESET

To move up number, have AUTO (UP) and press ADVANCE.
To move down number, have MANUAL (DOWN) and press ADVANCE.
Press start button to select.

### GENERAL SETTINGS

| #  | ADJUSTMENT                         | DEFAULT                 | MY SETTING              |
|:---|:-----------------------------------|:------------------------|:------------------------|
| 01 | AUTO OR FIXED REPLY                | 10% (AUTO REPLY)        | SCORES (FIXED REPLAY)   |
| 02 | REPLY LEVEL 1                      | 6,000,000               | 3,000,000               |
| 03 | REPLY LEVEL 2                      | 01                      | 5,000,000               |
| 04 | REPLY LEVEL 3                      | -                       | 7,000,000               |
| 05 | REPLY LEVEL 4                      | -                       | 9,000,000               |
| 06 | REPLY REWARD (SEE ABOVE)           | CREDIT (FREE GAME)      | EXTRA BALL              |
| 07 | SPECIAL REWARD                     | CREDIT (FREE GAME)      |                         |
| 08 | MATCH FEATURE (CHANCE FREE GAME)   | 8%                      | 20%                     |
| 09 | BALLS PER GAME                     | 3                       |                         |
| 10 | TILT WARNING PER GAME              | 3                       |                         |
| 11 | EXTRA BALLS PER GAME               | 4                       |                         |

### GAME DIFFICULTY SETTINGS

I just set #64 to medium which set all these automatically.

| #  | ADJUSTMENT                         | DEFAULT                 | MY MEDIUM SETTING       |
|:---|:-----------------------------------|:------------------------|:------------------------|
| 32 | % EXTRA BALL PER GAME              | 20%                     | 25% (MEDIUM)            |
| 33 | EXTRA BALL LIT MEMORY              | YES                     | YES (MEDIUM)            |
| 34 | SPOT MULTIBALL                     | YES                     | YES (MEDIUM)            |
| 35 | RUSH TIMER                         | 12 SECONDS              | 20 SECONDS (MEDIUM)     |
| 37 | CUSTOMER MEMORY                    | YES                     | YES (MEDIUM)            |
| 38 | FOOD ITEM MEMORY                   | YES                     | YES (MEDIUM)            |
| 39 | TODAY'S SPECIAL                    | EXTRA HARD              | MEDIUM                  |
| 40 | E-A-T LANE MEMORY                  | YES                     | YES (MEDIUM)            |
| 42 | CASH REGISTER RAMP                 | 6 SECONDS               | 12 SECONDS (MEDIUM)     |
| 43 | DINER RAMP (COFFEE CUP)            | EXTRA HARD              | MEDIUM                  |
| 46 | TOP EJECT AWARD                    | NO                      | YES (MEDIUM)            |
| 47 | CONSOLATION BALL TIME              | 40 SECONDS              | 40 SECONDS (MEDIUM)     |
| 48 | ATTRACT SOUNDS                     | ON                      | ON (MEDIUM)             |
| 62 | EXTRA EASY (SEE ABOVE)             | -                       | -                       |
| 63 | EASY (SEE ABOVE)                   | -                       | -                       |
| 64 | MEDIUM (SEE ABOVE)                 | -                       | YES                     |
| 65 | HARD (SEE ABOVE)                   | -                       | -                       |
| 66 | EXTRA HARD (SEE ABOVE)             | YES                     | -                       |

### OTHER SETTINGS

| #  | ADJUSTMENT                         | DEFAULT                 | MY SETTING              |
|:---|:-----------------------------------|:------------------------|:------------------------|
| 49 | CUSTOMER MESSAGE                   | YOU CAN EAT..           | DECOLA'S BAR...         |
| 62 | FACTORY SETTINGS                   |                         |                         |
| 69 | CLEAR AUDITS                       |                         |                         |
