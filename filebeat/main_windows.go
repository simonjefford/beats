// +build windows

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/elastic/beats/filebeat/beater"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/service"
)

var Name = "filebeat"

// The basic model of execution:
// - prospector: finds files in paths/globs to harvest, starts harvesters
// - harvester: reads a file, sends events to the spooler
// - spooler: buffers events until ready to flush to the publisher
// - publisher: writes to the network, notifies registrar
// - registrar: records positions of files read
// Finally, prospector uses the registrar information, on restart, to
// determine where in each file to restart a harvester.

func main() {
	var err error
	if len(os.Args) > 1 {
		cmd := strings.ToLower(os.Args[1])
		switch cmd {
		case "install":
			err = service.InstallWindowsService(Name, "filebeat log shipper")
			if err != nil {
				fmt.Printf("%v\n", err)
				os.Exit(1)
			}
			fmt.Println("installed")
			os.Exit(0)
		case "uninstall":
			err = service.UninstallWindowsService(Name)
			if err != nil {
				fmt.Printf("%v\n", err)
				os.Exit(1)
			}
			fmt.Println("uninstalled")
			os.Exit(0)
		}
	}

	if err := beat.Run(Name, "", beater.New()); err != nil {
		os.Exit(1)
	}
}
