// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package identifier

type HashSet map[Hash]struct{}

func NewHashSet() HashSet {
	return make(HashSet)
}

func (s HashSet) Add(v Hash) {
	s[v] = struct{}{}
}

func (s HashSet) Update(s2 HashSet) {
	for v := range s2 {
		s[v] = struct{}{}
	}
}

func (s HashSet) Count() int {
	return len(s)
}

func (s HashSet) Has(v Hash) bool {
	_, ok := s[v]
	return ok
}

func (s HashSet) Slice() []Hash {
	slice := make([]Hash, len(s))
	var i int
	for v := range s {
		slice[i] = v
		i += 1
	}
	return slice
}

func (s HashSet) Delete(v Hash) {
	delete(s, v)
}

func (s HashSet) Copy() HashSet {
	s2 := make(HashSet, len(s))
	for v := range s {
		s2[v] = struct{}{}
	}
	return s2
}