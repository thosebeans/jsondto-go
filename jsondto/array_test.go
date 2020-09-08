package jsondto

import (
    "testing"
    "errors"

    "github.com/stretchr/testify/assert"
)

func TestArrayUnmarshal(t *testing.T) {
    data := []byte(`
        [1,2,3]
    `)
    a := &Array{}
    assert.Nil(t, a.UnmarshalJSON(data))
    output,err := a.MarshalJSON()
    assert.Nil(t, err)
    assert.JSONEq(t, string(data), string(output))
}

func TestArrayErrors(t *testing.T) {
    a := &Array{}
    assert.IsType(t, errors.New(``), a.Append(nil))
    assert.IsType(t, &UnsupportedTypeError{}, a.Append(DummyMarshaler{}))
    assert.IsType(t, &BoundsError{}, a.Set(5, Int(1)))
}
