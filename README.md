
# Healthy CLI Tool

## Overview
**Healthy** is a simple CLI (Command Line Interface) tool that checks the health of your services by performing a basic TCP connection to a specified domain and port. It allows users to verify if a service is reachable and optionally loop this health check at intervals of 5 seconds.

## Features
- **Domain Health Check**: Check if a domain is reachable on a given port.
- **Custom Port**: Specify a port to check (defaults to port `80` if not provided).
- **Looping**: Option to continuously check the service every 5 seconds.
- **Simple Output**: Get clear status messages indicating whether the service is UP or DOWN.

## Installation
1. Install Go from the official site: https://golang.org/dl/.
2. Clone this repository or copy the code into a file called `main.go`.
3. Run the following command to install necessary dependencies:
   ```bash
   go get github.com/urfave/cli/v2
   ```
4. Build the executable:
   ```bash
   go build -o healthy main.go
   ```

## Usage
You can use the `healthy` CLI tool by passing the domain and optional port as flags. The tool also supports looping to continuously check the service every 5 seconds.

### Basic Health Check
To check a domain's health on the default port (80):
```bash
./healthy --domain example.com
```

### Custom Port
To check a specific port (e.g., port 8080):
```bash
./healthy --domain example.com --port 8080
```

### Loop Mode
To continuously check the health every 5 seconds:
```bash
./healthy --domain example.com --loop
```

You can also use the shorter flag forms:
```bash
./healthy -d example.com -p 8080 -l
```

### Flags
- `--domain` or `-d`: (Required) The domain to check.
- `--port` or `-p`: (Optional) The port number to check (defaults to `80`).
- `--loop` or `-l`: (Optional) Continuously check the service every 5 seconds (defaults to `true`).

## Example Output
```bash
[UP] example.com:80 is reachable
From: 192.168.1.100:34567
To: 93.184.216.34:80
--------------------
[DOWN] example.com:8080 is unreachable, error: dial tcp: i/o timeout
```

## Error Handling
If the service is unreachable, Healthy provides a detailed error message with the domain, port, and the specific error encountered (e.g., timeout, unreachable).

---

Enjoy monitoring your services with **Healthy**!
