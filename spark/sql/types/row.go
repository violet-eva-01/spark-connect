package types

import (
	"maps"
)

type Row interface {
	// At returns field's value at the given index within a [Row].
	// It returns nil for invalid indices.
	At(index int) any
	// Value returns field's value of the given column's name within a [Row].
	// It returns nil for invalid column's name.
	Value(name string) any
	// Values returns values of all fields within a [Row] as a slice of any.
	Values() []any
	// Len returns the number of fields within a [Row]
	Len() int
	FieldNames() []string
}

type rowImpl struct {
	values  []any
	offsets map[string]int
}

func (r *rowImpl) At(index int) any {
	if index < 0 || index > len(r.values) {
		return nil
	}
	return r.values[index]
}

func (r *rowImpl) Value(name string) any {
	idx, ok := r.offsets[name]
	if !ok {
		return nil
	}
	return r.values[idx]
}

func (r *rowImpl) Values() []any {
	return r.values
}

func (r *rowImpl) Len() int {
	return len(r.values)
}

func (r *rowImpl) FieldNames() []string {
	names := make([]string, len(r.offsets))
	// Sort the field names to make the output deterministic.
	for k, v := range maps.All(r.offsets) {
		names[v] = k
	}
	return names
}
