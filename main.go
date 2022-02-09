package main

import (
	"gosupervisor/configuration"
	"gosupervisor/supervisor"
	"log"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	configurationfile = kingpin.Flag("conf", "path to configuration file").Short('c').ExistingFile()
)

func main() {
	kingpin.Parse()
	parseConfig, err := configuration.LoadConfiguration(*configurationfile)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	supervisor := supervisor.NewSupervisor(parseConfig)

	supervisor.HttpServer.Run()

}
