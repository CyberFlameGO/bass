package bass

var ground = NewEnv()

func init() {
	for k, v := range primPreds {
		ground.Set(k, Func(string(k), v))
	}

	ground.Set("+", Func("+", func(nums ...int) int {
		sum := 0
		for _, num := range nums {
			sum += num
		}

		return sum
	}))

	ground.Set("cons", Func("cons", func(a, d Value) Value {
		return Pair{a, d}
	}))

	ground.Set("wrap", Func("wrap", func(c Combiner) Applicative {
		return Applicative{c}
	}))

	ground.Set("unwrap", Func("unwrap", func(a Applicative) Combiner {
		return a.Underlying
	}))

	ground.Set("op", Op("op", newOperative))

	ground.Set("eval", Func("eval", func(val Value, env *Env) (Value, error) {
		return val.Eval(env)
	}))

	ground.Set("def", Op("def", func(val List, env *Env) (Value, error) {
		switch d := val.Rest().(type) {
		case Empty:
			return nil, ErrBadSyntax
		case List:
			err := env.Define(val.First(), d.First())
			if err != nil {
				return nil, err
			}
		default:
			return nil, ErrBadSyntax
		}

		return val.First(), nil
	}))
}

func New() *Env {
	return NewEnv(ground)
}

type pred func(Value) bool

var primPreds = map[Symbol]pred{
	// primitive type checks
	"null?": func(val Value) bool {
		_, is := val.(Null)
		return is
	},
	"boolean?": func(val Value) bool {
		_, is := val.(Bool)
		return is
	},
	"number?": func(val Value) bool {
		_, is := val.(Int)
		return is
	},
	"string?": func(val Value) bool {
		_, is := val.(String)
		return is
	},
	"symbol?": func(val Value) bool {
		_, is := val.(Symbol)
		return is
	},
	"env?": func(val Value) bool {
		_, is := val.(*Env)
		return is
	},
	"list?": func(val Value) bool {
		_, is := val.(List)
		return is
	},
	"pair?": func(val Value) bool {
		_, is := val.(Pair)
		return is
	},
	"combiner?": func(val Value) bool {
		_, is := val.(Combiner)
		return is
	},
	"applicative?": func(val Value) bool {
		_, is := val.(Applicative)
		return is
	},
	"operative?": func(val Value) bool {
		switch x := val.(type) {
		case *Builtin:
			return x.Operative
		case *Operative:
			return true
		default:
			return false
		}
	},
	"empty?": func(val Value) bool {
		switch x := val.(type) {
		case Empty:
			return true
		case String:
			return len(x) == 0
		case Null:
			return true
		default:
			return false
		}
	},
}

func newOperative(val List, env *Env) (*Operative, error) {
	op := &Operative{
		Env: env,
	}

	switch val.(type) {
	case Empty:
		return nil, ArityError{
			Name: "op",
			Need: 3,
			Have: 0,
		}
	default:
		op.Formals = val.First()

		switch x := val.Rest().(type) {
		case Empty:
			return nil, ArityError{
				Name: "op",
				Need: 3,
				Have: 1,
			}
		case List:
			op.Eformal = x.First()
			switch b := x.Rest().(type) {
			case Empty:
				return nil, ArityError{
					Name: "op",
					Need: 3,
					Have: 2,
				}
			case List:
				op.Body = b.First()
			default:
				return nil, ErrBadSyntax
			}
		default:
			return nil, ErrBadSyntax
		}

		return op, nil
	}
}
