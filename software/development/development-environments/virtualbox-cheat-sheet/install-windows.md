# INSTALL WINDOWS CHEAT SHEET

_Basis steps to install Windows on VirtualBox._

## WINDOWS 11

**DOWNLOAD .iso IMAGE**

* GET .iso IMAGE
  * Download image from [official website](https://www.microsoft.com/en-us/software-download/windows11)

**VIRTUALBOX - CREATE NEW VM AND ATTACH .iso IMAGE**  

* CREATE VM
  * Name "VB-Windows-11" (Must be less than 15 characters for Windows Host Name)
  * Chose Windows 11 Pro 64-bit (8192 MB RAM, ~.100 GB Disk, .vdi, dynamically allocated)
* VM SETTINGS
  * Attach .iso image in Settings -> Storage
  * Settings->System->Processor
    * 4 Processors
    * Enable PAE/NX
  * Settings->Display->Screen
    * Video Memory: 256MB
    * Graphics Controller: VBoxSVGA
    * Enable 3D Acceleration enabled
    * Scale Factor 200% (This will help cut down on video RAM usage)

**START VM**

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

* LOGIN
  * Login as Jeff deCola

**INSTALL GUEST ADDITIONS ON WINDOWS**

* INSTALL CD FROM VM MENU
  * Devices->Insert Guest Additions CD image
* RUN GUEST ADDITIONS EXECUTABLE  
  * Open File Manger in Windows
  * Run `VBoxWindowsAdditions-amd64.exe`

**DISPLAY (AUTO RESIZE)**  

* VM MENU
  * This should now be available
  * View->Auto Resize Guest Display
  * Make sure your Host Settings->Display is 200% (Help with native resolution like 4K)

**DRAG AND DROP SETTINGS**

* VM MENU
  * Devices->Drag and Drop->Bidirectional

**VIRTUALBOX - NETWORK - BRIDGE MODE**

* SET BRIGE
  * The VM will receive it's own IP address if DHCP is enabled in the network.
  * Settings -> Network -> Adapter 1
    * `Bridged Adapter`
    * `Realtek Gaming GbE (GIGabit Ethernet) Family Controller`