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

// FilterCapacity sets the cuckoo filter capacity for the Filter's
// internal cuckoo filters
func FilterCapacity(capacity uint) FilterOption {
	return func(f *Filter) error {
		f.FilterCapacity = capacity
		return nil
	}
}

// NewFilter returns a Filter configured with the given FilterOptions
func NewFilter(options ...FilterOption) (*Filter, error) {
	filter := &Filter{
		FilterCapacity: 1000000,
	}

	var result *multierror.Error

	for _, option := range options {
		err := option(filter)
		result = multierror.Append(result, err)
	}

	return filter, result.ErrorOrNil()
}

// Exec executes the filter on the given input.
// Writes unique output to the unique stream.
// Writes repeated output to the repeated stream.
func (u *Filter) Exec(input io.Reader, unique, repeated io.Writer) error {
	cf := cuckoo.NewFilter(u.FilterCapacity)
	cf2 := cuckoo.NewFilter(u.FilterCapacity)

	reader := bufio.NewReader(input)

	for {
		text, readErr := reader.ReadBytes('\n')
		if readErr == io.EOF {
			if len(text) == 0 {
				return nil
			}

			text = append(text, '\n')
		} else if readErr != nil {
			return readErr
		}

		cmptxt := text
		if u.CaseI {
			cmptxt = []byte(strings.ToLower(string(text)))
		}

		if !cf.Lookup(cmptxt) {
			_, err := unique.Write(text)
			if err != nil {
				return err
			}
		} else {
			if !cf2.Lookup(cmptxt) {
				_, err := repeated.Write(text)
				if err != nil {
					return err
				}
			}

			cf2.InsertUnique(cmptxt)
		}

		cf.InsertUnique(cmptxt)

		if readErr == io.EOF {
			return nil
		}
	}
}
