package graph

import (
	"github.com/google/uuid"
	"time"
)

type Iterator interface {
	Next() bool
	Error() error
	Close() error
}

type Link struct {
	ID uuid.UUID
	URL string
	RetrievedAt time.Time
}

type LinkIterator interface {
	Iterator

	Link() *Link
}

type Graph interface {
	UpsertLink(link *Link) error
	FindLink(id uuid.UUID) (*Link, error)

	Links(fromID, toID uuid.UUID, retrievedBefore time.Time) (LinkIterator, error)
}