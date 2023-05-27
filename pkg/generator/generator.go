package generator

import (
	_ "embed"
	"fmt"
	"math/rand"
	"strings"
	"unicode"

	"github.com/shalimski/passphrase/pkg/data"
)

const (
	defaultCount     int = 4
	defaultSeparator     = ""
)

func Run(cfg *Config) {

	result := make([]string, 0, cfg.Count)

	for i := range data.Words {
		if i == cfg.Count {
			break
		}

		result = append(result, GetRandomWord(data.Words[i]))
	}

	fmt.Println(strings.Join(result, cfg.Separator))
}

func GetRandomWord(words []string) string {

	if len(words) == 0 {
		return ""
	}

	word := words[rand.Intn(len(words))]

	if len(word) == 0 {
		return ""
	}

	r := unicode.ToUpper(rune(word[0]))
	return string(r) + word[1:]
}

type Config struct {
	Separator string
	Count     int
}

func NewConfig(sep *string, count *int) *Config {

	var c int = defaultCount
	if count != nil && *count >= 1 && *count <= 4 {
		c = *count
	}

	var s string = defaultSeparator
	if sep != nil && len(*sep) <= 100 {
		s = *sep
	}

	return &Config{
		Separator: s,
		Count:     c,
	}
}
