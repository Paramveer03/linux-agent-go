# Linux Agent Go

## Project Overview
The Linux Agent Go is a lightweight, efficient agent designed for monitoring and managing Linux systems. It provides a range of capabilities to streamline operations and enhance security compliance settings.

## Features
- **Resource Monitoring**: Track CPU, memory, disk, and network usage.
- **Log Management**: Centralized logging and monitoring.
- **CIS Benchmark Checks**: Automatically assess compliance with Center for Internet Security benchmarks.
- **API Access**: Remote management through a RESTful API.
- **Custom Alerts**: Set up alerts based on system metrics and health checks.

## Installation Instructions
1. Clone the repository:
   ```sh
   git clone https://github.com/Paramveer03/linux-agent-go.git
   cd linux-agent-go
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Build the agent:
   ```sh
   go build -o linux-agent-go
   ```
4. Run the agent:
   ```sh
   ./linux-agent-go
   ```

## Usage Examples
To start monitoring your system:
```sh
./linux-agent-go --start
```

To check the current status:
```sh
./linux-agent-go --status
```

## CIS Checks
The Linux Agent Go provides built-in checks against the CIS benchmarks for various Linux distributions. To run the checks:
```sh
./linux-agent-go --cis-check
```

## API Documentation
The API provides endpoints for monitoring and managing the Linux agent:
- **GET /status**: Returns the current status of the agent.
- **POST /alerts**: Create a new alert.
- **GET /logs**: Retrieve logs.

For more detailed API usage, refer to [API Documentation](docs/api.md).