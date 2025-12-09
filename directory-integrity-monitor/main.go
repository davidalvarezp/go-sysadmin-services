package main

import (
    "log"
    "os"
    "path/filepath"
    "time"
)

var watchDirs = []string{"/etc", "/var/www"}
var scanInterval = 30 * time.Second
var snapshots = map[string]map[string]os.FileInfo{}

func snapshot(dir string) map[string]os.FileInfo {
    files := map[string]os.FileInfo{}
    filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
        if err == nil && !info.IsDir() {
            files[path] = info
        }
        return nil
    })
    return files
}

func detectChanges(oldSnap, newSnap map[string]os.FileInfo) {
    for file := range newSnap {
        if _, ok := oldSnap[file]; !ok {
            log.Printf("[ALERT] New file detected: %s", file)
        }
    }
    for file := range oldSnap {
        if _, ok := newSnap[file]; !ok {
            log.Printf("[ALERT] File removed: %s", file)
        }
    }
}

func main() {
    logFile := "/var/log/directory_integrity_monitor.log"
    f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    log.SetOutput(f)

    for _, dir := range watchDirs {
        snapshots[dir] = snapshot(dir)
    }

    log.Println("Directory Integrity Monitor started...")

    for {
        time.Sleep(scanInterval)
        for _, dir := range watchDirs {
            newSnap := snapshot(dir)
            detectChanges(snapshots[dir], newSnap)
            snapshots[dir] = newSnap
        }
    }
}
