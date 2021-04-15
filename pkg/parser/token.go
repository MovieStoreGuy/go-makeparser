package parser

type Token uint8

const (
	Illegal Token = iota
	EOF
	WhiteSpace

	Indent
	Comment

	Var
	Tar
)
