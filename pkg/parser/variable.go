package parser

type Variable struct {
	lineNum int
	err     error
	name    string
	value   string
}

func (v *Variable) Err() error      { return v.err }
func (v *Variable) LineNumber() int { return v.lineNum }
func (v *Variable) Name() string    { return v.name }
func (v *Variable) Value() string   { return v.value }
