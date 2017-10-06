# Flexswitch Support for Ingrasys S9100

## Overview
FlexSwitch by SnapRoute is a disaggregated micro-services based L2/L3 network stack, enabling organizations to achieve maximum agility, reliability, and security.

It has open source version on two sites: 
  - [OpenSnaproute] (Deprecated)
  - [OpxFlexswitch]

## Network stack support
Our patches for FlexSwitch network stack uses open source code released at [OpenSnaproute] to be the stack codebase, and leverage code from [OpxFlexswitch] to fix issues.

_**Note: The code on OpenSnaproute has been removed. We push the original code to the site to be our codebase.**_ 

## Asicd support
Snaproute's asic daemon serves as a hardware abstraction layer (HAL). Our implementation is based on open source code at [asicdopen]. We complete the SAI plugin implementaion using [SAI API 0.9.4]. We have verified it works on our white box Ingrasys S9100 and S8900 using Broadcom Tomahawk chip.

## Features
* BFP
* BGP
* DHCP relay
* ECMP
* LAG
* LLDP
* Loopback
* OSPF
* STP
* VRRP

## Build Guide
From this site, you have two ways to test FlexSwitch. One is to build the docker version package (Using Linux kernel driver), and the other is to build Ingrasys S9100 version package (Using asicd with SAI plugin).
You can use the following process to build packages. The deb files could be found under the folder *`reltools`*.

### Developing Environment Setup
#### Prerequisites
Prepare Ubuntu 14.04, and install GO 1.5.3 on it.

_Note: GO should be installed under `/user/local`._
#### Setup Environment
1. **Install git-lfs**
    ```
    $ curl -s https://packagecloud.io/install/repositories/github/git-lfs/.deb.sh | sudo bash
    $ sudo apt-get install git-lfs
    ```
    Check [git-lfs] if you are failed to install git-lfs.

2. **Download source code**

    Change to directory you want to use for downloading source code. The directory will be `[SrcHome]` for configuration in next section.
    ```
    $ git clone https://github.com/FlexSNR/reltools
    $ cd reltools
    $ fab setupDevEnv
    ```
    Type *`FlexSNR`* where asking Git username, and press enter for others.

3. **Configure environment variable**
    ```
    $ export PATH=$PATH:/usr/local/go/bin
    $ export SR_CODE_BASE=[SrcHome]
    $ export GOPATH=[SrcHome]/snaproute:[SrcHome]/external:[SrcHome]/generated
    ```

### Build docker version
```
$ python makePkg.py
```

### Build Ingrasys S9100 version
Change directory to `[SrcHome]`.
```
$ git clone https://github.com/FlexSNR/patch
$ cd patch
$ ./patch_util
$ cd ../reltools
$ python makePkg.py
```
_Note: To build successfully, you need to contact us for asic-related files_
## Deployment Guide
Our FlexSwitch is tested on [Open Network Linux] (ONL). You could follow the process the setup FlexSwitch on your machine.
### Setup ONIE
Follow the instruction to install ONIE on your machine.

ONIE Location:

