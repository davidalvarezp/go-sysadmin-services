package main

import (
    "fmt"
    "log"
    "net/http"
    "runtime"

    "github.com/shirou/gopsutil/disk"
    "github.com/shirou/gopsutil/mem"
)

func metricsHandler(w http.ResponseWriter, r *http.Request) {
    v, err := mem.VirtualMemory()
    if err != nil {
        log.Println("Error fetching memory stats:", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    d, err := disk.Usage("/")
    if err != nil {
        log.Println("Error fetching disk stats:", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    metrics := map[string]interface{}{
        "cpu_cores":   runtime.NumCPU(),
        "memory_total": v.Total,
        "memory_used":  v.Used,
        "disk_total":   d.Total,
        "disk_used":    d.Used,
    }

    for k, v := range metrics {
        fmt.Fprintf(w, "%s: %v\n", k, v)
    }
}

func main() {
    http.HandleFunc("/metrics", metricsHandler)
    log.Println("System Metrics Collector started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
