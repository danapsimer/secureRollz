package secureRollz

import "fmt"

type modifiedRoll struct {
	baseRoll
	roll     Roll
	total    RollValue
}

func (mr modifiedRoll) Results() []Roll {
	return []Roll{mr.roll}
}

func (mr modifiedRoll) Total() RollValue {
	return mr.total
}

type modifiedRoller struct {
	modifier RollValue
	roller Roller
}

func (mr modifiedRoller) String() string {
	sign := "+"
	m := mr.modifier
	if mr.modifier < 0 {
		sign = "-"
		m = -mr.modifier
	}
	return fmt.Sprintf("%s%s%d", mr.roller.String(), sign, m)
}

func (mr modifiedRoller) Roll() Roll {
	roll := mr.roller.Roll()
	return modifiedRoll{baseRoll{mr}, roll, roll.Total() + mr.modifier}
}

func ModifiedRoller(m RollValue, roller Roller) Roller {
	return modifiedRoller{m, roller}
}
