
# System Metrics Collector

A lightweight and high-performance **Go daemon** to collect system metrics on Linux servers.  
It monitors **CPU, memory, disk, and network usage** and exposes them via HTTP or logs for integration with monitoring tools.

---

## **Features**

- Collects system metrics in real-time:
  - CPU cores and usage
  - Memory total and usage
  - Disk usage
  - Network usage (optional)
- Exposes metrics via HTTP endpoint (`/metrics`) for Prometheus or other monitoring tools
- Lightweight and efficient, suitable for production servers
- Can run as a **systemd service**
- Logs all metrics and status periodically

---

## **Directory Structure**

```

system-metrics-collector/
├── main.go                     # Main Go source code
├── README.md                   # Documentation
└── systemd/
  └── system-metrics-collector.service  # systemd service file

````

---

## **Installation**

1. **Clone the repository**
```bash
git clone https://github.com/davidalvarezp/go-sysadmin-services.git
cd go-sysadmin-services/system-metrics-collector
````

2. **Build the daemon**

```bash
go build -o system-metrics-collector
sudo cp system-metrics-collector /usr/local/bin/
```

3. **Install systemd service**

```bash
sudo cp systemd/system-metrics-collector.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable system-metrics-collector
sudo systemctl start system-metrics-collector
```

4. **Check status**

```bash
sudo systemctl status system-metrics-collector
```

---

## **Usage**

* Access metrics via HTTP:

```bash
curl http://localhost:8080/metrics
```

Sample output:

```
cpu_cores: 8
memory_total: 16777216000
memory_used: 8234240000
disk_total: 500107862016
disk_used: 256000000000
```

* Integrate with monitoring dashboards (Prometheus, Grafana, etc.)

---

## **Configuration**

* HTTP port: Change the `:8080` in `main.go` if needed.
* Metrics collection interval: Can be modified by adding a ticker loop if periodic logging is desired.
* Extend metrics: Can add network I/O, process count, or custom checks using `gopsutil` library.

---

## **Contributing**

* Pull requests, bug fixes, and suggestions are welcome.
* Follow Go best practices and ensure the daemon remains lightweight.
* Keep modular code for integration with other sysadmin tools.
