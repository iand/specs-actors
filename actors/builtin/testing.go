package builtin

import (
	"strings"

	"golang.org/x/xerrors"
)

// Accumulates a sequence of errors, allowing the postponement of many checks for error conditions until
// the end of a sequence of operations.
type ErrAccumulator struct {
	errs errs
}

// Returns an error, if any were accumulated, else nil.
func (e *ErrAccumulator) AsError() error {
	// This comparison with nil, rather than just returning the "nil", is necessary because when
	// Errors(nil) is returned it becomes not equal to nil. Go figure.
	if e.errs == nil {
		return nil
	}
	return e.errs
}

// Adds an error to the accumulator.
func (e *ErrAccumulator) Add(err error) {
	e.errs = append(e.errs, err)
}

// Adds an error if predicate is false.
func (e *ErrAccumulator) Require(predicate bool, msg string, args ...interface{}) {
	if !predicate {
		e.Add(xerrors.Errorf(msg, args...))
	}
}

type errs []error

func (ae errs) Error() string {
	strs := make([]string, len(ae))
	for i, e := range ae {
		strs[i] = e.Error()
	}
	return strings.Join(strs, ";\n")
}
