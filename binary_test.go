package secureRollz_test

import (
	"github.com/danapsimer/secureRollz"
	"github.com/danapsimer/secureRollz/rolltest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryOpRollerAdd(t *testing.T) {
	roller := secureRollz.BinaryOpRoller(secureRollz.Add, secureRollz.DieRoller(6), secureRollz.Const(7))
	population := rolltest.RollerTest(t, roller, 2, 8, 13, "d6+7")
	mean, err := population.Mean()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 10.5, mean, 1.0)
	}
	stddev, err := population.StandardDeviation()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 1.7, stddev, 0.1)
	}
}

func TestBinaryOpRollerSubtract(t *testing.T) {
	roller := secureRollz.BinaryOpRoller(secureRollz.Subtract, secureRollz.DieRoller(6), secureRollz.Const(3))
	population := rolltest.RollerTest(t, roller, 2, -2, 3, "d6-3")
	mean, err := population.Mean()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 0.5, mean, 1.0)
	}
	stddev, err := population.StandardDeviation()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 1.7, stddev, 0.1)
	}
}

func TestBinaryOpRollerMultiply(t *testing.T) {
	roller := secureRollz.BinaryOpRoller(secureRollz.Multiply, secureRollz.DieRoller(6), secureRollz.Const(2))
	population := rolltest.RollerTest(t, roller, 2, 2, 12, "d6*2")
	mean, err := population.Mean()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 6.5, mean, 1.0)
	}
	stddev, err := population.StandardDeviation()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 3.4, stddev, 0.1)
	}
}

func TestBinaryOpRollerDivide(t *testing.T) {
	roller := secureRollz.BinaryOpRoller(secureRollz.Divide, secureRollz.DieRoller(6), secureRollz.Const(2))
	population := rolltest.RollerTest(t, roller, 2, 0, 3, "d6/2")
	mean, err := population.Mean()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 1.5, mean, 1.0)
	}
	stddev, err := population.StandardDeviation()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 1, stddev, 0.1)
	}
}
func BenchmarkModifiedRoller(b *testing.B) {
	roller := secureRollz.BinaryOpRoller(secureRollz.Add, secureRollz.DieRoller(6), secureRollz.Const(7))
	rolltest.RollerBenchmark(b, roller)
}
