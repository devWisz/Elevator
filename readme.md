# Elevator

Elevator is a lightweight CLI file downloader written in Go.It allows you to download anything you wanted to.


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

Download history is stored locally in:

```text
record.json
```

## License

Open source and available for anyone to use, modify, and contribute.
