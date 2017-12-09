package main

import (
	"net/http"
	"encoding/json"
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllBarang(rw http.ResponseWriter, rq *http.Request) {
	db, err := sql.Open("mysql",
		"root:14july98@tcp(127.0.0.1:3306)/inventaris")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	barang := barang{}

	rows, err := db.Query("select Nama_Barang, Kode_Barang, Kondisi, Lokasi from barang")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	//Scan data from table
	for rows.Next() {
		err := rows.Scan(&barang.Nama_Barang, &barang.Kode_Barang, &barang.Kondisi, &barang.Lokasi)
		if err != nil{
			log.Fatal(err)
		}
		json.NewEncoder(rw).Encode(&barang)
	}

//	type BarangInv []Barang

//	bi := BarangInv{
//	Barang{Nama_Barang: "DVD 3 Hari Untuk Selamanya",Kode_Barang: "KN-01", Kondisi: "Baik", Lokasi: "Di Atas TV Ruang Santai"},
//
//	json.NewEncoder(rw).Encode(bi)
// }	
}
