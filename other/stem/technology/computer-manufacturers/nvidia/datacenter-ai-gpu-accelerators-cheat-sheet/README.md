# DATACENTER AI/GPU ACCELERATORS

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Nvidia datacenter AI/GPU accelerators._

Table of Contents

* [ARCHITECTURE TIMELINE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/nvidia/datacenter-ai-gpu-accelerators-cheat-sheet#architecture-timeline)
* [HOW WORKSTATION DIFFERS FROM DATACENTER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/nvidia/datacenter-ai-gpu-accelerators-cheat-sheet#how-workstation-differs-from-datacenter)
* [ACCELERATORS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/nvidia/datacenter-ai-gpu-accelerators-cheat-sheet#accelerators)
* [PCIe vs SXM](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/nvidia/datacenter-ai-gpu-accelerators-cheat-sheet#pcie-vs-sxm)
* [DATACENTER vs WORKSTATION vs CONSUMER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/nvidia/datacenter-ai-gpu-accelerators-cheat-sheet#datacenter-vs-workstation-vs-consumer)
* [WHAT FITS IN MY DELL R730 RACK SERVER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/nvidia/datacenter-ai-gpu-accelerators-cheat-sheet#what-fits-in-my-dell-r730-rack-server)

## ARCHITECTURE TIMELINE

```text
2016  Pascal      P40, P100
2017  Volta       V100 (first tensor cores)
2018  Turing      T4 (INT8/INT4 inference)
2020  Ampere      A30, A40, A100 (3rd gen tensor, MIG)
2022  Ada         L4, L40S (FP8 support)
2023  Hopper      H100, H200 (transformer engine)
2025  Blackwell   B200, B300 (next-gen everything)
2026  Vera Rubin  announced, HBM4, 288 GB
```

## HOW WORKSTATION DIFFERS FROM DATACENTER

```text
                    Datacenter              Workstation             Consumer/Gaming
Cooling             Passive (chassis)       Active (fans on card)   Active (fans on card)
Display outputs     None                    Yes (DP, HDMI)          Yes (DP, HDMI)
Drivers             Compute-only            ISV certified           Game optimized
Fits in             Rack servers            Towers / desktops       Desktops
EULA                Datacenter OK           Commercial OK           No datacenter
ECC memory          Yes (hardware)          Yes (soft or hardware)  No
Example             A40, A100, L40S         RTX Pro 4000, Arc B70   RTX 3090, RTX 4090
```

## ACCELERATORS

| GPU | Year / Arch | VRAM / Type | Bandwidth | Tensor Cores / FP16 | TDP / PCIe | Used Price | Best For |
|-----|-------------|-------------|-----------|---------------------|------------|------------|----------|
| **Tesla P40** _(mine)_ | 2016 / Pascal | 24 GB GDDR5X | 346 GB/s | None / Slow (1/64) | 250W / 3.0 | ~$150 | Budget 8–30B inference |
| **Tesla T4** | 2018 / Turing | 16 GB GDDR6 | 300 GB/s | 320 (2nd gen) / 65 TFLOPS | 70W / 3.0 | ~$200–300 | Low-power inference |
| **Tesla V100 PCIe** | 2017 / Volta | 16 or 32 GB HBM2 | 900 GB/s | 640 (1st gen) / 28 TFLOPS | 250W / 3.0 | ~$300–600 | Tensor cores + HBM bandwidth |
| **A30** | 2020 / Ampere | 24 GB HBM2 | 933 GB/s | 224 (3rd gen) / 165 TFLOPS | 165W / 4.0 | ~$400–600 | MIG, low power, great bandwidth |
| **A40** | 2020 / Ampere | 48 GB GDDR6 | 696 GB/s | 336 (3rd gen) / 150 TFLOPS | 300W / 4.0 | ~$900 | Big VRAM on a budget |
| **A100 PCIe/SXM** | 2020 / Ampere | 40 or 80 GB HBM2e | 2,039 GB/s | 432 (3rd gen) / 312 TFLOPS | 250–400W / 4.0 | ~$5,000+ | Training + large inference |
| **L4** | 2023 / Ada | 24 GB GDDR6 | 300 GB/s | 240 (4th gen) / 121 TFLOPS | 72W / 4.0 | ~$2,000–3,000 | Modern T4 replacement |
| **L40S** | 2023 / Ada | 48 GB GDDR6 | 864 GB/s | 568 (4th gen) / 733 TFLOPS | 300W / 4.0 | ~$7,000+ | Big models + rendering |
| **H100 PCIe/SXM** | 2023 / Hopper | 80 GB HBM3 | 3,350 GB/s | 528 (4th gen) / 1,979 TFLOPS | 350–700W / 5.0 | ~$25,000+ | Frontier training + inference |
| **H200 SXM** | 2024 / Hopper | 141 GB HBM3e | 4,800 GB/s | 528 (4th gen) / 1,979 TFLOPS | 700W / 5.0 | ~$30,000+ | Massive models + KV cache |
| **B200** | 2025 / Blackwell | 192 GB HBM3e | 8,000 GB/s | Next gen / ~4,500 TFLOPS | 1,000W / 5.0 | N/A | Next-gen frontier |

## PCIe vs SXM

* **PCIe** — Standard slot. Fits normal servers like the R730.
GPU-to-GPU communication goes through PCIe bus.
* **SXM** — Custom socket on an HGX baseboard.
Requires purpose-built chassis (DGX, Dell XE9680, etc.).
NVLink for direct GPU-to-GPU at 900+ GB/s.
Typically 4U–8U, liquid cooled, expensive.

## DATACENTER vs WORKSTATION vs CONSUMER

This cheat sheet focuses on **datacenter** accelerators. Here's why the distinction matters:

* **Datacenter** (this list) — Passive cooling, no display outputs, server power delivery.
Designed for rack chassis with front-to-back airflow.
Examples: P40, A30, A40, A100, L40S, H100.
* **Workstation** — Active cooling (fans on card), display outputs, certified drivers.
Designed for tower chassis. Examples: Nvidia RTX Pro 4000, Intel Arc Pro B70.
* **Consumer/gaming** — Active cooling, display outputs, gaming drivers.
Cheapest path to big VRAM (RTX 3090 = 24 GB for ~$600 used) but
Nvidia EULA prohibits datacenter deployment. Won't fit in a rack server.

If you have a tower workstation alongside your rack server,
workstation and consumer cards are worth considering for local LLM inference.

## WHAT FITS IN MY DELL R730 RACK SERVER

The R730 has PCIe 3.0 x16 slots, passive GPU cooling via chassis fans,
and limited GPU power delivery. All cards must be passive cooled PCIe.

| GPU | VRAM | Bandwidth | TDP | Used Price | Notes |
|-----|------|-----------|-----|------------|-------|
| **Tesla P40** _(current)_ | 24 GB | 346 GB/s | 250W | ~$150 | No tensor cores, but 24 GB is solid |
| **Tesla T4** | 16 GB | 300 GB/s | 70W | ~$200–300 | Low power, tensor cores, but only 16 GB |
| **V100 PCIe 32 GB** | 32 GB | 900 GB/s | 250W | ~$500 | 2.6x bandwidth over P40, tensor cores, HBM2 |
| **A30** | 24 GB | 933 GB/s | 165W | ~$400–600 | Best bandwidth-per-dollar, MIG, low power |
| **A40** | 48 GB | 696 GB/s | 300W | ~$900 | Double the VRAM, check R730 power headroom |

The A30 and A40 are PCIe 4.0 cards but will run fine in the R730's 3.0 slots
at reduced bandwidth. Everything A100 and above typically needs a newer server
with more power delivery and often SXM sockets.

**Best upgrade paths from the P40:**

* **Budget** — A30 ($400–600). Same 24 GB, 2.7x the bandwidth, tensor cores, MIG, lower TDP.
* **More VRAM** — A40 ($900). 48 GB lets you run 70B models quantized. Verify R730 power budget at 300W.
* **Balanced** — V100 PCIe 32 GB ($500). 32 GB HBM2, 2.6x bandwidth, proven ecosystem.
