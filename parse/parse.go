package parse

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/danapsimer/secureRollz"
)

type term struct {
	op    secureRollz.BinaryOp
	right secureRollz.Roller
}

func (c *current) Errorf(format string, args ...interface{}) error {
	return errors.New(fmt.Sprintf("%s: %s", c.pos.String(), fmt.Sprintf(format, args...)))
}

var g = &grammar{
	rules: []*rule{
		{
			name: "Input",
			pos:  position{line: 19, col: 1, offset: 358},
			expr: &actionExpr{
				pos: position{line: 19, col: 10, offset: 367},
				run: (*parser).callonInput1,
				expr: &seqExpr{
					pos: position{line: 19, col: 10, offset: 367},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 19, col: 10, offset: 367},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 19, col: 12, offset: 369},
							label: "roll",
							expr: &ruleRefExpr{
								pos:  position{line: 19, col: 17, offset: 374},
								name: "Roll",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 19, col: 22, offset: 379},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 19, col: 24, offset: 381},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "Roll",
			pos:  position{line: 23, col: 1, offset: 411},
			expr: &actionExpr{
				pos: position{line: 23, col: 9, offset: 419},
				run: (*parser).callonRoll1,
				expr: &labeledExpr{
					pos:   position{line: 23, col: 9, offset: 419},
					label: "roll",
					expr: &ruleRefExpr{
						pos:  position{line: 23, col: 14, offset: 424},
						name: "Sum",
					},
				},
			},
		},
		{
			name: "Sum",
			pos:  position{line: 27, col: 1, offset: 454},
			expr: &actionExpr{
				pos: position{line: 27, col: 8, offset: 461},
				run: (*parser).callonSum1,
				expr: &seqExpr{
					pos: position{line: 27, col: 8, offset: 461},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 27, col: 8, offset: 461},
							label: "left",
							expr: &ruleRefExpr{
								pos:  position{line: 27, col: 13, offset: 466},
								name: "Product",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 27, col: 21, offset: 474},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 27, col: 23, offset: 476},
							label: "r",
							expr: &zeroOrMoreExpr{
								pos: position{line: 27, col: 25, offset: 478},
								expr: &ruleRefExpr{
									pos:  position{line: 27, col: 25, offset: 478},
									name: "SumTerm",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "SumTerm",
			pos:  position{line: 45, col: 1, offset: 974},
			expr: &actionExpr{
				pos: position{line: 45, col: 12, offset: 985},
				run: (*parser).callonSumTerm1,
				expr: &seqExpr{
					pos: position{line: 45, col: 12, offset: 985},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 45, col: 12, offset: 985},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 45, col: 15, offset: 988},
								name: "SumOp",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 45, col: 21, offset: 994},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 45, col: 23, offset: 996},
							label: "right",
							expr: &ruleRefExpr{
								pos:  position{line: 45, col: 29, offset: 1002},
								name: "Product",
							},
						},
					},
				},
			},
		},
		{
			name: "SumOp",
			pos:  position{line: 57, col: 1, offset: 1270},
			expr: &actionExpr{
				pos: position{line: 57, col: 10, offset: 1279},
				run: (*parser).callonSumOp1,
				expr: &labeledExpr{
					pos:   position{line: 57, col: 10, offset: 1279},
					label: "op",
					expr: &choiceExpr{
						pos: position{line: 57, col: 14, offset: 1283},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 57, col: 14, offset: 1283},
								val:        "+",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 57, col: 20, offset: 1289},
								val:        "-",
								ignoreCase: false,
							},
						},
					},
				},
			},
		},
		{
			name: "Product",
			pos:  position{line: 71, col: 1, offset: 1593},
			expr: &actionExpr{
				pos: position{line: 71, col: 12, offset: 1604},
				run: (*parser).callonProduct1,
				expr: &seqExpr{
					pos: position{line: 71, col: 12, offset: 1604},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 71, col: 12, offset: 1604},
							label: "left",
							expr: &ruleRefExpr{
								pos:  position{line: 71, col: 17, offset: 1609},
								name: "Factor",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 71, col: 24, offset: 1616},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 71, col: 26, offset: 1618},
							label: "r",
							expr: &zeroOrMoreExpr{
								pos: position{line: 71, col: 28, offset: 1620},
								expr: &ruleRefExpr{
									pos:  position{line: 71, col: 28, offset: 1620},
									name: "ProductTerm",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "ProductTerm",
			pos:  position{line: 89, col: 1, offset: 2114},
			expr: &actionExpr{
				pos: position{line: 89, col: 16, offset: 2129},
				run: (*parser).callonProductTerm1,
				expr: &seqExpr{
					pos: position{line: 89, col: 16, offset: 2129},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 89, col: 16, offset: 2129},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 89, col: 19, offset: 2132},
								name: "MulOp",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 89, col: 25, offset: 2138},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 89, col: 27, offset: 2140},
							label: "right",
							expr: &ruleRefExpr{
								pos:  position{line: 89, col: 33, offset: 2146},
								name: "Factor",
							},
						},
					},
				},
			},
		},
		{
			name: "MulOp",
			pos:  position{line: 101, col: 1, offset: 2436},
			expr: &actionExpr{
				pos: position{line: 101, col: 10, offset: 2445},
				run: (*parser).callonMulOp1,
				expr: &labeledExpr{
					pos:   position{line: 101, col: 10, offset: 2445},
					label: "op",
					expr: &choiceExpr{
						pos: position{line: 101, col: 14, offset: 2449},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 101, col: 14, offset: 2449},
								val:        "*",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 101, col: 20, offset: 2455},
								val:        "/",
								ignoreCase: false,
							},
						},
					},
				},
			},
		},
		{
			name: "Factor",
			pos:  position{line: 115, col: 1, offset: 2762},
			expr: &choiceExpr{
				pos: position{line: 115, col: 11, offset: 2772},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 115, col: 11, offset: 2772},
						name: "ReRoll",
					},
					&ruleRefExpr{
						pos:  position{line: 115, col: 20, offset: 2781},
						name: "ReRollable",
					},
				},
			},
		},
		{
			name: "ReRollable",
			pos:  position{line: 117, col: 1, offset: 2793},
			expr: &choiceExpr{
				pos: position{line: 117, col: 15, offset: 2807},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 117, col: 15, offset: 2807},
						name: "DiscardRoll",
					},
					&ruleRefExpr{
						pos:  position{line: 117, col: 29, offset: 2821},
						name: "Discardable",
					},
				},
			},
		},
		{
			name: "ReRoll",
			pos:  position{line: 119, col: 1, offset: 2834},
			expr: &actionExpr{
				pos: position{line: 119, col: 11, offset: 2844},
				run: (*parser).callonReRoll1,
				expr: &seqExpr{
					pos: position{line: 119, col: 11, offset: 2844},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 119, col: 11, offset: 2844},
							label: "roll",
							expr: &ruleRefExpr{
								pos:  position{line: 119, col: 16, offset: 2849},
								name: "ReRollable",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 119, col: 27, offset: 2860},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 119, col: 29, offset: 2862},
							val:        "r",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 119, col: 33, offset: 2866},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 119, col: 35, offset: 2868},
							label: "reRollN",
							expr: &zeroOrOneExpr{
								pos: position{line: 119, col: 43, offset: 2876},
								expr: &ruleRefExpr{
									pos:  position{line: 119, col: 43, offset: 2876},
									name: "Number",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 119, col: 51, offset: 2884},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 119, col: 53, offset: 2886},
							label: "dirOp",
							expr: &zeroOrOneExpr{
								pos: position{line: 119, col: 59, offset: 2892},
								expr: &ruleRefExpr{
									pos:  position{line: 119, col: 59, offset: 2892},
									name: "DirOp",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 119, col: 66, offset: 2899},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 119, col: 68, offset: 2901},
							label: "value",
							expr: &ruleRefExpr{
								pos:  position{line: 119, col: 74, offset: 2907},
								name: "Number",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 119, col: 81, offset: 2914},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Discardable",
			pos:  position{line: 147, col: 1, offset: 3585},
			expr: &choiceExpr{
				pos: position{line: 147, col: 16, offset: 3600},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 147, col: 16, offset: 3600},
						name: "MultiRoll",
					},
					&ruleRefExpr{
						pos:  position{line: 147, col: 28, offset: 3612},
						name: "Terminal",
					},
				},
			},
		},
		{
			name: "DiscardRoll",
			pos:  position{line: 149, col: 1, offset: 3622},
			expr: &actionExpr{
				pos: position{line: 149, col: 16, offset: 3637},
				run: (*parser).callonDiscardRoll1,
				expr: &seqExpr{
					pos: position{line: 149, col: 16, offset: 3637},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 149, col: 16, offset: 3637},
							label: "roll",
							expr: &ruleRefExpr{
								pos:  position{line: 149, col: 21, offset: 3642},
								name: "Discardable",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 149, col: 33, offset: 3654},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 149, col: 35, offset: 3656},
							val:        "D",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 149, col: 39, offset: 3660},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 149, col: 41, offset: 3662},
							label: "dirOp",
							expr: &zeroOrOneExpr{
								pos: position{line: 149, col: 47, offset: 3668},
								expr: &ruleRefExpr{
									pos:  position{line: 149, col: 47, offset: 3668},
									name: "DirOp",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 149, col: 54, offset: 3675},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 149, col: 56, offset: 3677},
							label: "n",
							expr: &ruleRefExpr{
								pos:  position{line: 149, col: 58, offset: 3679},
								name: "Number",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 149, col: 65, offset: 3686},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "DirOp",
			pos:  position{line: 169, col: 1, offset: 4136},
			expr: &actionExpr{
				pos: position{line: 169, col: 10, offset: 4145},
				run: (*parser).callonDirOp1,
				expr: &labeledExpr{
					pos:   position{line: 169, col: 10, offset: 4145},
					label: "op",
					expr: &choiceExpr{
						pos: position{line: 169, col: 14, offset: 4149},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 169, col: 14, offset: 4149},
								val:        "<",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 169, col: 20, offset: 4155},
								val:        ">",
								ignoreCase: false,
							},
						},
					},
				},
			},
		},
		{
			name: "MultiRoll",
			pos:  position{line: 183, col: 1, offset: 4439},
			expr: &actionExpr{
				pos: position{line: 183, col: 14, offset: 4452},
				run: (*parser).callonMultiRoll1,
				expr: &seqExpr{
					pos: position{line: 183, col: 14, offset: 4452},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 183, col: 14, offset: 4452},
							label: "n",
							expr: &ruleRefExpr{
								pos:  position{line: 183, col: 16, offset: 4454},
								name: "Number",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 183, col: 23, offset: 4461},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 183, col: 25, offset: 4463},
							label: "roll",
							expr: &ruleRefExpr{
								pos:  position{line: 183, col: 30, offset: 4468},
								name: "Terminal",
							},
						},
					},
				},
			},
		},
		{
			name: "Terminal",
			pos:  position{line: 195, col: 1, offset: 4731},
			expr: &choiceExpr{
				pos: position{line: 195, col: 13, offset: 4743},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 195, col: 13, offset: 4743},
						name: "DieRoll",
					},
					&ruleRefExpr{
						pos:  position{line: 195, col: 23, offset: 4753},
						name: "ConstRoll",
					},
					&ruleRefExpr{
						pos:  position{line: 195, col: 35, offset: 4765},
						name: "ParenRoll",
					},
				},
			},
		},
		{
			name: "ParenRoll",
			pos:  position{line: 197, col: 1, offset: 4776},
			expr: &actionExpr{
				pos: position{line: 197, col: 14, offset: 4789},
				run: (*parser).callonParenRoll1,
				expr: &seqExpr{
					pos: position{line: 197, col: 14, offset: 4789},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 197, col: 14, offset: 4789},
							val:        "(",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 197, col: 18, offset: 4793},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 197, col: 20, offset: 4795},
							label: "roll",
							expr: &ruleRefExpr{
								pos:  position{line: 197, col: 25, offset: 4800},
								name: "Roll",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 197, col: 30, offset: 4805},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 197, col: 32, offset: 4807},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "DieRoll",
			pos:  position{line: 205, col: 1, offset: 4980},
			expr: &actionExpr{
				pos: position{line: 205, col: 12, offset: 4991},
				run: (*parser).callonDieRoll1,
				expr: &seqExpr{
					pos: position{line: 205, col: 12, offset: 4991},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 205, col: 12, offset: 4991},
							val:        "d",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 205, col: 16, offset: 4995},
							label: "sides",
							expr: &ruleRefExpr{
								pos:  position{line: 205, col: 22, offset: 5001},
								name: "Number",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 205, col: 29, offset: 5008},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "ConstRoll",
			pos:  position{line: 213, col: 1, offset: 5165},
			expr: &actionExpr{
				pos: position{line: 213, col: 14, offset: 5178},
				run: (*parser).callonConstRoll1,
				expr: &seqExpr{
					pos: position{line: 213, col: 14, offset: 5178},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 213, col: 14, offset: 5178},
							label: "value",
							expr: &ruleRefExpr{
								pos:  position{line: 213, col: 20, offset: 5184},
								name: "Number",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 213, col: 27, offset: 5191},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Number",
			pos:  position{line: 221, col: 1, offset: 5367},
			expr: &actionExpr{
				pos: position{line: 221, col: 11, offset: 5377},
				run: (*parser).callonNumber1,
				expr: &oneOrMoreExpr{
					pos: position{line: 221, col: 12, offset: 5378},
					expr: &charClassMatcher{
						pos:        position{line: 221, col: 12, offset: 5378},
						val:        "[0-9]",
						ranges:     []rune{'0', '9'},
						ignoreCase: false,
						inverted:   false,
					},
				},
			},
		},
		{
			name:        "_",
			displayName: "\"whitespace\"",
			pos:         position{line: 230, col: 1, offset: 5564},
			expr: &zeroOrMoreExpr{
				pos: position{line: 230, col: 19, offset: 5582},
				expr: &charClassMatcher{
					pos:        position{line: 230, col: 19, offset: 5582},
					val:        "[ \\n\\t\\r]",
					chars:      []rune{' ', '\n', '\t', '\r'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 232, col: 1, offset: 5594},
			expr: &notExpr{
				pos: position{line: 232, col: 8, offset: 5601},
				expr: &anyMatcher{
					line: 232, col: 9, offset: 5602,
				},
			},
		},
	},
}

