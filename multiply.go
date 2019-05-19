package secureRollz

import "fmt"

type multiplyRoll struct {
	baseRoll
	roll  Roll
	total RollValue
}

func (mr multiplyRoll) Results() []Roll {
	return []Roll{mr.roll}
}

func (mr multiplyRoll) Total() RollValue {
	return mr.total
}

type multiplyRoller struct {
	multiplier RollValue
	roller     Roller
}

func (mr multiplyRoller) String() string {
	return fmt.Sprintf("%s*%d", mr.roller.String(), mr.multiplier)
}

func (mr multiplyRoller) Roll() Roll {
	roll := mr.roller.Roll()
	return multiplyRoll{baseRoll{mr}, roll, roll.Total() * mr.multiplier}
}

func MultiplyRoller(m RollValue, roller Roller) Roller {
	return multiplyRoller{m, roller}
}
