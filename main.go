package main

import (
	"github.com/liyuliang/rqueue/route"
	"flag"
	"fmt"
	"os"
)

func main() {

	route.Start(p)
}

var (
	a string
	p string
	g string
)

func init() {
	required := []string{"a", "g"}

	flag.StringVar(&a, "a", "", "auth token")
	flag.StringVar(&g, "g", "", "gateway url")
	flag.StringVar(&p, "p", "8888", "web port")

	flag.Parse()

	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })

	for _, req := range required {

		if !seen[req] {
			fmt.Fprintf(os.Stderr, "flag -%s is required \n", req)
			os.Exit(2)
		}
	}
}
