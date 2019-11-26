package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"sync/atomic"
)

type Backend struct {
	URL *url.URL
	Alive bool
	mux sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

type ServerPool struct {
	backends []*Backend
	current uint64
}

func main() {
	u, _ := url.Parse("http://localhost:8080")
	rp := httputil.NewSingleHostReverseProxy(u)
	// 初始化服务器， 并添加处理器
	http.HandlerFunc(rp.ServeHTTP)
}

func (s *ServerPool) NextIndex() int  {
	return int(atomic.AddUint64(&s.current, uint64(1)) % uint64(len(s.backends)))
}
// 返回下一个可用的服务器
func (s *ServerPool) GetNextPeer() *Backend {
	next := s.NextIndex()
	l := len(s.backends) + next // 从next开始遍历
	for i := next; i < l; i++ {
		idx := i % len(s.backends) // 通过取模计算获得索引
		// 如果找到一个可用的索引服务器，将它标记为当前的服务器， 如果不是初始的那个，就将他保存下来
		if i != next {
			atomic.StoreUint64(&s.current, uint64(idx)) //标记当前可用的服务器
		}
		return s.backends[idx]
	}

	return nil
}

func (b *Backend) SetAlive(alive bool)  {
	b.mux.Lock()
	b.Alive = alive
	b.mux.Unlock()
}

// 如果后端还活着， isAlive 返回true
func (b *Backend) IsAlive() (alive bool) {
	b.mux.Lock()
	alive = b.Alive
	b.mux.Unlock()
	return
}

// lb 对入向请求进行负载均衡
func lb(w http.ResponseWriter, r *http.Request)  {
	peer := ServerPool.GetNextPeer()
	if peer != nil {
		peer.ReverseProxy.ServeHTTP(w, r)
		return
	}

	http.Error(w, "Service not avaliable", http.StatusServiceUnavailable)
}


