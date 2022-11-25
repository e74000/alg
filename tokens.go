package alg

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type TokenID uint64

const (
	TidS TokenID = iota
	TidX
	TidSx
	TidExp
	TidTPT
	TidTP
	TidPT
	TidLn
	TidSin
	TidCos
	TidTan
	TidSec
	TidCsc
	TidCot
	TidSinh
	TidCosh
	TidTanh
	TidSech
	TidCsch
	TidCoth
	TidAdd
	TidSub
	TidMul
	TidDiv
	TidGreater
	TidLess
	TidGreaterEqual
	TidLessEqual
	TidEqual
	TidNotEqual
	TidRange
	TidSumSt
	TidSumEn
	TidProdSt
	TidProdEn
)

type Token struct {
	id  TokenID
	val float64
}

type Tokens []Token

var (
	isScalar  = regexp.MustCompile("^-?\\d*\\.?\\d*$")
	isScalarX = regexp.MustCompile("^-?\\d*\\.?\\d*x$")
)

// Tokenise parses a string to a slice of tokens
func Tokenise(s string) Tokens {
	split := strings.Split(s, " ")

	t := make(Tokens, 0, len(split))

	for _, sub := range split {
		if sub == "-" || sub == "." || sub == "x" {

		} else if isScalar.MatchString(sub) {
			v, err := strconv.ParseFloat(sub, 64)
			if err != nil {
				panic(err)
			}

			t = append(t, Token{id: TidS, val: v})
			continue
		} else if sub == "-x" {
			t = append(t, Token{id: TidSx, val: -1})
			continue
		} else if isScalarX.MatchString(sub) {
			v, err := strconv.ParseFloat(sub[:len(sub)-1], 64)
			if err != nil {
				panic(err)
			}

			t = append(t, Token{id: TidSx, val: v})
			continue
		}

		switch sub {
		case "x":
			t = append(t, Token{id: TidX})
		case "e":
			t = append(t, Token{id: TidExp})
		case "^":
			t = append(t, Token{id: TidTPT})
		case "ln":
			t = append(t, Token{id: TidLn})
		case "sin":
			t = append(t, Token{id: TidSin})
		case "cos":
			t = append(t, Token{id: TidCos})
		case "tan":
			t = append(t, Token{id: TidTan})
		case "sec":
			t = append(t, Token{id: TidSec})
		case "csc":
			t = append(t, Token{id: TidCsc})
		case "cot":
			t = append(t, Token{id: TidCot})
		case "+":
			t = append(t, Token{id: TidAdd})
		case "-":
			t = append(t, Token{id: TidSub})
		case "*":
			t = append(t, Token{id: TidMul})
		case "/":
			t = append(t, Token{id: TidDiv})
		case ">":
			t = append(t, Token{id: TidGreater})
		case "<":
			t = append(t, Token{id: TidLess})
		case ">=":
			t = append(t, Token{id: TidGreaterEqual})
		case "<=":
			t = append(t, Token{id: TidLessEqual})
		case "==":
			t = append(t, Token{id: TidEqual})
		case "!=":
			t = append(t, Token{id: TidNotEqual})
		case "<=>":
			t = append(t, Token{id: TidRange})
		case "+[":
			t = append(t, Token{id: TidSumSt})
		case "]+":
			t = append(t, Token{id: TidSumEn})
		case "*[":
			t = append(t, Token{id: TidProdSt})
		case "]*":
			t = append(t, Token{id: TidProdEn})
		default:
			panic(fmt.Sprintf("ERROR: Unknown token: %s", sub))
		}
	}

	return t
}

// pop removes the first element of a token slice
// It is unexported since it is just a tiny utility function.
func (t *Tokens) pop() Token {
	out := (*t)[0]
	*t = (*t)[1:]
	return out
}

