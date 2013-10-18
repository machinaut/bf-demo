package main

import (
	"github.com/machinaut/bf"
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hellogod")
	flag.Parse()
	if (flag.NArg() < 1) {
		fmt.Println("Needs file to run on command line.  Bugger off.")
		return;
	}
	prog, _ := os.Open(flag.Arg(0))
	bfi := bf.NewBFInterp(prog, os.Stdin, os.Stdout)
	bfi.Run()
}

