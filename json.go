package sets

import (
	"encoding/json"
)

func (s Set[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values())
}

func (s *Set[T]) UnmarshalJSON(data []byte) error {
	var v []T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	Init(s)
	_ = s.Add(v...)
	return nil
}
