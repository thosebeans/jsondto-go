package jsondto

import (
    "encoding/json"
)

// SyntaxError indicates a syntax-error in the input-data.
type SyntaxError struct {
    Err *json.SyntaxError
}

func (err *SyntaxError) Error() string {
    return `jsondto: SyntaxError: ` + err.Err.Error()
}

func (err *SyntaxError) Unwrap() error {
    return err.Err
}


// MultipleValuesError indicates multiple top-level JSON-values in the input-data.
type MultipleValuesError struct{
    Data []byte
}

func (err *MultipleValuesError) Error() string {
    return `jsondto: MultipleValuesError: multiple top-level values in input-data`
}

func wrapError(err *error) {
    switch e := (*err).(type) {
    case *json.SyntaxError:
        *err = &SyntaxError{Err:e}
    }
}
