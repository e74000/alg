package alg

type Greater struct {
	A, B, If, Else Term
}

func (e Greater) E(x float64) float64 {
	av := e.A.E(x)
	bv := e.B.E(x)

	if av > bv {
		return e.If.E(x)
	} else {
		return e.Else.E(x)
	}
}

func (e Greater) Dx() Term {
	return Greater{
		A:    e.A,
		B:    e.B,
		If:   e.If.Dx(),
		Else: e.Else.Dx(),
	}.T()
}

func (e Greater) T() Term {
	aOk, av := e.A.Is()
	bOk, bv := e.B.Is()

	if aOk && bOk {
		if av > bv {
			return e.If.T()
		} else {
			return e.Else.T()
		}
	}

	return Greater{
		A:    e.A.T(),
		B:    e.B.T(),
		If:   e.If.T(),
		Else: e.Else.T(),
	}
}

func (e Greater) Is() (bool, float64) {
	aOk, av := e.A.Is()
	bOk, bv := e.B.Is()

	if aOk && bOk {
		if av > bv {
			return e.If.Is()
		} else {
			return e.Else.Is()
		}
	}

	return false, 0
}

func (e Greater) Tokenise() Tokens {
	t := Tokens{{id: TidGreater}}
	t = append(t, e.A.Tokenise()...)
	t = append(t, e.B.Tokenise()...)
	t = append(t, e.If.Tokenise()...)
	t = append(t, e.Else.Tokenise()...)

	return t
}

type Less struct {
	A, B, If, Else Term
}

func (e Less) E(x float64) float64 {
	av := e.A.E(x)
	bv := e.B.E(x)

	if av < bv {
		return e.If.E(x)
	} else {
		return e.Else.E(x)
	}
}

func (e Less) Dx() Term {
	return Less{
		A:    e.A,
		B:    e.B,
		If:   e.If.Dx(),
		Else: e.Else.Dx(),
	}.T()
}

func (e Less) T() Term {
	aOk, av := e.A.Is()
	bOk, bv := e.B.Is()

	if aOk && bOk {
		if av < bv {
			return e.If.T()
		} else {
			return e.Else.T()
		}
	}

	return Less{
		A:    e.A.T(),
		B:    e.B.T(),
		If:   e.If.T(),
		Else: e.Else.T(),
	}
}

func (e Less) Is() (bool, float64) {
	aOk, av := e.A.Is()
	bOk, bv := e.B.Is()

	if aOk && bOk {
		if av < bv {
			return e.If.Is()
		} else {
			return e.Else.Is()
		}
	}

	return false, 0
}

func (e Less) Tokenise() Tokens {
	t := Tokens{{id: TidLess}}
	t = append(t, e.A.Tokenise()...)
	t = append(t, e.B.Tokenise()...)
	t = append(t, e.If.Tokenise()...)
	t = append(t, e.Else.Tokenise()...)

	return t
}

type GreaterEqual struct {
	A, B, If, Else Term
}

func (e GreaterEqual) E(x float64) float64 {
	av := e.A.E(x)
	bv := e.B.E(x)

	if av >= bv {
		return e.If.E(x)
	} else {
		return e.Else.E(x)
	}
}

func (e GreaterEqual) Dx() Term {
	return GreaterEqual{
		A:    e.A,
		B:    e.B,
		If:   e.If.Dx(),
		Else: e.Else.Dx(),
	}.T()
}

func (e GreaterEqual) T() Term {
	aOk, av := e.A.Is()
	bOk, bv := e.B.Is()

	if aOk && bOk {
		if av >= bv {
			return e.If.T()
		} else {
			return e.Else.T()
		}
	}

	return GreaterEqual{
		A:    e.A.T(),
		B:    e.B.T(),
		If:   e.If.T(),
		Else: e.Else.T(),
	}
}

func (e GreaterEqual) Is() (bool, float64) {
	aOk, av := e.A.Is()
	bOk, bv := e.B.Is()

	if aOk && bOk {
		if av >= bv {
			return e.If.Is()
		} else {
			return e.Else.Is()
		}
	}

	return false, 0
}

func (e GreaterEqual) Tokenise() Tokens {
	t := Tokens{{id: TidGreaterEqual}}
	t = append(t, e.A.Tokenise()...)
	t = append(t, e.B.Tokenise()...)
	t = append(t, e.If.Tokenise()...)
	t = append(t, e.Else.Tokenise()...)

	return t
}

