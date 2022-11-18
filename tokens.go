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

func Tokenise(s string) Tokens {
	split := strings.Split(s, " ")

	t := make(Tokens, 0, len(split))

	for _, sub := range split {
		if isScalar.MatchString(sub) {
			v, err := strconv.ParseFloat(sub, 64)
			if err != nil {
				panic(err)
			}

			t = append(t, Token{id: TidS, val: v})
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
		default:
			panic(fmt.Sprintf("ERROR: Unknown token: %s", sub))
		}
	}

	return t
}

func (t *Tokens) Pop() Token {
	out := (*t)[0]
	*t = (*t)[1:]
	return out
}

func (t *Tokens) Parse() Term {
	temp := t.Pop()

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
	case TidSumEn, TidSumSt, TidProdEn, TidProdSt:
		panic("ERROR: Sum/Prod are currently unsupported while I figure how to make them parse properly...")
	}

	panic("ERROR: Failed to parse token!")
	return nil
}

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
		default:
			panic("ERROR: Unknown token")
		}
	}
	return s
}
