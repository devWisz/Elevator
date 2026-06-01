# Elevator

Elevator is a lightweight command-line file downloader that allows you to download files directly from URLs, manage download history, and quickly re-download previously downloaded files.

The application is designed to provide a simple download experience while automatically handling file naming, extensions, save locations, and download records.

---

Use it directly : https://github.com/devWisz/Elevator/releases/tag/1.0

## Features :

### New Download

Download files directly from a URL.

* URL validation before downloading
* Automatic file name detection
* Custom file name support
* Custom save directory support
* Automatic file extension detection
* Duplicate file name protection
* Download progress indicator with speed information
* Interrupt-safe downloads with automatic cleanup

### Download History

Keep track of all completed downloads.

* Stores download records locally
* View previously downloaded files
* Displays:

  * File name
  * File size
  * Download date
* Re-download any file using its saved URL

### History Management

Manage stored download records easily.

* View complete download history
* Re-download previous files
* Clear all history records

### Smart File Handling

* Automatically creates missing directories
* Prevents accidental overwriting
* Detects content type from server response
* Assigns proper file extensions when needed

### Download Safety

* Cleans up partially downloaded files when interrupted
* Handles network failures gracefully
* Detects invalid URLs before starting downloads

### Restricted Platforms

Downloads from certain platforms are intentionally blocked.

Examples include:

* YouTube
* Vimeo
* TikTok
* Instagram
* Facebook

---

## Installation

### Clone the Repository

```bash
git clone https://github.com/devWisz/Elevator.git
```

```bash
cd Elevator
```

### Build

```bash
go build -o elevator
```

### Run

Linux / macOS

```bash
./elevator
```

Windows

```powershell
elevator.exe
```

---

## Creating a Windows Executable

To generate a Windows executable named `elevator.exe`:

### From Windows

```bash
go build -o elevator.exe
```

### Cross Compile from Linux/macOS

```bash
GOOS=windows GOARCH=amd64 go build -o elevator.exe
```

---

## Usage

After launching the application:

```text
1. New Download
2. View History & Re-download
3. Clear History
4. Exit
```

### New Download

1. Select `1`
2. Enter a valid file URL
3. Choose a save directory or press Enter for the default location
4. Choose a file name or press Enter for the suggested name
5. Wait for the download to complete

### View History & Re-download

1. Select `2`
2. Browse previously downloaded files
3. Enter the ID of any record
4. The file will be downloaded again using the original URL

### Clear History

1. Select `3`
2. All stored download records will be removed

### Exit

1. Select `4`
2. The application closes safely

---

## Download Records

Download history is stored locally in:

```text
record.json
```

Each record contains:

* Download ID
* File name
* File type
* File size
* Original URL
* Download timestamp

---

##Some Notes to be tak

* Internet connection is required.
* Some servers may restrict automated downloads.
* File sizes may display as "Unknown" if the server does not provide content length information.
* Existing files are never overwritten automatically.

---

## License

This project is open source and available for anyone to use, modify, and contribute to.