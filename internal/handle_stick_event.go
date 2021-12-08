package internal

import (
	"time"

	"github.com/ev3go/ev3dev"
	"github.com/sirupsen/logrus"
)

func handleStickEvent(motor *ev3dev.TachoMotor, data interface{}) {

	ratio, err := GetStickPosRatio(data)
	if err != nil {
		logrus.Errorf("Computation failed, err(%v)", err)
	}

	energy := int(ratio * motorAbsMaxSpeed)
	motor.SetSpeedSetpoint(energy).Command("run-forever")

	time.Sleep(time.Second / 2)
	motor.Command("stop")

	checkDevicesErrors(motor)
}
