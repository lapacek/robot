package main

import (
	"go-ev3-dualshock-3/internal"
)

func main() {

	robot := internal.Tracker{}
	robot.Open()
	robot.Run()
}
