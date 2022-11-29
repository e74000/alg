package alg

/*
Arithmetic defines all the simple arithmetic datatypes:
  Sum  => The sum of a list of terms
  Prod => The product of a list of terms
  Div  => A fraction between two terms
  Add  => Two terms added together
  Sub  => Subtract one term from another
  Mul  => Multiply two terms
*/

type Sum []Term

func (e Sum) Tokenise() Tokens {
	t := Tokens{{id: TidSumSt}}

	for _, term := range e {
		t = append(t, term.Tokenise()...)
	}

	t = append(t, Token{id: TidSumEn})

	return t
}

func (e Sum) E(x float64) float64 {
	var sum float64
	for _, term := range e {
		sum += term.E(x)
	}
	return sum
}

func (e Sum) Dx() Term {
	next := make(Sum, len(e))

	for i, term := range e {
		next[i] = term.Dx()
	}

	return next.T()
}

func (e Sum) T() Term {
	flat := e.flatten()
	sum := make(Sum, 0, len(flat))

	var total float64
	for _, term := range flat {
		ok, val := term.Is()
		if ok {
			total += val
		} else {
			sum = append(sum, term.T())
		}
	}

	if len(sum) == 0 {
		return S(total)
	}

	if total != 0 {
		sum = append(sum, S(total))
	}

	return sum
}

func (e Sum) Is() (bool, float64) {
	var cs float64
	for _, term := range e {
		ok, val := term.Is()
		if !ok {
			return false, 0
		}

		cs += val
	}

	return true, cs
}

func (e Sum) flatten() Sum {
	s1 := make(Sum, len(e))

	copy(s1, e)

	changed := true

	for changed {
		s2 := make(Sum, 0)
		changed = false
		for _, term := range s1 {
			if s, ok := term.(Sum); ok {
				s2 = append(s2, s...)
				changed = true
			} else if s, ok := term.(Add); ok {
				s2 = append(s2, s.A, s.B)
				changed = true
			} else if s, ok := term.(Sub); ok {
				s2 = append(s2, s.A, Mul{S(-1), s.B})
			} else {
				s2 = append(s2, term)
			}
		}

		s1 = make(Sum, len(s2))
		copy(s1, s2)
	}

	return s1
}

type Prod []Term

func (e Prod) Tokenise() Tokens {
	t := Tokens{{id: TidProdSt}}

	for _, term := range e {
		t = append(t, term.Tokenise()...)
	}

	t = append(t, Token{id: TidProdEn})

	return t
}

func (e Prod) E(x float64) float64 {
	var prod float64 = 1
	for _, term := range e {
		prod *= term.E(x)
	}
	return prod
}

func (e Prod) Dx() Term {
	sum := make(Sum, len(e))

	for i, iTerm := range e {
		prod := make(Prod, len(e))
		prod[0] = iTerm.Dx()
		count := 1
		for j := 0; j < len(e); j++ {
			if i == j {
				continue
			}

			prod[count] = e[j]
		}

		sum[i] = prod.T()
	}

	return sum.T()
}

func (e Prod) T() Term {
	flat := e.flatten()

	prod := make(Prod, 0, len(e))

	var total float64 = 1
	for _, term := range flat {
		ok, val := term.Is()
		if ok {
			total *= val
		} else {
			prod = append(prod, term.T())
		}
	}

	if total == 0 {
		return S(0)
	}

	if len(prod) == 0 {
		return S(total)
	}

	if total != 1 {
		prod = append(prod, S(total))
	}

	return prod
}

func (e Prod) Is() (bool, float64) {
	var sc float64 = 1
	stop := false

	for _, term := range e {
		ok, val := term.Is()
		if ok {
			sc *= val
		} else {
			stop = true
		}
	}

	if sc == 0 {
		return true, 0
	}

	if stop {
		return false, 0
	}

	return true, sc
}

func (e Prod) flatten() Prod {
	p1 := make(Prod, len(e))

	copy(p1, e)

	changed := true

	for changed {
		p2 := make(Prod, 0)
		changed = false
		for _, term := range p1 {
			if p, ok := term.(Sum); ok {
				p2 = append(p2, p...)
				changed = true
			} else if p, ok := term.(Mul); ok {
				p2 = append(p2, p.A, p.B)
				changed = true
			} else {
				p2 = append(p2, term)
			}
		}

		p1 = make(Prod, len(p2))
		copy(p1, p2)
	}

	return p1
}

type Div struct {
	N, D Term
}

