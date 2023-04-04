# WSL2 WITH XSERVER CHEAT SHEET

_`wsl2 with xserver._

Table of Contents

* [STEP 1: SET DISPLAY VARIABLE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wls2-with-xserver-cheat-sheet#step-1-set-display-variable)
* [STEP 2: INBOUND FIREWALL RULE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wls2-with-xserver-cheat-sheet#step-2-inbound-firewall-rule)
* [STEP 3: INSTALL XSERVER FOR WINDOWS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wls2-with-xserver-cheat-sheet#step-3-install-xserver-for-windows)

## STEP 1: SET DISPLAY VARIABLE

We need to set the DISPLAY variable to X.X.X.X:0 for GTKwave.

Add to .bashrc,

```bash
export DISPLAY=$(awk '/nameserver / {print $2; exit}' /etc/resolv.conf 2>/dev/null):0
export LIBGL_ALWAYS_INDIRECT=1
```

The end result will be something like DISPLAY=172.X.X.X:0

## STEP 2: SET INBOUND FIREWALL RULE

Set an inbound firewall rule to allow the
X11 on port 6000 with scope 172.16.0.0/12.

## STEP 3: INSTALL XSERVER FOR WINDOWS

Windows does not come with an xserver, so you must get one.

VcXsrv is a free Windows X Server. Download from
[here](https://sourceforge.net/projects/vcxsrv/).

When configuring disable access control.

Make sure this programs firewall rule is both for private
and public networks.
