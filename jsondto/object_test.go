package jsondto

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObjectUnmarshal(t *testing.T) {
    data := []byte(`
        {
            "a": 1,
            "b": null
        }
    `)
    o := &Object{}
    assert.Nil(t, o.UnmarshalJSON(data))
    output,err := o.MarshalJSON()
    assert.Nil(t, err)
    assert.JSONEq(t, string(data), string(output))
}

func TestObjectErrors(t *testing.T) {
    o := &Object{}
    assert.IsType(t, errors.New(``), o.Put(`1`, nil))
    assert.IsType(t, &UnsupportedTypeError{}, o.Put(`1`, DummyMarshaler{}))
}
