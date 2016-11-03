package editdistance

var DefaultLevenshtein = []EditOperation{
	Match{},
	Insertion{1},
	Deletion{1},
	Substitution{1},
}

var DefaultDamerau = []EditOperation{
	Match{},
	Insertion{1},
	Deletion{1},
	Substitution{1},
	Transposition{1},
}

var DefaultLCS = []EditOperation{
	Match{},
	Insertion{1},
	Deletion{1},
}