func (c *current) onInput1(roll interface{}) (interface{}, error) {
	return roll, nil
}

func (p *parser) callonInput1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInput1(stack["roll"])
}

func (c *current) onRoll1(roll interface{}) (interface{}, error) {
	return roll, nil
}

func (p *parser) callonRoll1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onRoll1(stack["roll"])
}

func (c *current) onSum1(left, r interface{}) (interface{}, error) {
	if r != nil {
		roller, ok := left.(secureRollz.Roller)
		if !ok {
			return nil, c.Errorf("left is not a Roller")
		}
		for _, trm := range r.([]interface{}) {
			term, ok := trm.(term)
			if !ok {
				return nil, c.Errorf("found element that is not a term")
			}
			roller = secureRollz.BinaryOpRoller(term.op, roller, term.right)
		}
		return roller, nil
	}
	return left, nil
}

func (p *parser) callonSum1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSum1(stack["left"], stack["r"])
}

func (c *current) onSumTerm1(op, right interface{}) (interface{}, error) {
	bop, ok := op.(secureRollz.BinaryOp)
	if !ok {
		return nil, c.Errorf("op is not a BinaryOp")
	}
	r, ok := right.(secureRollz.Roller)
	if !ok {
		return nil, c.Errorf("right is not a Roller")
	}
	return term{bop, r}, nil
}

