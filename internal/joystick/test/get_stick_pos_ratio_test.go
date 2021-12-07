package joystick_test

import (
	"testing"

	"github.com/lapacek/go-ev3-dualshock-3/internal/joystick"
)

func Test_joystick_get_stick_pos_ratio_tends_to_zero(t *testing.T) {

	ratio, err := joystick.GetStickPosRatio(-129.1)
	if err != nil {
		t.Errorf("computation failed, err(%v)", err)
	}

	var expected float32 = 0.0

	if ratio != expected {
		t.Errorf("computed speed(%v), expected speed(%v)", ratio, expected)
	}
}
