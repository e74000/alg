package alg

import (
	"reflect"
	"testing"
)

func TestTokenise(t *testing.T) {
	testCases := map[string]Term{
		"^ x 2.00 ":                      TP{X: X{}, P: 2},
		"^ 2.00 x ":                      PT{V: 2, X: X{}},
		"sin sin x ":                     Sin{X: Sin{X: X{}}},
		"+[ * 1.00 2.00 + 3.00 4.00 ]+ ": Sum{Mul{A: S(1), B: S(2)}, Add{A: S(3), B: S(4)}},
	}

	for s, term := range testCases {
		ts := term.Tokenise()
		tss := ts.String()

		if s != tss {
			t.Logf("Test failed on case: (%v)\nWanted: %s\nGot:    %s\n", term, s, tss)
			t.Fail()
		}
	}

}

func TestTokens_Parse(t *testing.T) {
	testCases := map[string]Term{
		"^ x 2.00":                      TPT{A: X{}, B: S(2)},
		"^ 2.00 x":                      TPT{A: S(2), B: X{}},
		"sin sin x":                     Sin{X: Sin{X: X{}}},
		"+[ * 1.00 2.00 + 3.00 4.00 ]+": Sum{Mul{A: S(1), B: S(2)}, Add{A: S(3), B: S(4)}},
		"^ sech x 2":                    TPT{A: Sech{X{}}, B: S(2)},
	}

	for s, term := range testCases {
		ts := Tokenise(s)
		tree := ts.Parse()

		trs := tree.Tokenise()

		if !reflect.DeepEqual(term, tree) {
			t.Logf("Test failed on case: (%s)Wanted: %v :: (%s)\nGot:    %v :: (%s)\n", s, term, ts.String(), tree, trs.String())
			t.Fail()
		}
	}
}

func TestTidy(t *testing.T) {
	messy := []Term{
		Prod{S(0), Add{S(0), Exp{Sx{-1}}}},
		Add{Prod{S(0), X{}}, X{}},
		Prod{Add{S(2), S(1)}},
		Prod{S(1), S(2), S(3), X{}, X{}},
		Div{Add{Prod{S(0), X{}}, X{}}, Sin{X{}}},
		Prod{X{}, X{}},
	}

	tidy := []Term{
		S(0),
		X{},
		S(3),
		Prod{X{}, X{}, S(6)},
		Div{X{}, Sin{X{}}},
		Mul{X{}, X{}},
	}

	for i := 0; i < len(messy); i++ {
		tidied := messy[i].T()

		ms := messy[i].Tokenise()
		ns := tidied.Tokenise()
		ts := tidy[i].Tokenise()

		if !reflect.DeepEqual(tidy[i], tidied) {
			t.Logf("Test failed on case: (%s)\nWanted: %s\nGot     %s\n", ms.String(), ts.String(), ns.String())
			t.Fail()
		}
	}
}

func TestSigmoid(t *testing.T) {
	sigmoid := Div{S(1), Add{S(1), Exp{Sx{-1}}}}
	prime := sigmoid.Dx()
	pTrue := Div{
		N: Exp{Sx{-1}},
		D: Mul{
			A: Add{S(1), Exp{Sx{-1}}},
			B: Add{S(1), Exp{Sx{-1}}},
		},
	}

	if !reflect.DeepEqual(prime, pTrue) {
		pts := prime.Tokenise()
		tts := pTrue.Tokenise()

		t.Logf("Sigmoid derivative has failed!\n %s != %s\n", pts.String(), tts.String())
		t.Fail()
	}
}