func (p *parser) callonSumTerm1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSumTerm1(stack["op"], stack["right"])
}

func (c *current) onSumOp1(op interface{}) (interface{}, error) {
	opb, ok := op.([]byte)
	if !ok {
		return nil, c.Errorf("op is not a []byte")
	}
	switch opb[0] {
	case '+':
		return secureRollz.Add, nil
	case '-':
		return secureRollz.Subtract, nil
	}
	return nil, c.Errorf("unknown sum operator: %s", string(opb))
}

func (p *parser) callonSumOp1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSumOp1(stack["op"])
}

func (c *current) onProduct1(left, r interface{}) (interface{}, error) {
	if r != nil {
		roller, ok := left.(secureRollz.Roller)
		if !ok {
			return nil, c.Errorf("left is not a Roller")
		}
		for _, trm := range r.([]interface{}) {
			term, ok := trm.(term)
			if !ok {
				return nil, c.Errorf("element of r is not a term")
			}
			roller = secureRollz.BinaryOpRoller(term.op, roller, term.right)
		}
		return roller, nil
	}
	return left, nil
}

func (p *parser) callonProduct1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onProduct1(stack["left"], stack["r"])
}

func (c *current) onProductTerm1(op, right interface{}) (interface{}, error) {
	bop, ok := op.(secureRollz.BinaryOp)
	if !ok {
		return nil, c.Errorf("op is not a secureRollz.BinaryOp")
	}
	r, ok := right.(secureRollz.Roller)
	if !ok {
		return nil, c.Errorf("right is not a secureRollz.Roller")
	}
	return term{bop, r}, nil
}

