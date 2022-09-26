# INSTALL WINDOWS CHEAT SHEET

_Basis steps to install Windows on VirtualBox._

Documentation and Reference

* [Install arch linux mini](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-arch-linux-mini.md)
* [Install debian mini](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-debian-mini.md)
* [Install ubuntu with GNOME desktop](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-ubuntu-with-gnome-desktop.md)

## WINDOWS 11 PRO

**DOWNLOAD .iso IMAGE**

* GET .iso IMAGE
  * Download image from [official website](https://www.microsoft.com/en-us/software-download/windows11)

**VIRTUALBOX - CREATE NEW VM AND ATTACH .iso IMAGE**  

* NEW VM
  * Name "VB-Windows-11" (Must be less than 15 characters for Windows Host Name)
  * Chose Windows 11 Pro 64-bit (8192 MB RAM, ~.100 GB Disk, .vdi, dynamically allocated)
* ATTACH IMAGE
  * Settings->Storage with Controller: IDE
  * Attach .iso image
* PROCESSOR
  * Settings->System->Processor
  * 4 Processors
  * Enable PAE/NX
* DISPLAY
  * Settings->Display->Screen
  * Video Memory: 256MB
  * Graphics Controller: VBoxSVGA
  * Enable 3D Acceleration enabled
  * Scale Factor 200% (This will help cut down on video RAM usage)
* SET BRIDGE
  * The VM will receive it's own IP address if DHCP is enabled in the network
  * Settings -> Network -> Adapter 1
    * `Bridged Adapter`
    * `Realtek Gaming GbE (GIGabit Ethernet) Family Controller`

**START VM**

* **START VM**
* STOP AT INSTALL NOW
  * Stop when you see "Install Now" Button

**OPEN REGEDIT IN VM (NOT IN HOST)**

* START REGEDIT
  * Click Shift-F10 to bring up cmd shell
  * `regedit`
* HKEY_LOCAL_MACHINE\SYSTEM\Setup
  * Right-click over Setup and add New `Key` named "LabConfig"
  * Right-click over LabConfig and new `DWORD (32-bit) Value` named "BypassTPMCheck"
  * Set "BypassTPMCheck" from 0 to 1
  * Repeat for
    * "BypassSecureBootCheck"
    * "BypassRamCheck"
* EXIT REGEDIT

**CONTINUE INSTALL**

* INSTALL NOW
  * Click "Install Now" button

**Q & A**

* IF IT HANGS
  * Just reboot
* ANSWER QUESTIONS
  * Choose "I don’t have a product key"
  * Choose "Windows 11 Pro"
  * Chose Custom Install
  * Host "VB-Windows-11" (Only 15 characters)
  * Offline account
  * User Jeff DeCola
  * etc...
* **CLOSE VM**

**VIRTUALBOX - REMOVE .iso IMAGE**

* VM SETTINGS  
  * Remove image in Settings -> Storage

**VIRTUALBOX - ATTACH GUEST ADDITIONS.iso IMAGE**

* VM SETTINGS
  * Attach guest additions .iso image in Settings -> Storage

**FIRST LOGIN AS JEFF**

* **START VM**
* LOGIN
  * Login as Jeff DeCola

**INSTALL GUEST ADDITIONS ON WINDOWS**

* INSTALL CD FROM VM MENU
  * This is probably not needed
  * Devices->Insert Guest Additions CD image
* RUN GUEST ADDITIONS EXECUTABLE  
  * Open File Manger in Windows
  * Run `VBoxWindowsAdditions-amd64.exe`
  * When done, it will ask to reboot

**SHARED SETTINGS**

* VM MENU
  * Devices->Shared Folders->Shared Folder Settings
    * Pick where you want this folder
  * Devices->Shared ClipBoard->Bidirectional
  * Devices->Drag and Drop->Bidirectional

**DISPLAY (AUTO RESIZE)**  

* VM MENU
  * This should now be available
  * View->Auto Resize Guest Display
  * Make sure your Host Settings->Display is 200% (Help with native resolution like 4K)

**VIRTUALBOX - REMOVE GUEST ADDITIONS.iso IMAGE**

* **CLOSE VM**
* VM SETTINGS
  * Remove guest additions .iso image in Settings -> Storage
* **START VM**

**YOUR HOME NETWORK**

* BRIDGE MODE
  * Since we are in bridge mode, I like to configure my home router to set the same ip address
