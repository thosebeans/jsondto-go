package jsondto

import (
	"bytes"
	"encoding/json"
	"io"
)

// Object represents a JSON-object.
type Object struct{
    m map[String]json.Marshaler
}

func (o *Object) init() {
    if o.m == nil {
        o.m = map[String]json.Marshaler{}
    }
}

func (o *Object) decodeJSON(dec *json.Decoder) (err error) {
    dec.UseNumber()
    o.Clear()
    defer wrapError(&err)
    var tok json.Token
    for {
        if tok,err = dec.Token(); err != nil {
            return err
        }
        if tok == json.Delim('}') {
            return nil
        }
        var k String = String(tok.(string))
        var v json.Marshaler
        if v,err = decodeJSON(dec); err != nil {
            return err
        }
        o.Put(k, v)
    }
}

// Clear removes all kv-pairs.
func (o *Object) Clear() {
    for o.Len() > 0 {
        for k := range o.m {
            o.Delete(k)
            break
        }
    }
}

// Delete removes a kv-pair.
func (o *Object) Delete(k String) {
    delete(o.m, k)
}

// Get returns a kv-pair.
func (o *Object) Get(k String) json.Marshaler {
    return o.m[k]
}

// Len returns the number of kv-pairs.
func (o *Object) Len() int {
    return len(o.m)
}

// MarshalJSON returns a JSON-object.
func (o *Object) MarshalJSON() ([]byte, error) {
    if o.Len() == 0 {
        return []byte(`{}`),nil
    }
    var buff *bytes.Buffer = bytes.NewBuffer(make([]byte, 0, uint(o.Len()) * 8))
    buff.WriteRune('{')
    var first bool = true
    for k,v := range o.m {
        if !first {
            buff.WriteRune(',')
        }
        var b []byte
        b,_ = k.MarshalJSON()
        buff.Write(b)
        buff.WriteRune(':')
        b,_ = v.MarshalJSON()
        buff.Write(b)
        first = false
    }
    buff.WriteRune('}')
    return buff.Bytes(),nil
}

// Put adds a kv-pair.
//
// errors:
//
//      *UnsupportedTypeError
//
func (o *Object) Put(k String, v json.Marshaler) error {
    o.init()
    if err := validateValue(v); err != nil {
        return err
    }
    o.m[k] = v
    return nil
}

// Range iterates over o, until f==false.
func (o *Object) Range(f func(k String, v json.Marshaler) bool) {
    for k,v := range o.m {
        if !f(k, v) { return }
    }
}

// UnmarshalJSON unmarshals a JSON-object into o.
//
// errors:
//
//      *MultipleValuesError
//      *SyntaxError
//      *UnmarshalTypeError
//
func (o *Object) UnmarshalJSON(d []byte) (err error) {
    var dec *json.Decoder = json.NewDecoder(bytes.NewBuffer(d))
    dec.UseNumber()
    defer wrapError(&err)
    var tok json.Token
    if tok,err = dec.Token(); err != nil {
        return err
    }
    if tok != json.Delim('{') {
        var err_ *UnmarshalTypeError = &UnmarshalTypeError{Type:`Object`}
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
            err_.Value = `array`
        }
        return err_
    }
    if err = o.decodeJSON(dec); err != nil {
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
