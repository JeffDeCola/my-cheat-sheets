# WSL2 WITH XSERVER CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_`wsl2 with xserver._

Table of Contents

* [WSL2 NOW COMES WITH WSLG](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-with-xserver-cheat-sheet#wsl2-now-comes-with-wslg)
  * [DISPLAY=:0](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-with-xserver-cheat-sheet#display0)
  * [SCALING](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-with-xserver-cheat-sheet#scaling)
* [USING YOU OWN XSERVER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-with-xserver-cheat-sheet#using-you-own-xserver)
  * [STEP 1: SET DISPLAY VARIABLE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-with-xserver-cheat-sheet#step-1-set-display-variable)
  * [STEP 2: SET INBOUND FIREWALL RULE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-with-xserver-cheat-sheet#step-2-set-inbound-firewall-rule)
  * [STEP 3: INSTALL XSERVER FOR WINDOWS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-with-xserver-cheat-sheet#step-3-install-xserver-for-windows)

## WSL2 NOW COMES WITH WSLG

wsl2 now comes with it's own xserver `wslg`.

### DISPLAY=:0

Set your display to,

```txt
export DISPLAY=:0
```

### SCALING

If you want to scale the display, create a `.wslgconfig` file in your
windows user directory.

```text
[system-distro-env]
;WESTON_RDP_HI_DPI_SCALING=false
WESTON_RDP_FRACTIONAL_HI_DPI_SCALING=true
WESTON_RDP_DEBUG_DESKTOP_SCALING_FACTOR=200
```

Restart wsl in powershell to use the new settings,

```bash
wsl --shutdown
```

## USING YOU OWN XSERVER

You can also use your won xserver such as VcXsrv.

### STEP 1: SET DISPLAY VARIABLE

We need to set the DISPLAY variable to X.X.X.X:0 for GTKwave.

Add to .bashrc,

```bash
export DISPLAY=$(awk '/nameserver / {print $2; exit}' /etc/resolv.conf 2>/dev/null):0
export LIBGL_ALWAYS_INDIRECT=1
```

The end result will be something like DISPLAY=172.X.X.X:0

### STEP 2: SET INBOUND FIREWALL RULE

Set an inbound firewall rule to allow the
X11 on port 6000 with scope 172.16.0.0/12.

### STEP 3: INSTALL XSERVER FOR WINDOWS

Windows does not come with an xserver, so you must get one.

VcXsrv is a free Windows X Server. Download from
[here](https://sourceforge.net/projects/vcxsrv/).

When configuring disable access control.

Make sure this programs firewall rule is both for private
and public networks.