https://github.com/opencomputeproject/onie/tree/master/machine/ingrasys/ingrasys_s9100
### Setup ONL
```
Install ONL by ONIE.
```
### Install dependency
```
$ apt-get install libjemalloc1, redis-tools, redis-server, libnl-3-200, libnl-genl-3-200
```
### Initialize switch ASIC
Please contact us.
### Install FlexSwitch
```
$ dpkg -i flexswitch_ingrasys_s9100-vagrant_1.0.0.171.44_amd64.deb
```
### Check FlexSwitch
```
$ ps aux | grep flex
root      3921  0.5  0.1 581028 15220 ?        Sl   02:07   1:38 /opt/flexswitch/bin/sysd -params=/opt/flexswitch/params
root      3928  5.2  2.3 1477512 187428 ?      Sl   02:07  14:58 /opt/flexswitch/bin/asicd -params=/opt/flexswitch/params
root      3931  0.1  0.4 619320 38132 ?        Sl   02:07   0:24 /opt/flexswitch/bin/lacpd -params=/opt/flexswitch/params
root      3939  0.1  0.4 544868 38136 ?        Sl   02:07   0:24 /opt/flexswitch/bin/stpd -params=/opt/flexswitch/params
root      3953  0.1  0.4 569036 37088 ?        Sl   02:07   0:25 /opt/flexswitch/bin/lldpd -params=/opt/flexswitch/params
root      3959  0.1  0.4 626180 37320 ?        Sl   02:07   0:28 /opt/flexswitch/bin/arpd -params=/opt/flexswitch/params
root      3966  0.1  0.1 745792 14396 ?        Sl   02:07   0:24 /opt/flexswitch/bin/ribd -params=/opt/flexswitch/params
root      3977  0.1  0.4 767580 38152 ?        Sl   02:07   0:23 /opt/flexswitch/bin/bfdd -params=/opt/flexswitch/params
root      3988  0.0  0.1 697076 15572 ?        Sl   02:07   0:08 /opt/flexswitch/bin/bgpd -params=/opt/flexswitch/params
root      4002  0.1  0.4 702936 37720 ?        Sl   02:07   0:24 /opt/flexswitch/bin/ospfv2d -params=/opt/flexswitch/params
root      4003  0.1  0.4 699024 36368 ?        Sl   02:07   0:24 /opt/flexswitch/bin/dhcpd -params=/opt/flexswitch/params
root      4015  0.1  0.4 684084 37404 ?        Sl   02:07   0:25 /opt/flexswitch/bin/dhcprelayd -params=/opt/flexswitch/params
root      4022  0.1  0.4 494296 36316 ?        Sl   02:07   0:25 /opt/flexswitch/bin/vrrpd -params=/opt/flexswitch/params
root      4029  0.1  0.4 489716 37708 ?        Sl   02:07   0:28 /opt/flexswitch/bin/vxland -params=/opt/flexswitch/params
root      4037  0.0  0.1 514240 13952 ?        Sl   02:07   0:04 /opt/flexswitch/bin/platformd -params=/opt/flexswitch/params
root      4045  0.2  0.5 649592 42824 ?        Sl   02:07   0:35 /opt/flexswitch/bin/ndpd -params=/opt/flexswitch/params
root      4053  0.0  0.1 446492 12956 ?        Sl   02:07   0:07 /opt/flexswitch/bin/fMgrd -params=/opt/flexswitch/params
root      4060  0.0  0.1 359244 11368 ?        Sl   02:07   0:04 /opt/flexswitch/bin/notifierd -params=/opt/flexswitch/params
root      4073  0.0  0.2 615540 19256 ?        Sl   02:07   0:04 /opt/flexswitch/bin/confd -params=/opt/flexswitch/params
root      4171  0.0  1.8 1040620 145860 ?      S    02:07   0:00 /opt/flexswitch/bin/asicd -params=/opt/flexswitch/params
root     14979  0.0  0.0  12732  2148 pts/3    S+   06:55   0:00 grep flex
```
### Test Flexswitch
Refererence Site:
1. [OPX Flx document]
2. [Snaproute Flexswitch document]
## License
Licenses for the software are available at [License](/LICENSE).

[OpenSnaproute]: <https://github.com/OpenSnaproute>
[OpxFlexswitch]: <https://github.com/open-switch>
[asicdopen]: <https://github.com/skotha-lnkd/asicdopen>
[SAI API 0.9.4]: <https://github.com/opencomputeproject/SAI>
[Open Network Linux]: <http://opennetlinux.org/>
[git-lfs]: <https://packagecloud.io/github/git-lfs/install>
[OPX Flx document]: <https://open-switch.github.io/flx-docs/developer.html>
[Snaproute Flexswitch document]: <http://docs.snaproute.com/index.html>
