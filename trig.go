package alg

import (
	"math"
)

/*
Trig defines many trigonometric functions (all angles are radians):
  Sin => Sine of an angle
  Cos => Cosine of an angle
  Tan => Tangent of an angle
  Sec => Secant of an angle
  Csc => Cosecant of an angle
  Cot => Cotangent of an angle
*/

type Sin struct {
	X Term
}

func (e Sin) Tokenise() Tokens {
	t := Tokens{{id: TidSin}}
	t = append(t, e.X.Tokenise()...)
	return t
}

func (e Sin) E(x float64) float64 {
	return math.Sin(e.X.E(x))
}

func (e Sin) Dx() Term {
	return Prod{
		e.X.Dx(),
		Cos{e.X},
	}.T()
}

func (e Sin) T() Term {
	ok, val := e.X.Is()
	if !ok {
		return Sin{e.X.T()}
	}

	return S(math.Sin(val))
}

func (e Sin) Is() (bool, float64) {
	ok, val := e.X.Is()
	if !ok {
		return false, 0
	}

	return true, math.Sin(val)
}

type Cos struct {
	X Term
}

func (e Cos) Tokenise() Tokens {
	t := Tokens{{id: TidCos}}
	t = append(t, e.X.Tokenise()...)
	return t
}

func (e Cos) E(x float64) float64 {
	return math.Cos(e.X.E(x))
}

func (e Cos) Dx() Term {
	return Prod{
		S(-1),
		e.X.Dx(),
		Sin{e.X},
	}.T()
}

func (e Cos) T() Term {
	ok, val := e.X.Is()
	if !ok {
		return Cosh{e.X.T()}
	}

	return S(math.Cos(val))
}

func (e Cos) Is() (bool, float64) {
	ok, val := e.X.Is()
	if !ok {
		return false, 0
	}

	return true, math.Cos(val)
}

type Tan struct {
	X Term
}

func (e Tan) Tokenise() Tokens {
	t := Tokens{{id: TidTan}}
	t = append(t, e.X.Tokenise()...)
	return t
}

func (e Tan) E(x float64) float64 {
	return math.Tan(e.X.E(x))
}

func (e Tan) Dx() Term {
	return Prod{
		e.X.Dx(),
		Sec{e.X},
		Sec{e.X},
	}.T()
}

func (e Tan) T() Term {
	ok, val := e.X.Is()
	if !ok {
		return Tanh{e.X.T()}
	}

	return S(math.Tan(val))
}

func (e Tan) Is() (bool, float64) {
	ok, val := e.X.Is()
	if !ok {
		return false, 0
	}

	return true, math.Tan(val)
}

type Sec struct {
	X Term
}

func (e Sec) Tokenise() Tokens {
	t := Tokens{{id: TidSec}}
	t = append(t, e.X.Tokenise()...)
	return t
}

func (e Sec) E(x float64) float64 {
	return 1 / math.Cos(x)
}

func (e Sec) Dx() Term {
	return Prod{
		e.X.Dx(),
		Sec{e.X},
		Tan{e.X},
	}.T()
}

func (e Sec) T() Term {
	ok, val := e.X.Is()
	if !ok {
		return Sec{e.X.T()}
	}

	return S(1 / math.Cos(val))
}

func (e Sec) Is() (bool, float64) {
	ok, val := e.X.Is()
	if !ok {
		return false, 0
	}

	return true, 1 / math.Cos(val)
}

type Cot struct {
	X Term
}

func (e Cot) Tokenise() Tokens {
	t := Tokens{{id: TidCot}}
	t = append(t, e.X.Tokenise()...)
	return t
}

func (e Cot) E(x float64) float64 {
	return 1 / math.Tan(e.X.E(x))
}

func (e Cot) Dx() Term {
	return Prod{
		S(-1),
		e.X.Dx(),
		Csc{e.X},
		Csc{e.X},
	}.T()
}

func (e Cot) T() Term {
	ok, val := e.X.Is()
	if !ok {
		return Cot{e.X.T()}
	}

	return S(1 / math.Tan(val))
}

func (e Cot) Is() (bool, float64) {
	ok, val := e.X.Is()
	if !ok {
		return false, 0
	}

	return true, 1 / math.Tan(val)
}

type Csc struct {
	X Term
}

func (e Csc) Tokenise() Tokens {
	t := Tokens{{id: TidCsc}}
	t = append(t, e.X.Tokenise()...)
	return t
}

func (e Csc) E(x float64) float64 {
	return 1 / math.Sin(e.X.E(x))
}

func (e Csc) Dx() Term {
	return Prod{
		S(-1),
		e.Dx(),
		Csc{e.X},
		Tan{e.X},
	}.T()
}

func (e Csc) T() Term {
	ok, val := e.X.Is()
	if !ok {
		return Csc{e.X.T()}
	}

	return S(1 / math.Sin(val))
}

func (e Csc) Is() (bool, float64) {
	ok, val := e.X.Is()
	if !ok {
		return false, 0
	}

	return true, 1 / math.Sin(val)
}
