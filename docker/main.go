package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
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
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debugf("pac server is listening %s", listen)
  	err := http.ListenAndServe(listen, nil)
  	if err != nil {
		logrus.Fatal(err)
	}
}