func (p *parser) callonProductTerm1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onProductTerm1(stack["op"], stack["right"])
}

func (c *current) onMulOp1(op interface{}) (interface{}, error) {
	opb, ok := op.([]byte)
	if !ok {
		return nil, c.Errorf("op is not a []byte")
	}
	switch opb[0] {
	case '*':
		return secureRollz.Multiply, nil
	case '/':
		return secureRollz.Divide, nil
	}
	return nil, c.Errorf("unknown mul operator: %s", string(opb))
}

func (p *parser) callonMulOp1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMulOp1(stack["op"])
}

func (c *current) onReRoll1(roll, reRollN, dirOp, value interface{}) (interface{}, error) {
	rrn := 1
	if reRollN != nil {
		var ok bool
		rrn, ok = reRollN.(int)
		if !ok {
			return nil, c.Errorf("reRollN is not a number")
		}
	}
	lessThan := true
	if dirOp != nil {
		var ok bool
		lessThan, ok = dirOp.(bool)
		if !ok {
			return nil, c.Errorf("dirOp is not a bool")
		}
	}
	v, ok := value.(int)
	if !ok {
		return nil, c.Errorf("value is not a number")
	}
	r, ok := roll.(secureRollz.Roller)
	if !ok {
		return nil, c.Errorf("roll is not a Roller")
	}
	return secureRollz.ReRollRoller(rrn, secureRollz.RollValue(v), lessThan, r), nil
}

func (p *parser) callonReRoll1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onReRoll1(stack["roll"], stack["reRollN"], stack["dirOp"], stack["value"])
}

func (c *current) onDiscardRoll1(roll, dirOp, n interface{}) (interface{}, error) {
	lowest := true
	if dirOp != nil {
		var ok bool
		lowest, ok = dirOp.(bool)
		if !ok {
			return nil, c.Errorf("dirOp is not a bool")
		}
	}
	nn, ok := n.(int)
	if !ok {
		return nil, c.Errorf("n is not a number")
	}
	r, ok := roll.(secureRollz.Roller)
	if !ok {
		return nil, c.Errorf("roll is not a Roller")
	}
	return secureRollz.DiscardRoller(nn, lowest, r), nil
}

