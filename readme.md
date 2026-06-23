# Elevator

Elevator is a lightweight CLI file downloader written in Go. It lets you download files from URLs, keep track of download history, and re-download files whenever needed.

Demo Video link : https://youtu.be/zoV_UFQ96WE

Release: https://github.com/devWisz/Elevator/releases/tag/1.0

## Features

1. Download Files
2. Download History
3. Smart Handling

### Restrictions

Downloads from some platforms are intentionally blocked, including:

* YouTube
* Vimeo
* TikTok
* Instagram
* Facebook

## Installation

```bash
git clone https://github.com/devWisz/Elevator.git
cd Elevator
go build -o elevator
```

### Run

Linux/macOS

```bash
./elevator
```

Windows

```bash
elevator.exe
```

## Build for Windows

```bash
go build -o elevator.exe
```

Cross-compile from Linux/macOS:

```bash
GOOS=windows GOARCH=amd64 go build -o elevator.exe
```

## Usage

```text
1. New Download
2. View History & Re-download
3. Clear History
4. Exit
```

Download history is stored locally in:

```text
record.json
```

## License

Open source and available for anyone to use, modify, and contribute.
