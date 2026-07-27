package sets

import (
	"go.yaml.in/yaml/v4"
)

func (s Set[T]) MarshalYAML() (any, error) {
	return yaml.Marshal(s.Values())
}

func (s *Set[T]) UnmarshalYAML(node *yaml.Node) error {
	var v []T
	if err := node.Decode(&v); err != nil {
		return err
	}
	*s = New[T]().Add(v...)
	return nil
}
