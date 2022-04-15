package bass

import (
	"context"
	"strconv"
)

type Int int

func (value Int) Repr() string {
	return strconv.Itoa(int(value))
}

func (value Int) Equal(other Value) bool {
	var o Int
	return other.Decode(&o) == nil && value == o
}

func (value Int) Decode(dest any) error {
	switch x := dest.(type) {
	case *Int:
		*x = value
		return nil
	case *Value:
		*x = value
		return nil
	case *Bindable:
		*x = value
		return nil
	case *int:
		*x = int(value)
		return nil
	default:
		return DecodeError{
			Source:      value,
			Destination: dest,
		}
	}
}

// Eval returns the value.
func (value Int) Eval(_ context.Context, _ *Scope, cont Cont) ReadyCont {
	return cont.Call(value, nil)
}

var _ Bindable = Int(0)

func (binding Int) Bind(_ context.Context, _ *Scope, cont Cont, val Value, _ ...Annotated) ReadyCont {
	return cont.Call(binding, BindConst(binding, val))
}

func (Int) EachBinding(func(Symbol, Range) error) error {
	return nil
}
