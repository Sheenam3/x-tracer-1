package main

import (
	"flag"
	"github.com/Sheenam3/x-tracer-1/database"
	"github.com/Sheenam3/x-tracer-1/events"
	"github.com/Sheenam3/x-tracer-1/pkg"
	"github.com/Sheenam3/x-tracer-1/ui"
)

func main() {

	database.Init()

	ui.SubscribeListeners()
	pkg.SubscribeListeners()

	go events.Run()

	port := flag.String("port", "6666", "")
	pkg.SetPort(*port)
	go pkg.StartServer()

	ui.InitGui()

}