func (e Div) Tokenise() Tokens {
	t := Tokens{{id: TidDiv}}
	t = append(t, e.N.Tokenise()...)
	t = append(t, e.D.Tokenise()...)
	return t
}

func (e Div) E(x float64) float64 {
	return e.N.E(x) / e.D.E(x)
}

func (e Div) Dx() Term {
	tidied := e.T()

	switch tidied.(type) {
	case Prod, S:
		return tidied.Dx()
	}

	return Div{
		N: Sum{
			Prod{e.N.Dx(), e.D},
			Prod{S(-1), e.N, e.D.Dx()},
		},
		D: Prod{e.D, e.D},
	}.T()
}

func (e Div) T() Term {
	nOk, nVal := e.N.Is()
	dOk, dVal := e.D.Is()

	if nOk && dOk {
		return S(nVal / dVal)
	} else if nOk {
		return Div{S(nVal), e.D}
	} else if dOk {
		return Prod{e.N, S(1 / dVal)}
	}

	return Div{
		N: e.N.T(),
		D: e.D.T(),
	}
}

func (e Div) Is() (bool, float64) {
	nOk, nVal := e.N.Is()
	dOk, dVal := e.D.Is()

	if nOk && dOk {
		return true, nVal / dVal
	}

	return false, 0
}

type Add struct {
	A Term
	B Term
}

func (e Add) E(x float64) float64 {
	return e.A.E(x) + e.B.E(x)
}

func (e Add) Dx() Term {
	return Add{
		e.A.Dx(),
		e.B.Dx(),
	}.T()
}

func (e Add) T() Term {
	aok, av := e.A.Is()
	bok, bv := e.B.Is()

	if aok && bok {
		return S(av + bv)
	} else if aok && av == 0 {
		return e.B
	} else if bok && bv == 0 {
		return e.A
	}

	return Add{e.A.T(), e.B.T()}
}

func (e Add) Is() (bool, float64) {
	aok, av := e.A.Is()
	bok, bv := e.B.Is()

	if aok && bok {
		return true, av + bv
	}

	return false, 0
}

func (e Add) Tokenise() Tokens {
	t := Tokens{{id: TidAdd}}
	t = append(t, e.A.Tokenise()...)
	t = append(t, e.B.Tokenise()...)
	return t
}

type Sub struct {
	A Term
	B Term
}

func (e Sub) E(x float64) float64 {
	return e.A.E(x) - e.B.E(x)
}

func (e Sub) Dx() Term {
	return Sub{
		e.A.Dx(),
		e.B.Dx(),
	}.T()
}

func (e Sub) T() Term {
	aok, av := e.A.Is()
	bok, bv := e.B.Is()

	if aok && bok {
		return S(av - bv)
	} else if aok && av == 0 {
		return Mul{S(-1), e.B}
	} else if bok && bv == 0 {
		return e.A
	}

	return Sub{e.A.T(), e.B.T()}
}

func (e Sub) Is() (bool, float64) {
	aok, av := e.A.Is()
	bok, bv := e.B.Is()

	if aok && bok {
		return true, av - bv
	}

	return false, 0
}

func (e Sub) Tokenise() Tokens {
	t := Tokens{{id: TidSub}}
	t = append(t, e.A.Tokenise()...)
	t = append(t, e.B.Tokenise()...)
	return t
}

type Mul struct {
	A Term
	B Term
}

func (e Mul) E(x float64) float64 {
	return e.A.E(x) * e.B.E(x)
}

func (e Mul) Dx() Term {
	return Add{
		Mul{e.A.Dx(), e.B},
		Mul{e.B.Dx(), e.A},
	}.T()
}

func (e Mul) T() Term {
	aok, av := e.A.Is()
	bok, bv := e.B.Is()

	if aok && bok {
		return S(av + bv)
	} else if aok && av == 1 {
		return e.B
	} else if bok && bv == 1 {
		return e.A
	} else if aok && av == 0 || bok && bv == 0 {
		return S(0)
	}

	return Mul{e.A.T(), e.B.T()}
}

func (e Mul) Is() (bool, float64) {
	aok, av := e.A.Is()
	bok, bv := e.B.Is()

	if aok && bok {
		return true, av + bv
	} else if aok && av == 0 || bok && bv == 0 {
		return true, 0
	}

	return false, 0
}

func (e Mul) Tokenise() Tokens {
	t := Tokens{{id: TidMul}}
	t = append(t, e.A.Tokenise()...)
	t = append(t, e.B.Tokenise()...)
	return t
}
