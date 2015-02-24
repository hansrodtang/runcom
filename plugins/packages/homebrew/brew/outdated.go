package brew

func Outdated() []string {
	s, _ := list(brewCommand, "outdated")
	return s
}
