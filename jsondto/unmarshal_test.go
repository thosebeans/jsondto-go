package jsondto

import (
	"encoding/json"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalSimple(t *testing.T) {
    data := map[string]json.Marshaler{
        `null`: Null{},
        `true`: Bool(true),
        `12`:   Int(12),
        `3.14`: Float(3.14),
        `"ab"`: String(`ab`),
    }
    for input,expectedVal := range data {
        if val,err := UnmarshalJSON([]byte(input)); err != nil {
            t.Fatal(`unexpected error: `, err)
        } else if expectedVal != val {
            t.Fatal(`values dont match: `, val, ` != `, expectedVal)
        }
    }
}

func TestUnmarshalErrors(t *testing.T) {
    var err error
    _,err = UnmarshalJSON(nil)
    assert.Equal(t, io.EOF, err)
    _,err = UnmarshalJSON([]byte(`{"asd":}`))
    assert.IsType(t, &SyntaxError{}, err)
}
