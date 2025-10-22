package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

var (
	baseUrl = "http://localhost:808"
)

type LoadBalancer struct {
	RevProxy httputil.ReverseProxy
}

type Endpoints struct {
	List []*url.URL
}

func (e *Endpoints) Shuffle() {
	temp := e.List[0]
	e.List = e.List[1:]
	e.List = append(e.List, temp)
}

func makeLoadBalancer(amount int) {
	var lb LoadBalancer
	var ep Endpoints
	r := http.NewServeMux()
	server := http.Server{
		Addr:    ":8090",
		Handler: r,
	}

	// Create endpoints
	for i := range amount {
		ep.List = append(ep.List, createEndpoint(baseUrl, i))
	}

	r.HandleFunc(
		"/loadbalance",
		makeRequest(&lb, &ep),
	)

	server.ListenAndServe()
}

func makeRequest(lb *LoadBalancer, ep *Endpoints) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		lb.RevProxy = *httputil.NewSingleHostReverseProxy(ep.List[0])
		ep.Shuffle()
		lb.RevProxy.ServeHTTP(w, r)
	}
}

func createEndpoint(endpoint string, idx int) *url.URL {
	link := endpoint + strconv.Itoa(idx)
	url, _ := url.Parse(link)
	return url
}
