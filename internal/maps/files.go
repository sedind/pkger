// Code generated by github.com/gobuffalo/mapgen. DO NOT EDIT.

package maps

import (
	"encoding/json"
	"fmt"
	"sort"
	"sync"

	"github.com/markbates/pkger/pkging"
)

// Files wraps sync.Map and uses the following types:
// key:   pkging.Path
// value: pkging.File
type Files struct {
	data *sync.Map
	once *sync.Once
}

func (m *Files) Data() *sync.Map {
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

func (m *Files) MarshalJSON() ([]byte, error) {
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

func (m *Files) UnmarshalJSON(b []byte) error {
	mm := map[string]pkging.File{}

	if err := json.Unmarshal(b, &mm); err != nil {
		return err
	}
	for k, v := range mm {
		var pt pkging.Path
		if err := json.Unmarshal([]byte(k), &pt); err != nil {
			return err
		}
		m.Store(pt, v)
	}
	return nil
}

// Delete the key from the map
func (m *Files) Delete(key pkging.Path) {
	m.Data().Delete(key)
}

// Load the key from the map.
// Returns pkging.File or bool.
// A false return indicates either the key was not found
// or the value is not of type pkging.File
func (m *Files) Load(key pkging.Path) (pkging.File, bool) {
	i, ok := m.Data().Load(key)
	if !ok {
		return nil, false
	}
	s, ok := i.(pkging.File)
	return s, ok
}

// Range over the pkging.File values in the map
func (m *Files) Range(f func(key pkging.Path, value pkging.File) bool) {
	m.Data().Range(func(k, v interface{}) bool {
		key, ok := k.(pkging.Path)
		if !ok {
			return false
		}
		value, ok := v.(pkging.File)
		if !ok {
			return false
		}
		return f(key, value)
	})
}

// Store a pkging.File in the map
func (m *Files) Store(key pkging.Path, value pkging.File) {
	m.Data().Store(key, value)
}

// Keys returns a list of keys in the map
func (m *Files) Keys() []pkging.Path {
	var keys []pkging.Path
	m.Range(func(key pkging.Path, value pkging.File) bool {
		keys = append(keys, key)
		return true
	})
	sort.Slice(keys, func(a, b int) bool {
		return keys[a].String() <= keys[b].String()
	})
	return keys
}

func (m *Files) String() string {
	return fmt.Sprintf("%v", m.Keys())
}
