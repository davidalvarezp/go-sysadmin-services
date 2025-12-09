# Ransomware Daemon

A high-performance **Go daemon** designed for **real-time ransomware detection and prevention** on Linux servers.  
This service monitors critical directories, detects suspicious file activity, and automatically responds to prevent data loss.

---

## **Features**

- Real-time monitoring of critical directories (`/home`, `/var/www`, `/etc`)
- Detection of suspicious file extensions (e.g., `.locked`, `.crypt`, `.encrypt`)
- Automatic backup of affected files to a safe directory
- Temporary isolation of directories under attack
- Logging of all events for auditing
- Can be run as a **systemd daemon** for persistent operation

---

## **Directory Structure**

```

ransomware-daemon/
├── main.go                # Main Go source code
├── README.md              # This documentation
└── systemd/
  └── ransomware-daemon.service  # systemd service file

````

---

## **Installation**

1. **Clone the repository**
```bash
git clone https://github.com/davidalvarezp/go-sysadmin-services.git
cd go-sysadmin-services/ransomware-daemon
````

2. **Build the daemon**

```bash
go build -o ransomware-daemon
sudo cp ransomware-daemon /usr/local/bin/
```

3. **Set up backup directory**

```bash
sudo mkdir -p /backup/ransomware_safe
sudo chmod 700 /backup/ransomware_safe
```

4. **Install systemd service**

```bash
sudo cp systemd/ransomware-daemon.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable ransomware-daemon
sudo systemctl start ransomware-daemon
```

5. **Check status**

```bash
sudo systemctl status ransomware-daemon
```

---

## **Configuration**

* **Watched directories:** Modify the `watchDirs` slice in `main.go` to add or remove directories.
* **Suspicious extensions:** Modify the `suspiciousExt` slice in `main.go` to detect additional file types.
* **Backup location:** Modify the `backupDir` variable in `main.go` to change the backup path.

---

## **Usage**

Once installed as a systemd service, the daemon will:

1. Monitor the specified directories in real time.
2. Backup any files with suspicious activity.
3. Temporarily isolate directories under attack.
4. Log all events to `/var/log/ransomware_daemon.log`.

---

## **Logging**

All actions are logged in:

```bash
/var/log/ransomware_daemon.log
```

Sample log entry:

```
2025-12-09 14:05:12 [ALERT] Suspicious file detected: /home/user/document.locked
2025-12-09 14:05:12 [BACKUP] /home/user/document.locked -> /backup/ransomware_safe/20251209_140512/document.locked
2025-12-09 14:05:12 [AISLADO] /home/user
```

---

## **Contributing**

* Submit pull requests for improvements, bug fixes, or new features.
* Ensure changes are compatible with Linux servers and do not break existing backup logic.
* Follow Go best practices and maintain modularity.
