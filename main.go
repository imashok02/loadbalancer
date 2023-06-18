package main

import (
	"fmt"
	"net/http"
	"tippermoney/personal/load-balancer/loadbalancer"
	"tippermoney/personal/load-balancer/nodebox"
)

func main() {
	servers := []nodebox.Node{
		nodebox.NewServer("https://www.google.com"),
		nodebox.NewServer("https://www.bing.com"),
		nodebox.NewServer("https://www.duckduckgo.com"),
	}

	lb := loadbalancer.NewLoadBalancer("8000", servers)

	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.serverProxy(rw, req)
	}

	http.HandleFunc("/", handleRedirect)

	fmt.Printf("Servind requests at port at localhost:%s\n", lb.port)

	http.ListenAndServe(":"+lb.port, nil)
}

func handleErr(err error) {
	if err != nil {
		fmt.Println("Errored ==> ", err)
	}
}
