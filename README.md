# GoNS
DNS Lookup CLI written in Golang

## Overview

GoNS is a simple, fast, and flexible DNS lookup CLI tool built using Go (Golang). It allows users to query DNS records for a domain, specifically for A (IPv4) and MX (Mail Exchange) records. With the option to specify custom DNS servers, this tool is designed for network administrators, developers, and anyone who needs a quick DNS lookup from the command line.

## Purpose

The goal of GoNS is to provide a lightweight and efficient way for users to query DNS records directly from the command line. It supports A and MX record lookups and allows basic customization, such as choosing a DNS server to perform the query.

## Features

### 1. DNS Lookup Capabilities

GoNS supports two main types of DNS lookups:

- **A Record Lookup (IPv4)**: Resolve a domain to one or more IPv4 addresses.
- **MX Record Lookup**: Resolve the domain to its associated mail exchange servers.

### 2. Command-Line Interface (CLI) Features

- **Command Structure**: The tool follows the simple format:

```bash
gons <domain> --type=<record-type>
```

- **Record Types**: Query for **A** or **MX** records. Example:
- `--type=A` for A records (IPv4).
- `--type=MX` for MX records.

- **Custom DNS Server**: Users can specify a custom DNS server using the `--server` flag. Example:

```bash
--server=8.8.8.8
```


- **Plain Text Output**: The tool outputs results in a human-readable format:
- **A Record**: `IP Address: 192.168.1.1`
- **MX Record**: `MX Record: mail.example.com, Priority: 10`

- **Error Handling**: Displays user-friendly error messages for invalid domains or failed queries.
