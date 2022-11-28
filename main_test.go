package alg

import "testing"

func TestTokenise(t *testing.T) {
	testCases := map[string]Term{
		"^ x 2.00 ":                      TP{X: X{}, P: 2},
		"^ 2.00 x ":                      PT{V: 2, X: X{}},
		"sin sin x ":                     Sin{Sin{X: X{}}},
		"+[ * 1.00 2.00 + 3.00 4.00 ]+ ": Sum{Mul{S(1), S(2)}, Add{S(3), S(4)}},
	}

	for s, term := range testCases {
		ts := term.Tokenise()

		if s != ts.String() {
			t.Logf("Test failed on case: (%v)\nWanted: %s\nGot:    %s\n", term, s, ts.String())
			t.Fail()
		}
	}

}
