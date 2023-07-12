package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

var number uint64 = 0

func main() {
	//m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//m.Lock()
		//number++
		//m.Unlock()
		atomic.AddUint64(&number, 1)
		w.Write([]byte(fmt.Sprintf("Você teve %d acessos a essa página.", number)))
		log.Printf("Você teve %d acessos a essa página.", number)
	})
	http.ListenAndServe(":3000", nil)
}
