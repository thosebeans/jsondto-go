package jsondto

type DummyMarshaler struct {}

func (m DummyMarshaler) MarshalJSON() ([]byte,error) { return nil,nil }
