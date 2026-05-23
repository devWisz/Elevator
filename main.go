package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"
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
	os.Remove(historyFile)
	fmt.Println("Historry cleared successfully")
		case "4":
			fmt.Println("Thank you for using Elevator..")
		default:
			fmt.Println("error in choosing opption.Please select 1-4/")
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

			if isOutofScope(parseURL.String()){
				fmt.Println("Error : Sorry!! Downlaods are prohibited from this platform")
				return
			}


			defaultDir :=getDefaultSaveDir()
			fmt.Println("Save Directory(Press Enter for: %s)\nPath>",defaultDir)

			saveDir,_ := reader.ReadString('\n')
			saveDir = strings.TrimSpace(saveDir)

			if saveDir ==""{

				saveDir = defaultDir 
			}

			saveDir = resolvePath(saveDir)
		}

		
		func handleHistory(reader *bufio.Reader){
history := loadHistory()
if len(history ==0){

	fmt.Println("\n NO download history available")
	return
}

fmt.Printf("\n%-4s | %-30s | %-10s | %-10s\n", "ID","File Name","Size","Data")
fmt.Println(strings.Repeat("-",70))

for _,r := range history {
	  fmt.Printf("%-4d | %-30.30s | %=10s | %s\n",
	  r.ID, r.FileName . r.FileSize, r.DownloadAt.Format("2006-01-02"))
} 

fmt.Print("\nEnter ID to ReSDownload (or Enter to cancel): ")
input , _ := reader.ReadString('\n')
input = strings.TrimSpace(input)
if input == ""{
	return
	
}
		}

		func executeDownload(){

if err := os.MkdirAll(savePath,0755);err != nil {
	return nil, fmt.Errorf("System Error(Could not access folder) : %w",err)
}
 
client := &http.Client {
	Timeout: 60*time.Minute,
	Transport : &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify:true},
	},
}

resp, err := client.Get(targertURL)
if err != nil {
	return nil, 
	fmt.Errorf("Error in Network: %w",err)
}

defer resp.Body.Close()


if resp.StatusCode !=http.StatusOK {
	return nil, fmt.Errorf("Error in Server: %s",resp.Status)
}
		}

		func resolvePath( p string ) string{

			home, _ := os.UserHomeDir(

				if strings.HasPrefix(p, "-")
				p = filepath.Join(Home,p[1:])
			)
		}

		lowerP := strings.ToLower(p)
        if lowerP =="desktop"{
			p= filepath.Join(home,"Desktop")
		} else if lowerP =="downloads"{
			p = filepath.Join(home,"Downloads")
		} else if lowerP =="documents"{
			p=filePath.Join(home,documents)
		}


		func isOutOfScope(rawURL string) bool {
			u, _ := url.Parse(rawURL)
			host := strings.ToLower(u.Host)
			for _, domain := range outOfscopeDomains {
				if strings.Contains(host,domain){
					return true
				}
			}
			
		 return false
		}

		
		func fixExtension(filename string , contentType string)string {

			exts ,_ : mime.ExtensionByType(ContentType)
           if len(exts)==0 {
	       return filename
}

if filepath.Ext(filename) != ""{

	return filename
}

return filename +exts[0]
		}

		func getFilenamefromURL(rawURL string)string{
 
			u,_ := url.Parse(rawURL)
			name := filepath.Base(u.Path)
			if name == "." || name  "/" || len(name) <3 {
				return "download_file"
			}

			return name

		}

		type progressWriter struct {
			total int64
			written int64
			startTime time.Time
		}


		func (pw *progressWriter) Writer(p []byte) (int,error){

			n:=len(p)
			pw.written += int64(n)
			percent := 0.0
			if pw.total >0 {

				percent = float64(pw.written)/float64(pw.total)*100
			}

			elapsed := time.Since(pw.startTime).Seconds()
			speed := float64(pw.written)/elapsed
			fmt.Println("\rProgress: %1f%%(%s/s)",percent,formalBytes(int64(speed)))
		return n,nil
		}

		func loadHistory() [] DownloadRecord{

			var history []DownloadRecord 
			data, err := os.ReadFile(historyFile)
			if err != nil {return history }
 json.Unmarshal(data, &history)

 return history


		} 