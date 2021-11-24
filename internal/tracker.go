package internal

import (
	"fmt"
	"github.com/ev3go/ev3dev"
	_ "github.com/ev3go/ev3dev"
	"github.com/sirupsen/logrus"
	"gobot.io/x/gobot"
	_ "gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/joystick"
	"log"
	"time"
)

type Tracker struct {
	// TODO: nullptr
	robot gobot.Robot

	outA *ev3dev.TachoMotor
}

func (t *Tracker) Open () {
	logrus.Debug("Opening...")

	// Stolen here: https://github.com/ev3go/ev3dev/blob/master/examples/demo/demo.go
	//
	// Get the handle for the medium motor on outA.
	outA, err := ev3dev.TachoMotorFor("ev3-ports:outA", "lego-ev3-m-motor")
	if err != nil {
		log.Fatalf("failed to find medium motor on outA: %v", err)
	}
	err = outA.SetStopAction("brake").Err()
	if err != nil {
		log.Fatalf("failed to set brake stop for medium motor on outA: %v", err)
	}
}

func (t *Tracker) Run () {

	logrus.Debug("Starting...")

	joystickAdaptor := joystick.NewAdaptor()
	joystick := joystick.NewDriver(joystickAdaptor,
		//"./platforms/joystick/configs/dualshock3.json",
		"../cmd/dualshock3.json",
	)

	work := func() {
		//joystick.On(joystick.Event("square_press"), func(data interface{}) {
		//	fmt.Println("square_press")
		//})
		//joystick.On(joystick.Event("square_release"), func(data interface{}) {
		//	fmt.Println("square_release")
		//})
		//joystick.On(joystick.Event("triangle_press"), func(data interface{}) {
		//	fmt.Println("triangle_press")
		//})
		//joystick.On(joystick.Event("triangle_release"), func(data interface{}) {
		//	fmt.Println("triangle_release")
		//})
		//joystick.On(joystick.Event("left_x"), func(data interface{}) {
		//	fmt.Println("left_x", data)
		//})
		//joystick.On(joystick.Event("left_y"), func(data interface{}) {
		//	fmt.Println("left_y", data)
		//})
		joystick.On(joystick.Event("right_x"), t.handleStickAction )
		//joystick.On(joystick.Event("right_y"), func(data interface{}) {
		//	fmt.Println("right_y", data)
		//})
	}

	robot := gobot.NewRobot("joystickBot",
		[]gobot.Connection{joystickAdaptor},
		[]gobot.Device{joystick},
		work,
	)

	err := robot.Start()

	if err != nil {
		logrus.Error("Error occured: ", err)
	}

	logrus.Debug("Stoping...")
}

// STRATEGY
// TODO: Make a unit test for this strategy.
func (t *Tracker) handleStickAction(data interface{}) {
	fmt.Println("right_x", data)

	input, ok := data.(int)
	if !ok {
		logrus.Errorf("Invalid input(%v)", data)
	}

	// max is 32000
	// example: 15.5k / 8k = 1
	// minimum motor speed quantum is 25
	speed := (input / 8000) * 25

	//t.outA.SetSpeedSetpoint(50 * maxMedium / 100).Command("run-forever")
	t.outA.SetSpeedSetpoint(speed).Command("run-forever")

	time.Sleep(time.Second / 2)
	t.outA.Command("stop")

	checkErrors(t.outA)
}

func checkErrors(devs ...ev3dev.Device) {
	for _, d := range devs {
		err := d.(*ev3dev.TachoMotor).Err()
		if err != nil {
			drv, dErr := ev3dev.DriverFor(d)
			if dErr != nil {
				drv = fmt.Sprintf("(missing driver name: %v)", dErr)
			}
			addr, aErr := ev3dev.AddressOf(d)
			if aErr != nil {
				drv = fmt.Sprintf("(missing port address: %v)", aErr)
			}
			log.Fatalf("motor error for %s:%s on port %s: %v", d, drv, addr, err)
		}
	}
}