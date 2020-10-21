package main

import (
	"fmt"
	"goinreverse/heartbeat"
	"goinreverse/parsers"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"sync"
	"time"
)

var ServerPort = os.Getenv("SERVER_PORT")
var config = parsers.ParseConfig()

type Server struct {
	URL      *url.URL
	Name     string
	isOnline bool
	mux      sync.RWMutex
	Proxy    *httputil.ReverseProxy
}

func (s *Server) SetOnline(status bool) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.isOnline = status
}

func (s *Server) IsOnline() bool {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.isOnline
}


type ServerPool struct {
	servers []*Server
	current int
}

//AddServer - Adds individual server to the pool of available servers
func (sp *ServerPool) AddServer(server *Server) {
	log.Printf("Added server %s to pool", server.Name)
	sp.servers = append(sp.servers, server)
}

//HealthCheck - Loops servers in pool and pings all the servers
func (sp *ServerPool) HealthCheck() {
	for _, server := range sp.servers {
		status := heartbeat.PingServer(server.URL)
		//TODO: Ping each server 3 times? To determine if healthy
		server.SetOnline(status)
	}
}

func (sp *ServerPool) getNextAvailableServerIndex() int {
	return (sp.current + 1) % len(sp.servers)
}

//Looks at server pool determines if server is healthy and uses it as proxy
func (sp *ServerPool) getHealthyServer() *Server {
	nextServer := sp.getNextAvailableServerIndex()
	if sp.servers[nextServer].isOnline && sp.current != nextServer {
		sp.current = nextServer
		return sp.servers[nextServer]
	}

	if sp.servers[sp.current].isOnline {
		return sp.servers[sp.current]
	}

	return nil
}

func loadBalanceHandler(w http.ResponseWriter, r *http.Request) {
	proxyPeer := serverPool.getHealthyServer()
	if proxyPeer != nil {
		proxyPeer.Proxy.ServeHTTP(w, r)
		return
	}
	http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
}

//healthCheck - Periodically set by {config.Healthecks} calls HealthCheck method
func healthCheck() {
	t := time.NewTicker(time.Second * time.Duration(config.Healthcheck))
	//TODO: maybe? rewrite in a more clear way ( snippet from Google )
	for {
		select {
		case <- t.C:
			log.Println("Starting health check...")
			serverPool.HealthCheck()
			log.Println("Health check completed!")
		}
	}
}

var serverPool ServerPool

func main() {
	servers := config.Servers

	for _, server := range servers {
		serverUrl, err := url.Parse(server.Url)
		if err != nil {
			log.Fatal("Unable to parse url, error:", err)
		}

		director := func(req *http.Request) {
			req.Header.Add("X-Forwarded-Host", req.Host)
			req.Header.Add("X-Origin-Host", serverUrl.Host)
			req.URL.Scheme = "http"
			req.URL.Host = serverUrl.Host
		}

		proxy := &httputil.ReverseProxy{Director: director}

		serverPool.AddServer(&Server {
			URL:      serverUrl,
			Name:     server.Name,
			isOnline: true,
			Proxy:    proxy,
		})
	}

	server := http.Server {
		Addr: fmt.Sprintf(":%s", ServerPort),
		Handler: http.HandlerFunc(loadBalanceHandler),
	}

	go healthCheck()

	log.Printf("Load Balancer started and listening on port: %s\n", ServerPort)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server stopped unexpectedly, error:", err)
	}
}