func (p *parser) callonDiscardRoll1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDiscardRoll1(stack["roll"], stack["dirOp"], stack["n"])
}

func (c *current) onDirOp1(op interface{}) (interface{}, error) {
	opb, ok := op.([]byte)
	if !ok {
		return nil, c.Errorf("op is not a []byte")
	}
	switch opb[0] {
	case '<':
		return true, nil
	case '>':
		return false, nil
	}
	return nil, c.Errorf("unknown direction operator, %s", string(opb))
}

func (p *parser) callonDirOp1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDirOp1(stack["op"])
}

func (c *current) onMultiRoll1(n, roll interface{}) (interface{}, error) {
	nn, ok := n.(int)
	if !ok {
		return nil, c.Errorf("n is not a number")
	}
	r, ok := roll.(secureRollz.Roller)
	if !ok {
		return nil, c.Errorf("roll is not a Roller")
	}
	return secureRollz.MultiRoller(nn, r), nil
}

func (p *parser) callonMultiRoll1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMultiRoll1(stack["n"], stack["roll"])
}

func (c *current) onParenRoll1(roll interface{}) (interface{}, error) {
	roller, ok := roll.(secureRollz.Roller)
	if !ok {
		return nil, c.Errorf("roll is not a Roller")
	}
	return secureRollz.ParenRoller(roller), nil
}

func (p *parser) callonParenRoll1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onParenRoll1(stack["roll"])
}

func (c *current) onDieRoll1(sides interface{}) (interface{}, error) {
	s, ok := sides.(int)
	if !ok {
		return nil, c.Errorf("sides is not a number: %T", sides)
	}
	return secureRollz.DieRoller(s), nil
}

func (p *parser) callonDieRoll1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDieRoll1(stack["sides"])
}

func (c *current) onConstRoll1(value interface{}) (interface{}, error) {
	v, ok := value.(int)
	if !ok {
		return nil, c.Errorf("value is not a number: %T", value)
	}
	return secureRollz.Const(secureRollz.RollValue(v)), nil
}

func (p *parser) callonConstRoll1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConstRoll1(stack["value"])
}

func (c *current) onNumber1() (interface{}, error) {
	ns := string(c.text)
	n, err := strconv.Atoi(ns)
	if err != nil {
		return nil, c.Errorf("invalid number: %s - %s", ns, err.Error())
	}
	return n, nil
}

func (p *parser) callonNumber1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onNumber1()
}

var (
	// errNoRule is returned when the grammar to parse has no rule.
	errNoRule = errors.New("grammar has no rule")

	// errInvalidEncoding is returned when the source is not properly
	// utf8-encoded.
	errInvalidEncoding = errors.New("invalid encoding")

	// errNoMatch is returned if no match could be found.
	errNoMatch = errors.New("no match found")
)

// Option is a function that can set an option on the parser. It returns
// the previous setting as an Option.
type Option func(*parser) Option

// Debug creates an Option to set the debug flag to b. When set to true,
// debugging information is printed to stdout while parsing.
//
// The default is false.
func Debug(b bool) Option {
	return func(p *parser) Option {
		old := p.debug
		p.debug = b
		return Debug(old)
	}
}

// Memoize creates an Option to set the memoize flag to b. When set to true,
// the parser will cache all results so each expression is evaluated only
// once. This guarantees linear parsing time even for pathological cases,
// at the expense of more memory and slower times for typical cases.
//
// The default is false.
func Memoize(b bool) Option {
	return func(p *parser) Option {
		old := p.memoize
		p.memoize = b
		return Memoize(old)
	}
}

// Recover creates an Option to set the recover flag to b. When set to
// true, this causes the parser to recover from panics and convert it
// to an error. Setting it to false can be useful while debugging to
// access the full stack trace.
//
// The default is true.
func Recover(b bool) Option {
	return func(p *parser) Option {
		old := p.recover
		p.recover = b
		return Recover(old)
	}
}

// ParseFile parses the file identified by filename.
func ParseFile(filename string, opts ...Option) (interface{}, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ParseReader(filename, f, opts...)
}

// ParseReader parses the data from r using filename as information in the
// error messages.
func ParseReader(filename string, r io.Reader, opts ...Option) (interface{}, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return Parse(filename, b, opts...)
}

// Parse parses the data from b using filename as information in the
// error messages.
func Parse(filename string, b []byte, opts ...Option) (interface{}, error) {
	return newParser(filename, b, opts...).parse(g)
}

// position records a position in the text.
type position struct {
	line, col, offset int
}

func (p position) String() string {
	return fmt.Sprintf("%d:%d [%d]", p.line, p.col, p.offset)
}

// savepoint stores all state required to go back to this point in the
// parser.
type savepoint struct {
	position
	rn rune
	w  int
}

type current struct {
	pos  position // start position of the match
	text []byte   // raw text of the match
}

// the AST types...

type grammar struct {
	pos   position
	rules []*rule
}

type rule struct {
	pos         position
	name        string
	displayName string
	expr        interface{}
}

