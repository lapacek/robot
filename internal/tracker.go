package internal

import (
	"fmt"
	"time"

	"github.com/ev3go/ev3dev"
	"github.com/sirupsen/logrus"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/joystick"
)

type work_t func()

type Tracker struct {

	// TODO: nullptr
	joystick        *joystick.Driver
	joystickAdaptor *joystick.Adaptor
	robot           *gobot.Robot

	// the physical outputs
	outA *ev3dev.TachoMotor
	outB *ev3dev.TachoMotor

	work work_t
}

func (t *Tracker) open() bool {
  
	var err error
	logrus.Debug("Opening...")
	defer func () {
		if err != nil {
			logrus.Fatal("Component can`t be open.")
		}
		logrus.Debug("Opened")
	}()

	// stolen here: https://github.com/ev3go/ev3dev/blob/master/examples/demo/demo.go
	//
	// get the handle for the medium motor on outA.
	outA, err := ev3dev.TachoMotorFor("ev3-ports:outA", "lego-ev3-m-motor")
	if err != nil {
		logrus.Errorf("Failed to find medium motor on outA: %v", err)

		return false
	}

	err = outA.SetStopAction("brake").Err()
	if err != nil {
		logrus.Errorf("Failed to set brake stop for medium motor on outA: %v", err)

		return false
	}

	// get the handle for the left large motor on outB.
	outB, err := ev3dev.TachoMotorFor("ev3-ports:outB", "lego-ev3-l-motor")
	if err != nil {
		logrus.Errorf("Failed to find left large motor on outB: %v", err)

		return false
	}

	err = outB.SetStopAction("brake").Err()
	if err != nil {
		logrus.Errorf("Failed to set brake stop for left large motor on outB: %v", err)

		return false
	}

	t.joystickAdaptor = joystick.NewAdaptor()
	t.joystick = joystick.NewDriver(t.joystickAdaptor,
		"../config/dualshock3.json",
	)

	t.work = func() {

		logrus.Debug("Working...")
		defer logrus.Debug("Work stoped")

		t.joystick.On(t.joystick.Event("right_x"), t.handleStickAction)
		t.joystick.On(t.joystick.Event("right_y"), t.handleStickAction)
		t.joystick.On(t.joystick.Event("left_x"), t.handleStickAction)
		t.joystick.On(t.joystick.Event("left_y"), t.handleStickAction)
	}

	return true
}

func (t *Tracker) Run() {

	if !t.open() {
		logrus.Fatal("Component is not opened.")
	}

	logrus.Debug("Starting...")
	defer logrus.Debug("Stoped")

	t.robot = gobot.NewRobot("joystickBot",
		[]gobot.Connection{t.joystickAdaptor},
		[]gobot.Device{t.joystick},
		t.work,
	)
  
	err := t.robot.Start()
	if err != nil {
		logrus.Error("Error occured: ", err)
	}
}

func (t *Tracker) handleStickAction(data interface{}) {
	logrus.Debugf("right_x data(%v)", data)

	input, ok := data.(int)
	if !ok {
		logrus.Errorf("Invalid input(%v)", data)
	}

	// max is 32000
	// example: 15.5k / 8k = 1
	// minimum motor speed quantum is 25
	speed := (input / 8000) * 25

	t.outA.SetSpeedSetpoint(speed).Command("run-forever")

	time.Sleep(time.Second / 2)
	t.outA.Command("stop")

	checkErrors(t.outA)
}

func (t *Tracker) close() {

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

			logrus.Fatalf("motor error for %s:%s on port %s: %v", d, drv, addr, err)
		}
	}
}
