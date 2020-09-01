package jsondto

import (
    "encoding/json"
    "errors"
    "reflect"
)

func validateValue(v json.Marshaler) error {
    switch v.(type) {
    case Null,Bool,Int,Float,String,*Object,*Array:
        return nil
    case nil:
        return errors.New(`jsondto: Use of nil as json.Marshaler`)
    }
    return &UnsupportedTypeError{T:reflect.TypeOf(v)}
}
