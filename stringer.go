package sets

import (
	"fmt"
	"strings"
)

func (s Set[T]) String() string {
	t := strings.Builder{}
	t.WriteRune('{')
	for k := range s {
		t.WriteString(fmt.Sprint(k))
	}
	t.WriteRune('}')
	return t.String()
}