type choiceExpr struct {
	pos          position
	alternatives []interface{}
}

type actionExpr struct {
	pos  position
	expr interface{}
	run  func(*parser) (interface{}, error)
}

type seqExpr struct {
	pos   position
	exprs []interface{}
}

type labeledExpr struct {
	pos   position
	label string
	expr  interface{}
}

type expr struct {
	pos  position
	expr interface{}
}

type andExpr expr
type notExpr expr
type zeroOrOneExpr expr
type zeroOrMoreExpr expr
type oneOrMoreExpr expr

type ruleRefExpr struct {
	pos  position
	name string
}

type andCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type notCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type litMatcher struct {
	pos        position
	val        string
	ignoreCase bool
}

type charClassMatcher struct {
	pos        position
	val        string
	chars      []rune
	ranges     []rune
	classes    []*unicode.RangeTable
	ignoreCase bool
	inverted   bool
}

type anyMatcher position

// errList cumulates the errors found by the parser.
type errList []error

func (e *errList) add(err error) {
	*e = append(*e, err)
}

func (e errList) err() error {
	if len(e) == 0 {
		return nil
	}
	e.dedupe()
	return e
}

func (e *errList) dedupe() {
	var cleaned []error
	set := make(map[string]bool)
	for _, err := range *e {
		if msg := err.Error(); !set[msg] {
			set[msg] = true
			cleaned = append(cleaned, err)
		}
	}
	*e = cleaned
}

func (e errList) Error() string {
	switch len(e) {
	case 0:
		return ""
	case 1:
		return e[0].Error()
	default:
		var buf bytes.Buffer

		for i, err := range e {
			if i > 0 {
				buf.WriteRune('\n')
			}
			buf.WriteString(err.Error())
		}
		return buf.String()
	}
}

// parserError wraps an error with a prefix indicating the rule in which
// the error occurred. The original error is stored in the Inner field.
type parserError struct {
	Inner  error
	pos    position
	prefix string
}

// Error returns the error message.
func (p *parserError) Error() string {
	return p.prefix + ": " + p.Inner.Error()
}

// newParser creates a parser with the specified input source and options.
func newParser(filename string, b []byte, opts ...Option) *parser {
	p := &parser{
		filename: filename,
		errs:     new(errList),
		data:     b,
		pt:       savepoint{position: position{line: 1}},
		recover:  true,
	}
	p.setOptions(opts)
	return p
}

// setOptions applies the options to the parser.
func (p *parser) setOptions(opts []Option) {
	for _, opt := range opts {
		opt(p)
	}
}

type resultTuple struct {
	v   interface{}
	b   bool
	end savepoint
}

type parser struct {
	filename string
	pt       savepoint
	cur      current

	data []byte
	errs *errList

	recover bool
	debug   bool
	depth   int

	memoize bool
	// memoization table for the packrat algorithm:
	// map[offset in source] map[expression or rule] {value, match}
	memo map[int]map[interface{}]resultTuple

	// rules table, maps the rule identifier to the rule node
	rules map[string]*rule
	// variables stack, map of label to value
	vstack []map[string]interface{}
	// rule stack, allows identification of the current rule in errors
	rstack []*rule

	// stats
	exprCnt int
}

// push a variable set on the vstack.
func (p *parser) pushV() {
	if cap(p.vstack) == len(p.vstack) {
		// create new empty slot in the stack
		p.vstack = append(p.vstack, nil)
	} else {
		// slice to 1 more
		p.vstack = p.vstack[:len(p.vstack)+1]
	}

	// get the last args set
	m := p.vstack[len(p.vstack)-1]
	if m != nil && len(m) == 0 {
		// empty map, all good
		return
	}

	m = make(map[string]interface{})
	p.vstack[len(p.vstack)-1] = m
}

// pop a variable set from the vstack.
func (p *parser) popV() {
	// if the map is not empty, clear it
	m := p.vstack[len(p.vstack)-1]
	if len(m) > 0 {
		// GC that map
		p.vstack[len(p.vstack)-1] = nil
	}
	p.vstack = p.vstack[:len(p.vstack)-1]
}

func (p *parser) print(prefix, s string) string {
	if !p.debug {
		return s
	}

	fmt.Printf("%s %d:%d:%d: %s [%#U]\n",
		prefix, p.pt.line, p.pt.col, p.pt.offset, s, p.pt.rn)
	return s
}

func (p *parser) in(s string) string {
	p.depth++
	return p.print(strings.Repeat(" ", p.depth)+">", s)
}

func (p *parser) out(s string) string {
	p.depth--
	return p.print(strings.Repeat(" ", p.depth)+"<", s)
}

func (p *parser) addErr(err error) {
	p.addErrAt(err, p.pt.position)
}

