package parser

// Target is abstraction of make file instruction
type Target struct {
	dependsOn []*Target
	recipes   []string
	lineNum   int
	err       error
}

func (t *Target) Err() error      { return t.err }
func (t *Target) LineNumber() int { return t.lineNum }

func (t *Target) Recipes() []string {
	recipes := make([]string, 0, len(t.recipes))
	copy(t.recipes, recipes)
	return recipes
}

func (t *Target) DependsOn() []*Target {
	depends := make([]*Target, 0, len(t.dependsOn))
	copy(t.dependsOn, depends)
	return depends
}
