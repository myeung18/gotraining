package memory

import (
	graph2 "github.com/ardanlabs/gotraining/topics/go/handson/ch6/graph"
)

type linkIterator struct {
	s *InMemoryGraph

	links  []*graph2.Link
	curIdx int
}

func (l *linkIterator) Next() bool {
	if l.curIdx >= len(l.links) {
		return false
	}
	l.curIdx++
	return true
}

func (l *linkIterator) Error() error {
	return nil
}

func (l *linkIterator) Close() error {
	return nil
}

func (l *linkIterator) Link() *graph2.Link {
	l.s.mu.RLock()

	link := new(graph2.Link)
	*link = *l.links[l.curIdx-1]
	l.s.mu.RUnlock()
	return link
}

func helper() {

}
