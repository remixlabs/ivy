package value

import (
	"robpike.io/ivy/config"
)

type String string

func NewString(s string) String {
	return String(s)
}

func (s String) String() string {
	return "\"" + string(s) + "\""
}

func (s String) Sprint(conf *config.Config) string {
	return string(s)
}

func (s String) ProgString() string {
	return string(s)
}

func (s String) Eval(Context) Value {
	return s
}

func (s String) Inner() Value {
	return s
}

func (s String) toArray() []Value {
	v := make([]Value, len(s))
	for i, c := range s {
		v[i] = Char(c)
	}
	return v
}

func (s String) toType(conf *config.Config, which valueType) Value {
	switch which {
	case stringType:
		return s
	case vectorType:
		return NewVector(s.toArray())
	case matrixType:
		return NewMatrix([]Value{one}, s.toArray())
	}
	Errorf("cannot convert string to %s", which)
	return nil
}
