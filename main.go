package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	port := 9009

	http.HandleFunc("/barang", func(rw http.ResponseWriter, rq *http.Request) {
		switch rq.Method {
			case "GET":
				s := rq.URL.Query().Get("bidang")
				if (s!="") {
					//do nothing
				} else {
					GetAllBarang(rw, rq)
				}
			default:
				http.Error(rw,"invalid",405)
		}
	})

	http.HandleFunc("/kineklub", func(rw http.ResponseWriter, rq *http.Request) {
		switch rq.Method {
			case "GET":
				GetBarangKineklub(rw, rq)
			default:
				http.Error(rw,"invalid",405)
		}
	})

	log.Printf("Server starting on port %v\n",port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port),nil))
}

