# WORKSTATION AI/GPU CARDS

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Workstation GPUs for AI inference, content creation, and professional applications._

Table of Contents

* [HOW WORKSTATION DIFFERS FROM DATACENTER](#how-workstation-differs-from-datacenter)
* [NVIDIA WORKSTATION GPUs](#nvidia-workstation-gpus)
* [INTEL WORKSTATION GPUs](#intel-workstation-gpus)
* [CONSUMER GPUs USED AS WORKSTATION](#consumer-gpus-used-as-workstation)

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

## NVIDIA WORKSTATION GPUs

Formerly the Quadro line, now branded RTX Pro (Blackwell) and RTX A-series (Ampere/Ada).

| GPU | Year / Arch | VRAM / Type | Bandwidth | Tensor Cores | TDP | Price | Best For |
|-----|-------------|-------------|-----------|--------------|-----|-------|----------|
| **RTX A4000** | 2021 / Ampere | 16 GB GDDR6 | 448 GB/s | 192 (3rd gen) | 140W | ~$900 used | Budget workstation, CAD |
| **RTX A5000** | 2021 / Ampere | 24 GB GDDR6 | 768 GB/s | 256 (3rd gen) | 230W | ~$1,500 used | Mid-range AI + rendering |
| **RTX A6000** | 2020 / Ampere | 48 GB GDDR6 | 768 GB/s | 336 (3rd gen) | 300W | ~$2,500 used | Big VRAM workstation king |
| **RTX 4000 Ada** | 2023 / Ada | 20 GB GDDR6 | 360 GB/s | 240 (4th gen) | 130W | ~$1,250 | Low-power Ada, single slot |
| **RTX 5000 Ada** | 2023 / Ada | 32 GB GDDR6 | 576 GB/s | 400 (4th gen) | 250W | ~$4,000 | 32 GB + FP8 support |
| **RTX 6000 Ada** | 2022 / Ada | 48 GB GDDR6 | 960 GB/s | 568 (4th gen) | 300W | ~$6,500 | Flagship Ada workstation |
| **RTX Pro 4000** | 2025 / Blackwell | 24 GB GDDR7 | — | 5th gen | 200W | ~$1,800 | Current gen mid-range |
| **RTX Pro 6000** | 2025 / Blackwell | 96 GB GDDR7 | — | 5th gen | 350W | ~$8,000+ | Flagship Blackwell workstation |

All Nvidia workstation cards have active cooling, display outputs, ISV-certified
drivers (SolidWorks, Revit, Maya, etc.), and soft or hardware ECC.

## INTEL WORKSTATION GPUs

Intel Arc Pro B-Series (Battlemage / Xe2 architecture, launched March 2026).

| GPU | VRAM / Type | Bandwidth | AI Compute | TDP | Price | Best For |
|-----|-------------|-----------|------------|-----|-------|----------|
| **Arc Pro B40** | 12 GB GDDR6 | 192 GB/s | 99 INT8 TOPS | 75W | ~$230 | Entry-level, single slot |
| **Arc Pro B50** | 16 GB GDDR6 | 256 GB/s | 137 INT8 TOPS | 150W | ~$400 | Budget AI + display |
| **Arc Pro B60** | 24 GB GDDR6 | 480 GB/s | 197 INT8 TOPS | 150W | ~$660 | Mid-range, 24 GB sweet spot |
| **Arc Pro B65** | 32 GB GDDR6 | 608 GB/s | 197 INT8 TOPS | 200W | TBD (< $949) | Big VRAM, cost-optimized |
| **Arc Pro B70** | 32 GB GDDR6 | 608 GB/s | 367 INT8 TOPS | 160–290W | $949 | Flagship, local LLM inference |

Intel cards use oneAPI / OpenVINO / SYCL (not CUDA). Software ecosystem is maturing
but not as plug-and-play as Nvidia. All have active cooling, 4× DP 2.1 outputs,
soft ECC, SR-IOV, and multi-GPU scaling on Linux.

**Why Intel is interesting for AI:**

* 32 GB VRAM at $949 (B70) undercuts Nvidia workstation cards significantly
* Multi-GPU pooling across 4 cards = 128 GB combined VRAM
* SR-IOV for GPU sharing across VMs (homelab friendly)
* CUDA is the trade-off — you're on Intel's software stack

## CONSUMER GPUs USED AS WORKSTATION

Not workstation cards, but widely used for local AI inference in desktop/tower builds.
No ISV-certified drivers, no ECC, Nvidia EULA prohibits datacenter use.

| GPU | Year / Arch | VRAM | Bandwidth | TDP | Used Price | Notes |
|-----|-------------|------|-----------|-----|------------|-------|
| **RTX 3090** | 2020 / Ampere | 24 GB GDDR6X | 936 GB/s | 350W | ~$600 | The homelab legend, best VRAM/$ |
| **RTX 4090** | 2022 / Ada | 24 GB GDDR6X | 1,008 GB/s | 450W | ~$1,800 | Fastest consumer, 24 GB |
| **RTX 5090** | 2025 / Blackwell | 32 GB GDDR7 | 1,790 GB/s | 575W | ~$2,000+ | 32 GB, new gen, hard to find |
| **RX 7900 XTX** | 2023 / RDNA 3 | 24 GB GDDR6 | 960 GB/s | 355W | ~$700 | AMD option, ROCm on Linux |

The RTX 3090 at ~$600 used remains the sweet spot for local LLM work.
24 GB VRAM, ~3x faster than a Tesla P40, fits any standard desktop case.
