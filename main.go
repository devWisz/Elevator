package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)


type DownloadRecord struct {
	ID            int       `json:"id"`
	FileName      string    `json:"file_name"`
	FileType      string    `json:"file_type"`
	FileSize      string    `json:"file_size"`
	OriginalURL   string    `json:"original_url"`
	DownloadedAt  time.Time `json:"downloaded_at"`
}

const historyFile = "record.json"

var outOfScopeDomains = []string{

	"youtube.com", "youtu.be", "vimeo.com",
	"tiktok.com", "instagram.com", "facebook.com",
}

func main() {
	fmt.Println("Elevator")
	fmt.Println("Download Anything")
	fmt.Println(strings.Repeat("-", 45))

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nMain Menu:")
		fmt.Println("1. New Download")
		fmt.Println("2. View History & Re-download")
		fmt.Println("3. Clear History")
		fmt.Println("4. Exit")
		fmt.Print("\nSelect a option among 1-4 :")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			handleNewDownload(reader)
		case "2":
			handleHistory(reader)
		case "3":
			os.Remove(historyFile)
			fmt.Println("History cleared successfully")
		case "4":
			fmt.Println("Thank you for using Elevator..")
			return
		default:
			fmt.Println("error in choosing opption.Please select 1-4")
		}
	}
}

func handleNewDownload(reader *bufio.Reader) {
	fmt.Print("\nEnter URL: ")
	rawURL, _ := reader.ReadString('\n')
	rawURL = strings.TrimSpace(rawURL)

	if rawURL == "" {
		return
	}

	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil || !strings.Contains(rawURL, ".") {
		fmt.Println(" Error: Invalid URL format.")
		return
	}

	if isOutOfScope(parsedURL.String()) {
		fmt.Println("Error : Sorry!! Downloads are prohibited from this platform")
		return
	}

	defaultDir := getDefaultSaveDir()
	fmt.Printf("Save Directory (Press Enter for: %s)\nPath > ", defaultDir)
	
	saveDir, _ := reader.ReadString('\n')
	saveDir = strings.TrimSpace(saveDir)
	
	if saveDir == "" {
		saveDir = defaultDir
	}

	saveDir = resolvePath(saveDir)

	suggestedName := getFilenameFromURL(parsedURL.String())
	fmt.Printf("File Name [%s]: ", suggestedName)
	newName, _ := reader.ReadString('\n')
	newName = strings.TrimSpace(newName)
	if newName == "" {
		newName = suggestedName
	}

	record, err := executeDownload(parsedURL.String(), saveDir, newName)
	if err != nil {
		fmt.Printf("Fail to download: %v\n", err)
	} else {
		saveToHistory(record)
	}
}

func handleHistory(reader *bufio.Reader) {
	history := loadHistory()
	if len(history) == 0 {
		fmt.Println("\nNo download history available.")
		return
	}

	fmt.Printf("\n%-4s | %-30s | %-10s | %-10s\n", "ID", "File Name", "Size", "Date")
	fmt.Println(strings.Repeat("-", 70))
	for _, r := range history {
		fmt.Printf("%-4d | %-30.30s | %-10s | %s\n",
			r.ID, r.FileName, r.FileSize, r.DownloadedAt.Format("2006-01-02"))
	}

	fmt.Print("\nEnter ID to Re-download (or Enter to cancel): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		return
	}

	var selected *DownloadRecord
	for _, r := range history {
		if fmt.Sprintf("%d", r.ID) == input {
			selected = &r
			break
		}
	}

	if selected != nil {
		fmt.Printf("\nRestarting download of: %s\n", selected.FileName)
		newRec, err := executeDownload(selected.OriginalURL, ".", selected.FileName)
		if err == nil {
			saveToHistory(newRec)
		}
	}
}



