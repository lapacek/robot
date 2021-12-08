package internal

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// stores the value of the approximated maximum/minimum stick position
const limitPos float32 = 32000.0

// provides the ratio of the stick position to the stick limit position, where
// stick position is floating point number included in the [-32768.0, 32767.0]
func GetStickPosRatio(input interface{}) (float32, error) {

	stickPos, ok := input.(float32)
	if !ok {
		msg := fmt.Sprintf("Invalid input(%v)", input)
		logrus.Errorf(msg)

		return 0, fmt.Errorf(msg)
	}

	return stickPos / limitPos, nil
}