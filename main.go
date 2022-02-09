package main

import (
	"gosupervisor/configuration"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	configurationfile = kingpin.Flag("conf", "path to configuration file").Short('c').ExistingFile()
)

func main() {
	kingpin.Parse()
	configuration.LoadConfiguration(*configurationfile)

}
