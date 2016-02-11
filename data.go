package editdistance

import (
	"fmt"
	"io"
)

type Matrix [][]int

// Distance reads the edit distance off the Levenshtein matrix.
func (matrix Matrix) Distance() int {
	return matrix[len(matrix)-1][len(matrix[0])-1]
}

func (matrix Matrix) Write(p StringSequencePair, writer io.Writer) {
	fmt.Fprintf(writer, "    ")
	for j := 0; j < p.TargetLen(); j++ {
		targetRune := p.TargetAt(j)
		fmt.Fprintf(writer, "  %s", targetRune)
	}
	fmt.Fprintf(writer, "\n")
	fmt.Fprintf(writer, "  %2d", matrix[0][0])
	for j := 0; j < p.TargetLen(); j++ {
		fmt.Fprintf(writer, " %2d", matrix[0][j+1])
	}
	fmt.Fprintf(writer, "\n")
	for i := 0; i < p.SourceLen(); i++ {
		sourceRune := p.SourceAt(i)
		fmt.Fprintf(writer, "%s %2d", sourceRune, matrix[i+1][0])
		for j := 0; j < p.TargetLen(); j++ {
			fmt.Fprintf(writer, " %2d", matrix[i+1][j+1])
		}
		fmt.Fprintf(writer, "\n")
	}
}

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

func NewRunes(source, target string) Runes {
	return Runes{
		Source: []rune(source),
		Target: []rune(target),
	}
}

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
