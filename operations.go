package editdistance

type (
	Match struct {
		Cost int
	}

	Insertion struct {
		Cost int
	}

	Deletion struct {
		Cost int
	}

	Substitution struct {
		Cost int
	}

	Transposition struct {
		Cost int
	}
)

var (
	_ EditOperation = Match{}
	_ EditOperation = Insertion{}
	_ EditOperation = Deletion{}
	_ EditOperation = Substitution{}
	_ EditOperation = Transposition{}
)

func (o Match) Apply(data SequencePair, matrix Matrix, i, j int) (int, bool) {
	if i > 0 && j > 0 {
		if data == nil {
			return matrix[i - 1][j - 1] + o.Cost, true
		}

		if data.Equal(i - 1, j - 1) {
			return matrix[i - 1][j - 1] + o.Cost, true
		}
	}

	return 0, false
}

func (o Match) Backtrack(matrix Matrix, i, j int) (int, int) {
	return i - 1, j - 1
}

func (o Match) String() string {
	return "match"
}

func (o Insertion) Apply(data SequencePair, matrix Matrix, i, j int) (int, bool) {
	if j > 0 {
		return matrix[i][j-1] + o.Cost, true
	}

	return 0, false
}

func (o Insertion) Backtrack(matrix Matrix, i, j int) (int, int) {
	return i, j - 1
}

func (o Insertion) String() string {
	return "ins"
}

func (o Deletion) Apply(data SequencePair, matrix Matrix, i, j int) (int, bool) {
	if i > 0 {
		return matrix[i-1][j] + o.Cost, true
	}

	return 0, false
}

func (o Deletion) Backtrack(matrix Matrix, i, j int) (int, int) {
	return i - 1, j
}

func (o Deletion) String() string {
	return "del"
}

func (o Substitution) Apply(data SequencePair, matrix Matrix, i, j int) (int, bool) {
	if i > 0 && j > 0 {
		return matrix[i-1][j-1] + o.Cost, true
	}

	return 0, false
}

func (o Substitution) Backtrack(matrix Matrix, i, j int) (int, int) {
	return i - 1, j - 1
}

func (o Substitution) String() string {
	return "del"
}

func (o Transposition) Apply(data SequencePair, matrix Matrix, i, j int) (int, bool) {
	if i > 1 && j > 1 {
		if data == nil {
			return matrix[i-2][j-2] + o.Cost, true
		}

		if data.Equal(i-1, j-2) && data.Equal(i-2, j-1) {
			return matrix[i-2][j-2] + o.Cost, true
		}
	}

	return 0, false
}

func (o Transposition) Backtrack(matrix Matrix, i, j int) (int, int) {
	return i - 2, j - 2
}

func (o Transposition) String() string {
	return "trp"
}
