package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./pac"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("request: %s %s", r.Method, r.URL.Path)
		if r.URL.Path == "/pac.js" {
			w.Header().Set("Content-Type", "text/javascript; charset=UTF-8")
		}
		fs.ServeHTTP(w, r)
	})

	listen := "0.0.0.0:10801"
	fmt.Printf("pac server is listening %s\n", listen)
	err := http.ListenAndServe(listen, nil)
	if err != nil {
		log.Fatal(err)
	}
}
