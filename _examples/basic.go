package main

import (
	tl "github.com/mhmorgan/termlog"
)

func main() {
	tl.Info("Hello, world!")
	tl.Emph("This is a test")
	tl.SetTimestampEnabled(true)
	tl.Warn("This is a warning")
	tl.Errf("Something went %v!", "wrong")
}
