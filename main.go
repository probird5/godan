package main

import (
  "encoding/json"
  "fmt"
  "os"
)

type responseStruct struct {
  City  string  `json:"city"`
  RegionCode string `json:"region_code"`
  OS string `json:"os"`
  Isp string `json:"isp"`
  Ports []int `json:"ports"`
  Hostnames []string `json:"hostnames"`
  Data []struct {
   Timestamp string `json:"timestamp"`
  } `json:"data"`
}


func main() {
    // Open the JSON file
    jsonFile, err := os.Open("output.json")
    if err != nil {
        fmt.Println("Error opening JSON file:", err)
        return
    }
    defer jsonFile.Close()

    // Decode the JSON data into the responseStruct
    var info responseStruct
    if err := json.NewDecoder(jsonFile).Decode(&info); err != nil {
        fmt.Println("Error decoding JSON:", err)
        return
    }

    // Custom output formatting
    fmt.Printf("City: %s\n", info.City)
    fmt.Printf("Region Code: %s\n", info.RegionCode)
    fmt.Printf("Operating System: %s\n", info.OS) // If you want to display the OS
    fmt.Printf("ISP: %s\n", info.Isp)
    fmt.Print("Ports: ")
    for i, port := range info.Ports {
        fmt.Print(port)
        if i < len(info.Ports)-1 {
            fmt.Print(", ")
        }
    }
    fmt.Println()

    fmt.Print("Hostnames: ")
    for i, hostname := range info.Hostnames {
        fmt.Print(hostname)
        if i < len(info.Hostnames)-1 {
            fmt.Print(", ")
        }
    }
    fmt.Println()

    // Printing details from nested structs or slices within the struct
    for _, data := range info.Data {
        fmt.Println("Data Record:")
        fmt.Printf("    Timestamp: %s\n", data.Timestamp)
    }
}
