package app

import (
	"fmt"

	"github.com/ev3go/ev3dev"
	"github.com/sirupsen/logrus"
)

func checkDevicesErrors(devs ...ev3dev.Device) {

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
