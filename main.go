package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	name   string
	pathMd string
)

func init() {
	flag.StringVar(&name, "name", "nil", "module name - required")
	flag.StringVar(&pathMd, "path", "nil", "module path")
	flag.Parse()

	if name == "nil" {
		fmt.Println("Module name is not defined")
		flag.Usage()
		os.Exit(1)
	}
	if pathMd == "nil" || pathMd == "." {
		pathMd = ""
	}
}

func main() {
	maker, err := NewModuleCreator(name, pathMd)
	if err != nil {
		fmt.Println("Create Module Error:", err.Error())
		return
	}
	errs := maker.ExecAsync()

	for _, err := range errs {
		if err == nil {
			continue
		}
		fmt.Println("Error: ", err.Error())
	}

	fmt.Println("*** Exit program ***")
}
