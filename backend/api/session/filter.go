package session

import (
	"net/url"
	"time"

	"github.com/gorilla/schema"

	"github.com/aborilov/tech-challenge-time/backend/v1/pkg/session/model"
)

type filter struct {
	StartAfter time.Time `schema:"start_after"`
	EndBefore  time.Time `schema:"end_before"`
}

func (f filter) toServiceFilter() model.Filter {
	return model.Filter{
		StartAfter: f.StartAfter,
		EndBefore:  f.EndBefore,
	}
}

func parseFilter(v url.Values) (model.Filter, error) {
	f := filter{}
	schemaDecoder := schema.NewDecoder()
	schemaDecoder.IgnoreUnknownKeys(true)
	if err := schemaDecoder.Decode(&f, v); err != nil {
		return model.Filter{}, err
	}
	return f.toServiceFilter(), nil
}
