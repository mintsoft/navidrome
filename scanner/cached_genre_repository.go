package scanner

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/jellydator/ttlcache/v2"
	"github.com/navidrome/navidrome/log"
	"github.com/navidrome/navidrome/model"
)

var instance *cachedGenreRepo
var once sync.Once

func newCachedGenreRepository(ctx context.Context, repo model.GenreRepository) model.GenreRepository {
	once.Do(func() {

		r := &cachedGenreRepo{
			GenreRepository: repo,
			ctx:             ctx,
		}
		genres, err := repo.GetAll()

		if err != nil {
			log.Error(ctx, "Could not load genres from DB", err)
			//	return repo
			return
		}

		r.cache = ttlcache.NewCache()
		for _, g := range genres {
			_ = r.cache.Set(strings.ToLower(g.Name), g.ID)
		}

		cacheMetrics := r.cache.GetMetrics()
		log.Info(ctx, "GenreCache Contains : Inserted: %d, Retrievals: %d, Hits %d, Misses %d, Evicted %d", cacheMetrics.Inserted, cacheMetrics.Retrievals, cacheMetrics.Hits, cacheMetrics.Misses, cacheMetrics.Evicted)

		instance = r
	})

	return instance.GenreRepository
}

type cachedGenreRepo struct {
	model.GenreRepository
	cache *ttlcache.Cache
	ctx   context.Context
}

func (r *cachedGenreRepo) Put(g *model.Genre) error {
	id, err := r.cache.GetByLoader(strings.ToLower(g.Name), func(key string) (interface{}, time.Duration, error) {
		err := r.GenreRepository.Put(g)
		return g.ID, 24 * time.Hour, err
	})
	g.ID = id.(string)
	return err
}
