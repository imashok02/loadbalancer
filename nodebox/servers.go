package nodebox

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Server struct {
	address string
	proxy   *httputil.ReverseProxy
}

func NewServer(address string) *Server {
	serverUri, _ := url.Parse(address)

	// handleErr(err)

	return &Server{
		address: address,
		proxy:   httputil.NewSingleHostReverseProxy(serverUri),
	}
}

type Node interface {
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, req *http.Request)
}

func (server *Server) Address() string {
	return server.address
}

func (server *Server) IsAlive() bool {
	return true
}

func (server *Server) Serve(rw http.ResponseWriter, req *http.Request) {
	server.proxy.ServeHTTP(rw, req)
}
