package main

import (
	"bufio"
	"os"

	"github.com/seiflotfy/cuckoofilter"
)

func main() {
	cf := cuckoofilter.NewDefaultCuckooFilter()

	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}

		if !cf.Lookup(text) {
			os.Stdout.Write(text)
		}

		cf.InsertUnique(text)
	}

}
