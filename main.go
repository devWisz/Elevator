package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
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
	fmt.Println("Elevator")
	fmt.Println("Download Anything fast")
	fmt.Println(strings.Repeat("-", 45))

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nMain Menu:")
		fmt.Println("1. New Download")
		fmt.Println("2. View History & Re-download")
		fmt.Println("3. Clear History")
		fmt.Println("4. Exit")
		fmt.Print("\n Choose a option among 1-4 > ")

		choice,_ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choic {
		case "1": 
		abc 
		case "2":
			efg 
		case "3":
			hij 
		case "4":
			klm 
		default:
			fmt.Println("error in choosing opption")
		}

		