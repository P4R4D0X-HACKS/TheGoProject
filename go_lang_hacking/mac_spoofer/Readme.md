# Go MAC Address Changer
======================

This Go program allows you to change the MAC address of a specified network interface on your system. It lists all available network interfaces and provides usage instructions.

## Features
--------

-   Lists all network interfaces and their current MAC addresses.
-   Allows changing the MAC address of an active network interface.
-   Displays help information with usage instructions.

## Prerequisites
-------------

-   Go must be installed on your machine. You can download it from the [official Go website](https://golang.org/dl/).
-   The program requires `sudo` privileges to change the MAC address.

## Installation
------------

1.  Clone this repository to your local machine:
```bash
git clone https://github.com/P4R4D0X-HACKS/TheGoProject
cd your-repo
```
2. Build or run the program:
```bash
go run main.go
```

## Usage
-----

To change the MAC address of a network interface, use the following command format:

```bash
go run main.go -iface <interface> -newMac <address>
```

### Flags

-   `-iface <interface>`: Specify the interface for which you want to change the MAC address (e.g., `eth0`).
-   `-newMac <address>`: Specify the new MAC address to assign (e.g., `00:11:22:33:44:55`).
-   `-help`: Display help information.

**Examples**
```bash
go run main.go -iface eth0 -newMac 00:11:22:33:44:55
```

```bash
go run main.go -help
```
