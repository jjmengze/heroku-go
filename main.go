package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
                p:=os.Getenv("PORT")
		w.Write([]byte(p))
	})

	http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), nil)
}
