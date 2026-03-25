# DATACENTER AI/GPU ACCELERATORS

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Nvidia datacenter AI/GPU accelerators._

Table of Contents

* [ARCHITECTURE TIMELINE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/nvidia/datacenter-ai-gpu-accelerators-cheat-sheet#architecture-timeline)
* [ACCELERATORS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/nvidia/datacenter-ai-gpu-accelerators-cheat-sheet#accelerators)
* [PCIe vs SXM](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/nvidia/datacenter-ai-gpu-accelerators-cheat-sheet#pcie-vs-sxm)
* [WHAT FITS IN MY DELL R730 RACK SERVER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/nvidia/datacenter-ai-gpu-accelerators-cheat-sheet#what-fits-in-my-dell-r730-rack-server)

## ARCHITECTURE TIMELINE

```text
2016  Pascal      P40, P100
2017  Volta       V100 (first tensor cores)
2018  Turing      T4 (INT8/INT4 inference)
2020  Ampere      A100 (3rd gen tensor, MIG)
2022  Ada         L4, L40S (FP8 support)
2023  Hopper      H100, H200 (transformer engine)
2025  Blackwell   B200, B300 (next-gen everything)
2026  Vera Rubin  announced, HBM4, 288 GB
```

## ACCELERATORS

| GPU | Year / Arch | VRAM / Type | Bandwidth | Tensor Cores / FP16 | TDP / PCIe | Used Price | Best For |
|-----|-------------|-------------|-----------|---------------------|------------|------------|----------|
| **Tesla P40** *(mine)* | 2016 / Pascal | 24 GB GDDR5X | 346 GB/s | None / Slow (1/64) | 250W / 3.0 | ~$150–200 | Budget 8–30B inference |
| **Tesla T4** | 2018 / Turing | 16 GB GDDR6 | 300 GB/s | 320 (2nd gen) / 65 TFLOPS | 70W / 3.0 | ~$200–300 | Low-power inference |
| **Tesla V100 PCIe** | 2017 / Volta | 16 or 32 GB HBM2 | 900 GB/s | 640 (1st gen) / 28 TFLOPS | 250W / 3.0 | ~$300–600 | Best R730 upgrade path |
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

## WHAT FITS IN MY DELL R730 RACK SERVER

The R730 has PCIe 3.0 slots and standard power.

* **P40** 24 GB for $300, great VRAM-per-dollar. No tensor cores, old arch, 250W
* **T4** 70W, tensor cores, fits any slot but only 16 GB VRAM
* **V100 PCIe 32 GB**  2.6x bandwidth over P40, tensor cores, 32 GB HBM2 250W

Everything A100 and above needs a newer server, more power, and often SXM sockets.
