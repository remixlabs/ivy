package value

import (
	"fmt"

	"robpike.io/ivy/config"
)

// map is used purely for output
// it is unusable in regular ivy syntax
type Map map[string]Value

func NewMap(m map[string]Value) Map {
	return Map(m)
}

func (m Map) String() string {
	return "{ map }"
}

func (m Map) Sprint(conf *config.Config) string {
	s := "{ "
	for k, v := range m {
		s += fmt.Sprint(k, ":", v.Sprint(conf), " ")
	}
	return s + " }"
}

func (m Map) Eval(Context) Value {
	return m
}

func (m Map) Inner() Value {
	return m
}

func (m Map) ProgString() string {
	// There is no such thing as a vector in program listings; they
	// are represented as a sliceExpr.
	panic("map.ProgString - cannot happen")
}

func (m Map) toType(conf *config.Config, which valueType) Value {
	switch which {
	case mapType:
		return m
	}
	Errorf("cannot convert map to %s", which)
	return nil
}
