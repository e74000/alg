package alg

// Term defines a tree data type that represents algebra
// E() recursively evaluates the tern
// Dx() recursively returns the derivative of the term
// T() Simplifies the term, it is called after each DDx() call,
// Is() Recursively works out if the term is equivalent to a number,
// Tokenise() Recursively converts the tree to a list of tokens in prefix notation
type Term interface {
	E(x float64) float64
	Dx() Term
	T() Term
	Is() (bool, float64)
	Tokenise() Tokens
}

// Terms is a list of all the different term types.
// It is included to allow you to gob encode terms.
var Terms = []Term{
	new(S),
	new(X),
	new(Sx),
	new(Exp),
	new(Ln),
	new(TPT),
	new(TP),
	new(PT),
	new(Sum),
	new(Prod),
	new(Div),
	new(Add),
	new(Sub),
	new(Mul),
	new(Sin),
	new(Cos),
	new(Tan),
	new(Sec),
	new(Csc),
	new(Cot),
	new(Sinh),
	new(Cosh),
	new(Tanh),
	new(Sech),
	new(Csch),
	new(Coth),
	new(Greater),
	new(Less),
	new(GreaterEqual),
	new(LessEqual),
	new(Equal),
	new(NotEqual),
	new(Range),
}
