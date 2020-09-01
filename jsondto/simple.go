package jsondto

import (
    "strconv"
)

// Null represents a null value.
type Null struct{}

func (v Null) MarshalJSON() ([]byte, error) {
    return []byte(`null`),nil
}

// Bool represents a bool value.
type Bool bool

func (v Bool) MarshalJSON() ([]byte, error) {
    if v {
        return []byte(`true`),nil
    }
    return []byte(`false`),nil
}

// Int represents an integer value.
type Int int64

func (v Int) MarshalJSON() ([]byte, error) {
    return []byte(strconv.FormatInt(int64(v), 10)),nil
}

// Float represents a floating-point value.
type Float float64

func (v Float) MarshalJSON() ([]byte, error) {
    return []byte(strconv.FormatFloat(float64(v), 'f', -1, 64)),nil
}

// String represents a string value.
type String string

func (v String) MarshalJSON() ([]byte, error) {
    return []byte(strconv.Quote(string(v))),nil
}
