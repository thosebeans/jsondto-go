package jsondto

import (
    "bytes"
    "encoding/json"
    "io"
)

func decodeJSON(dec *json.Decoder) (val json.Marshaler, err error) {
    dec.UseNumber()
    var tok json.Token
    if tok,err = dec.Token(); err != nil {
        return nil,err
    }
    switch t := tok.(type) {
    case nil:
        return Null{},nil
    case bool:
        return Bool(t),nil
    case json.Number:
        var i int64
        if i,err = t.Int64(); err == nil {
            return Int(i),nil
        }
        var f float64
        f,_ = t.Float64()
        return Float(f),nil
    case string:
        return String(t),nil
    case json.Delim:
    }
    return nil,nil
}

// UnmarshalJSON unmarshals d into a value-type.
// 
// errors:
//
//      *SyntaxError        
//      *MultipleValuesError
//
func UnmarshalJSON(d []byte) (val json.Marshaler, err error) {
    var dec *json.Decoder = json.NewDecoder(bytes.NewBuffer(d))
    defer wrapError(&err)
    if val,err = decodeJSON(dec); err != nil {
        return nil,err
    }
    if _,err = dec.Token(); err == io.EOF {
        return val,nil
    }
    if err == nil {
        return nil,&MultipleValuesError{Data:d}
    }
    return nil,err
}
