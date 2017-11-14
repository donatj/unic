package main

import (
	"bufio"
	"io"
	"os"

	"github.com/seiflotfy/cuckoofilter"
)

func main() {
	cf := cuckoofilter.NewDefaultCuckooFilter()

	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			os.Stderr.WriteString(err.Error())
			os.Exit(2)
		}

		if !cf.Lookup(text) {
			os.Stdout.Write(text)
		}

		cf.InsertUnique(text)
	}

}