type LessEqual struct {
	A, B, If, Else Term
}

func (e LessEqual) E(x float64) float64 {
	av := e.A.E(x)
	bv := e.B.E(x)

	if av <= bv {
		return e.If.E(x)
	} else {
		return e.Else.E(x)
	}
}

func (e LessEqual) Dx() Term {
	return LessEqual{
		A:    e.A,
		B:    e.B,
		If:   e.If.Dx(),
		Else: e.Else.Dx(),
	}.T()
}

func (e LessEqual) T() Term {
	aOk, av := e.A.Is()
	bOk, bv := e.B.Is()

	if aOk && bOk {
		if av <= bv {
			return e.If.T()
		} else {
			return e.Else.T()
		}
	}

	return LessEqual{
		A:    e.A.T(),
		B:    e.B.T(),
		If:   e.If.T(),
		Else: e.Else.T(),
	}
}

func (e LessEqual) Is() (bool, float64) {
	aOk, av := e.A.Is()
	bOk, bv := e.B.Is()

	if aOk && bOk {
		if av <= bv {
			return e.If.Is()
		} else {
			return e.Else.Is()
		}
	}

	return false, 0
}

func (e LessEqual) Tokenise() Tokens {
	t := Tokens{{id: TidLessEqual}}
	t = append(t, e.A.Tokenise()...)
	t = append(t, e.B.Tokenise()...)
	t = append(t, e.If.Tokenise()...)
	t = append(t, e.Else.Tokenise()...)

	return t
}

type Equal struct {
	A, B, If, Else Term
}

func (e Equal) E(x float64) float64 {
	if e.A.E(x) == e.B.E(x) {
		return e.If.E(x)
	} else {
		return e.Else.E(x)
	}
}

func (e Equal) Dx() Term {
	return Equal{
		A:    e.A,
		B:    e.B,
		If:   e.If.Dx(),
		Else: e.Else.Dx(),
	}
}

func (e Equal) T() Term {
	aOk, av := e.A.Is()
	bOk, bv := e.B.Is()

	if aOk && bOk {
		if av == bv {
			return e.If.T()
		} else {
			return e.Else.T()
		}
	}

	return Equal{
		A:    e.A.T(),
		B:    e.B.T(),
		If:   e.If.T(),
		Else: e.Else.T(),
	}
}

func (e Equal) Is() (bool, float64) {
	aOk, av := e.A.Is()
	bOk, bv := e.B.Is()

	if aOk && bOk {
		if av == bv {
			return e.If.Is()
		} else {
			return e.Else.Is()
		}
	}

	return false, 0
}

func (e Equal) Tokenise() Tokens {
	t := Tokens{{id: TidEqual}}
	t = append(t, e.A.Tokenise()...)
	t = append(t, e.B.Tokenise()...)
	t = append(t, e.If.Tokenise()...)
	t = append(t, e.Else.Tokenise()...)
	return t
}

type NotEqual struct {
	A, B, If, Else Term
}

func (e NotEqual) E(x float64) float64 {
	if e.A.E(x) != e.B.E(x) {
		return e.If.E(x)
	} else {
		return e.Else.E(x)
	}
}

func (e NotEqual) Dx() Term {
	return NotEqual{
		A:    e.A,
		B:    e.B,
		If:   e.If.Dx(),
		Else: e.Else.Dx(),
	}
}

func (e NotEqual) T() Term {
	aOk, av := e.A.Is()
	bOk, bv := e.B.Is()

	if aOk && bOk {
		if av != bv {
			return e.If.T()
		} else {
			return e.Else.T()
		}
	}

	return NotEqual{
		A:    e.A.T(),
		B:    e.B.T(),
		If:   e.If.T(),
		Else: e.Else.T(),
	}
}

func (e NotEqual) Is() (bool, float64) {
	aOk, av := e.A.Is()
	bOk, bv := e.B.Is()

	if aOk && bOk {
		if av != bv {
			return e.If.Is()
		} else {
			return e.Else.Is()
		}
	}

	return false, 0
}

func (e NotEqual) Tokenise() Tokens {
	t := Tokens{{id: TidNotEqual}}
	t = append(t, e.A.Tokenise()...)
	t = append(t, e.B.Tokenise()...)
	t = append(t, e.If.Tokenise()...)
	t = append(t, e.Else.Tokenise()...)
	return t
}
