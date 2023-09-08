package parser

func getIndent(line string) int {
	indent := 0
	for _, c := range line {
		switch c {
		case ' ', '\t':
			indent++
		default:
			return indent
		}
	}

	return indent
}

func removeIndent(line string) (string, int) {
	indent := getIndent(line)
	line = line[indent:]

	return line, indent
}
