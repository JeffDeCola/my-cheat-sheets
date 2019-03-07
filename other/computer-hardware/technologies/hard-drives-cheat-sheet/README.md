# HARD DRIVES CHEAT SHEET

`hard drives` _are hard drives._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## TYPES

* DESKTOP - GOOD
* NAS (Network-Attached Storage) - BETTER
* ENTERPRISE NAS - BEST

## RAID

Redundant array of Inexpensive Disks.

As an example, assume each disk is a 1TB drive.

| # | D1 | D2 | D3 | D4 | D5 | D6 | D7 | D8 | TOTAL |   % |          COMMENT |             
|--:|---:|---:|---:|---:|---:|---:|---:|---:|------:|----:|-----------------:|
|   |    |    |    |    |    |    |    |    |       |     |                  |
| 0 | A1 | A2 |    |    |    |    |    |    |   2TB | 100 |      Performance |
|   | A3 | A4 |    |    |    |    |    |    |       |     |         Striping |
|   |    |    |    |    |    |    |    |    |       |     |                  |
|   |    |    |    |    |    |    |    |    |       |     |                  |
| 1 | A1 | A1 |    |    |    |    |    |    |   1TB |  50 |      Reliability |
|   | A2 | A2 |    |    |    |    |    |    |       |     |           Mirror |
|   |    |    |    |    |    |    |    |    |       |     |                  |