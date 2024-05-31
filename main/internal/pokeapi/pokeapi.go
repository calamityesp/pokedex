package pokeapi

import (
	"net/http"
	"time"

	pokecache "github.com/calamityesp/pokedex/main/internal"
)

const baseUrl = "https://pokeapi.co/api/v2/"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCacheHelper(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
