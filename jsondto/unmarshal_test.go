package jsondto

import (
    "testing"
    "encoding/json"
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
