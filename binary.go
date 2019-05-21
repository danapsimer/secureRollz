package secureRollz

import (
	"fmt"
	"log"
)

type BinaryOp int

const (
	Unknown BinaryOp = iota
	Add
	Subtract
	Multiply
	Divide
)

func (bo BinaryOp) String() string {
	sign := ""
	switch bo {
	case Add:
		sign = "+"
	case Subtract:
		sign = "-"
	case Multiply:
		sign = "*"
	case Divide:
		sign = "/"
	}
	return sign
}

type binaryRoll struct {
	baseRoll
	left, right Roll
	total       RollValue
}

func (mr binaryRoll) Results() []Roll {
	return []Roll{mr.left, mr.right}
}

func (mr binaryRoll) Total() RollValue {
	return mr.total
}

func (mr binaryRoll) String() string {
	return fmt.Sprintf("%s{total:%d,left:%s,right:%s}",mr.Source().String(),mr.total,mr.left,mr.right)
}

type binaryRoller struct {
	op          BinaryOp
	left, right Roller
}

func (mr binaryRoller) String() string {
	return fmt.Sprintf("%s%s%s", mr.left.String(), mr.op, mr.right.String())
}

func (mr binaryRoller) Roll() Roll {
	left := mr.left.Roll()
	right := mr.right.Roll()
	switch mr.op {
	case Add:
		return binaryRoll{baseRoll{mr}, left, right, left.Total() + right.Total()}
	case Subtract:
		return binaryRoll{baseRoll{mr}, left, right, left.Total() - right.Total()}
	case Multiply:
		return binaryRoll{baseRoll{mr}, left, right, left.Total() * right.Total()}
	case Divide:
		return binaryRoll{baseRoll{mr}, left, right, left.Total() / right.Total()}
	}
	log.Panicf("unknown binary operator: %d", mr.op)
	return nil
}

func BinaryOpRoller(op BinaryOp, left, right Roller) Roller {
	if op == Unknown {
		panic("binary op should be set.")
	}
	return binaryRoller{op, left, right}
}
