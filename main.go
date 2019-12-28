package main

import (
	"github.com/liyuliang/rqueue/route"
	"github.com/liyuliang/rqueue/system"
	"flag"
	"fmt"
	"os"
)

func main() {

	system.Init(redis)
	route.Start(p)
}

var (
	redis string
	p     string
)

func init() {
	required := []string{"redis", "p",}

	flag.StringVar(&p, "p", "8888", "web port")
	flag.StringVar(&redis, "redis", "", "redis connect address")

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
