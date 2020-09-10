package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hashicorp/hcl"
)

const hclExt = ".hcl"

func main() {
	if err := run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}

func run(arg []string) error {
	opt := NewOption()
	paths, err := opt.Parse(arg)
	if err != nil {
		return err
	}
	for _, v := range  paths{
		if err := check(v, opt.Verbose); err != nil {
			return err
		}
	}
	return nil
}

func check(dir string, verbose bool)  error {
	return filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if filepath.Ext(path) == hclExt {
			f, err :=os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()
			if err := parseHCL(f); err != nil {
				if verbose {
					fmt.Print("❎ ")
				}
				fmt.Printf("%s:%v\n",f.Name(), err)
			} else if verbose{
				fmt.Printf("☑️ %s\n", path)
			}
		}
		return nil
	})
}


func parseHCL(r io.Reader) error {
	b, err :=ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	_, err = hcl.ParseBytes(b)
	return err
}

