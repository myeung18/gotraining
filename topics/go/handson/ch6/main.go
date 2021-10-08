package main

import (
	"fmt"
	graph2 "github.com/ardanlabs/gotraining/topics/go/handson/ch6/graph"
	memory2 "github.com/ardanlabs/gotraining/topics/go/handson/ch6/store/memory"
	"github.com/google/uuid"
	"time"
)

func main() {

	grp := memory2.NewInMemoryGrph()

	l := new(graph2.Link)
	l.URL = "www.google.com"
	l.RetrievedAt = time.Now().Add(time.Duration(-30) * time.Minute)
	grp.UpsertLink(l)

	it, err := grp.Links(uuid.Nil, uuid.MustParse("ffffffff-ffff-ffff-ffff-ffffffffffff"), time.Now())
	if err != nil {
		panic(err)
	}

	for it.Next() {
		fmt.Println(it.Link())
	}

}
