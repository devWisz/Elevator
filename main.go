package main

import (
	"bufio"
	"crypto/tls"
    "fmt"
)

type DownloadRecord struct {
	ID            int       `json:"id"`
	FileName      string    `json:"file_name"`
	FileType      string    `json:"file_type"`
	FileSize      string    `json:"file_size"`
}


func main() {
	fmt.Println("Elevate")
	fmt.Println("Download Anything Fast")
}