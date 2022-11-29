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

		if !reflect.DeepEqual(term, tree) {
			t.Logf("Test failed on case: (%s)\nWanted: %v\nGot:    %v\n", s, term, tree)
			t.Fail()
		}
	}
}
