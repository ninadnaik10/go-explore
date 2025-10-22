package main

import "loadbalancer/servers"

func main() {
	go servers.RunServers(5)
	makeLoadBalancer(5)
}
