package data

import (
	_ "embed"
	"strings"
)

const endline = "\n"

var (
	//go:embed sub_nouns.txt
	sub_nouns string

	//go:embed verbs.txt
	verbs string

	//go:embed adjectives.txt
	adjectives string

	//go:embed obj_nouns.txt
	obj_nouns string

	Words = [4][]string{
		strings.Split(sub_nouns, endline),
		strings.Split(verbs, endline),
		strings.Split(adjectives, endline),
		strings.Split(obj_nouns, endline),
	}
)
