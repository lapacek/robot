package app

import (
	"github.com/ev3go/ev3dev"
	"github.com/sirupsen/logrus"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/joystick"
)

// represents an absolute value of an approximate max speed of a motor
const motorAbsMaxSpeed = 100

// ev3dev related constants
const (
	// output ports
	ev3OutAPortName = "ev3-ports:outA"
	ev3OutBPortName = "ev3-ports:outB"

	// devices
	ev3LargeMotorName = "lego-ev3-l-motor"

	// actions
	ev3BreakActionName = "break"
)

// dualshock3 controller driver related constants
//
// * you can see more here ../config/dualshock3.json
const (
	joystickRightYAxisName = "right_y"
	joystickLeftYAxisName  = "left_y"
)

type workT func()

type Manager struct {
	name string
	work workT

	joystick        *joystick.Driver
	joystickAdaptor *joystick.Adaptor
	robot           *gobot.Robot

	// ev3 physical devices
	outA *ev3dev.TachoMotor
	outB *ev3dev.TachoMotor
}

func NewManager(name string) *Manager {
	t := Manager{}
	t.name = name

	return &t
}

func (t *Manager) Run() {

	if !t.open() {
		logrus.Fatal("Component is not opened.")
	}

	logrus.Debug("Starting...")
	defer logrus.Debug("Stoped")

	t.robot = gobot.NewRobot(
		t.name,
		[]gobot.Connection{t.joystickAdaptor},
		[]gobot.Device{t.joystick},
		t.work,
	)

	err := t.robot.Start()
	if err != nil {
		logrus.Errorf("Error occured, err(%v)", err)
	}
}

func (t *Manager) open() bool {

	logrus.Debug("Opening...")

	if !t.initMotors() {
		return false
	}

	t.initJoystick()
	t.initWork()

	logrus.Debug("Opened")

	return true
}

func (t *Manager) initMotors() bool {

	outA, err := ev3dev.TachoMotorFor(ev3OutAPortName, ev3LargeMotorName)
	if err != nil {
		logrus.Errorf("Failed to find right large motor on outA: %v", err)

		return false
	}

	err = outA.SetStopAction(ev3BreakActionName).Err()
	if err != nil {
		logrus.Errorf("Failed to set brake stop for large motor on outA: %v", err)

		return false
	}

	outB, err := ev3dev.TachoMotorFor(ev3OutBPortName, ev3LargeMotorName)
	if err != nil {
		logrus.Errorf("Failed to find left large motor on outB: %v", err)

		return false
	}

	err = outB.SetStopAction(ev3BreakActionName).Err()
	if err != nil {
		logrus.Errorf("Failed to set brake stop for left large motor on outB: %v", err)

		return false
	}

	return true
}

func (t *Manager) initJoystick() {

	t.joystickAdaptor = joystick.NewAdaptor()
	t.joystick = joystick.NewDriver(t.joystickAdaptor,
		"../config/dualshock3.json",
	)
}

func (t *Manager) initWork() {

	t.work = func() {

		logrus.Debug("Working...")
		defer logrus.Debug("Work stopped.")

		var err error

		err = t.joystick.On(t.joystick.Event(joystickRightYAxisName), t.rightStickAction)
		if err != nil {
			logrus.Errorf("Failed to register a joystick right_y event handler, err(%v)", err)
		}

		err = t.joystick.On(t.joystick.Event(joystickLeftYAxisName), t.leftStickAction)
		if err != nil {
			logrus.Errorf("Failed to register a joystick right_y event handler, err(%v)", err)
		}
	}
}

func (t *Manager) rightStickAction(data interface{}) {

	logrus.Tracef("Joystick event received, right y(%v)", data)

	onStickEvent(t.outB, data)
}

func (t *Manager) leftStickAction(data interface{}) {

	logrus.Tracef("Joystick event received, left y(%v)", data)

	onStickEvent(t.outA, data)
}

func (t *Manager) close() {

	err := t.robot.Stop()
	if err != nil {
		logrus.Errorf("Failed to stop a robot, err(%v)", err)
	}
}
