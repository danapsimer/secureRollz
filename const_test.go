package secureRollz_test

import (
	"github.com/danapsimer/secureRollz"
	"github.com/danapsimer/secureRollz/rolltest"
	"testing"
)

func TestConst(t *testing.T) {
	roller := secureRollz.Const(4)
	rolltest.RollerTest(t, roller, 0, 4, 4, "4")
}

func BenchmarkConst(b *testing.B) {
	rolltest.RollerBenchmark(b, secureRollz.Const(4))
}
