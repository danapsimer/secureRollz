package secureRollz_test

import (
	"bytes"
	"fmt"
	"github.com/montanaflynn/stats"
	"github.com/stretchr/testify/assert"
	"io"
	"math"
	"secureRollz"
	"testing"
)

func testRoller(t *testing.T, roller secureRollz.Roller, numResults int, min, max secureRollz.RollValue, source string) stats.Float64Data {
	samples := int(max-min+1)*1000
	population := stats.Float64Data(make([]float64,samples))
	histo := make([]int, max-min+1)
	if assert.NotNil(t, roller) {
		for i := 0; i < samples; i++ {
			roll := roller.Roll()
			if assert.NotNil(t, roll) {
				assert.Equal(t, source, roll.Source().String())
				assert.Equal(t, numResults, len(roll.Results()))
				total := roll.Total()
				if assert.Conditionf(t, func() (bool) {
					return min <= total && total <= max
				}, "Roller returned a number out of range: expected a number between %d, and %d inclusive but got %d", min, max, total) {
					population[i] = float64(total)
					histo[total-min] += 1
				}
			}
		}
	}
	printHistogram(t, min, max, histo, 80)
	return population
}

var fractionalBlocks = []rune(" \u258F\u258E\u258D\u258C\u258B\u258A\u2589\u2588")
func printHistogram(t *testing.T, minV, maxV secureRollz.RollValue, histo []int, maxWidth int) {
	w := bytes.NewBuffer(make([]byte, 0, 512))
	var max int
	for _, h := range histo {
		if h > max {
			max = h
		}
	}
	scale := float64(maxWidth) / float64(max)
	for i := range histo {
		io.WriteString(w, fmt.Sprintf(" %-8d: ", minV + secureRollz.RollValue(i)))
		blocksCalc := scale * float64(histo[i])
		blocks := int(math.Floor(blocksCalc))
		blocksFractionValue := int(math.Ceil((blocksCalc - float64(blocks)) / 0.125))
		for i = 0; i < blocks; i++ {
			io.WriteString(w, "\u2588")
		}
		io.WriteString(w, string(fractionalBlocks[blocksFractionValue]))
		io.WriteString(w, "\n")
	}
	t.Logf("\n%s", w.String())
}

var result secureRollz.Roll


func rollerBenchmark(b *testing.B, roller secureRollz.Roller) {
	var roll secureRollz.Roll
	for n := 0; n < b.N; n++ {
		roll = roller.Roll()
	}
	result = roll
}
