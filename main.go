package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/hcl"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	f := os.Stdin
	if len(os.Args) > 1 {
		var err error
		f, err = os.Open(os.Args[1])
		if err != nil {
			return fmt.Errorf("open file error, %v", err)
		}
	}
	b, err :=ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	_, err = hcl.ParseBytes(b)
	return err
}

