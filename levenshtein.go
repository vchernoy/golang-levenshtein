package editdistance

import (
	"fmt"
	"io"
	"math"
	"os"
)

// DistanceForStrings returns the edit distance between source and target.
func DistanceForStrings(source []rune, target []rune, ops []EditOperation) int {
	return DistanceFor(Runes{Source: source, Target: target}, ops)
}

// DistanceForStrings returns the edit distance between source and target.
func DistanceFor(p SequencePair, ops []EditOperation) int {
	return DistanceForMatrix(MatrixFor(p, ops))
}

// DistanceForMatrix reads the edit distance off the given Levenshtein matrix.
func DistanceForMatrix(matrix [][]int) int {
	return matrix[len(matrix)-1][len(matrix[0])-1]
}

// MatrixForStrings generates a 2-D array representing the dynamic programming
// table used by the Levenshtein algorithm, as described e.g. here:
// http://www.let.rug.nl/kleiweg/lev/
// The reason for putting the creation of the table into a separate function is
// that it cannot only be used for reading of the edit distance between two
// strings, but also e.g. to backtrace an edit script that provides an
// alignment between the characters of both strings.
func MatrixForStrings(source []rune, target []rune, ops []EditOperation) [][]int {
	return MatrixFor(Runes{Source: source, Target: target}, ops)
}

func MatrixFor(p SequencePair, ops []EditOperation) [][]int {
	// Make a 2-D matrix. Rows correspond to prefixes of source, columns to
	// prefixes of target. Cells will contain edit distances.
	// Cf. http://www.let.rug.nl/~kleiweg/lev/levenshtein.html
	height := p.SourceLen() + 1
	width := p.TargetLen() + 1
	matrix := make([][]int, height)

	// Initialize trivial distances (from/to empty string). That is, fill
	// the left column and the top row with row/column indices.
	for i := 0; i < height; i++ {
		matrix[i] = make([]int, width)
		matrix[i][0] = i
	}
	for j := 1; j < width; j++ {
		matrix[0][j] = j
	}

	// Fill in the remaining cells: for each prefix pair, choose the
	// (edit history, operation) pair with the lowest cost.
	for i := 1; i < height; i++ {
		for j := 1; j < width; j++ {
			lowestCost := math.MaxInt32
			for _, op := range ops {
				if cost, ok := op.Apply(p, matrix, i, j); ok && cost < lowestCost {
					lowestCost = cost
				}
			}

			matrix[i][j] = lowestCost
		}
	}
	//LogMatrix(source, target, matrix)
	return matrix
}

// EditScriptForStrings returns an optimal edit script to turn source into
// target.
func EditScriptForStrings(source []rune, target []rune, ops []EditOperation) EditScript {
	return EditScriptFor(Runes{Source: source, Target: target}, ops)
}

// EditScriptFor returns an optimal edit script to turn source into
// target.
func EditScriptFor(p SequencePair, ops []EditOperation) EditScript {
	return EditScriptForMatrix(MatrixFor(p, ops), ops)
}

// EditScriptForMatrix returns an optimal edit script based on the given
// Levenshtein matrix.
func EditScriptForMatrix(matrix [][]int, ops []EditOperation) EditScript {
	return backtrace(len(matrix[0])-1, len(matrix)-1, matrix, ops)
}

// WriteMatrix writes a visual representation of the given matrix for the given
// strings to the given writer.
func WriteMatrix(source []rune, target []rune, matrix [][]int, writer io.Writer) {
	fmt.Fprintf(writer, "    ")
	for _, targetRune := range target {
		fmt.Fprintf(writer, "  %c", targetRune)
	}
	fmt.Fprintf(writer, "\n")
	fmt.Fprintf(writer, "  %2d", matrix[0][0])
	for j := range target {
		fmt.Fprintf(writer, " %2d", matrix[0][j+1])
	}
	fmt.Fprintf(writer, "\n")
	for i, sourceRune := range source {
		fmt.Fprintf(writer, "%c %2d", sourceRune, matrix[i+1][0])
		for j := range target {
			fmt.Fprintf(writer, " %2d", matrix[i+1][j+1])
		}
		fmt.Fprintf(writer, "\n")
	}
}

func WriteMatrixFor(p StringSequencePair, matrix [][]int, writer io.Writer) {
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

// LogMatrix writes a visual representation of the given matrix for the given
// strings to os.Stderr. This function is deprecated, use
// WriteMatrix(source, target, matrix, os.Stderr) instead.
func LogMatrix(source []rune, target []rune, matrix [][]int) {
	WriteMatrix(source, target, matrix, os.Stderr)
}

func backtrace(i int, j int, matrix [][]int, ops []EditOperation) EditScript {
	for _, op := range ops {
		ib, jb := op.Backtrack(matrix, i, j)

		if ib < 0 || jb < 0 {
			continue
		}

		if cost, ok := op.Apply(nil, matrix, ib, jb); ok && cost == matrix[i][j] {
			return append(backtrace(ib, jb, matrix, ops), op)
		}
	}

	return EditScript{}
}
