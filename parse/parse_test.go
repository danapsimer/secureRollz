package parse_test

import (
	"github.com/danapsimer/secureRollz"
	"github.com/danapsimer/secureRollz/parse"
	"github.com/danapsimer/secureRollz/rolltest"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func testParse(t *testing.T, expr, expected string, numResults int, min, max secureRollz.RollValue, mean, stddev float64) {
	t.Run(expected, func(t *testing.T) {
		roller, err := parse.ParseReader("string", strings.NewReader(expr))
		if assert.NoError(t, err) {
			if assert.NotNil(t, roller) {
				if dieRoller, ok := roller.(secureRollz.Roller); ok {
					assert.Equal(t, expected, dieRoller.String())
					rolltest.RollerTestStats(t, dieRoller, numResults, min, max, mean, stddev, expected, false)
				}
			}
		}
	})
}

func TestParse(t *testing.T) {
	testParse(t, "3 + 1", "3+1", 2, 4, 4, 4, 0)
	testParse(t, "3 - 1", "3-1", 2, 2, 2, 2, 0)
	testParse(t, "3 * 4", "3*4", 2, 12, 12, 12, 0)
	testParse(t, "3 * 4 + 5", "3*4+5", 2, 17, 17, 17, 0)
	testParse(t, "3 * (4 + 5)", "3*(4+5)", 2, 27, 27, 27, 0)
	testParse(t, "4 / 2", "4/2", 2, 2, 2, 2, 0)
	testParse(t, "d6", "d6", 0, 1, 6, 3.5, 1.7)
	testParse(t, "d6 + 7", "d6+7", 2, 8, 14, 10.5, 1.7)
	testParse(t, "d6 - 1", "d6-1", 2, 0, 5, 2.5, 1.7)
	testParse(t, "d6 * 2", "d6*2", 2, 2, 12, 7, 3.4)
	testParse(t, "d6 / 2", "d6/2", 2, 0, 3, 1.5, 0.9)
	testParse(t, "d6 + d12", "d6+d12", 2, 2, 18, 10.5, 3.8)
	testParse(t, "d20 - d4", "d20-d4", 2, -3, 19, 8, 5.8)
	testParse(t, "d6 * d8", "d6*d8", 2, 1, 48, 16, 11.7)
	testParse(t, "d6 * d6 + d8 - 5", "d6*d6+d8-5", 2, -3, 39, 12, 9.2)
	testParse(t, "d20 / d4", "d20/d4", 2, 0, 20, 5.0, 4.7)
	testParse(t, "3d6", "3d6", 3, 3, 18, 10, 3)
	testParse(t, "3d6 + 1", "3d6+1", 2, 4, 19, 11, 3)
	testParse(t, "3d6 * 2", "3d6*2", 2, 6, 36, 21, 5.9)
	testParse(t, "2d20D1", "2d20D1", 1, 1, 20, 13, 4.7) // Advantage
	testParse(t, "2d20D>1", "2d20D>1", 1, 1, 20, 7, 4.7) // Disadvantage
	testParse(t, "4d6D1", "4d6D1", 3, 3, 18, 12, 2.8)
	testParse(t, "4d6D<1", "4d6D1", 3, 3, 18, 12, 2.8)
	testParse(t, "4d6D<2", "4d6D2", 2, 2, 12, 9, 2)
	testParse(t, "4d6D>1", "4d6D>1", 3, 3, 18, 9, 2.8)
	testParse(t, "4d6r<2", "4d6r<2", 4, 4, 24, 16, 3.2)
	testParse(t, "4d6r>5", "4d6r>5", 4, 4, 24, 12, 3.2)
	testParse(t, "4d6D1r<2", "4d6D1r<2", 3, 3, 18, 12, 2.5)
	testParse(t, "4d6D1 r < 2", "4d6D1r<2", 3, 3, 18, 13, 2.5)
	testParse(t, "(4d6r<2)D1", "(4d6r<2)D1", 3, 3, 18, 13, 2.5)
	testParse(t, "( 4 d6 r < 2 ) D 1", "(4d6r<2)D1", 3, 3, 18, 13, 2.5)
}
