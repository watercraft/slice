package slice

import (
	"sort"
)

type StringMap map[string]interface{}

func (sm StringMap) Keys() []string {
	ret := make([]string, len(sm))
	i := 0
	for k := range sm {
		ret[i] = k
		i++
	}
	return ret
}

func (sm StringMap) SortedKeys() []string {
	ret := sm.Keys()
	sort.Strings(ret)
	return ret
}

func (sm StringMap) SortedValues() []interface{} {
	keys := sm.SortedKeys()
	ret := make([]interface{}, len(sm))
	for i, k := range keys {
		ret[i] = sm[k]
	}
	return ret
}
