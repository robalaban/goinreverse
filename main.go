package main

import (
	"goinreverse/parsers"
	"log"
	"net/http/httputil"
	"net/url"
	"sync"
)

var config = parsers.ParseConfig()

type Server struct {
	URL *url.URL
	Name string
	Status bool
	mux sync.RWMutex
	Proxy *httputil.ReverseProxy
}

type ServerPool struct {
	servers []*Server
	current uint32
}

//AddServer - Adds individual server to the pool of available servers
func (s *ServerPool) AddServer(server *Server) {
	log.Printf("Added server %s to pool", server.Name)
	s.servers = append(s.servers, server)
}

var serverPool ServerPool

func main() {
	servers := config.Servers

	for _, server := range servers {
		serverUrl, err := url.Parse(server.Url)
		if err != nil {
			log.Fatal("Unable to parse url, error:", err)
		}

		proxy := httputil.NewSingleHostReverseProxy(serverUrl)

		serverPool.AddServer(&Server {
			URL: serverUrl,
			Name: server.Name,
			Status: true,
			Proxy: proxy,
		})
	}

	log.Print(config.Servers)
}
