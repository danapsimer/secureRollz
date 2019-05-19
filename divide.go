package secureRollz

import "fmt"

type divideRoll struct {
	baseRoll
	result Roll
	total  RollValue
}

func (dr divideRoll) Results() []Roll {
	return []Roll{dr.result}
}

func (dr divideRoll) Total() RollValue {
	return dr.total
}

type divideRoller struct {
	divisor RollValue
	roller  Roller
}

func (dr divideRoller) String() string {
	return fmt.Sprintf("%s/%d", dr.roller.String(), dr.divisor)
}

func (dr divideRoller) Roll() Roll {
	roll := dr.roller.Roll()
	total := roll.Total() / dr.divisor
	return divideRoll{baseRoll{dr}, roll, total}
}

func DivideRoller(divisor RollValue, roller Roller) Roller {
	return divideRoller{divisor, roller}
}
