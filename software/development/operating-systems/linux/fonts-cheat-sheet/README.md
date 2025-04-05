# FONTS CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_How to install fonts on linux._

Table of Contents

* [INSTALL INCONSOLATA FONT](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/fonts-cheat-sheet#install-inconsolata-font)

## INSTALL INCONSOLATA FONT

As an example, lets install the inconsolata font,

Get package,

```bash
sudo apt-get install fonts-inconsolata -y
sudo fc-cache -fv
```

But it doesn’t have everything so Download .tff font from google,

```bash
wget -O ~/Downloads/Open_Sans.zip https://fonts.google.com/download?family=Open%20Sans
wget -O ~/Downloads/Inconsolata.zip https://fonts.google.com/download?family=Inconsolata
```

Unzip and place in fonts folder,

```bash
cd /usr/share/fonts
sudo mkdir googlefonts
sudo chmod -R --reference=opentype googlefonts
cd googlefonts
sudo unzip -d . ~/Downloads/Inconsolata.zip
```

Now register the fonts,

```bash
sudo fc-cache -fv
```

Check you have them,

```bash
find /usr/share/fonts | grep -i inconsolata
fc-list | grep -i inconsolata
fc-match Inconsolata
```
