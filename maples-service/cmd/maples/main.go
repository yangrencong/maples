// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: cda8b667e3
// Version Date: 2021-09-27T00:46:54Z

package main

import (
	"flag"

	// This Service
	"maples/maples-service/handlers"
	"maples/maples-service/svc/server"
)

func main() {
	// Update addresses if they have been overwritten by flags
	flag.Parse()

	cfg := server.DefaultConfig
	cfg = handlers.SetConfig(cfg)

	server.Run(cfg)
}