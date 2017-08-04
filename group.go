package pop

import (
	"fmt"
	"strings"
)

type GroupClause struct {
	Field string
}

type groupClauses []GroupClause

func (c GroupClause) String() string {
	return fmt.Sprintf("%q", c.Field)
}

func (c groupClauses) String() string {
	cs := []string{}
	for _, cl := range c {
		cs = append(cs, cl.String())
	}
	return strings.Join(cs, ", ")
}
