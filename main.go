package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stdout, "Failed due to: %s", err)
		os.Exit(1)
	}
}

func run(args []string, out io.Writer) error {
	fs := flag.NewFlagSet(args[0], flag.ContinueOnError)
	end := fs.String("end", "now", "Suffix to print in greeting")
	fs.Parse(args[1:])
	fmt.Fprint(out, concat("Go build it ", *end))
	return nil
}

func concat(a, b string) string {
	return a + b
}
