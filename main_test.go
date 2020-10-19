package main

import (
	"testing"
)

func TestServerPool_AddServer(t *testing.T) {
	serverName := "andromeda"
	var sp ServerPool
	s := &Server{Name: string(serverName)}
	sp.AddServer(s)
	if sp.servers[0].Name != serverName {
		t.Errorf("Server name not added to pool expected: %s\n", serverName)
	}
}
