package src

func wrapStr(str string) string {
	if len(str) == 1 {
		return " " + str + " "
	}
	return str
}
