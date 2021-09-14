package pgkit

func SuffixRawLocking(isRowLocking bool)  string {
	if isRowLocking {
		return "FOR UPDATE"
	}
	return ""
}

