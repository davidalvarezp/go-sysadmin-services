
# Go Sysadmin Services

A collection of high-performance Go daemons and tools for **system administration**, **monitoring**, and **security automation**.  
These tools are designed to be lightweight, concurrent, and suitable for deployment on Linux servers.

## Features

- Real-time file monitoring and ransomware prevention
- System metrics collection (CPU, memory, disk, network)
- Automated backups and file integrity checks
- Network scanning and IP blocking
- Logging and alerting integration with Slack, Email, and SIEM

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/go-sysadmin-services.git
cd go-sysadmin-services

# Build individual service (example: ransomware-daemon)
cd ransomware-daemon
go build -o ransomware-daemon
sudo cp ransomware-daemon /usr/local/bin/
````

## Usage

```bash
# Run the daemon
sudo systemctl start ransomware-daemon
sudo systemctl enable ransomware-daemon
```

## Services Included

1. **Ransomware Daemon**
2. **System Metrics Collector**
3. **Directory Integrity Monitor**

## Contributing

Pull requests, issues, and suggestions are welcome!
Please follow Go best practices and maintain modularity for each service.

## License

MIT License