func executeDownload(targetURL, savePath, fileName string) (*DownloadRecord, error) {

	if err := os.MkdirAll(savePath, 0755); err != nil {
		return nil, fmt.Errorf("system error (could not access folder): %w", err)
	}

	client := &http.Client{
		Timeout: 60 * time.Minute,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := client.Get(targetURL)
	if err != nil {
		return nil, fmt.Errorf("network error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server error: %s", resp.Status)
	}

	contentType := resp.Header.Get("Content-Type")
	finalName := fixExtension(fileName, contentType)
	absPath := filepath.Join(savePath, finalName)


	if _, err := os.Stat(absPath); err == nil {
		ext := filepath.Ext(finalName)
		base := strings.TrimSuffix(finalName, ext)
		absPath = filepath.Join(savePath, fmt.Sprintf("%s_%d%s", base, time.Now().Unix(), ext))
	}

	out, err := os.Create(absPath)
	if err != nil {
		return nil, fmt.Errorf("could not create file on disk: %w", err)
	}
	defer out.Close()


	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		out.Close()
		os.Remove(absPath)
		fmt.Println("\n Download interrupted. Cleanup complete.")
		os.Exit(1)
	}()

	pw := &progressWriter{
		total:     resp.ContentLength,
		startTime: time.Now(),
	}

	_, err = io.Copy(out, io.TeeReader(resp.Body, pw))
	if err != nil {
		return nil, fmt.Errorf("transfer failed: %w", err)
	}

	fmt.Printf("\n DOWNLOAD COMPLETE\n Location: %s\n", absPath)

	return &DownloadRecord{
		FileName:      filepath.Base(absPath),
		FileType:      contentType,
		FileSize:      formatBytes(resp.ContentLength),
		OriginalURL:   targetURL,
		DownloadedAt:  time.Now(),
	}, nil
}

func resolvePath(p string) string {
	home, _ := os.UserHomeDir()


	if strings.HasPrefix(p, "~") {
		p = filepath.Join(home, p[1:])
	}

	
	lowerP := strings.ToLower(p)
	if lowerP == "desktop" {
		p = filepath.Join(home, "Desktop")
	} else if lowerP == "downloads" {
		p = filepath.Join(home, "Downloads")
	} else if lowerP == "documents" {
		p = filepath.Join(home, "Documents")
	}

	if !filepath.IsAbs(p) {
		p = filepath.Join(home, p)
	}

	return filepath.Clean(p)
}

func getDefaultSaveDir() string {
	home, _ := os.UserHomeDir()
	
	
	downloads := filepath.Join(home, "Downloads")
	
	
	if _, err := os.Stat(downloads); err == nil {
		return downloads
	}
	return home
}

func isOutOfScope(rawURL string) bool {
	u, _ := url.Parse(rawURL)
	host := strings.ToLower(u.Host)
	for _, domain := range outOfScopeDomains {
		if strings.Contains(host, domain) {
			return true
		}
	}
	return false
}

func fixExtension(filename string, contentType string) string {
	exts, _ := mime.ExtensionsByType(contentType)
	if len(exts) == 0 {
		return filename
	}
	if filepath.Ext(filename) != "" {
		return filename
	}
	return filename + exts[0]
}

func getFilenameFromURL(rawURL string) string {
	u, _ := url.Parse(rawURL)
	name := filepath.Base(u.Path)
	if name == "." || name == "/" || len(name) < 3 {
		return "download_file"
	}
	return name
}

type progressWriter struct {
	total     int64
	written   int64
	startTime time.Time
}

func (pw *progressWriter) Write(p []byte) (int, error) {
	n := len(p)
	pw.written += int64(n)
	percent := 0.0
	if pw.total > 0 {
		percent = float64(pw.written) / float64(pw.total) * 100
	}
	elapsed := time.Since(pw.startTime).Seconds()
	speed := float64(pw.written) / elapsed
	fmt.Printf("\rProgress: %.1f%% (%s/s) ", percent, formatBytes(int64(speed)))
	return n, nil
}

func formatBytes(b int64) string {
	if b <= 0 { return "Unknown" }
	const unit = 1024
	if b < unit { return fmt.Sprintf("%d B", b) }
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}

func loadHistory() []DownloadRecord {
	var history []DownloadRecord
	data, err := os.ReadFile(historyFile)
	if err != nil { return history }
	json.Unmarshal(data, &history)
	return history
}

func saveToHistory(record *DownloadRecord) {
	history := loadHistory()
	maxID := 0
	for _, r := range history {
		if r.ID > maxID { maxID = r.ID }
	}
	record.ID = maxID + 1
	history = append(history, *record)
	data, _ := json.MarshalIndent(history, "", "  ")
	os.WriteFile(historyFile, data, 0644)
}