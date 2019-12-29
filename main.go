package main

import (
	"github.com/liyuliang/rqueue/route"
	"github.com/liyuliang/rqueue/system"
	"flag"
	"fmt"
	"os"
)

func main() {

	system.Init(u)
	route.Start(p)
}

var (
	u string
	p string
)

func init() {

	var required []string

	flag.StringVar(&p, "p", "8888", "web port")
	flag.StringVar(&u, "u", "redis://127.0.0.1:6379/0", "using the redis -u <uri> option and a valid URI")

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
