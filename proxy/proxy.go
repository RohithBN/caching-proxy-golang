package proxy

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"

	"github.com/RohithBN/caching-proxy/cache"
)

type ProxyObject struct {
	Origin string
	Port   int
	Cache  map[string]*cache.CacheObject
	mutex  sync.RWMutex 
}

func NewProxy(Origin string, Port int) *ProxyObject {
	_, err := url.Parse(Origin)
	if err != nil {
		panic(fmt.Sprintf("Invalid origin URL: %v", err))
	}

	return &ProxyObject{
		Origin: Origin,
		Port:   Port,
		Cache:  make(map[string]*cache.CacheObject),
	}
}

func (p *ProxyObject) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	p.mutex.RLock()
	cacheEntry, exists := p.Cache[r.URL.String()]
	p.mutex.RUnlock()

	if !exists {
		fmt.Printf("Cache MISS for %s\n", r.URL.String())
		w.Header().Set("X-Cache", "MISS")

		if err := p.handleCacheMiss(w, r); err != nil {
			http.Error(w, fmt.Sprintf("Error forwarding request: %v", err), http.StatusBadGateway)
			return
		}
	} else {
		fmt.Printf("Cache HIT for %s\n", r.URL.String())
		w.Header().Set("X-Cache", "HIT")

		// Copy original headers
		for k, v := range cacheEntry.Response.Header {
			w.Header()[k] = v
		}

		w.WriteHeader(cacheEntry.Response.StatusCode)

		if _, err := w.Write(cacheEntry.ResponseBody); err != nil {
			fmt.Printf("Error writing cached response: %v\n", err)
		}
	}
}

func (p *ProxyObject) handleCacheMiss(w http.ResponseWriter, r *http.Request) error {
	proxyReq, err := http.NewRequest(http.MethodGet, p.Origin+r.URL.String(), nil)
	if err != nil {
		return fmt.Errorf("error creating proxy request: %v", err)
	}

	proxyReq.Header = r.Header

	client := &http.Client{}
	response, err := client.Do(proxyReq)
	if err != nil {
		return fmt.Errorf("error making request to origin: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	p.mutex.Lock()
	p.Cache[r.URL.String()] = cache.NewCacheObject(response, body)
	p.mutex.Unlock()

	for k, v := range response.Header {
		w.Header()[k] = v
	}

	w.WriteHeader(response.StatusCode)
	if _, err := w.Write(body); err != nil {
		return fmt.Errorf("error writing response: %v", err)
	}

	return nil
}

func (p *ProxyObject) ClearCache() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.Cache = make(map[string]*cache.CacheObject)
	fmt.Println("Cache cleared successfully")
}
