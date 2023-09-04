package main

import (
	"flag"
	"fmt"
	"github.com/wrkode/kasba/cmd"
	"github.com/wrkode/kasba/internal/util"
	"os"
)

var Version = ""

func main() {
	flag.Parse()

	if *util.VersionFlag {
		fmt.Println("Version:", Version)
		os.Exit(0)
	}
	cmd.Run()
}
