// Code generated by github.com/gobuffalo/mapgen. DO NOT EDIT.

package pkger

import (
	"encoding/json"
	"fmt"
	"sort"
	"sync"

	"github.com/gobuffalo/here"
)

// infosMap wraps sync.Map and uses the following types:
// key:   string
// value: here.Info
type infosMap struct {
	data *sync.Map
	once *sync.Once
}

func (m *infosMap) Data() *sync.Map {
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

func (m *infosMap) MarshalJSON() ([]byte, error) {
	mm := map[string]interface{}{}
	m.data.Range(func(key, value interface{}) bool {
		mm[fmt.Sprintf("%s", key)] = value
		return true
	})
	return json.Marshal(mm)
}

func (m *infosMap) UnmarshalJSON(b []byte) error {
	mm := map[string]here.Info{}

	if err := json.Unmarshal(b, &mm); err != nil {
		return err
	}
	for k, v := range mm {
		m.Store(k, v)
	}
	return nil
}

// Delete the key from the map
func (m *infosMap) Delete(key string) {
	m.Data().Delete(key)
}

// Load the key from the map.
// Returns here.Info or bool.
// A false return indicates either the key was not found
// or the value is not of type here.Info
func (m *infosMap) Load(key string) (here.Info, bool) {
	i, ok := m.Data().Load(key)
	if !ok {
		return here.Info{}, false
	}
	s, ok := i.(here.Info)
	return s, ok
}

// LoadOrStore will return an existing key or
// store the value if not already in the map
func (m *infosMap) LoadOrStore(key string, value here.Info) (here.Info, bool) {
	i, _ := m.Data().LoadOrStore(key, value)
	s, ok := i.(here.Info)
	return s, ok
}

// LoadOr will return an existing key or
// run the function and store the results
func (m *infosMap) LoadOr(key string, fn func(*infosMap) (here.Info, bool)) (here.Info, bool) {
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

// Range over the here.Info values in the map
func (m *infosMap) Range(f func(key string, value here.Info) bool) {
	m.Data().Range(func(k, v interface{}) bool {
		key, ok := k.(string)
		if !ok {
			return false
		}
		value, ok := v.(here.Info)
		if !ok {
			return false
		}
		return f(key, value)
	})
}

// Store a here.Info in the map
func (m *infosMap) Store(key string, value here.Info) {
	m.Data().Store(key, value)
}

// Keys returns a list of keys in the map
func (m *infosMap) Keys() []string {
	var keys []string
	m.Range(func(key string, value here.Info) bool {
		keys = append(keys, key)
		return true
	})
	sort.Strings(keys)
	return keys
}
