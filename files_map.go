// Code generated by github.com/gobuffalo/mapgen. DO NOT EDIT.

package pkger

import (
	"encoding/json"
	"sort"
	"sync"
)

// filesMap wraps sync.Map and uses the following types:
// key:   Path
// value: *File
type filesMap struct {
	data *sync.Map
	once *sync.Once
}

func (m *filesMap) Data() *sync.Map {
	if m.once == nil {
		m.once = &sync.Once{}
	}
	m.once.Do(func() {
		if m.data == nil {
			m.data = &sync.Map{}
		}
	})
	return m.data
}

func (m *filesMap) MarshalJSON() ([]byte, error) {
	var err error
	mm := map[string]interface{}{}
	m.Data().Range(func(key, value interface{}) bool {
		var b []byte
		b, err = json.Marshal(key)
		if err != nil {
			return false
		}
		mm[string(b)] = value
		return true
	})

	if err != nil {
		return nil, err
	}

	return json.Marshal(mm)
}

func (m *filesMap) UnmarshalJSON(b []byte) error {
	mm := map[string]*File{}

	if err := json.Unmarshal(b, &mm); err != nil {
		return err
	}
	for k, v := range mm {
		var pt Path
		if err := json.Unmarshal([]byte(k), &pt); err != nil {
			return err
		}
		m.Store(pt, v)
	}
	return nil
}

// Delete the key from the map
func (m *filesMap) Delete(key Path) {
	m.Data().Delete(key)
}

// Load the key from the map.
// Returns *File or bool.
// A false return indicates either the key was not found
// or the value is not of type *File
func (m *filesMap) Load(key Path) (*File, bool) {
	i, ok := m.Data().Load(key)
	if !ok {
		return nil, false
	}
	s, ok := i.(*File)
	return s, ok
}

// LoadOrStore will return an existing key or
// store the value if not already in the map
func (m *filesMap) LoadOrStore(key Path, value *File) (*File, bool) {
	i, _ := m.Data().LoadOrStore(key, value)
	s, ok := i.(*File)
	return s, ok
}

// LoadOr will return an existing key or
// run the function and store the results
func (m *filesMap) LoadOr(key Path, fn func(*filesMap) (*File, bool)) (*File, bool) {
	i, ok := m.Load(key)
	if ok {
		return i, ok
	}
	i, ok = fn(m)
	if ok {
		m.Store(key, i)
		return i, ok
	}
	return i, false
}

// Range over the *File values in the map
func (m *filesMap) Range(f func(key Path, value *File) bool) {
	m.Data().Range(func(k, v interface{}) bool {
		key, ok := k.(Path)
		if !ok {
			return false
		}
		value, ok := v.(*File)
		if !ok {
			return false
		}
		return f(key, value)
	})
}

// Store a *File in the map
func (m *filesMap) Store(key Path, value *File) {
	m.Data().Store(key, value)
}

// Keys returns a list of keys in the map
func (m *filesMap) Keys() []Path {
	var keys []Path
	m.Range(func(key Path, value *File) bool {
		keys = append(keys, key)
		return true
	})
	sort.Slice(keys, func(a, b int) bool {
		return keys[a].String() <= keys[b].String()
	})
	return keys
}