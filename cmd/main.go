package main

import (
	"flag"

	"github.com/shalimski/passphrase/pkg/generator"
)

func main() {

	sep := flag.String("s", "", "words separator. (default '')")
	count := flag.Int("c", 4, "count of words. May be from 1 to 4.")
	flag.Parse()

	cfg := generator.NewConfig(sep, count)
	generator.Run(cfg)

}
