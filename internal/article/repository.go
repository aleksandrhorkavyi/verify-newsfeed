package article

import (
	"bufio"
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
)

type Repository interface {
	DoSomething() map[string]any
	One() map[string]any
}

type repository struct {
	es *elasticsearch.TypedClient
}

func NewRepository(es *elasticsearch.TypedClient) Repository {
	return &repository{es: es}
}

func (r repository) One() map[string]any {
	resp, err := r.es.Search().
		Index("my_index").
		Request(&search.Request{
			Query: &types.Query{
				Match: map[string]types.MatchQuery{
					"short": {Query: "short"},
				},
			},
		}).
		Do(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	return map[string]any{
		"r": resp,
	}
}

func (r repository) DoSomething() map[string]any {
	dir, _ := os.Getwd()

	file, _ := os.Open(path.Join(dir, "/examples/elastic/json/map.json"))

	resp, err := r.es.Search().Index("my_index").Raw(bufio.NewReader(file)).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return map[string]any{
		"you are suck": "ya",
		"s":            resp,
	}
}
