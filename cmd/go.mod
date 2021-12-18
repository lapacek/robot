module github.com/lapacek/go-ev3-dualshock-3/cmd

require (
	github.com/ev3go/ev3dev v0.0.0-20210313113244-a5fda5c6a492 // indirect
	github.com/lapacek/go-ev3-dualshock-3/internal v0.0.0
	github.com/sirupsen/logrus v1.8.1 // indirect
	gobot.io/x/gobot v1.15.0 // indirect
)

replace github.com/lapacek/go-ev3-dualshock-3/internal => ../internal

go 1.16
