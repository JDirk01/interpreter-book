package lexer

type Lexer struct {
	input		string //input
	position	int    //current position in input
	readPosition	inr    //current reading position (pos after current pos).
	ch		byte   //current char
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
