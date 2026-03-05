package system

import (
    "fmt"
    "os"
    "os/exec"
    "runtime"
)

// GetHostInfo retrieves and displays system information
func GetHostInfo() {
    hostname, err := os.Hostname()
    if err != nil {
        fmt.Println("Error getting hostname:", err)
        return
    }

    osVersion, err := exec.Command("uname", "-o").Output()
    if err != nil {
        fmt.Println("Error getting OS version:", err)
        return
    }

    kernelVersion, err := exec.Command("uname", "-r").Output()
    if err != nil {
        fmt.Println("Error getting kernel version:", err)
        return
    }

    uptime, err := exec.Command("uptime", "-p").Output()
    if err != nil {
        fmt.Println("Error getting uptime:", err)
        return
    }

    fmt.Printf("Hostname: %s\n", hostname)
    fmt.Printf("OS Version: %s\n", osVersion)
    fmt.Printf("Kernel Version: %s\n", kernelVersion)
    fmt.Printf("Uptime: %s\n", uptime)
}

func main() {
    GetHostInfo()
}