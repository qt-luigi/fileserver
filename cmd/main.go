package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/qt-luigi/fileserver"
)

var argPort int
var argDir string

func init() {
	flag.IntVar(&argPort, "p", 8080, "port number")
	flag.StringVar(&argDir, "d", "./", "file directory")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usageMsg)
	}
}

func main() {
	flag.Parse()

	if _, err := os.Stat(argDir); err != nil {
		log.Fatal(err)
	}
	log.Printf("file directory: %s\n", argDir)

	port := fmt.Sprintf(":%d", argPort)

	ipv4, err := fileserver.LocalIPv4()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("access URL: http://%s%s/\n", ipv4, port)

	log.Fatal(http.ListenAndServe(port, http.FileServer(http.Dir(argDir))))
}
