package editdistance

import (
	"math"
	"fmt"
)


type (
	EditScript []EditOperation

	EditOperation interface {
		fmt.Stringer
		Apply(data SequencePair, matrix Matrix, i, j int) (int, bool)
		Backtrack(matrix Matrix, i, j int) (int, int)
	}

	Matrix [][]int

	SequencePair interface {
		SourceLen() int
		TargetLen() int
		Equal(i, j int) bool
	}
)

func NewMatrix(p SequencePair, ops []EditOperation) Matrix {
	// Make a 2-D matrix. Rows correspond to prefixes of source, columns to
	// prefixes of target. Cells will contain edit distances.
	// Cf. http://www.let.rug.nl/~kleiweg/lev/levenshtein.html
	height := p.SourceLen() + 1
	width := p.TargetLen() + 1
	matrix := make(Matrix, height)

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

// Distance reads the edit distance off the Levenshtein matrix.
func (m Matrix) Distance() int {
	return m[len(m)-1][len(m[0])-1]
}

// NewEditScript returns an optimal edit script based on the given
// Levenshtein matrix.
func (m Matrix) EditScript(p SequencePair, ops []EditOperation) EditScript {
	return backtrace(p, len(m)-1, len(m[0])-1, m, ops)
}

func backtrace(p SequencePair, i int, j int, matrix Matrix, ops []EditOperation) EditScript {
	for _, op := range ops {
		ib, jb := op.Backtrack(matrix, i, j)

		if ib < 0 || jb < 0 {
			continue
		}

		if cost, ok := op.Apply(p, matrix, i, j); ok && cost == matrix[i][j] {
			return append(backtrace(p, ib, jb, matrix, ops), op)
		}
	}

	return EditScript{}
}
