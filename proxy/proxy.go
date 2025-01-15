package proxy

import (
	"fmt"
	"io"
	"net/http"

	"github.com/RohithBN/caching-proxy/cache"
)

type ProxyObject struct {
	Origin string
	Port   int
	Cache  map[string]*cache.CacheObject
}

func NewProxy(Origin string, Port int) *ProxyObject {
	return &ProxyObject{
		Origin: Origin,
		Port:   Port,
		Cache:  make(map[string]*cache.CacheObject),
	}
}

func (p *ProxyObject) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, ok := p.Cache[r.URL.String()]
		if !ok {
			fmt.Println("Cache miss for", r.URL.String())
			response, err := http.Get(p.Origin + r.URL.String())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			p.Cache[r.URL.String()] = cache.NewCacheObject(response, body)
		} else {
			fmt.Println("Cache hit for", r.URL.String())
			w.Write(p.Cache[r.URL.String()].ResponseBody)
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func (p *ProxyObject)ClearCache() {
	for k := range p.Cache{
		delete(p.Cache, k)
	}
}
