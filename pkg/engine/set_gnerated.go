// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package engine

type StringSet map[string]struct{}

func NewStringSet() StringSet {
	return make(StringSet)
}

func (s StringSet) Add(v string) {
	s[v] = struct{}{}
}

func (s StringSet) Update(s2 StringSet) {
	for v := range s2 {
		s[v] = struct{}{}
	}
}

func (s StringSet) Count() int {
	return len(s)
}

func (s StringSet) Has(v string) bool {
	_, ok := s[v]
	return ok
}

func (s StringSet) Slice() []string {
	slice := make([]string, len(s))
	var i int
	for v := range s {
		slice[i] = v
		i += 1
	}
	return slice
}

func (s StringSet) Delete(v string) {
	delete(s, v)
}

func (s StringSet) Copy() StringSet {
	s2 := make(StringSet, len(s))
	for v := range s {
		s2[v] = struct{}{}
	}
	return s2
}

type DeveloperSet map[Developer]struct{}

func NewDeveloperSet() DeveloperSet {
	return make(DeveloperSet)
}

func (s DeveloperSet) Add(v Developer) {
	s[v] = struct{}{}
}

func (s DeveloperSet) Update(s2 DeveloperSet) {
	for v := range s2 {
		s[v] = struct{}{}
	}
}

func (s DeveloperSet) Count() int {
	return len(s)
}

func (s DeveloperSet) Has(v Developer) bool {
	_, ok := s[v]
	return ok
}

func (s DeveloperSet) Slice() []Developer {
	slice := make([]Developer, len(s))
	var i int
	for v := range s {
		slice[i] = v
		i += 1
	}
	return slice
}

func (s DeveloperSet) Delete(v Developer) {
	delete(s, v)
}

func (s DeveloperSet) Copy() DeveloperSet {
	s2 := make(DeveloperSet, len(s))
	for v := range s {
		s2[v] = struct{}{}
	}
	return s2
}

type CommitSet map[Commit]struct{}

func NewCommitSet() CommitSet {
	return make(CommitSet)
}

func (s CommitSet) Add(v Commit) {
	s[v] = struct{}{}
}

func (s CommitSet) Update(s2 CommitSet) {
	for v := range s2 {
		s[v] = struct{}{}
	}
}

func (s CommitSet) Count() int {
	return len(s)
}

func (s CommitSet) Has(v Commit) bool {
	_, ok := s[v]
	return ok
}

func (s CommitSet) Slice() []Commit {
	slice := make([]Commit, len(s))
	var i int
	for v := range s {
		slice[i] = v
		i += 1
	}
	return slice
}

func (s CommitSet) Delete(v Commit) {
	delete(s, v)
}

func (s CommitSet) Copy() CommitSet {
	s2 := make(CommitSet, len(s))
	for v := range s {
		s2[v] = struct{}{}
	}
	return s2
}
