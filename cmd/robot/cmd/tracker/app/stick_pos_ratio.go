package app

import (
	"fmt"
	"math"

	"github.com/sirupsen/logrus"
)

// stores the value of the approximated maximum/minimum stick position
const limitPos float64 = 32000.0

// StickPosRatio provides the ratio of the stick position to the stick limit position, where
// stick position is floating point number included in the [-32768.0, 32767.0]
func StickPosRatio(input interface{}) (float64, error) {

	stickPos, ok := input.(float64)
	if !ok {
		msg := fmt.Sprintf("Invalid input(%v)", input)
		logrus.Errorf(msg)

		return 0, fmt.Errorf(msg)
	}

	r := stickPos / limitPos

	return math.Round(r*100) / 100, nil
}
