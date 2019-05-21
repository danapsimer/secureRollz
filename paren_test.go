package secureRollz_test

import (
	"github.com/danapsimer/secureRollz"
	"github.com/danapsimer/secureRollz/rolltest"
	"testing"
)

func TestParenRoller(t *testing.T) {
	roller := secureRollz.ParenRoller(secureRollz.DieRoller(20))
	rolltest.RollerTestStats(t, roller, 0, 1, 20, 10.5, 5.7, "(d20)", true)
}
