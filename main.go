package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/4726/kubernetes-example/app"
	"github.com/4726/kubernetes-example/config"
)

const defaultConfigPath = "config.yml"

var configPath string

var usage = `
Usage: kubernetes-example <command> [arguments]

Commands:
	-c, -config <file path> path to config file
	-h, -help	prints usage string
`

func init() {
	flag.StringVar(&configPath, "config", defaultConfigPath, "path to config file")
	flag.StringVar(&configPath, "c", defaultConfigPath, "path to config file")
	flag.Usage = func() {
		fmt.Printf("%s\n", usage)
		os.Exit(0)
	}
}

func main() {
	flag.Parse()
	conf, err := config.FromFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	a := app.New(conf)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(c)

	go func() {
		<-c
		a.Close()
		os.Exit(1)
	}()

	a.Run()
}
