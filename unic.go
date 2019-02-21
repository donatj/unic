package unic

import (
	"bufio"
	"io"
	"strings"

	multierror "github.com/hashicorp/go-multierror"
	cuckoo "github.com/seiflotfy/cuckoofilter"
)

// Filter is a unique filter utilizing Cuckoo Filters
type Filter struct {
	CaseI          bool
	FilterCapacity uint
}

// FilterOption sets an option of the passed Filter
type FilterOption func(*Filter) error

// FilterCaseInsensitive configures the Filter to be Case Insensitive
func FilterCaseInsensitive(f *Filter) error {
	f.CaseI = true
	return nil
}

// NewFilter returns a Filter configured with the given FilterOptions
func NewFilter(options ...FilterOption) (*Filter, error) {
	filter := &Filter{}

	var result *multierror.Error

	for _, option := range options {
		err := option(filter)
		result = multierror.Append(result, err)
	}

	return filter, result.ErrorOrNil()
}

// Exec executes the filter on the given input
// writing unique output to the unique stream and
// repeated output to the repeated stream
func (u *Filter) Exec(input io.Reader, unique, repeated io.Writer) error {
	cf := cuckoo.NewFilter(u.FilterCapacity)
	cf2 := cuckoo.NewFilter(u.FilterCapacity)

	reader := bufio.NewReader(input)

	for {
		text, err := reader.ReadBytes('\n')
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		cmptxt := text
		if u.CaseI {
			cmptxt = []byte(strings.ToLower(string(text)))
		}

		if !cf.Lookup(cmptxt) {
			unique.Write(text)
		} else {
			if !cf2.Lookup(cmptxt) {
				repeated.Write(text)
			}

			cf2.InsertUnique(cmptxt)
		}

		cf.InsertUnique(cmptxt)
	}
}
