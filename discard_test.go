package secureRollz_test

import (
	"github.com/danapsimer/secureRollz"
	"github.com/danapsimer/secureRollz/rolltest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiscardRoller(t *testing.T) {
	roller := secureRollz.DiscardRoller(1, true, secureRollz.MultiRoller(4, secureRollz.DieRoller(6)))
	population := rolltest.RollerTest(t, roller, 3, 3, 18, "4d6D1")
	mean, err := population.Mean()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 12, mean, 1.0)
	}
	stddev, err := population.StandardDeviation()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 2.8, stddev, 0.1)
	}
}

func TestDiscardRollerHighest(t *testing.T) {
	roller := secureRollz.DiscardRoller(1, false, secureRollz.MultiRoller(4, secureRollz.DieRoller(6)))
	population := rolltest.RollerTest(t, roller, 3, 3, 18, "4d6D>1")
	mean, err := population.Mean()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 9, mean, 1.0)
	}
	stddev, err := population.StandardDeviation()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 2.8, stddev, 0.1)
	}
}

func TestDiscardRollerCombined(t *testing.T) {
	roller := secureRollz.DiscardRoller(1, true,
		secureRollz.DiscardRoller(1, false,
			secureRollz.MultiRoller(5, secureRollz.DieRoller(6))))
	population := rolltest.RollerTest(t, roller, 3, 3, 18, "5d6D>1D1")
	mean, err := population.Mean()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 10, mean, 1.0)
	}
	stddev, err := population.StandardDeviation()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 2.8, stddev, 0.2)
	}
}

func BenchmarkDiscardRoller(b *testing.B) {
	roller := secureRollz.DiscardRoller(1, false, secureRollz.MultiRoller(4, secureRollz.DieRoller(6)))
	rolltest.RollerBenchmark(b, roller)
}
