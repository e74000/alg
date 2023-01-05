package alg

import (
	"errors"
	"fmt"
	"github.com/e74000/bimap"
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

var mTokenString = map[TokenID]string{
	TidX:            "x",
	TidExp:          "e",
	TidTPT:          "^",
	TidLn:           "ln",
	TidSin:          "sin",
	TidCos:          "cos",
	TidTan:          "tan",
	TidSec:          "sec",
	TidCsc:          "csc",
	TidCot:          "cot",
	TidSinh:         "sinh",
	TidCosh:         "cosh",
	TidTanh:         "tanh",
	TidSech:         "sech",
	TidCsch:         "csch",
	TidCoth:         "coth",
	TidAdd:          "+",
	TidSub:          "-",
	TidMul:          "*",
	TidDiv:          "/",
	TidGreater:      ">",
	TidLess:         "<",
	TidGreaterEqual: ">=",
	TidLessEqual:    "<=",
	TidEqual:        "==",
	TidNotEqual:     "!=",
	TidRange:        "<=>",
	TidSumSt:        "+[",
	TidSumEn:        "]+",
	TidProdSt:       "*[",
	TidProdEn:       "]*",
}

var bmTokenString = bimap.MapToBimap(mTokenString)

type Tokens []Token

var (
	isScalar  = regexp.MustCompile("^-?\\d*\\.?\\d*$")
	isScalarX = regexp.MustCompile("^-?\\d*\\.?\\d*x$")
)

// Tokenise parses a string to a slice of tokens
func Tokenise(s string) (Tokens, error) {
	split := strings.Split(s, " ")

	t := make(Tokens, 0, len(split))

	for _, sub := range split {
		if sub == "" {
			continue
		} else if sub == "-" || sub == "." || sub == "x" {

		} else if isScalar.MatchString(sub) {
			v, err := strconv.ParseFloat(sub, 64)
			if err != nil {
				return nil, err
			}

			t = append(t, Token{id: TidS, val: v})
			continue
		} else if sub == "-x" {
			t = append(t, Token{id: TidSx, val: -1})
			continue
		} else if isScalarX.MatchString(sub) {
			v, err := strconv.ParseFloat(sub[:len(sub)-1], 64)
			if err != nil {
				return nil, err
			}

			t = append(t, Token{id: TidSx, val: v})
			continue
		}

		t = append(t, Token{id: bmTokenString.GetRev(sub)})
	}

	return t, nil
}

// pop removes the first element of a token slice
// It is unexported since it is just a tiny utility function.
func (t *Tokens) pop() Token {
	out := (*t)[0]
	*t = (*t)[1:]
	return out
}

// Parse converts a token slice to a tree with prefix notation
// This is a massive automatically-ish generated switch statement - I would rather if it could be avoided, but I can't figure out any better methods right now...
func (t *Tokens) Parse() (Term, error) {
	temp := t.pop()

	switch temp.id {
	case TidS:
		return S(temp.val), nil
	case TidX:
		return X{}, nil
	case TidSx:
		return Sx{temp.val}, nil
	case TidExp:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return Exp{
			X: arg1,
		}, nil
	case TidLn:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return Ln{
			X: arg1,
		}, nil
	case TidTPT:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg2, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return TPT{
			A: arg1,
			B: arg2,
		}, nil
	case TidTP:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return TP{
			X: arg1,
			P: temp.val,
		}, nil
	case TidPT:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return PT{
			V: temp.val,
			X: arg1,
		}, nil
	case TidSin:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return Sin{
			X: arg1,
		}, nil
	case TidCos:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return Cos{
			X: arg1,
		}, nil
	case TidTan:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return Tan{
			X: arg1,
		}, nil
	case TidSec:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return Sec{
			X: arg1,
		}, nil
	case TidCsc:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return Csc{
			X: arg1,
		}, nil
	case TidCot:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return Cot{
			X: arg1,
		}, nil
	case TidCosh:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return Cosh{
			X: arg1,
		}, nil
	case TidSinh:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return Sinh{
			X: arg1,
		}, nil
	case TidTanh:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return Tanh{
			X: arg1,
		}, nil
	case TidSech:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return Sech{
			X: arg1,
		}, nil
	case TidCsch:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return Csch{
			X: arg1,
		}, nil
	case TidCoth:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return Coth{
			X: arg1,
		}, nil
	case TidAdd:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg2, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return Add{
			A: arg1,
			B: arg2,
		}, nil
	case TidSub:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg2, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return Sub{
			A: arg1,
			B: arg2,
		}, nil
	case TidMul:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg2, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return Mul{
			A: arg1,
			B: arg2,
		}, nil
	case TidDiv:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg2, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return Div{
			N: arg1,
			D: arg2,
		}, nil
	case TidGreater:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg2, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg3, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg4, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return Greater{
			A:    arg1,
			B:    arg2,
			If:   arg3,
			Else: arg4,
		}, nil
	case TidLess:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg2, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg3, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg4, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return Less{
			A:    arg1,
			B:    arg2,
			If:   arg3,
			Else: arg4,
		}, nil
	case TidGreaterEqual:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg2, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg3, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg4, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return GreaterEqual{
			A:    arg1,
			B:    arg2,
			If:   arg3,
			Else: arg4,
		}, nil
	case TidLessEqual:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg2, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg3, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg4, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return LessEqual{
			A:    arg1,
			B:    arg2,
			If:   arg3,
			Else: arg4,
		}, nil
	case TidEqual:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg2, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg3, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg4, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return Equal{
			A:    arg1,
			B:    arg2,
			If:   arg3,
			Else: arg4,
		}, nil
	case TidNotEqual:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg2, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg3, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg4, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return NotEqual{
			A:    arg1,
			B:    arg2,
			If:   arg3,
			Else: arg4,
		}, nil
	case TidRange:
		arg1, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg2, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg3, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg4, err := t.Parse()
		if err != nil {
			return nil, err
		}
		arg5, err := t.Parse()
		if err != nil {
			return nil, err
		}
		return Range{
			X:    arg1,
			A:    arg2,
			B:    arg3,
			If:   arg4,
			Else: arg5,
		}, nil
	case TidSumSt:
		sum := make(Sum, 0)
		closed := false
		for len(*t) != 0 {
			if (*t)[0].id == TidSumEn {
				closed = true
				t.pop()
				break
			}

			arg, err := t.Parse()
			if err != nil {
				return nil, err
			}

			sum = append(sum, arg)
		}

		if !closed {
			return nil, errors.New("unmatched brackets")
		}

		return sum, nil
	case TidProdSt:
		prod := make(Prod, 0)
		closed := false
		for len(*t) != 0 {
			if (*t)[0].id == TidProdEn {
				closed = true
				t.pop()
				break
			}

			arg, err := t.Parse()
			if arg != nil {
				return nil, err
			}

			prod = append(prod, arg)
		}

		if !closed {
			return nil, errors.New("unmatched brackets")
		}

		return prod, nil
	case TidSumEn, TidProdEn:
		return nil, errors.New("unmatched brackets")
	}

	return nil, errors.New("failed to parse token")
}

// String converts a token slice to a string
func (t *Tokens) String() string {
	s := ""

	for _, token := range *t {
		switch token.id {
		case TidS:
			s += fmt.Sprintf("%.2f", token.val)
		case TidSx:
			s += fmt.Sprintf("%.2fx", token.val)
		default:
			s += bmTokenString.GetFor(token.id)
		}

		s += " "
	}
	return s
}
