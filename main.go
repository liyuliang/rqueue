package main

import (
	"github.com/liyuliang/rqueue/route"
	"github.com/liyuliang/rqueue/system"
	"github.com/liyuliang/utils/format"
	"flag"
	"fmt"
	"os"
)

func main() {

	system.Init(format.ToMap(map[string]string{
		system.SystemRedisUri: u,
		system.SystemPopNum:   format.IntToStr(n),
		system.SystemTplDir:   tplDir,
		system.SystemIsDebug:  debug,

		//"uuidNum":  format.IntToStr(un),
	}))

	route.Start(p)
}

var (
	u      string
	p      string
	n      int
	un     int
	tplDir string
	debug  string
)

func init() {

	var required []string

	flag.StringVar(&p, "p", "8888", "web port")
	flag.StringVar(&u, "u", "redis://127.0.0.1:6379/0", "using the redis -u <uri> option and a valid URI")
	flag.StringVar(&tplDir, "tplDir", "tpl", "spider tpl directory")
	flag.IntVar(&n, "n", 50, "default queue pop number")
	flag.IntVar(&un, "un", 50, "uuid pool size")
	flag.StringVar(&debug, "debug", "false", "is running debug module")

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
