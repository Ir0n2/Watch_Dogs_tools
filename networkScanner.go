package main

import (
    "fmt"
    "net"
    //"os"
    "sync"
    "time"
)

// Configuration for network scanning
const (
    networkRange = "109.233.191." // Set this to the IP prefix of your network (like "192.168.1.")
    startIP      = 1
    endIP        = 255
    port         = 80          // Define a single port or change this to scan multiple ports
    timeout      = 1 * time.Second
)

// Scans a single IP for open ports and identifies services
func scanIP(ip string, port int, wg *sync.WaitGroup) {
    defer wg.Done()

    address := fmt.Sprintf("%s:%d", ip, port)
    conn, err := net.DialTimeout("tcp", address, timeout)
    if err == nil {
        fmt.Printf("Device found at IP: %s\n", ip)
	fmt.Printf("Port %d open on %s\n", port, ip)
        
        // Optional: Banner grabbing to identify service
        banner := make([]byte, 1024)
        conn.SetReadDeadline(time.Now().Add(timeout))
        n, _ := conn.Read(banner)
        if n > 0 {
            fmt.Printf("Service on port %d: %s\n", port, string(banner[:n]))
        } else {
            fmt.Printf("Service on port %d: Unknown\n", port)
        }
        
        conn.Close()
    }
}

// Main function to initiate the scan
func main() {
    var wg sync.WaitGroup

    fmt.Printf("Starting scan on network: %s0-255\n", networkRange)
    for i := startIP; i <= endIP; i++ {
        ip := fmt.Sprintf("%s%d", networkRange, i)
        wg.Add(1)
        go scanIP(ip, port, &wg)
    }

    wg.Wait()
    fmt.Println("Scan complete")
}

