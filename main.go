package main

import (
	"bufio"
	"flag"
	"io"
	"os"
	"strings"

	"github.com/seiflotfy/cuckoofilter"
)

var (
	iflag = flag.Bool("i", false, "Case insensitive comparison of lines.")
)

func init() {
	flag.Parse()
}

func main() {
	cf := cuckoo.NewFilter(1000000)

	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			os.Stderr.WriteString(err.Error())
			os.Exit(2)
		}

		cmptxt := text
		if *iflag {
			cmptxt = []byte(strings.ToLower(string(text)))
		}

		if !cf.Lookup(cmptxt) {
			os.Stdout.Write(text)
		}

		cf.InsertUnique(cmptxt)
	}

}
