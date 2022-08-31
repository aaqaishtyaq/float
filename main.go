package main

import (
	"embed"
	"flag"
	"fmt"
	"os"
)

var (
	port        string
	contentFile string
	buildString = "dev"
)

//go:embed static
var assets embed.FS

func init() {
	flag.StringVar(&port, "port", "8080", "Port to be used by HTTP Server")
	flag.StringVar(&contentFile, "file", "/etc/float/float.yml", "Config file with the float page content")
	flag.Parse()
}

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "version":
			fmt.Printf("Float version: %s\n", buildString)
			os.Exit(0)
		}
	}

	serveHTTP()
}
