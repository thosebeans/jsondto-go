package jsondto

// Array represents a JSON-array.
type Array struct{}

// MarshalJSON returns a JSON-array.
func (a *Array) MarshalJSON() ([]byte, error) {
    return nil,nil
}
