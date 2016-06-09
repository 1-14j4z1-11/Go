package eval

import (
	"fmt"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%f", l)
}

func (u unary) String() string {
	return fmt.Sprintf("%c%s", u.op, u.x.String())
}

func (b binary) String() string {
	return fmt.Sprintf("%s %c %s", b.x.String(), b.op, b.y.String())
}

func (c call) String() string {
	str := c.fn + "("

	for i, arg := range c.args {
		str += arg.String()

		if i != len(c.args) - 1 {
			str += ", "
		}
	}

	return str + ")"
}
