package secureRollz

import "fmt"

type constRoll struct {
	baseRoll
	value RollValue
}

func (cr constRoll) String() string {
	return cr.Source().String()
}

func (cr constRoll) Results() []Roll {
	return []Roll{}
}

func (cr constRoll) Total() RollValue {
	return cr.value
}

type constRoller struct {
	value RollValue
}

func (cr constRoller) String() string {
	return fmt.Sprintf("%d", cr.value)
}

func (cr constRoller) Roll() Roll {
	return constRoll{baseRoll{cr}, cr.value}
}

func Const(value RollValue) Roller {
	return constRoller{value}
}



