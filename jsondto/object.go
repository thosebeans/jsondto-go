package jsondto

// Object represents a JSON-object.
type Object struct{}

// MarshalJSON returns a JSON-object.
func (o *Object) MarshalJSON() ([]byte, error) {
    return nil,nil
}
