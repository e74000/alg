package alg

// Term defines a tree data type that represents algebra
// E() recursively evaluates the tern
// Dx() recursively returns the derivative of the term
// T() Simplifies the term, it is called after each DDx() call,
// Is() Recursively works out if the term is equivalent to a number,
type Term interface {
	E(x float64) float64
	Dx() Term
	T() Term
	Is() (bool, float64)
	Tokenise() Tokens
}
