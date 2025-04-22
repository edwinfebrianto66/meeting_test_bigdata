package main

import (
    "encoding/csv"
    "fmt"
    "net/http"
    "os"
    "strconv"
    "time"
    "encoding/json"
    "runtime"
)

type Response struct {
    TotalRows       int     `json:"total_rows"`
    AvgAge          float64 `json:"avg_age"`
    TotalSalary     int     `json:"total_salary"`
    ExecutionTimeMs int64   `json:"execution_time_ms"`
    MemoryUsedMB    float64 `json:"memory_used_mb"`
}

func testBigData(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    file, err := os.Open("data/dataset.csv")
    if err != nil {
        http.Error(w, "Failed to open CSV", 500)
        return
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        http.Error(w, "Failed to read CSV", 500)
        return
    }

    totalRows := 0
    totalAge := 0
    totalSalary := 0

    for i, row := range records {
        if i == 0 {
            continue
        }
        age, _ := strconv.Atoi(row[2])
        salary, _ := strconv.Atoi(row[3])
        totalAge += age
        totalSalary += salary
        totalRows++
    }

    elapsed := time.Since(start).Milliseconds()
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    memoryMB := float64(m.Alloc) / 1024.0 / 1024.0

    response := Response{
        TotalRows:       totalRows,
        AvgAge:          float64(totalAge) / float64(totalRows),
        TotalSalary:     totalSalary,
        ExecutionTimeMs: elapsed,
        MemoryUsedMB:    memoryMB,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/test-bigdata", testBigData)
    fmt.Println("Go service running on port 4001")
    http.ListenAndServe(":4001", nil)
}
