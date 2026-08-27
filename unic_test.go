package unic

import (
	"errors"
	"testing"
)

func TestNewFilterJoinsOptionErrors(t *testing.T) {
	first := errors.New("first")
	second := errors.New("second")
	called := 0

	filter, err := NewFilter(
		func(*Filter) error {
			called++
			return first
		},
		func(f *Filter) error {
			called++
			f.CaseI = true
			return nil
		},
		func(*Filter) error {
			called++
			return second
		},
	)

	if filter == nil {
		t.Fatal("NewFilter returned a nil filter")
	}
	if !filter.CaseI {
		t.Error("NewFilter did not apply options after an error")
	}
	if called != 3 {
		t.Errorf("NewFilter called %d options, want 3", called)
	}
	if !errors.Is(err, first) {
		t.Error("NewFilter error does not contain the first option error")
	}
	if !errors.Is(err, second) {
		t.Error("NewFilter error does not contain the second option error")
	}
}

func TestNewFilterReturnsNilWithoutOptionErrors(t *testing.T) {
	_, err := NewFilter(FilterCaseInsensitive)
	if err != nil {
		t.Errorf("NewFilter returned an unexpected error: %v", err)
	}
}
