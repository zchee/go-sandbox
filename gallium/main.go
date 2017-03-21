package main

import (
	"os"
	"runtime"

	"github.com/alexflint/gallium"
)

func main() {
	runtime.LockOSThread()         // must be the first statement in main - see below
	gallium.Loop(os.Args, onReady) // must be called from main function
}

func onReady(app *gallium.App) {
	app.NewWindow("https://godoc.org/", "godoc.org")
}
