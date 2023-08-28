package main

import (
	"fmt"
	"github.com/wrkode/kasba/cmd"
)

var Version = ""

func main() {
	fmt.Println(Version)
	cmd.Execute()
}
