# COMMON COMMANDS CHEAT SHEET

`Common commands` _lists my favorite linux commands._

Sections,

* [SYSTEM INFO](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/common-commands-cheat-sheet#system-info)
* [HARDWARE INFO](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/common-commands-cheat-sheet#hardware-info)
* [USER INFO](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/common-commands-cheat-sheet#user-info)
* [MONITORING](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/common-commands-cheat-sheet#monitoring)
* [FILE AND DIRECTORY](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/common-commands-cheat-sheet#file-and-directory)
  * [FILE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/common-commands-cheat-sheet#file)
  * [FILE TRANSFER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/common-commands-cheat-sheet#file-transfer)
  * [DIRECTORY](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/common-commands-cheat-sheet#directory)
  * [ARCHIVE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/common-commands-cheat-sheet#archive)
  * [SEARCH](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/common-commands-cheat-sheet#search)
* [PROCESSES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/common-commands-cheat-sheet#processes)
* [NETWORKING](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/common-commands-cheat-sheet#networking)

* [SEARCH](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/common-commands-cheat-sheet#monitoring)
* [MONITORING](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/common-commands-cheat-sheet#monitoring)
* [MONITORING](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/common-commands-cheat-sheet#monitoring)
* [MONITORING](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/common-commands-cheat-sheet#monitoring)

I got a lot of this [here](https://www.linuxtrainingacademy.com/linux-commands-cheat-sheet/).

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## SYSTEM INFO

```bash
# Display Linux system information
uname -a

# Display kernel release information
uname -r

# Show how long the system has been running + load
uptime

# Show system host name
hostname

# Display the IP addresses of the host
hostname -I

# Show system reboot history
last reboot

# Show the current date and time
date

# Show this month's calendar
cal

# Display who is online
w

# Who you are logged in as
whoami
```

## HARDWARE INFO

```bash
# Display CPU information
cat /proc/cpuinfo

# Display memory information
cat /proc/meminfo

# Display free and used memory ( -h for human readable, -m for MB, -g for GB.)
free -h

# Show info about disk sda
hdparm -i /dev/sda

# Show free and used space on mounted filesystems
df -h

# Show free and used inodes on mounted filesystems
df -i

# Display disks partitions sizes and types
fdisk -l
```

## USER INFO

```bash
# Display the user and group ids of your current user.
id

# Display the last users who have logged onto the system.
last

# Show who is logged into the system.
who
w

# Create a group named "test".
groupadd test

# Create an account named john, with a comment of "John Smith" and create the user's home directory.
useradd -c "John Smith" -m john

# Delete the john account.
userdel john

# Add the john account to the sales group
usermod -aG sales john
```

## MONITORING

```bash
# Display and manage the top processes
top
htop

# Display processor related statistics
mpstat 1

# List all open files on the system
lsof

# List files opened by user jeff
lsof -u jeff

# Display free and used memory ( -h for human readable, -m for MB, -g for GB.)
free -h
```

## FILE AND DIRECTORY

### FILE

```bash
# List all files in a long listing (detailed) format
ls -al

# Change file permissions (User, Group World) (Read Write Execute)
chmod 777 filename

# Create an empty file
touch file

# View the contents of file
cat file

# Remove (delete) file
rm file

# Force removal of file without prompting for confirmation (Dangerous)
rm -f file

# Copy file1 to file2
cp file1 file2

# Rename or move file1 to file2. If file2 is an existing directory,
# move file1 into directory file2
mv file1 file2

# Create symbolic link to linkname
ln -s /path/to/file linkname

# Browse through a text file
less file

# Display the first 10 lines of file
head file

# Display the last 10 lines of file
tail file

# Display the last 10 lines of file and "follow" the file as it grows.
tail -f file
```
### FILE TRANSFER

```bash
# Secure copy file.txt to the /tmp folder on server
scp file.txt server:/tmp

# Copy *.html files from server to the local /tmp folder.
scp server:/var/www/*.html /tmp

# Copy all files and directories recursively from server to the current system's /tmp folder.
scp -r server:/var/www /tmp

# Synchronize /home to /backups/home
rsync -a /home /backups/

# Synchronize files/directories between the local and remote system with compression enabled
rsync -avz /home server:/backups/
```

### DIRECTORY

```bash
# Display the present working directory
pwd

# To go up one level of the directory tree.
cd ..

# Go to the $HOME directory
cd

# Change to the /etc directory
cd /etc

# Create a directory
mkdir directory

# Remove the directory/file and its contents recursively (Dangerous)
rm -r directory

# Forcefully remove directory recursively (Dangerous)
rm -rf directory

# Copy source_directory recursively to destination. 
# If destination exists, copy source_directory into destination, 
# otherwise create destination with the contents of source_directory.
cp -r source_directory destination
```

### ARCHIVE

```bash
# Create tar named archive.tar containing directory.
tar cf archive.tar directory

# Extract the contents from archive.tar.
tar xf archive.tar

# Create a gzip compressed tar file name archive.tar.gz.
tar czf archive.tar.gz directory

# Extract a gzip compressed tar file.
tar xzf archive.tar.gz

# Create a tar file with bzip2 compression
tar cjf archive.tar.bz2 directory

# Extract a bzip2 compressed tar file.
tar xjf archive.tar.bz2
```

### SEARCH

```bash
# Search for pattern in file
grep pattern file

# Search recursively for pattern in directory
grep -r pattern directory

# Find files and directories by name
locate name

# Find files in /home/john that start with "prefix".
find /home/john -name 'prefix*'

# Find files larger than 100MB in /home
find /home -size +100M
```

## PROCESSES

```bash
# Display your currently running processes
ps

# Display all the currently running processes on the system.
ps -ef

# Display process information for processname
ps -ef | grep processname

# Kill process with process ID of pid
kill pid

# Kill all processes named processname
killall processname

# Start program in the background
program &

# Display stopped or background jobs
bg

# Brings the most recent background job to foreground
fg

# Brings job n to the foreground
fg n
```

## NETWORKING

```bash
# Display all network interfaces and ip address
ifconfig -a

# Display eth0 address and details
ifconfig eth0

# Query or control network driver and hardware settings
ethtool eth0

# Send ICMP echo request to host
ping host

# Display whois information for domain
whois domain

# Display DNS information for domain
dig domain

# Reverse lookup of IP_ADDRESS
dig -x IP_ADDRESS

# Display DNS ip address for domain
host domain

# Display the network address of the host name.
hostname -i

# Display all local ip addresses
hostname -I

# Download http://domain.com/file
wget http://domain.com/file

# Display listening tcp and udp ports and corresponding programs
netstat -nutlp

# Capture and display all packets on interface eth0
tcpdump -i eth0

# Monitor all traffic on port 80 ( HTTP )
tcpdump -i eth0 'port 80'

# Connect to host as your local username.
ssh host

# Connect to host as user
ssh user@host

# Connect to host using port
ssh -p port user@host
```
