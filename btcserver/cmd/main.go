package main

import (
	"flag"
	"github.com/generativelabs/btcserver/cmd/server"
)

func main() {

	migrateFlag := flag.Bool("migrate", false, "migrate")
	flag.Parse()

	//log.Info().Msgf("👷 Start migrate : %+v", *migrateFlag)

	server.Run(*migrateFlag)
}
