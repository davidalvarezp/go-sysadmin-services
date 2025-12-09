# Directory Integrity Monitor

A **Go daemon** for monitoring changes in critical directories on Linux servers.  
It detects **unexpected file additions, deletions, or renames** and logs alerts for sysadmins, helping to prevent tampering or ransomware activity.

---

## **Features**

- Monitors specified directories in real-time (or periodically)
- Detects:
  - New files
  - Deleted files
  - Renamed files
- Logs all detected changes for auditing
- Can run as a **systemd service**
- Lightweight and efficient, designed for production servers

---

## **Directory Structure**

```

directory-integrity-monitor/
├── main.go
├── README.md
└── systemd/
└── directory-integrity-monitor.service

````

---

## **Installation**

1. **Clone the repository**
```bash
git clone https://github.com/davidalvarezp/go-sysadmin-services.git
cd go-sysadmin-services/directory-integrity-monitor
````

2. **Build the daemon**

```bash
go build -o directory-integrity-monitor
sudo cp directory-integrity-monitor /usr/local/bin/
```

3. **Install systemd service**

```bash
sudo cp systemd/directory-integrity-monitor.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable directory-integrity-monitor
sudo systemctl start directory-integrity-monitor
```

4. **Check status**

```bash
sudo systemctl status directory-integrity-monitor
```

---

## **Usage**

* The daemon monitors directories specified in the `watchDirs` variable in `main.go`.
* Logs detected changes to `/var/log/directory_integrity_monitor.log`.

Sample log entries:

```
2025-12-09 15:10:45 [ALERT] New file detected: /etc/new_config.conf
2025-12-09 15:11:12 [ALERT] File removed: /var/www/index.html
```

---

## **Configuration**

* **Watched directories:** Modify the `watchDirs` slice in `main.go` to add/remove directories.
* **Scan interval:** Modify `scanInterval` variable for periodic scanning.
* **Log file location:** Can be updated in `main.go` to a preferred path.

---

## **Contributing**

* Pull requests and suggestions are welcome.
* Ensure changes do not impact performance on large directories.
* Follow Go best practices and keep code modular for integration with other sysadmin tools.
