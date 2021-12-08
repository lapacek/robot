package main

import (
	"github.com/lapacek/go-ev3-dualshock-3/internal"
)

func main() {

	robot := internal.NewTracker("Tracker")
	robot.Run()
}
