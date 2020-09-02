package jsondto

import (
	"bytes"
	"encoding/json"
	"io"
)

// Array represents a JSON-array.
type Array struct{
    s []json.Marshaler
}

func (a *Array) init() {
    if a.s == nil {
        a.s = []json.Marshaler{}
    }
}

func (a *Array) decodeJSON(dec *json.Decoder) (err error) {
    dec.UseNumber()
    defer wrapError(&err)
    a.Clear()
    for {
        if !dec.More() {
            if _,err = dec.Token(); err != nil {
                return err
            }
            return nil
        }
        var v json.Marshaler
        if v,err = decodeJSON(dec); err != nil {
            return err
        }
        a.Append(v)
    }
}

// Append adds v to a.
//
// errors:
//
//      *UnsupportedTypeError
//
func (a *Array) Append(v ...json.Marshaler) error {
    for _,v_ := range v {
        if err := validateValue(v_); err != nil {
            return err
        }
    }
    a.s = append(a.s, v...)
    return nil
}

// At returns the value at index i.
func (a *Array) At(i int) json.Marshaler {
    if i < 0 || i >= a.Len() {
        return nil
    }
    return a.s[i]
}

// Cap returns the capacity of a.
func (a *Array) Cap() int {
    return cap(a.s)
}

// Clear removes all values.
func (a *Array) Clear() {
    if a.Len() == 0 { return }
    for i := range a.s {
        a.s[i] = nil
    }
    a.s = a.s[:0]
}

// Insert adds v at i.
//
// errors:
//
//      *BoundsError
//      *UnsupportedTypeError
//
func (a *Array) Insert(i int, v json.Marshaler) error {
    if i < 0 || i > a.Len() {
        return &BoundsError{Index:i}
    }
    if err := validateValue(v); err != nil {
        return err
    }
    if i == a.Len() {
        return a.Append(v)
    }
    a.Append(Null{})
    copy(a.s[i+1:], a.s[i:])
    return a.Set(i, v)
}

// Len returns the number of values.
func (a *Array) Len() int {
    return len(a.s)
}

// MarshalJSON returns a JSON-array.
func (a *Array) MarshalJSON() ([]byte, error) {
    a.init()
    return json.Marshal(a.s)
}

// Remove removes the element at index i.
//
// errors:
//
//      *BoundsError
//
func (a *Array) Remove(i int) error {
    if a.At(i) == nil {
        return &BoundsError{Index:i}
    }
    if i != a.Len() - 1 {
        copy(a.s[i:], a.s[i+1:])
    }
    a.s = a.s[:a.Len()-1]
    return nil
}

// Reserve ensures a capacity of, at least, c.
func (a *Array) Reserve(c int) {
    if c > a.Cap() {
        var s []json.Marshaler = make([]json.Marshaler, a.Len(), c)
        copy(s, a.s)
        a.s = s
    }
}

// Set sets the value at index i.
//
// errors:
//
//      *BoundsError
//      *UnsupportedTypeError
//
func (a *Array) Set(i int, v json.Marshaler) error {
    if a.At(i) == nil {
        return &BoundsError{Index:i}
    }
    if err := validateValue(v); err != nil {
        return err
    }
    a.s[i] = v
    return nil
}

// Trim reduces the capacity to len.
func (a *Array) Trim() {
    if a.Len() == a.Cap() { return }
    var s []json.Marshaler = make([]json.Marshaler, a.Len())
    copy(s, a.s)
    a.s = s
}

// UnmarshalJSON unmarshals a JSON-array into a.
//
// errors:
//
//      *MultipleValuesError
//      *SyntaxError
//      *UnmarshalTypeError
//
func (a *Array) UnmarshalJSON(d []byte) (err error) {
    var dec *json.Decoder = json.NewDecoder(bytes.NewBuffer(d))
    dec.UseNumber()
    defer wrapError(&err)
    var tok json.Token
    if tok,err = dec.Token(); err != nil {
        return err
    }
    if tok != json.Delim('[') {
        var err_ *UnmarshalTypeError = &UnmarshalTypeError{Type:`Array`}
        switch t := tok.(type) {
        case nil:
            err_.Value = `null`
        case bool:
            err_.Value = `bool`
        case json.Number:
            if _,e := t.Int64(); e == nil {
                err_.Value = `int`
            }
            err_.Value = `float`
        case string:
            err_.Value = `string`
        case json.Delim:
            err_.Value = `object`
        }
        return err_
    }
    if err = a.decodeJSON(dec); err != nil {
        return err
    }
    if tok,err = dec.Token(); err == nil {
        return &MultipleValuesError{Data:d}
    }
    if err != io.EOF {
        return err
    }
    return nil
}
