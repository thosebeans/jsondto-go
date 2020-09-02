package jsondto

import (
    "encoding/json"
    "reflect"
)

// UnmarshalTypeError indicates a JSON-value, inappropriate for a specific go-type.
type UnmarshalTypeError struct {
    Value string // json-value-type: null,bool,int,float,string,object,array
    Type  string // go-type: Object,Array
}

func (err *UnmarshalTypeError) Error() string {
    return `jsondto: UnmarshalTypeError: attempted unmarshal of ` +
                err.Value +
                ` into ` +
                err.Type
}

// UnsupportedTypeError indicates an illegal type.
type UnsupportedTypeError struct {
    T reflect.Type
}

func (err *UnsupportedTypeError) Error() string {
    return `jsondto: UnsupportedTypeError: illegal type ` + err.T.String()
}

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