func (p *parser) addErrAt(err error, pos position) {
	var buf bytes.Buffer
	if p.filename != "" {
		buf.WriteString(p.filename)
	}
	if buf.Len() > 0 {
		buf.WriteString(":")
	}
	buf.WriteString(fmt.Sprintf("%d:%d (%d)", pos.line, pos.col, pos.offset))
	if len(p.rstack) > 0 {
		if buf.Len() > 0 {
			buf.WriteString(": ")
		}
		rule := p.rstack[len(p.rstack)-1]
		if rule.displayName != "" {
			buf.WriteString("rule " + rule.displayName)
		} else {
			buf.WriteString("rule " + rule.name)
		}
	}
	pe := &parserError{Inner: err, pos: pos, prefix: buf.String()}
	p.errs.add(pe)
}

// read advances the parser to the next rune.
func (p *parser) read() {
	p.pt.offset += p.pt.w
	rn, n := utf8.DecodeRune(p.data[p.pt.offset:])
	p.pt.rn = rn
	p.pt.w = n
	p.pt.col++
	if rn == '\n' {
		p.pt.line++
		p.pt.col = 0
	}

	if rn == utf8.RuneError {
		if n == 1 {
			p.addErr(errInvalidEncoding)
		}
	}
}

// restore parser position to the savepoint pt.
func (p *parser) restore(pt savepoint) {
	if p.debug {
		defer p.out(p.in("restore"))
	}
	if pt.offset == p.pt.offset {
		return
	}
	p.pt = pt
}

// get the slice of bytes from the savepoint start to the current position.
func (p *parser) sliceFrom(start savepoint) []byte {
	return p.data[start.position.offset:p.pt.position.offset]
}

func (p *parser) getMemoized(node interface{}) (resultTuple, bool) {
	if len(p.memo) == 0 {
		return resultTuple{}, false
	}
	m := p.memo[p.pt.offset]
	if len(m) == 0 {
		return resultTuple{}, false
	}
	res, ok := m[node]
	return res, ok
}

func (p *parser) setMemoized(pt savepoint, node interface{}, tuple resultTuple) {
	if p.memo == nil {
		p.memo = make(map[int]map[interface{}]resultTuple)
	}
	m := p.memo[pt.offset]
	if m == nil {
		m = make(map[interface{}]resultTuple)
		p.memo[pt.offset] = m
	}
	m[node] = tuple
}

func (p *parser) buildRulesTable(g *grammar) {
	p.rules = make(map[string]*rule, len(g.rules))
	for _, r := range g.rules {
		p.rules[r.name] = r
	}
}

func (p *parser) parse(g *grammar) (val interface{}, err error) {
	if len(g.rules) == 0 {
		p.addErr(errNoRule)
		return nil, p.errs.err()
	}

	// TODO : not super critical but this could be generated
	p.buildRulesTable(g)

	if p.recover {
		// panic can be used in action code to stop parsing immediately
		// and return the panic as an error.
		defer func() {
			if e := recover(); e != nil {
				if p.debug {
					defer p.out(p.in("panic handler"))
				}
				val = nil
				switch e := e.(type) {
				case error:
					p.addErr(e)
				default:
					p.addErr(fmt.Errorf("%v", e))
				}
				err = p.errs.err()
			}
		}()
	}

	// start rule is rule [0]
	p.read() // advance to first rune
	val, ok := p.parseRule(g.rules[0])
	if !ok {
		if len(*p.errs) == 0 {
			// make sure this doesn't go out silently
			p.addErr(errNoMatch)
		}
		return nil, p.errs.err()
	}
	return val, p.errs.err()
}

