// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package manage

type QueueSet map[*Queue]struct{}

func NewQueueSet() QueueSet {
	return make(QueueSet)
}

func (s QueueSet) Add(v *Queue) {
	s[v] = struct{}{}
}

func (s QueueSet) Update(s2 QueueSet) {
	for v := range s2 {
		s[v] = struct{}{}
	}
}

func (s QueueSet) Count() int {
	return len(s)
}

func (s QueueSet) Has(v *Queue) bool {
	_, ok := s[v]
	return ok
}

func (s QueueSet) Slice() []*Queue {
	slice := make([]*Queue, len(s))
	var i int
	for v := range s {
		slice[i] = v
		i += 1
	}
	return slice
}

func (s QueueSet) Delete(v *Queue) {
	delete(s, v)
}

func (s QueueSet) Copy() QueueSet {
	s2 := make(QueueSet, len(s))
	for v := range s {
		s2[v] = struct{}{}
	}
	return s2
}
