package servers

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type ServerList struct {
	Ports []int
}

func (s *ServerList) Populate(amount int) {
	if amount >= 10 {
		log.Fatal("Amount of ports can't exceed 10")
	}
	for x := range amount {
		s.Ports = append(s.Ports, x)
	}
}

func RunServers(amount int) {
	sl := ServerList{}
	sl.Populate(amount)

	wg := sync.WaitGroup{}
	wg.Add(amount)
	defer wg.Wait()

	for range amount {
		go makeServer(&sl, wg)
	}
}

func makeServer(sl *ServerList, wg sync.WaitGroup) {
	defer wg.Done()
	port := sl.Ports[0]
	sl.Ports = sl.Ports[1:]

	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Server %d", port)
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":808%d", port),
		Handler: r,
	}

	server.ListenAndServe()

}
