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

const historyFile = "record.json"

var outOfScopeDomains = []string {
	"youtube.com","viemo.com","tiktok.com","instagram.com","facebook.com",

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
		handleNewDownload(reader)
		case "2":
		handleHistory(reader)
		case "3":
	
		case "4":
			klm 
		default:
			fmt.Println("error in choosing opption")
		}

		func handleNewDownload(reader *bufio.Reader){
			fmt.Print("\nEnter URL : ")
			rawURL, _ := reader.ReadString('\n')
			rawURL = strings.TrimSpace(rawURL)

			if rawURL == ""{
				return
			}

			parsedURL, err := url.ParseRequestURL(rawURL)
			if err != nil || !strings.Contains(rawURL, "."){
				fmt.Println("Error !! Invalid URL form.")
				return
			}

		}
		
		func handleHistory(reader *bufio.Reader){
history := loadHistory()
		}

		func executeDownload(){

if err := os.MkdirAll(savePath,0755);err != nil {
	return nil, fmt.Errorf("System Error(Could not access folder) : %w",err)
}
 
client := &http.Client {
	Timeout: 60*time.Minute,
	Transport : &http.Transport{
	}
}
		}

		func resolvePath( p string ) string{

		}