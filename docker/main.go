package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("./pac"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/pac.js" {
			w.Header().Set("Content-Type", "text/javascript; charset=UTF-8")
		}
		fs.ServeHTTP(w, r)
	})

	listen := fmt.Sprintf("%s:%s", "0.0.0.0", os.Getenv("LISTEN_PORT"))
	fmt.Printf("pac server is listening %s\n", listen)
  	err := http.ListenAndServe(listen, nil)
  	if err != nil {
		log.Fatal(err)
	}
}