func (p *parser) parseRule(rule *rule) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRule " + rule.name))
	}

	if p.memoize {
		res, ok := p.getMemoized(rule)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
	}

	start := p.pt
	p.rstack = append(p.rstack, rule)
	p.pushV()
	val, ok := p.parseExpr(rule.expr)
	p.popV()
	p.rstack = p.rstack[:len(p.rstack)-1]
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}

	if p.memoize {
		p.setMemoized(start, rule, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseExpr(expr interface{}) (interface{}, bool) {
	var pt savepoint
	var ok bool

	if p.memoize {
		res, ok := p.getMemoized(expr)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
		pt = p.pt
	}

	p.exprCnt++
	var val interface{}
	switch expr := expr.(type) {
	case *actionExpr:
		val, ok = p.parseActionExpr(expr)
	case *andCodeExpr:
		val, ok = p.parseAndCodeExpr(expr)
	case *andExpr:
		val, ok = p.parseAndExpr(expr)
	case *anyMatcher:
		val, ok = p.parseAnyMatcher(expr)
	case *charClassMatcher:
		val, ok = p.parseCharClassMatcher(expr)
	case *choiceExpr:
		val, ok = p.parseChoiceExpr(expr)
	case *labeledExpr:
		val, ok = p.parseLabeledExpr(expr)
	case *litMatcher:
		val, ok = p.parseLitMatcher(expr)
	case *notCodeExpr:
		val, ok = p.parseNotCodeExpr(expr)
	case *notExpr:
		val, ok = p.parseNotExpr(expr)
	case *oneOrMoreExpr:
		val, ok = p.parseOneOrMoreExpr(expr)
	case *ruleRefExpr:
		val, ok = p.parseRuleRefExpr(expr)
	case *seqExpr:
		val, ok = p.parseSeqExpr(expr)
	case *zeroOrMoreExpr:
		val, ok = p.parseZeroOrMoreExpr(expr)
	case *zeroOrOneExpr:
		val, ok = p.parseZeroOrOneExpr(expr)
	default:
		panic(fmt.Sprintf("unknown expression type %T", expr))
	}
	if p.memoize {
		p.setMemoized(pt, expr, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseActionExpr(act *actionExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseActionExpr"))
	}

	start := p.pt
	val, ok := p.parseExpr(act.expr)
	if ok {
		p.cur.pos = start.position
		p.cur.text = p.sliceFrom(start)
		actVal, err := act.run(p)
		if err != nil {
			p.addErrAt(err, start.position)
		}
		val = actVal
	}
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}
	return val, ok
}

func (p *parser) parseAndCodeExpr(and *andCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndCodeExpr"))
	}

	ok, err := and.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, ok
}

func (p *parser) parseAndExpr(and *andExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(and.expr)
	p.popV()
	p.restore(pt)
	return nil, ok
}

func (p *parser) parseAnyMatcher(any *anyMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAnyMatcher"))
	}

	if p.pt.rn != utf8.RuneError {
		start := p.pt
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseCharClassMatcher(chr *charClassMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseCharClassMatcher"))
	}

	cur := p.pt.rn
	// can't match EOF
	if cur == utf8.RuneError {
		return nil, false
	}
	start := p.pt
	if chr.ignoreCase {
		cur = unicode.ToLower(cur)
	}

	// try to match in the list of available chars
	for _, rn := range chr.chars {
		if rn == cur {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of ranges
	for i := 0; i < len(chr.ranges); i += 2 {
		if cur >= chr.ranges[i] && cur <= chr.ranges[i+1] {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of Unicode classes
	for _, cl := range chr.classes {
		if unicode.Is(cl, cur) {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	if chr.inverted {
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseChoiceExpr(ch *choiceExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseChoiceExpr"))
	}

	for _, alt := range ch.alternatives {
		p.pushV()
		val, ok := p.parseExpr(alt)
		p.popV()
		if ok {
			return val, ok
		}
	}
	return nil, false
}

func (p *parser) parseLabeledExpr(lab *labeledExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLabeledExpr"))
	}

	p.pushV()
	val, ok := p.parseExpr(lab.expr)
	p.popV()
	if ok && lab.label != "" {
		m := p.vstack[len(p.vstack)-1]
		m[lab.label] = val
	}
	return val, ok
}

func (p *parser) parseLitMatcher(lit *litMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLitMatcher"))
	}

	start := p.pt
	for _, want := range lit.val {
		cur := p.pt.rn
		if lit.ignoreCase {
			cur = unicode.ToLower(cur)
		}
		if cur != want {
			p.restore(start)
			return nil, false
		}
		p.read()
	}
	return p.sliceFrom(start), true
}

func (p *parser) parseNotCodeExpr(not *notCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotCodeExpr"))
	}

	ok, err := not.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, !ok
}

func (p *parser) parseNotExpr(not *notExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(not.expr)
	p.popV()
	p.restore(pt)
	return nil, !ok
}

func (p *parser) parseOneOrMoreExpr(expr *oneOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseOneOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			if len(vals) == 0 {
				// did not match once, no match
				return nil, false
			}
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseRuleRefExpr(ref *ruleRefExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRuleRefExpr " + ref.name))
	}

	if ref.name == "" {
		panic(fmt.Sprintf("%s: invalid rule: missing name", ref.pos))
	}

	rule := p.rules[ref.name]
	if rule == nil {
		p.addErr(fmt.Errorf("undefined rule: %s", ref.name))
		return nil, false
	}
	return p.parseRule(rule)
}

func (p *parser) parseSeqExpr(seq *seqExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseSeqExpr"))
	}

	var vals []interface{}

	pt := p.pt
	for _, expr := range seq.exprs {
		val, ok := p.parseExpr(expr)
		if !ok {
			p.restore(pt)
			return nil, false
		}
		vals = append(vals, val)
	}
	return vals, true
}

func (p *parser) parseZeroOrMoreExpr(expr *zeroOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseZeroOrOneExpr(expr *zeroOrOneExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrOneExpr"))
	}

	p.pushV()
	val, _ := p.parseExpr(expr.expr)
	p.popV()
	// whether it matched or not, consider it a match
	return val, true
}

func rangeTable(class string) *unicode.RangeTable {
	if rt, ok := unicode.Categories[class]; ok {
		return rt
	}
	if rt, ok := unicode.Properties[class]; ok {
		return rt
	}
	if rt, ok := unicode.Scripts[class]; ok {
		return rt
	}

	// cannot happen
	panic(fmt.Sprintf("invalid Unicode class: %s", class))
}
