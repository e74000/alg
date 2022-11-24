package alg

import "math"

/*
Hyperbolic defines many trigonometric functions:
  Sinh => Hyperbolic sin
  Cosh => Hyperbolic cosine
  Tanh => Hyperbolic tangent
  Sech => Hyperbolic secant
  Csch => Hyperbolic cosecant
  Coth => Hyperbolic cotangent
*/

type Sinh struct {
	X Term
}

func (e Sinh) E(x float64) float64 {
	return math.Sinh(e.X.E(x))
}

func (e Sinh) Dx() Term {
	return Prod{
		e.X.Dx(),
		Cosh{
			e.X,
		}.T(),
	}
}

func (e Sinh) T() Term {
	ok, v := e.X.Is()

	if ok {
		return S(math.Sinh(v))
	}

	return Sinh{e.X.T()}
}

func (e Sinh) Is() (bool, float64) {
	ok, v := e.X.Is()

	if ok {
		return true, math.Sinh(v)
	}

	return false, 0
}

func (e Sinh) Tokenise() Tokens {
	t := Tokens{{id: TidSinh}}
	t = append(t, e.X.Tokenise()...)
	return t
}

type Cosh struct {
	X Term
}

func (e Cosh) E(x float64) float64 {
	return math.Cosh(e.X.E(x))
}

func (e Cosh) Dx() Term {
	return Prod{
		e.X.Dx(),
		Sinh{
			e.X,
		}.T(),
	}
}

func (e Cosh) T() Term {
	ok, v := e.X.Is()

	if ok {
		return S(math.Cosh(v))
	}

	return Cosh{e.X.T()}
}

func (e Cosh) Is() (bool, float64) {
	ok, v := e.X.Is()

	if ok {
		return true, math.Cosh(v)
	}

	return false, 0
}

func (e Cosh) Tokenise() Tokens {
	t := Tokens{{id: TidCosh}}
	t = append(t, e.X.Tokenise()...)
	return t
}

type Tanh struct {
	X Term
}

func (e Tanh) E(x float64) float64 {
	return math.Tanh(e.X.E(x))
}

func (e Tanh) Dx() Term {
	return Mul{
		A: e.X.Dx(),
		B: TP{
			X: Sech{e.X},
			P: 2,
		},
	}.T()
}

func (e Tanh) T() Term {
	ok, v := e.X.Is()
	if ok {
		return S(math.Tanh(v))
	}

	return Tanh{e.X.T()}
}

func (e Tanh) Is() (bool, float64) {
	ok, v := e.X.Is()
	if ok {
		return true, math.Tanh(v)
	}

	return false, 0
}

func (e Tanh) Tokenise() Tokens {
	t := Tokens{{id: TidTanh}}
	t = append(t, e.X.Tokenise()...)
	return t
}

type Coth struct {
	X Term
}

func (e Coth) E(x float64) float64 {
	return 1 / math.Tanh(e.X.E(x))
}

func (e Coth) Dx() Term {
	return Prod{
		S(-1),
		e.X.Dx(),
		TP{
			X: Csch{e.X},
			P: 2,
		},
	}.T()
}

func (e Coth) T() Term {
	ok, v := e.X.Is()
	if ok {
		return S(1 / math.Tanh(v))
	}

	return Coth{e.X.T()}
}

func (e Coth) Is() (bool, float64) {
	ok, v := e.X.Is()
	if ok {
		return true, 1 / math.Tanh(v)
	}

	return false, 0
}

func (e Coth) Tokenise() Tokens {
	t := Tokens{{id: TidCoth}}
	t = append(t, e.X.Tokenise()...)
	return t
}

type Sech struct {
	X Term
}

func (e Sech) E(x float64) float64 {
	return 1 / math.Cosh(e.X.E(x))
}

func (e Sech) Dx() Term {
	return Prod{
		S(-1),
		e.X.Dx(),
		Tanh{e.X},
		Sech{e.X},
	}
}

func (e Sech) T() Term {
	ok, v := e.X.Is()
	if ok {
		return S(1 / math.Cosh(v))
	}

	return Sech{e.X.T()}
}

func (e Sech) Is() (bool, float64) {
	ok, v := e.X.Is()
	if ok {
		return true, 1 / math.Cosh(v)
	}

	return false, 0
}

func (e Sech) Tokenise() Tokens {
	t := Tokens{{id: TidSech}}
	t = append(t, e.X.Tokenise()...)
	return t
}

type Csch struct {
	X Term
}

func (e Csch) E(x float64) float64 {
	return 1 / math.Sinh(e.X.E(x))
}

func (e Csch) Dx() Term {
	return Prod{
		S(-1),
		e.X.Dx(),
		Coth{e.X},
		Csch{e.X},
	}
}

func (e Csch) T() Term {
	ok, v := e.X.Is()
	if ok {
		return S(1 / math.Sinh(v))
	}

	return Csch{e.X.T()}
}

func (e Csch) Is() (bool, float64) {
	ok, v := e.X.Is()
	if ok {
		return true, 1 / math.Sinh(v)
	}

	return false, 0
}

func (e Csch) Tokenise() Tokens {
	t := Tokens{{id: TidCsch}}
	t = append(t, e.X.Tokenise()...)
	return t
}
