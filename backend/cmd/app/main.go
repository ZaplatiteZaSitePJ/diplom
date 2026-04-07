package main

import (
	"flag"
	"inno-accounting/internal/api"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	stringPath string
)

func init() {
	flag.StringVar(&stringPath, "path", "configs/api.toml", "dev congig path")
}

func main() {
	flag.Parse()
	config := api.NewConfig()
	if _, err := toml.DecodeFile(stringPath, config); err != nil {
		log.Print("Error via decoding .toml config, using default ", err)
	}

	server := api.New(config)

	err := server.Start(); if err != nil{
		log.Fatal(err)
	}

	
}