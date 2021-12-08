package internal_test

import (
	"testing"

	"github.com/lapacek/go-ev3-dualshock-3/internal"
)

func Test_joystick_get_stick_pos_ratio_tends_to_zero(t *testing.T) {

	ratio, err := internal.GetStickPosRatio(-129.1)
	if err != nil {
		t.Errorf("Computation failed, err(%v)", err)
	}

	var expected float32 = 0.0

	if ratio != expected {
		t.Errorf("Computed pos ratio(%v), expected pos ratio(%v)", ratio, expected)
	}
}
