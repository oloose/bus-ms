package main

import (
	"os"
	"time"

	busServer "github.com/oloose/bus-ms/internal/server"

	"net/http"

	"github.com/urfave/cli"
	"log"
)

var busEnv *BusEnv

type BusEnv struct {
	filePath string
}

func main() {
	busEnv = &BusEnv{}

	/*
	 *	Setup CLI
	 */
	app := cli.NewApp()
	app.Name = "Bus Microservice"
	app.Usage = "Start and manage the bus micro service"
	app.Version = "1.0.0"
	app.Compiled = time.Now()
	// define flags/parameters
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "bus-file, f",
			Value:       "busplan.json",
			Usage:       "Set file url to bus plan",
			Destination: &busEnv.filePath,
		},
	}
	// define commands
	app.Commands = []cli.Command{
		{
			Name:   "start",
			Usage:  "Starts the bus micro service",
			Action: StartBusServer,
		},
	}

	// run cli
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Error during start up: '%s'", err)
	}
}

func StartBusServer(c *cli.Context) error {
	//check if file with path busEnv.filePath exists on machine
	_, err := os.Stat(busEnv.filePath)
	if err != nil {
		// check if file exists on url path
		response, err := http.Head(busEnv.filePath)
		if err != nil || response.StatusCode != http.StatusOK {
			return err
		}
	}

	server := busServer.NewServer(busEnv.filePath)
	server.Start()

	return nil
}
