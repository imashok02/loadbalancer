package loadbalancer

import (
	"fmt"
	"net/http"
	"tippermoney/personal/load-balancer/nodebox"
)

type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []nodebox.Node
}

func NewLoadBalancer(port string, nodes []nodebox.Node) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: 0,
		servers:         nodes,
	}
}

func (lb *LoadBalancer) getNextAvailableServer() nodebox.Node {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]

	for !server.IsAlive() {
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}

	lb.roundRobinCount++

	return server
}

func (lb *LoadBalancer) serverProxy(rw http.ResponseWriter, req *http.Request) {
	targetServer := lb.getNextAvailableServer()

	fmt.Printf("Forwading request to address %q\n", targetServer.Address())
	targetServer.Serve(rw, req)
}
