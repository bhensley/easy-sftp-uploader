package main

import (
	"fmt"

	"github.com/bhensley/easy-sftp-uploader/config"
)

func main() {
	fmt.Println("Starting App")
	config, _ := config.LoadConfiguration("config.json")
	fmt.Printf("Config: {0}", config.Hosts)
}
