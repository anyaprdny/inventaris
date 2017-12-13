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
				if rq.URL.Query().Get("bidang") != "" {
					bid := rq.URL.Query().Get("bidang")
					if (bid != "") {
					
						if bid == "kine" {
							GetBarangKineklub(rw, rq)
						} else if bid == "video" {
							GetBarangVideografi(rw, rq)
						} else if bid == "doksos" {
							GetBarangDoksos(rw, rq)
						} else if bid == "danus" {
							GetBarangDanus(rw, rq)
						} else if bid == "bk" {
							GetBarangBK(rw, rq)
						} else if bid == "foto" {
							GetBarangFoto(rw, rq)
						} else if bid == "brt" {
							GetBarangBRT(rw, rq)
						} else {
							GetBarangBidang(rw, rq, bid)
						}
					} else {
						GetAllBarang(rw, rq)
					}
				} else if rq.URL.Query().Get("nama") != "" {
					nama := rq.URL.Query().Get("nama")
					GetNamaBarang(rw, rq, nama) 
				} else if rq.URL.Query().Get("kode") != "" {
					kode := rq.URL.Query().Get("kode")
					GetKodeBarang(rw, rq, kode)
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
				if rq.URL.Query().Get("kode") != ""{
					kode := rq.URL.Query().Get("kode")
					if kode != "" {
				 		GetKondisiByKodeBarang(rw, rq, kode)
					} else {
						GetAllBarang(rw, rq)
					}
				} else if rq.URL.Query().Get("kondisi") != "" {
					kondisi := rq.URL.Query().Get("kondisi")
					if kondisi != "" {
						GetKondisiBarang(rw, rq, kondisi)
					} else {
						GetAllBarang(rw, rq)
					}
				}
				
			case "PUT":
				kode := rq.URL.Query().Get("update")
				UpdateKondisi(rw, rq, kode)
			default:
				http.Error(rw,"invalid",405)
		}
	})

	//Lokasi Barang
	http.HandleFunc("/lokasi_barang/", func(rw http.ResponseWriter, rq *http.Request) {
		switch rq.Method {
			case "GET":
				lokasi := rq.URL.Query().Get("lokasi")
				if (lokasi != "") {
					GetLokasiBarang(rw, rq, lokasi)
				} else {
					GetAllBarang(rw, rq)
				}
			case "PUT":
				kode := rq.URL.Query().Get("update")
				UpdateLokasi(rw, rq, kode)
			default:
				http.Error(rw,"invalid",405)
		}
	})
	
	//Kode Barang
	http.HandleFunc("/kodebarang/", func(rw http.ResponseWriter, rq *http.Request) {
		switch rq.Method {
			case "GET":
				kode := rq.URL.Query().Get("kode")
				if (kode != "") {
					GetKodeBarang(rw, rq, kode)
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
				item := rq.URL.Query().Get("item")
				if (item !="") {
					GetNamaBarang(rw, rq, item)
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

