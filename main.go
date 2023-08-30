package main

import (
	"flag"
	"fmt"
	"github.com/wrkode/kasba/cmd"
	"github.com/wrkode/kasba/internal/util"
	"os"
)

func main() {
	flag.Parse()

	if *util.VersionFlag {
		fmt.Println("Version:", cmd.Version)
		os.Exit(0)
	}
	cmd.Run()
}
