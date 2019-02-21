package main

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/donatj/unic"
)

var (
	iflag = flag.Bool("i", false, "Case insensitive comparison of lines.")
	dflag = flag.Bool("d", false, "Only output lines that are repeated in the input.")
)

func init() {
	flag.Parse()
}

func main() {
	options := []unic.FilterOption{}

	if *iflag {
		options = append(options, unic.FilterCaseInsensitive)
	}

	filter, err := unic.NewFilter(options...)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(2)
	}

	{
		var err error
		if !*dflag {
			err = filter.Exec(os.Stdin, os.Stdout, ioutil.Discard)
		} else {
			err = filter.Exec(os.Stdin, ioutil.Discard, os.Stdout)
		}

		if err != nil {
			os.Stderr.WriteString(err.Error() + "\n")
			os.Exit(2)
		}
	}
}