// Parse converts a token slice to a tree with prefix notation
func (t *Tokens) Parse() Term {
	temp := t.pop()

	switch temp.id {
	case TidS:
		return S(temp.val)
	case TidX:
		return X{}
	case TidSx:
		return Sx{temp.val}
	case TidExp:
		return Exp{
			X: t.Parse(),
		}
	case TidLn:
		return Ln{
			X: t.Parse(),
		}
	case TidTPT:
		return TPT{
			A: t.Parse(),
			B: t.Parse(),
		}
	case TidTP:
		return TP{
			X: t.Parse(),
			P: temp.val,
		}
	case TidPT:
		return PT{
			V: temp.val,
			X: t.Parse(),
		}
	case TidSin:
		return Sin{
			X: t.Parse(),
		}
	case TidCos:
		return Cos{
			X: t.Parse(),
		}
	case TidTan:
		return Tan{
			X: t.Parse(),
		}
	case TidSec:
		return Sec{
			X: t.Parse(),
		}
	case TidCsc:
		return Csc{
			X: t.Parse(),
		}
	case TidCot:
		return Cot{
			X: t.Parse(),
		}
	case TidAdd:
		return Add{
			A: t.Parse(),
			B: t.Parse(),
		}
	case TidSub:
		return Sub{
			A: t.Parse(),
			B: t.Parse(),
		}
	case TidMul:
		return Mul{
			A: t.Parse(),
			B: t.Parse(),
		}
	case TidDiv:
		return Div{
			N: t.Parse(),
			D: t.Parse(),
		}
	case TidGreater:
		return Greater{
			A:    t.Parse(),
			B:    t.Parse(),
			If:   t.Parse(),
			Else: t.Parse(),
		}
	case TidLess:
		return Less{
			A:    t.Parse(),
			B:    t.Parse(),
			If:   t.Parse(),
			Else: t.Parse(),
		}
	case TidGreaterEqual:
		return GreaterEqual{
			A:    t.Parse(),
			B:    t.Parse(),
			If:   t.Parse(),
			Else: t.Parse(),
		}
	case TidLessEqual:
		return LessEqual{
			A:    t.Parse(),
			B:    t.Parse(),
			If:   t.Parse(),
			Else: t.Parse(),
		}
	case TidEqual:
		return Equal{
			A:    t.Parse(),
			B:    t.Parse(),
			If:   t.Parse(),
			Else: t.Parse(),
		}
	case TidNotEqual:
		return NotEqual{
			A:    t.Parse(),
			B:    t.Parse(),
			If:   t.Parse(),
			Else: t.Parse(),
		}
	case TidRange:
		return Range{
			X:    t.Parse(),
			A:    t.Parse(),
			B:    t.Parse(),
			If:   t.Parse(),
			Else: t.Parse(),
		}
	case TidSumSt:
		sum := make(Sum, 0)
		closed := false
		for len(*t) != 0 {
			if (*t)[0].id == TidSumEn {
				closed = true
				t.pop()
			}

			sum = append(sum, t.Parse())
		}

		if !closed {
			panic("Unmatched brackets")
		}

		return sum
	case TidProdSt:
		prod := make(Prod, 0)
		closed := false
		for len(*t) != 0 {
			if (*t)[0].id == TidProdEn {
				closed = true
				t.pop()
			}

			prod = append(prod, t.Parse())
		}

		if !closed {
			panic("Unmatched brackets")
		}

		return prod
	case TidSumEn, TidProdEn:
		panic("Unmatched brackets!")
	}

	panic("ERROR: Failed to parse token!")
	return nil
}

// String converts a token slice to a string
func (t *Tokens) String() string {
	s := ""

	for _, token := range *t {
		switch token.id {
		case TidS:
			s += fmt.Sprintf("%.2f ", token.val)
		case TidSx:
			s += fmt.Sprintf("%.2fx", token.val)
		case TidX:
			s += "x "
		case TidExp:
			s += "e "
		case TidLn:
			s += "ln "
		case TidTPT, TidTP, TidPT:
			s += "^ "
		case TidSin:
			s += "sin "
		case TidCos:
			s += "cos "
		case TidTan:
			s += "tan "
		case TidSec:
			s += "sec "
		case TidCsc:
			s += "csc "
		case TidCot:
			s += "cot "
		case TidAdd:
			s += "+ "
		case TidSub:
			s += "- "
		case TidMul:
			s += "* "
		case TidDiv:
			s += "/ "
		case TidGreater:
			s += "> "
		case TidLess:
			s += "< "
		case TidGreaterEqual:
			s += ">= "
		case TidLessEqual:
			s += "<= "
		case TidEqual:
			s += "== "
		case TidNotEqual:
			s += "!= "
		case TidRange:
			s += "<=> "
		case TidSumSt:
			s += "+[ "
		case TidSumEn:
			s += "]+ "
		case TidProdSt:
			s += "*[ "
		case TidProdEn:
			s += "]* "
		default:
			panic("ERROR: Unknown token")
		}
	}
	return s
}
