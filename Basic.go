package alg

import (
	"math"
)

/*
Basic defines the set of "simple" algebraic datatypes:
  S   => A scalar, a float constant
  X   => A variable.
  Sx  => A scalar multiplied by a variable
  TPT => Term to the power of a term
  TP  => Term to the power of a scalar
  PT  => Scalar to the power of a term
  Exp => The exponent of a term
  Ln  => The natural logarithm of a term
*/

type S float64

func (e S) E(_ float64) float64 {
	return float64(e)
}

func (e S) Dx() Term {
	return S(0)
}

func (e S) T() Term {
	return e
}

func (e S) Is() (bool, float64) {
	return true, float64(e)
}

func (e S) Tokenise() Tokens {
	return Tokens{{id: TidS, val: float64(e)}}
}

type X struct{}

func (e X) E(x float64) float64 {
	return x
}

func (e X) Dx() Term {
	return S(1)
}

func (e X) T() Term {
	return e
}

func (e X) Is() (bool, float64) {
	return false, 0
}

func (e X) Tokenise() Tokens {
	return Tokens{{id: TidX}}
}

type Sx struct {
	S float64
}

func (e Sx) E(x float64) float64 {
	return e.S * x
}

func (e Sx) Dx() Term {
	return S(e.S)
}

func (e Sx) T() Term {
	if e.S == 0 {
		return S(0)
	}

	return e
}

func (e Sx) Is() (bool, float64) {
	if e.S == 0 {
		return true, 0
	}

	return false, 0
}

func (e Sx) Tokenise() Tokens {
	return Tokens{{id: TidSx, val: e.S}}
}

type Exp struct {
	X Term
}

func (e Exp) E(x float64) float64 {
	return math.Exp(e.X.E(x))
}

func (e Exp) Dx() Term {
	return Prod{e.X.Dx(), Exp{e.X}}.T()
}

func (e Exp) T() Term {
	ok, val := e.X.Is()
	if ok {
		return S(math.Exp(val))
	}

	return Exp{e.X.T()}
}

func (e Exp) Is() (bool, float64) {
	ok, val := e.X.Is()
	if ok {
		return true, math.Exp(val)
	}

	return false, 0
}

func (e Exp) Tokenise() Tokens {
	t := Tokens{{id: TidExp}}
	t = append(t, e.X.Tokenise()...)
	return t
}

type TPT struct {
	A, B Term
}

func (e TPT) E(x float64) float64 {
	return math.Pow(e.A.E(x), e.B.E(x))
}

func (e TPT) Dx() Term {
	return Prod{
		TPT{e.A, Sum{e.B, S(-1)}},
		Sum{
			Prod{
				e.B.Dx(),
				e.A,
				Ln{e.A},
			},
			Prod{
				e.B,
				e.A.Dx(),
			},
		},
	}
}

func (e TPT) T() Term {
	aOk, aVal := e.A.Is()
	bOk, bVal := e.B.Is()

	if aOk && bOk {
		return S(math.Pow(aVal, bVal))
	} else if aOk {
		if aVal == 0 || aVal == 1 {
			return S(aVal)
		}
		return PT{aVal, e.B.T()}
	} else if bOk {
		if bVal == 0 {
			return S(1)
		} else if bVal == 1 {
			return e.A
		}
		return TP{e.A.T(), bVal}
	}

	return TPT{
		e.A.T(),
		e.B.T(),
	}
}

func (e TPT) Is() (bool, float64) {
	aOk, aVal := e.A.Is()
	bOk, bVal := e.B.Is()

	if aOk && bOk {
		return true, math.Pow(aVal, bVal)
	} else if aOk {
		if aVal == 0 || aVal == 1 {
			return true, aVal
		}
	} else if bOk {
		if bVal == 0 {
			return true, 1
		}
	}

	return false, 0
}

func (e TPT) Tokenise() Tokens {
	t := Tokens{{id: TidTPT}}
	t = append(t, e.A.Tokenise()...)
	t = append(t, e.B.Tokenise()...)
	return t
}

type TP struct {
	X Term
	P float64
}

func (e TP) E(x float64) float64 {
	return math.Pow(e.X.E(x), e.P)
}

func (e TP) Dx() Term {
	return Prod{
		S(e.P),
		TP{
			X: e.X,
			P: e.P - 1,
		},
	}.T()
}

func (e TP) T() Term {
	ok, val := e.X.Is()
	if ok {
		return S(math.Pow(val, e.P))
	}

	if e.P == 0 {
		return S(0)
	}

	if e.P == 1 {
		return e.X.T()
	}

	return TP{e.X.T(), e.P}
}

func (e TP) Is() (bool, float64) {
	ok, val := e.X.Is()
	if ok {
		return true, math.Pow(val, e.P)
	}

	if e.P == 0 {
		return true, 0
	}

	return false, 0
}

func (e TP) Tokenise() Tokens {
	t := Tokens{{id: TidTPT}}
	t = append(t, e.X.Tokenise()...)
	t = append(t, Token{id: TidS, val: e.P})
	return t
}

type PT struct {
	V float64
	X Term
}

func (e PT) E(x float64) float64 {
	return math.Pow(e.V, e.X.E(x))
}

func (e PT) Dx() Term {
	return Prod{
		S(math.Log(e.V)),
		e.X.Dx(),
		PT{
			V: e.V,
			X: e.X,
		},
	}.T()
}

func (e PT) T() Term {
	ok, val := e.X.Is()

	if ok {
		return S(math.Pow(e.V, val))
	}

	return PT{e.V, e.X.T()}
}

func (e PT) Is() (bool, float64) {
	ok, val := e.X.Is()

	if ok {
		return true, math.Pow(e.V, val)
	}

	return false, 0
}

func (e PT) Tokenise() Tokens {
	t := Tokens{{id: TidTPT}}
	t = append(t, Token{id: TidS, val: e.V})
	t = append(t, e.X.Tokenise()...)
	return t
}

type Ln struct {
	X Term
}

func (e Ln) E(x float64) float64 {
	return math.Log(e.X.E(x))
}

func (e Ln) Dx() Term {
	return Div{
		e.X.Dx(),
		e.X,
	}.T()
}

func (e Ln) T() Term {
	ok, val := e.X.Is()
	if ok {
		return S(val)
	}

	return Ln{e.X.T()}
}

func (e Ln) Is() (bool, float64) {
	ok, val := e.X.Is()
	if ok {
		return true, val
	}

	return false, 0
}

func (e Ln) Tokenise() Tokens {
	t := Tokens{{id: TidLn}}
	t = append(t, e.X.Tokenise()...)
	return t
}
