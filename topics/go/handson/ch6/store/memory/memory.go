package memory

import (
	graph2 "github.com/ardanlabs/gotraining/topics/go/handson/ch6/graph"
	"github.com/google/uuid"
	"sync"
	"time"
)

type InMemoryGraph struct {
	mu sync.RWMutex

	links map[uuid.UUID]*graph2.Link

	linkURLIndex map[string]*graph2.Link
}

func NewInMemoryGrph() *InMemoryGraph {
	return &InMemoryGraph{
		links : make(map[uuid.UUID]*graph2.Link),
		linkURLIndex: make(map[string]*graph2.Link),
	}
}

func (i *InMemoryGraph) UpsertLink(link *graph2.Link) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	//exist then update
	if existing := i.linkURLIndex[link.URL]; existing != nil {
		link.ID = existing.ID

		return nil
	}

	for {
		link.ID = uuid.New()
		if i.links[link.ID] == nil {
			break;
		}
	}

	lcpy := new(graph2.Link)
	*lcpy = *link
	i.linkURLIndex[lcpy.URL] = lcpy
	i.links[lcpy.ID] = lcpy
	//add new
	return nil
}

func (i *InMemoryGraph) FindLink(id uuid.UUID) (*graph2.Link, error) {
	return nil, nil
}

func (i *InMemoryGraph) Links(fromID, toID uuid.UUID, retrievedBefore time.Time) (graph2.LinkIterator, error) {
	from, to := fromID.String(), toID.String()

	i.mu.RLock()
	var list []*graph2.Link
	for linkID, link := range i.links {
		if id := linkID.String(); id >= from && id <= to && link.RetrievedAt.Before(retrievedBefore) {
			list = append(list, link)
		}
	}

	i.mu.RUnlock()
	//always return a concrete impl.
	return &linkIterator{s: i, links: list}, nil
}
