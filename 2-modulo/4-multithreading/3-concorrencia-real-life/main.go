package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync/atomic"
)

var number uint64 = 0

func main() {
	// m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		// number++
		atomic.AddUint64(&number, 1)
		// m.Unlock()
		w.Write([]byte("Voce teve acesso a essa pagina " + strconv.FormatUint(number, 10) + " vezes"))
	})

	fmt.Println("Servidor rodando na porta 8080")
	http.ListenAndServe(":8080", nil)
}
