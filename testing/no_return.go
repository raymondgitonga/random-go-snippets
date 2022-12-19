package testing

import (
	"log"
)

type values struct {
	a      int
	b      int
	logger log.Logger
}

func (v *values) calculate() {
	c := v.a * v.b

	v.logger.Println(c)
}
