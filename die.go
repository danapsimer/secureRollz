package secureRollz

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type dieRoll struct {
	baseRoll
	value RollValue
}

func (dr dieRoll) String() string {
	return fmt.Sprintf("%s{%d}",dr.Source().String(), dr.value)
}

func (dr dieRoll) Results() []Roll {
	return []Roll{}
}

func (dr dieRoll) Total() RollValue {
	return dr.value
}

type dieRoller struct {
	sides int
}

func (dr dieRoller) String() string {
	return fmt.Sprintf("d%d", dr.sides)
}

func (dr dieRoller) Roll() Roll {
	roll, _ := rand.Int(rand.Reader, big.NewInt(int64(dr.sides)))
	return dieRoll{baseRoll{dr},RollValue(roll.Int64()+1)}
}

func DieRoller(sides int) Roller {
	return dieRoller{ sides }
}
