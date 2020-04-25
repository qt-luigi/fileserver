# fileserver

fileserver is a static file HTTP server.

## Installation & Execution

When you have installed the Go, Please execute following commands:

```sh
go get -u github.com/qt-luigi/fileserver/
cd filesever/cmd
go build -o fileserver main.go arg.go (for Windows: -o fileserver.exe)
./fileserver
```

## Usage

```sh
$ fileserver -h
fileserver is a static file HTTP server.

Usage:
	fileserver [-p port] [-d directory]

Arguments are:

	-p port
		port number
	-d directory
		file directory
```

## License

MIT

## Author

Ryuji Iwata

## Note

This tool is mainly using by myself. :-)

