package main

import (
    "log"
    "os"
    "os/exec"
    "path/filepath"
    "time"

    "github.com/fsnotify/fsnotify"
)

var watchDirs = []string{"/home", "/var/www", "/etc"}
var suspiciousExt = []string{".locked", ".crypt", ".encrypt"}
var backupDir = "/backup/ransomware_safe"

func backupFile(filePath string) {
    timestamp := time.Now().Format("20060102_150405")
    destDir := filepath.Join(backupDir, timestamp)
    os.MkdirAll(destDir, 0755)
    destFile := filepath.Join(destDir, filepath.Base(filePath))
    exec.Command("cp", "-p", filePath, destFile).Run()
    log.Printf("[BACKUP] %s -> %s", filePath, destFile)
}

func isolateDir(dir string) {
    exec.Command("chmod", "-R", "000", dir).Run()
    log.Printf("[AISLADO] %s", dir)
}

func isSuspicious(file string) bool {
    for _, ext := range suspiciousExt {
        if filepath.Ext(file) == ext {
            return true
        }
    }
    return false
}

func handleEvent(event fsnotify.Event) {
    if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
        if isSuspicious(event.Name) {
            log.Printf("[ALERT] Suspicious file detected: %s", event.Name)
            backupFile(event.Name)
            isolateDir(filepath.Dir(event.Name))
        }
    }
}

func main() {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }
    defer watcher.Close()

    for _, dir := range watchDirs {
        err = watcher.Add(dir)
        if err != nil {
            log.Fatal(err)
        }
    }

    log.Println("Ransomware Daemon started...")

    done := make(chan bool)
    go func() {
        for {
            select {
            case event := <-watcher.Events:
                go handleEvent(event)
            case err := <-watcher.Errors:
                log.Println("Error:", err)
            }
        }
    }()

    <-done
}
