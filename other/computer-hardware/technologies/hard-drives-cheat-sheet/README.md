# HARD DRIVES CHEAT SHEET

`hard drives` _are hard drives._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## TYPES

* DESKTOP - GOOD
* NAS (Network-Attached Storage) - BETTER
* ENTERPRISE NAS - BEST

## RAID 0, 1, 5, 6

Redundant array of Inexpensive Disks.

As an example, assume each disk is a 1 TB drive.

| # | D1 | D2 | D3 | D4 | D5 | D6 | D7 | D8 |  TOT |   % |            COMMENT |
|--:|---:|---:|---:|---:|---:|---:|---:|---:|-----:|----:|-------------------:|
|   |    |    |    |    |    |    |    |    |      |     |                    |
| 0 | A1 | A2 |    |    |    |    |    |    | 2 TB | 100 |        Performance |
|   | A3 | A4 |    |    |    |    |    |    |      |     |           Striping |
|   |    |    |    |    |    |    |    |    |      |     |                    |
|   |    |    |    |    |    |    |    |    |      |     |                    |
| 1 | A1 | A1 |    |    |    |    |    |    | 1 TB |  50 |           Reliable |
|   | A2 | A2 |    |    |    |    |    |    |      |     |             Mirror |
|   |    |    |    |    |    |    |    |    |      |     |                    |
|   |    |    |    |    |    |    |    |    |      |     |                    |
| 5 | A1 | A2 |  P |    |    |    |    |    | 2 TB |  66 | Reliable (loose 1) |
|   | A3 | P  | A4 |    |    |    |    |    |      |     |  Striping 1 Parity |
|   |    |    |    |    |    |    |    |    |      |     |                    |
|   |    |    |    |    |    |    |    |    |      |     |                    |
|   | A1 | A2 | A3 |  P |    |    |    |    | 3 TB |  75 | Reliable (loose 1) |
|   | A4 | A5 | P  | A6 |    |    |    |    |      |     |  Striping 1 Parity |
|   |    |    |    |    |    |    |    |    |      |     |                    |
|   |    |    |    |    |    |    |    |    |      |     |                    |
|   | A1 | A2 | A3 | A4 |  P |    |    |    | 4 TB |  80 | Reliable (loose 1) |
|   | A5 | A6 | A7 |  P | A8 |    |    |    |      |     |  Striping 1 Parity |
|   |    |    |    |    |    |    |    |    |      |     |                    |
|   |    |    |    |    |    |    |    |    |      |     |                    |
| 6 | A1 | A2 |  P |  P |    |    |    |    | 2 TB |  50 | Reliable (loose 2) |
|   | A3 |  P |  P | A4 |    |    |    |    |      |     |  Striping 2 Parity |
|   |    |    |    |    |    |    |    |    |      |     |                    |
|   |    |    |    |    |    |    |    |    |      |     |                    |
|   | A1 | A2 | A3 |  P | P  |    |    |    | 3 TB |  60 | Reliable (loose 2) |
|   | A4 | A5 |  P |  P | A6 |    |    |    |      |     |  Striping 2 Parity |
|   |    |    |    |    |    |    |    |    |      |     |                    |
|   |    |    |    |    |    |    |    |    |      |     |                    |
|   | A1 | A2 | A3 | A4 |  P |  P |    |    | 4 TB |  66 | Reliable (loose 2) |
|   | A5 | A6 | A7 |  P |  P | A8 |    |    |      |     |  Striping 2 Parity |
|   |    |    |    |    |    |    |    |    |      |     |                    |

## RAID 10, 50

You can combine the above raid combinations.

RAID 10 is Really RAID 1 and RAID 0 combined.
Raid 50 is Raid 5 and RAID 0 combined.

| # | D1 | D2 | D3 | D4 | D5 | D6 | D7 | D8 |  TOT |   % |            COMMENT |
|--:|---:|---:|---:|---:|---:|---:|---:|---:|-----:|----:|-------------------:|
|   |    |    |    |    |    |    |    |    |      |     |                    |
|10 | A1 | A1 | A2 | A2 |    |    |    |    | 2 TB |  50 | Perform / Reliable |
|   | A3 | A3 | A4 | A4 |    |    |    |    |      |     |  Striping / Mirror |
|   |    |    |    |    |    |    |    |    |      |     |                    |
|   |    |    |    |    |    |    |    |    |      |     |                    |
|   | A1 | A1 | A2 | A2 | A3 | A3 |    |    | 3 TB |  50 | Perform / Reliable |
|   | A4 | A4 | A5 | A5 | A6 | A6 |    |    |      |     |  Striping / Mirror |
|   |    |    |    |    |    |    |    |    |      |     |                    |
|50 | A1 | A2 |  P | A3 | A4 |  P |    |    | 4 TB |  66 | Perform / Reliable |
|   | A5 |  P | A6 | A7 |  P | A8 |    |    |      |     |  Strip 2 Par / Mir |
|   |    |    |    |    |    |    |    |    |      |     |                    |
|   | A1 | A2 | A3 |  P | A4 | A5 | A6 | P  | 6 TB |  75 | Perform / Reliable |
|   | A7 | A8 |  P | A9 | A10| A11|  P | A12|      |     |  Strip 2 Par / Mir |
|   |    |    |    |    |    |    |    |    |      |     |                    |

