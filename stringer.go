package sets

import (
	"fmt"
	"strings"
)

func (s Set[T]) String() string {
	t := strings.Builder{}
	t.WriteRune('{')
	flg := false
	for k := range s {
		if flg {
			t.WriteString(", ")
		}
		t.WriteString(fmt.Sprint(k))
		flg = true
	}
	t.WriteRune('}')
	return t.String()
}
