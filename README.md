# MY CHEAT SHEETS

[![CodeClimate Issues](https://codeclimate.com/github/JeffDeCola/my-cheat-sheets/badges/issue_count.svg)](https://codeclimate.com/github/JeffDeCola/my-cheat-sheets/issues)
[![MIT License](http://img.shields.io/:license-mit-blue.svg)](http://jeffdecola.mit-license.org)

_A place to keep all my cheat sheets
for the complete development of
ASIC/FPGA hardware or a software app/service.
These cheat sheets contain info I gathered from other sources and
figured out over time. It's an ongoing process._

Table of Contents,

* **[SOFTWARE CHEAT SHEETS](https://github.com/JeffDeCola/my-cheat-sheets#software-cheat-sheets) -
  _For the goal of creating an App/Service_**
  * [DEVELOPMENT](https://github.com/JeffDeCola/my-cheat-sheets#development)
    * DEVELOPMENT ENVIRONMENTS
    * LANGUAGES
    * OPERATING SYSTEMS
    * SOFTWARE ARCHITECTURES
    * SOURCE / VERSION CONTROL
  * [OPERATIONS TOOLS](https://github.com/JeffDeCola/my-cheat-sheets#operations-tools)
    * CONFIGURATION MANAGEMENT
    * CONTINUOUS INTEGRATION / CONTINUOUS DEPLOYMENT
    * ORCHESTRATION
    * TELEMETRY
  * [SERVICE ARCHITECTURES](https://github.com/JeffDeCola/my-cheat-sheets#service-architectures)
    * SaaS - SOFTWARE AS A SERVICE
    * FaaS - FUNCTIONS AS A SERVICE
    * PaaS - PLATFORM AS A SERVICE
    * CaaS - CONTAINERS AS A SERVICE
    * IaaS - INFRASTRUCTURE AS A SERVICE
  * [SERVICE PROVIDERS](https://github.com/JeffDeCola/my-cheat-sheets#service-providers)
    * AMAZON AWS
    * GOOGLE GCP
    * MICROSOFT AZURE
* **[HARDWARE CHEAT SHEETS](https://github.com/JeffDeCola/my-cheat-sheets#hardware-cheat-sheets) -
  _For the goal of creating an ASIC/FPGA_**
  * [DEVELOPMENT](https://github.com/JeffDeCola/my-cheat-sheets#development-1)
    * FPGA DEVELOPMENT BOARDS
    * HARDWARE ARCHITECTURES
    * LANGUAGES
  * [TOOLS](https://github.com/JeffDeCola/my-cheat-sheets#tools)
    * SIMULATION
    * SYNTHESIS
    * TIMING
* **[OTHER CHEAT SHEETS](https://github.com/JeffDeCola/my-cheat-sheets#other-cheat-sheets) -
  _Random things I'm interested in_**
  * [COMPUTER HARDWARE](https://github.com/JeffDeCola/my-cheat-sheets#computer-hardware)
  * [HISTORY](https://github.com/JeffDeCola/my-cheat-sheets#history)
  * [MATHEMATICS](https://github.com/JeffDeCola/my-cheat-sheets#mathematics)
  * [MINING CRYPTOCURRENCY](https://github.com/JeffDeCola/my-cheat-sheets#mining-cryptocurrency)
  * [NETWORKS](https://github.com/JeffDeCola/my-cheat-sheets#networks)
  * [RANDOM STUFF](https://github.com/JeffDeCola/my-cheat-sheets#random-stuff)
  * [SCIENCE](https://github.com/JeffDeCola/my-cheat-sheets#science)
  * [SINGLE BOARD COMPUTERS (SBC)](https://github.com/JeffDeCola/my-cheat-sheets#single-board-computers-sbc)

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/) updated with
[ci-README.md](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/ci-README.md)

## SOFTWARE CHEAT SHEETS

Various Apps and tools for the goal of creating an App/Service.

![IMAGE - creating services environment overview - IMAGE](docs/pics/creating-services-environment-overview.jpg)

### DEVELOPMENT

* DEVELOPMENT ENVIRONMENTS

  * [software install methods](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/software-install-methods-cheat-sheet)
  * [vagrant](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/vagrant-cheat-sheet)
    ([my-vagrant-boxes](https://github.com/JeffDeCola/my-vagrant-boxes))
  * [virtualbox](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/virtualbox-cheat-sheet)
  * [visual studio code](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/visual-studio-code-cheat-sheet)
    * [verilog-HDL extension](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/visual-studio-code-cheat-sheet/verilog-hdl-extension.md)

* LANGUAGES

  * [go](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet),
    [go tools](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-tools-cheat-sheet)
    ([my-go-examples](https://github.com/JeffDeCola/my-go-examples),
    [my-go-packages](https://github.com/JeffDeCola/my-go-packages),
    [my-go-tools](https://github.com/JeffDeCola/my-go-tools))
  * [LaTeX](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet)
    * [LaTeX math mode](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet/latex-math-mode.md)
    * [LaTeX graphs](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet/latex-graphs.md)
      ([my-latex-graphs](https://github.com/JeffDeCola/my-latex-graphs))
  * [python](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/python-cheat-sheet)
    ([my-python-examples](https://github.com/JeffDeCola/my-python-examples))
  * [php](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/php-cheat-sheet)
    ([my-php-containers](https://github.com/JeffDeCola/my-php-containers))

* OPERATING SYSTEMS

  * LINUX
    ([my-linux-shell-scripts](https://github.com/JeffDeCola/my-linux-shell-scripts))

    * [common commands](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/common-commands-cheat-sheet)
    * [dns](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/dns-cheat-sheet),
      [create dns server using bind](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/dns-cheat-sheet/create-dns-server-using-bind.md)
    * [dual-boot-ubuntu-and-windows](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/dual-boot-ubuntu-and-windows)
    * [fonts](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/fonts-cheat-sheet)
    * [init SysV script](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/init-sysv-script-cheat-sheet)
    * [LS_COLORS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/ls_colors-cheat-sheet)
    * [network device configuration](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/network-device-configuration-cheat-sheet)
    * [redirect output](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/redirect-output-cheat-sheet)
    * [ssh and keys](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/ssh-and-keys-cheat-sheet)
    * [systemd](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/systemd-cheat-sheet)

  * LINUX DISTRIBUTIONS

    * _coming soon_

  * WINDOWS

    * [wsl (bash on ubuntu on windows)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl-bash-on-ubuntu-on-windows-cheat-sheet)

  * MAC OS

    * [launchd](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/macos/launchd-cheat-sheet)
    * [LSCOLORS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/macos/lscolors-cheat-sheet)

* SOFTWARE ARCHITECTURES

  * API

    * [RESTful](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/api/RESTful-cheat-sheet)
      ([RESTful-API-test](https://github.com/JeffDeCola/RESTful-API-test))
    * [youtube content id api](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/api/youtube-content-id-api-cheat-sheet)
    * [youtube data api v3](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/api/youtube-data-api-v3-cheat-sheet)
    * [youtube player api](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/api/youtube-player-api-cheat-sheet)

  * AUTHORIZATION

    * [OAuth 2.0 authorization](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/authorization/OAuth-2.0-authorization-cheat-sheet)
  
  * BLOCKCHAIN

    * [blockchain](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/blockchain/blockchain-cheat-sheet)

  * CRYPTOGRAPHY

    * [asymmetric-cryptography](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/cryptography/asymmetric-cryptography-cheat-sheet)
    * [hashing](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/cryptography/hashing-cheat-sheet)
    * [symmetric-cryptography](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/cryptography/symmetric-cryptography-cheat-sheet)

  * DATABASE

    * [postgreSQL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/database/postgresql-cheat-sheet)
    * [redis](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/database/redis-cheat-sheet)
    * [relational vs non-relational](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/database/relational-versus-non-relational-cheat-sheet)

  * MESSAGING

    * [NATS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/messaging/NATS-cheat-sheet)
    * [protobuf](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/messaging/protobuf-cheat-sheet)

  * MICROSERVICES

    * [microservices](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/microservices/microservices-cheat-sheet)
      ([catch-microservice](https://github.com/JeffDeCola/catch-microservice),
      [data-crunch-engine](https://github.com/JeffDeCola/data-crunch-engine))

* SOURCE / VERSION CONTROL

  * [git](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/source-version-control/git-cheat-sheet)

### OPERATIONS TOOLS

* CONFIGURATION MANAGEMENT

  * [ansible](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/configuration-management/ansible-cheat-sheet)

* CONTINUOUS INTEGRATION / CONTINUOUS DEPLOYMENT

  * [concourse](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/continuous-integration-continuous-deployment/concourse-cheat-sheet)
    ([my-concourse-ci-tasks](https://github.com/JeffDeCola/my-concourse-ci-tasks))

* ORCHESTRATION

  * BUILDS / DEPLOYMENT / CONTAINERS

    * [docker](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/orchestration/builds-deployment-containers/docker-cheat-sheet)
      ([my-docker-image-builds](https://github.com/JeffDeCola/my-docker-image-builds))
    * [packer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/orchestration/builds-deployment-containers/packer-cheat-sheet)
      ([my-packer-image-builds](https://github.com/JeffDeCola/my-packer-image-builds))
    * [terraform](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/orchestration/builds-deployment-containers/terraform-cheat-sheet)

  * CLUSTER MANAGERS / RESOURCE MANAGEMENT / SCHEDULING

    * [kubernetes](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/orchestration/cluster-managers-resource-management-scheduling/kubernetes-cheat-sheet)
    * [marathon](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/orchestration/cluster-managers-resource-management-scheduling/marathon-cheat-sheet)
      ([hello-go-deploy-marathon](https://github.com/JeffDeCola/hello-go-deploy-marathon))
    * [mesos](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/orchestration/cluster-managers-resource-management-scheduling/mesos-cheat-sheet)

  * DISCOVERY / CONFIGURATION

    * [consul](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/orchestration/discovery-configuration/consul-cheat-sheet)

* TELEMETRY

  * [grafana](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/telemetry/grafana-cheat-sheet)
  * [stackdriver](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/telemetry/stackdriver-cheat-sheet)

### SERVICE ARCHITECTURES

* Overview

  * [overview of service architectures](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/overview/overview-of-service-architectures-cheat-sheet)

* SaaS - SOFTWARE AS A SERVICE

  * [keybase](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/software-as-a-service/keybase-cheat-sheet)
  * [slack](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/software-as-a-service/slack-cheat-sheet)

* FaaS - FUNCTION AS A SERVICE

  * [aws lambda](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/function-as-a-service/aws-lambda-cheat-sheet)
    (_hello-go-deploy-aws-lambda repo coming soon_)
  * [google cloud functions (gcf)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/function-as-a-service/google-cloud-functions-cheat-sheet)
    (_hello-go-deploy-gcf repo coming soon_)
  * [microsoft azure functions](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/function-as-a-service/microsoft-azure-functions-cheat-sheet)
    (_hello-go-deploy-azure-functions repo coming soon_)

* PaaS - PLATFORM AS A SERVICE

  * [aws elastic beanstalk](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/platform-as-a-service/aws-elastic-beanstalk-cheat-sheet)
    ([hello-go-deploy-aws-elastic-beanstalk](https://github.com/JeffDeCola/hello-go-deploy-aws-elastic-beanstalk))
  * [google app engine (gae)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/platform-as-a-service/google-app-engine-cheat-sheet)
    ([hello-go-deploy-gae](https://github.com/JeffDeCola/hello-go-deploy-gae))
  * [microsoft azure app service](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/platform-as-a-service/microsoft-azure-app-service-cheat-sheet)
    ([hello-go-deploy-azure-app-service](https://github.com/JeffDeCola/hello-go-deploy-azure-app-service))

* CaaS - CONTAINERS AS A SERVICE

  * [amazon elastic container service (ecs)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/containers-as-a-service/amazon-elastic-container-service-cheat-sheet)
    ([hello-go-deploy-amazon-ecs](https://github.com/JeffDeCola/hello-go-deploy-amazon-ecs))
  * [amazon elastic container service for kubernetes (eks)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/containers-as-a-service/amazon-elastic-container-service-for-kubernetes-cheat-sheet)
    ([hello-go-deploy-amazon-eks](https://github.com/JeffDeCola/hello-go-deploy-amazon-eks))
  * [google kubernetes engine (gke)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/containers-as-a-service/google-kubernetes-engine-cheat-sheet)
    ([hello-go-deploy-gke](https://github.com/JeffDeCola/hello-go-deploy-gke))
  * [microsoft azure kubernetes service (aks)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/containers-as-a-service/microsoft-azure-kubernetes-service-cheat-sheet)
    ([hello-go-deploy-aks](https://github.com/JeffDeCola/hello-go-deploy-aks))

* IaaS - INFRASTRUCTURE AS A SERVICE

  * [amazon elastic compute cloud (ec2)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/infrastructure-as-a-service/amazon-elastic-compute-cloud-cheat-sheet)
    ([hello-go-deploy-amazon-ec2](https://github.com/JeffDeCola/hello-go-deploy-amazon-ec2))
  * [google compute engine (gce)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/infrastructure-as-a-service/google-compute-engine-cheat-sheet)
    ([hello-go-deploy-gce](https://github.com/JeffDeCola/hello-go-deploy-gce))
  * [microsoft azure virtual machines (vm)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/infrastructure-as-a-service/microsoft-azure-virtual-machines-cheat-sheet)
    ([hello-go-deploy-azure-vm](https://github.com/JeffDeCola/hello-go-deploy-azure-vm))

### SERVICE PROVIDERS

* [amazon web services (aws)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-providers/amazon-web-services-cheat-sheet)
* [google cloud platform (gcp)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-providers/google-cloud-platform-cheat-sheet)
  * [google cloud pub/sub](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-providers/google-cloud-platform-cheat-sheet/google-cloud-pub-sub.md)
  * [google cloud spanner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-providers/google-cloud-platform-cheat-sheet/google-cloud-spanner.md)
  * [google cloud storage](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-providers/google-cloud-platform-cheat-sheet/google-cloud-storage.md)
  * [google source repositories (git)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-providers/google-cloud-platform-cheat-sheet/google-source-repositories-git.md)
  * [google stackdriver monitoring](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-providers/google-cloud-platform-cheat-sheet/google-stackdriver-monitoring.md)
* [microsoft azure](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-providers/microsoft-azure-cheat-sheet)

## HARDWARE CHEAT SHEETS

Various Apps and tools for the goal of creating an ASIC/FPGA.

![IMAGE - creating an ASIC environment overview - IMAGE](docs/pics/creating-an-asic-environment-overview.jpg)

### DEVELOPMENT

* FPGA DEVELOPMENT BOARDS
  
  * [Digilent ARTY-S7](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/fpga-development-boards/digilent-arty-s7-cheat-sheet)

* HARDWARE ARCHITECTURES

  * [ARM](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/hardware-architectures/arm-cheat-sheet)

* LANGUAGES

  * [SystemVerilog](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet),
    [VS Code verilog-HDL extension](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/visual-studio-code-cheat-sheet/verilog-hdl-extension.md)
    ([my-systemverilog-examples](https://github.com/JeffDeCola/my-systemverilog-examples))

### TOOLS

* SIMULATION

  * [GTKWave](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/tools/simulation/gtkwave-cheat-sheet)
  * [iverilog](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/tools/simulation/iverilog-cheat-sheet)

* SYNTHESIS

  * [Xilinx Vivado](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/tools/synthesis/xilinx-vivado-cheat-sheet)

* TIMING

  * _coming soon_

## OTHER CHEAT SHEETS

Random other things I'm interested in.

### COMPUTER HARDWARE

* APPLE

  * [macbook pro models](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/computer-hardware/apple/macbook-pro-models-cheat-sheet)

* PC

  * [cpu](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/computer-hardware/pc/cpu-cheat-sheet)
  * [ram](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/computer-hardware/pc/ram-cheat-sheet)

* TECHNOLOGIES

  * [data storage](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/computer-hardware/technologies/data-storage-cheat-sheet)
  * [hard drives](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/computer-hardware/technologies/hard-drives-cheat-sheet)
  * [microSD cards](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/computer-hardware/technologies/microSD-cards-cheat-sheet)
  * [monitors](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/computer-hardware/technologies/monitors-cheat-sheet)
  * [ssd](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/computer-hardware/technologies/ssd-cheat-sheet)
  * [usb](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/computer-hardware/technologies/usb-cheat-sheet)

### HISTORY

* AMERICAN

  * [Pre-Colonial (Before 1600)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/history/american/pre-colonial-cheat-sheet)
  * [Colonial America (1600-1799)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/history/american/colonial-america-cheat-sheet)
  * [A New Nation (1800-1849)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/history/american/a-new-nation-cheat-sheet)
  * [Civil War & Reconstruction (1850-1899)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/history/american/civil-war-and-reconstruction-cheat-sheet)
  * [Progressive Era & World Wars (1900-1949)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/history/american/progressive-era-and-world-wars-cheat-sheet)
  * [Mid-Century & Cold War (1950-1999)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/history/american/mid-century-and-cold-war-cheat-sheet)
  * [New Millennium (2000 Onward)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/history/american/new-millennium-cheat-sheet)

* WARS

  * [The American Revolutionary War (1775-1783)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/history/wars/the-american-revolutionary-war-cheat-sheet)
  * [The War of 1812 (1812-1815)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/history/wars/the-war-of-1812-cheat-sheet)
  * [The American Civil War (1861-1865)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/history/wars/the-american-civil-war-cheat-sheet)
  * [WW1 (1914-1918)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/history/wars/ww1-cheat-sheet)
  * [WW2 (1939-1945)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/history/wars/ww2-cheat-sheet)

### MATHEMATICS

_This entire sections is a massive work in progress, it will take me
quite a few years to finish._

* APPLIED

  * [computer science (software)](https://github.com/JeffDeCola/my-cheat-sheets#software-cheat-sheets)
  * [control theory](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/applied/control-theory-cheat-sheet)
  * [cryptography](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/applied/cryptography-cheat-sheet)
  * [electrical engineering](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/applied/electrical-engineering-cheat-sheet)
  * [game theory](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/applied/game-theory-cheat-sheet)
  * [optimization](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/applied/optimization-cheat-sheet)
  * [physical sciences (chemistry & physics)](https://github.com/JeffDeCola/my-cheat-sheets#science)
  * [probability and statistics](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/applied/probability-and-statistics-cheat-sheet)

* NOTABLE EQUATIONS

  * [notable equations](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/notable-equations/notable-equations-cheat-sheet)

* PURE

  * CHANGES

    * [calculus](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/changes/calculus-cheat-sheet)
    * [chaos theory](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/changes/chaos-theory-cheat-sheet)
    * [complex analysis](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/changes/complex-analysis-cheat-sheet)
    * [dynamical systems](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/changes/dynamical-systems-cheat-sheet)
    * [vector calculus](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/changes/vector-calculus-cheat-sheet)

  * FOUNDATIONS

    * [category theory](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/foundations/category-theory-cheat-sheet)
    * [fundamental rules](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/foundations/fundamental-rules-cheat-sheet)
    * [mathematics logic](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/foundations/mathematics-logic-cheat-sheet)
    * [set theory](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/foundations/set-theory-cheat-sheet)
    * [theory of computation (complexity theory)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/foundations/theory-of-computation-complexity-theory-cheat-sheet)

  * NUMBER SYSTEMS

    * [complex numbers](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/number-systems/complex-numbers-cheat-sheet)
    * [integers](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/number-systems/integers-cheat-sheet)
    * [natural numbers](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/number-systems/natural-numbers-cheat-sheet)
    * [prime numbers](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/number-systems/prime-numbers-cheat-sheet)
    * [real numbers](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/number-systems/real-numbers-cheat-sheet)

  * SPACES

    * [differential geometry](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/spaces/differential-geometry-cheat-sheet)
    * [fractal geometry](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/spaces/fractal-geometry-cheat-sheet)
    * [geometry](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/spaces/geometry-cheat-sheet)
    * [measure theory](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/spaces/measure-theory-cheat-sheet)
    * [topology](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/spaces/topology-cheat-sheet)
    * [trigonometry](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/spaces/trigonometry-cheat-sheet)

  * STRUCTURES

    * [algebra](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/structures/algebra-cheat-sheet)
    * [combinations](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/structures/combinations-cheat-sheet)
    * [group theory](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/structures/group-theory-cheat-sheet)
    * [linear algebra](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/structures/linear-algebra-cheat-sheet)
    * [number theory](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/structures/number-theory-cheat-sheet)
    * [order theory](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mathematics/pure/structures/order-theory-cheat-sheet)

### MINING CRYPTOCURRENCY

* CRYPTOCURRENCY

  * [cryptocurrency](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/cryptocurrency/cryptocurrency-cheat-sheet)
    ([jeffCoin](https://github.com/JeffDeCola/jeffCoin),
    [crypto-miner-manager](https://github.com/JeffDeCola/crypto-miner-manager),
    [crypto-wallet-status](https://github.com/JeffDeCola/crypto-wallet-status))

* FULL NODE

  * [mine ZEC (Zcash) on macOS cpu using zcashd](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/full-node/mine-ZEC-macOS-cpu-zcashd-cheat-sheet)

* MULTIPLE POOLS

  * [mine MULTIPLE COINS on windows gpu using multipoolminer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/multiple-pools/mine-MULTICOINS-windows-gpu-multipoolminer-cheat-sheet)
  * [mine MULTIPLE COINS on windows gpu using awesome miner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/multiple-pools/mine-MULTICOINS-windows-gpu-awesome-miner-cheat-sheet)

* POOLS

  * [mine BEAM (BEAM) on windows gpu using lolMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-BEAM-windows-gpu-lolMiner-cheat-sheet)
  * [mine BTG (Bitcoin Gold) on windows gpu using gminer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-BTG-windows-gpu-gminer-cheat-sheet)
  * [mine RVN (Ravencoin) on windows gpu using t-rex](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-RVN-windows-gpu-t-rex-cheat-sheet)
  * [mine ZCL (Zclassic) on windows gpu using lolMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-ZCL-windows-gpu-lolMiner-cheat-sheet)
  * [mine ZEC (Zcash) on windows gpu using funakoshiMiner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-ZEC-windows-gpu-funakoshiMiner-cheat-sheet)
  * [mine ZEL (ZelCash) on windows gpu using gminer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/mining-cryptocurrency/pools/mine-ZEL-windows-gpu-gminer-cheat-sheet)

### NETWORKS

* CELLULAR

  * [telecommunication standards](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/networks/cellular/telecommunication-standards-cheat-sheet)

* WAN / LAN

  * [ethernet wan / lan](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/networks/wan-lan/ethernet-wan-lan-cheat-sheet)
  * [wireless lan](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/networks/wan-lan/wireless-lan-cheat-sheet)

### RANDOM STUFF

* [famous quotes](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-stuff/famous-quotes-cheat-sheet)
* [typical highway interchanges](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-stuff/typical-highway-interchanges-cheat-sheet)

### SCIENCE

* BIOLOGICAL / LIFE SCIENCE

  * [biology](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/science/biological-life-science/biology-cheat-sheet)
  * [social](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/science/biological-life-science/social-cheat-sheet)

* EARTH & SPACE SCIENCE

  * [astronomy](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/science/earth-and-space-science/astronomy-cheat-sheet)

* PHYSICAL SCIENCE

  * [chemistry](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/science/physical-science/chemistry-cheat-sheet)
  * [physics](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/science/physical-science/physics-cheat-sheet)

### SINGLE BOARD COMPUTERS (SBC)

* HUMMINGBOARD

  * [install and configure OS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/hummingboard/install-and-configure-os-cheat-sheet)
  * [specifications](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/hummingboard/specifications-cheat-sheet)

* RASPBERRY PI (RasPi)

  * [install and configure OS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/raspberry-pi/install-and-configure-os-cheat-sheet)
  * [specifications](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/raspberry-pi/specifications-cheat-sheet)
