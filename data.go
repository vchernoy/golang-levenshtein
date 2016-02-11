package editdistance

type SequencePair interface {
	SourceLen() int
	TargetLen() int
	Equal(i, j int) bool
}

type StringSequencePair interface {
	SequencePair
	SourceAt(i int) string
	TargetAt(i int) string
}

type Runes struct {
	Source []rune
	Target []rune
}

var _ StringSequencePair = Runes{}

func (r Runes) SourceLen() int {
	return len(r.Source)
}
func (r Runes) TargetLen() int {
	return len(r.Target)
}
func (r Runes) Equal(i, j int) bool {
	return r.Source[i] == r.Target[j]
}
func (r Runes) SourceAt(i int) string {
	return string(r.Source[i])
}
func (r Runes) TargetAt(i int) string {
	return string(r.Target[i])
}

type Words struct {
	Source []string
	Target []string
}

var _ StringSequencePair = Words{}

func (w Words) SourceLen() int {
	return len(w.Source)
}
func (w Words) TargetLen() int {
	return len(w.Target)
}
func (w Words) Equal(i, j int) bool {
	return w.Source[i] == w.Target[j]
}
func (w Words) SourceAt(i int) string {
	return w.Source[i]
}
func (w Words) TargetAt(i int) string {
	return w.Target[i]
}
