package app

import (
	"testing"
)

func TestJoystickZeroPos(t *testing.T) {

	ratio, err := StickPosRatio(-129.1)
	if err != nil {
		t.Errorf("Computation failed, err(%v)", err)
	}

	var expected float64 = 0.0

	if ratio != expected {
		t.Errorf("Computed stick pos ratio(%v), expected pos ratio(%v)", ratio, expected)
	}
}
