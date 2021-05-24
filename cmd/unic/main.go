package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/donatj/unic"
)

var (
	iflag = flag.Bool("i", false, "Case insensitive comparison of lines.")
	dflag = flag.Bool("d", false, "Only output lines that are repeated in the input.")

	vflag = flag.Bool("version", false, "Show build info.")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println(buildString)
	}

	flag.Parse()

	if *vflag {
		fmt.Println(buildString)
		os.Exit(0)
	}
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
