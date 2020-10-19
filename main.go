package main

import (
	"goinreverse/parsers"
	"log"
)

var config = parsers.ParseConfig()

func main() {
	log.Print(config.Servers)
}
