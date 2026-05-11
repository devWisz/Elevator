package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"  
	"net/url"
	"os"
)

type DownloadRecord struct {
	ID            int       `json:"id"`
	FileName      string    `json:"file_name"`
	FileType      string    `json:"file_type"`
	FileSize      string    `json:"file_size"`
	OriginalURL   string    `json:"original_url"`
	LocalLocation string    `json:"local_location"`
	DownloadedAt  time.Time `json:"downloaded_at"`
}


func main() {
	fmt.Println("Elevator | Download anything fast")
	fmt.Println("Status: The best")
	fmt.Println(strings.Repeat("-", 45))

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nMain Menu:")
		fmt.Println("1. New Download")
		fmt.Println("2. View History & Re-download")
		fmt.Println("3. Clear History")
		fmt.Println("4. Exit")
		fmt.Print("\nSelection > ")

		