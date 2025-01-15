// main.go
package main

import (
    "fmt"
    "net/http"
    "math/rand"
    "time"
    
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    requestCounter = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"path"},
    )
    
    responseTime = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "http_response_time_seconds",
            Help:    "Response time in seconds",
            Buckets: []float64{0.1, 0.5, 1, 2, 5},
        },
        []string{"path"},
    )

    randomGauge = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Name: "random_number",
            Help: "A random number between 0 and 100",
        },
    )
)

func init() {
    prometheus.MustRegister(requestCounter)
    prometheus.MustRegister(responseTime)
    prometheus.MustRegister(randomGauge)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    
    requestCounter.WithLabelValues("/").Inc()
    
    // Simulate some work
    time.Sleep(time.Duration(rand.Float64() * float64(time.Second)))
    
    fmt.Fprintf(w, "Hello, Metrics World!")
    
    duration := time.Since(start).Seconds()
    responseTime.WithLabelValues("/").Observe(duration)
}

func updateRandomMetric() {
    for {
        randomGauge.Set(rand.Float64() * 100)
        time.Sleep(5 * time.Second)
    }
}

func main() {
    go updateRandomMetric()
    
    http.HandleFunc("/", mainHandler)
    http.Handle("/metrics", promhttp.Handler())
    
    fmt.Println("Server starting on :8080")
    http.ListenAndServe(":8080", nil)
}