package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	tracker "github.com/teleivo/nti"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stdout, "Failed due to: %s\n", err)
		os.Exit(1)
	}
}

func run(args []string, out io.Writer) error {
	fs := flag.NewFlagSet(args[0], flag.ContinueOnError)
	end := fs.String("end", "now", "Suffix to print in greeting")
	_ = end
	err := fs.Parse(args[1:])
	if err != nil {
		return err
	}

	im, err := tracker.NewImport()
	if err != nil {
		return err
	}
	enc := json.NewEncoder(out)
	err = enc.Encode(im)
	if err != nil {
		return err
	}

	return nil
}

func concat(a, b string) string {
	return a + b
}
