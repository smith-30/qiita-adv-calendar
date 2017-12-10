package main

import (
	"os"

	"github.com/smith-30/qiita-adv-calendar/helper/env"
)

func main() {
	env.LoadEnv()

	cli := &CLI{}
	os.Exit(cli.Run(os.Args))
}
