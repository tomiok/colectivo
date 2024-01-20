package colectivo

import "errors"

type Optional struct {
	V any
}

func (o *Optional) Get() any {
	return o.V
}

func (o *Optional) GetErr() (any, error) {
	if o.IsPresent() {
		return o.V, nil
	}

	return nil, errors.New("empty")
}

func (o *Optional) OrElse(a any) any {
	if o.IsPresent() {
		return o.V
	}
	o.V = a
	return a
}

func (o *Optional) IsPresent() bool {
	return o.V != nil
}

func Nil() Optional {
	return Optional{
		V: nil,
	}
}
