package main

import (
	"fmt"
	"log"
	"net/http"
)


func main(){
	port := 9009

	http.HandleFunc("/barang/", func(rw http.ResponseWriter, rq *http.Request) {
		switch rq.Method {
			case "GET":
				s := rq.URL.Query().Get("bidang")
				if (s != "") {
					if s == "kineklub" {
						GetBarangKineklub(rw, rq)
					} else if s == "videografi" {
						GetBarangVideografi(rw, rq)
					} else if s == "doksos" {
						GetBarangDoksos(rw, rq)
					} else if s == "danus" {
						GetBarangDanus(rw, rq)
					} else if s == "bk" {
						GetBarangBK(rw, rq)
					} else if s == "fotografi" {
						GetBarangFotografi(rw, rq)
					} else if s == "brt" {
						GetBarangBRT(rw, rq)
					} 
				} else {
					GetAllBarang(rw, rq)
				}
			default:
				http.Error(rw,"invalid",405)
		}
	})
	
	//Kondisi Barang
	http.HandleFunc("/kondisi_barang/", func(rw http.ResponseWriter, rq *http.Request) {
		switch rq.Method {
			case "GET":
				s := rq.URL.Query().Get("kondisi")
				if s != "" {
					GetKondisiBarang(rw, rq, s)
				} else {
					GetAllBarang(rw, rq)
				}
			default:
				http.Error(rw,"invalid",405)
		}
	})

	//Lokasi Barang
	http.HandleFunc("/lokasi_barang/", func(rw http.ResponseWriter, rq *http.Request) {
		switch rq.Method {
			case "GET":
				s := rq.URL.Query().Get("lokasi")
				if (s != "") {
					log.Printf(s)
					GetLokasiBarang(rw, rq, s)
				} else {
					GetAllBarang(rw, rq)
				}
			default:
				http.Error(rw,"invalid",405)
		}
	})
	
	//Kode Barang
	http.HandleFunc("/kodebarang/", func(rw http.ResponseWriter, rq *http.Request) {
		switch rq.Method {
			case "GET":
				s := rq.URL.Query().Get("kode")
				if (s!= "") {
					GetKodeBarang(rw, rq, s)
				} else {
					GetAllBarang(rw,rq)
				}
			default:
				http.Error(rw,"Invalid Request",405)
		}
	})
	
	//Nama Barang
	http.HandleFunc("/barang/nama/", func(rw http.ResponseWriter, rq *http.Request) {
		switch rq.Method {
			case "GET":
				s:= rq.URL.Query().Get("item")
				if (s!="") {
					GetNamaBarang(rw, rq, s)
				} else {
					GetAllBarang(rw, rq)
				}
			default:
				http.Error(rw, "Invalid Request",405)
			}
	})
				
	//Create Barang
	http.HandleFunc("/barang/tambah", func(rw http.ResponseWriter, rq *http.Request) { 
		switch rq.Method {
			case "POST":
				PostBarang(rw, rq)
			default:
				http.Error(rw,"invalid",405)
		}
	})

	//Update Kondisi Barang
	http.HandleFunc("/barang/update/", func(rw http.ResponseWriter, rq *http.Request)  {
		switch rq.Method {
			case "PUT":
				s := rq.URL.Query().Get("Kode_Barang")
				UpdateBarang(rw, rq, s)

			default:
				http.Error(rw,"invalid",405)
		}
	})

	//Delete Barang
	http.HandleFunc("/barang/hapus/", func(rw http.ResponseWriter, rq *http.Request) {
		switch rq.Method {
			case "DELETE":
				s := rq.URL.Query().Get("Kode_Barang")
				DeleteBarang(rw, rq, s)
			default:
				http.Error(rw,"invalid",405)
		}
	})

	log.Printf("Server starting on port %v\n",port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port),nil))
}

