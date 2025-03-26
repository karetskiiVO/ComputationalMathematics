package main

import "slices"

type UniversalMap[K, V any] struct {
	equal func(k1, k2 K) bool
	hashf func(k K) uintptr

	values map[uintptr]*struct {
		keys   []K
		values []V
	}
}

func NewUniversalMap[K, V any](
	equal func(k1, k2 K) bool,
	hashf func(k K) uintptr,
) *UniversalMap[K, V] {
	return &UniversalMap[K, V]{
		equal: equal,
		hashf: hashf,
		values: make(map[uintptr]*struct {
			keys   []K
			values []V
		}),
	}
}

// просрачивается при вставке
func (m *UniversalMap[K, V]) At(key K) (*V, bool) {
	sl, ok := m.values[m.hashf(key)]
	if !ok {
		return nil, false
	}

	idx := slices.IndexFunc(sl.keys, func(val K) bool {
		return m.equal(val, key)
	})

	if idx == -1 {
		return nil, false
	}
	return &sl.values[idx], true
}

func (m *UniversalMap[K, V]) Insert(key K, value V) {
	sl, ok := m.values[m.hashf(key)]
	if !ok {
		m.values[m.hashf(key)] = &struct {
			keys   []K
			values []V
		}{}
		sl = m.values[m.hashf(key)]
	}

	idx := slices.IndexFunc(sl.keys, func(val K) bool {
		return m.equal(val, key)
	})
	if idx != -1 {
		sl.values[idx] = value
	} else {
		sl.keys = append(sl.keys, key)
		sl.values = append(sl.values, value)
	}
}
