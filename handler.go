package main

import (
	"net/http"
	"encoding/json"
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

//=== Get All Barang ===
//Fungsi untuk mendapatkan seluruh barang yang ada dalam inventaris 
func GetAllBarang(rw http.ResponseWriter, rq *http.Request) {
	db, err := sql.Open("mysql",
		"root:14july98@tcp(127.0.0.1:3306)/inventaris")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	barang := barang{}

	rows, err := db.Query("select Kode_Barang, Nama_Barang, Kondisi, Lokasi from barang")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	//Scan data from table
	for rows.Next() {
		err := rows.Scan(&barang.Kode_Barang, &barang.Nama_Barang, &barang.Kondisi, &barang.Lokasi)
		if err != nil{
			log.Fatal(err)
		}
		json.NewEncoder(rw).Encode(&barang)
	}
}

//=== Get Barang Bidang ===
//Fungsi untuk menampilkan barang inventaris sesuai bidang
func GetBarangBidang(rw http.ResponseWriter, rq *http.Request, bidang string) {
	db, err := sql.Open("mysql",
		"root:14july98@tcp(127.0.0.1:3306)/inventaris")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	barang := barang{}

	rows, err := db.Query("select Kode_Barang, Nama_Barang, Kondisi, Lokasi from barang where Kode_Barang like?","%"+bidang+"%")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//scan data from table
	for rows.Next() {
		err := rows.Scan(&barang.Kode_Barang, &barang.Nama_Barang, &barang.Kondisi, &barang.Lokasi)
		if err != nil{
			log.Fatal(err)
		}
		json.NewEncoder(rw).Encode(&barang)
	}
}
//=== Get Nama Barang ===
//Fungsi untuk menampilkan barang sesuai dengan nama barang
func GetNamaBarang(rw http.ResponseWriter, rq *http.Request, nama string) {
	db, err := sql.Open("mysql",
		"root:14july98@tcp(127.0.0.1:3306)/inventaris")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	barang := barang{}

	rows, err := db.Query("select Kode_Barang, Nama_Barang, Kondisi, Lokasi from barang where Nama_Barang like?","%"+nama+"%")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//Scan data from table
	for rows.Next() {
		err := rows.Scan(&barang.Kode_Barang,&barang.Nama_Barang, &barang.Kondisi, &barang.Lokasi)
		if err != nil{
			log.Fatal(err)
		}
		json.NewEncoder(rw).Encode(&barang)
	}
}

//=== Get Kondisi Barang ===
//Fungsi untuk menampilkan barang sesuai kondisi barang
func GetKondisiBarang(rw http.ResponseWriter, rq *http.Request, kondisi string) {
	db, err := sql.Open("mysql",
		"root:14july98@tcp(127.0.0.1:3306)/inventaris")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	barang := barang{}
	
	rows, err := db.Query("select Kode_Barang, Nama_Barang, Kondisi, Lokasi from barang where Kondisi like?","%"+kondisi+"%")
	
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()
	
	//scan data from table
	for rows.Next() {
		err := rows.Scan(&barang.Kode_Barang, &barang.Nama_Barang, &barang.Kondisi, &barang.Lokasi)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(rw).Encode(&barang)
	}
}

//=== Get Lokasi Barang ===
//Fungsi untuk menampilkan barang sesuai dengan lokasi barang
func GetLokasiBarang(rw http.ResponseWriter, rq *http.Request, lokasi string) {
	db, err := sql.Open("mysql",
		"root:14july98@tcp(127.0.0.1:3306)/inventaris")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	barang := barang{}

	rows, err := db.Query("select Kode_Barang, Nama_Barang, Kondisi, Lokasi from barang where Lokasi like?","%"+lokasi+"%") 
	log.Printf(lokasi)
	
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()

	//scan data from table
	for rows.Next() {
		err := rows.Scan(&barang.Kode_Barang, &barang.Nama_Barang, &barang.Kondisi, &barang.Lokasi)
		
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(rw).Encode(&barang)
	}
}

func GetBarangKineklub(rw http.ResponseWriter, rq *http.Request) {
	db, err := sql.Open("mysql",
		"root:14july98@tcp(127.0.0.1:3306)/inventaris")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	barang := barang{}

	//search with query bidang kineklub
	rows, err := db.Query("select Kode_Barang, Nama_Barang, Kondisi, Lokasi from barang where Kode_Barang like '%KN%'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//Scan data from table
	for rows.Next() {
		err := rows.Scan(&barang.Kode_Barang, &barang.Nama_Barang, &barang.Kondisi, &barang.Lokasi)
		if err != nil{
			log.Fatal(err)
		}
		json.NewEncoder(rw).Encode(&barang)
	}
}

func GetBarangVideografi(rw http.ResponseWriter, rq *http.Request) {
	db, err := sql.Open("mysql",
		"root:14july98@tcp(127.0.0.1:3306)/inventaris")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	barang := barang{}

	//search with query bidang videografi
	rows, err := db.Query("select Kode_Barang, Nama_Barang, Kondisi, Lokasi from barang where Kode_Barang like '%VD%'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//Scan data from table
	for rows.Next() {
		err := rows.Scan(&barang.Kode_Barang, &barang.Nama_Barang, &barang.Kondisi, &barang.Lokasi)
		if err != nil{
			log.Fatal(err)
		}
		json.NewEncoder(rw).Encode(&barang)
	}

}

func GetBarangDoksos(rw http.ResponseWriter, rq *http.Request) {
	db, err := sql.Open("mysql",
		"root:14july98@tcp(127.0.0.1:3306)/inventaris")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	barang := barang{}

	//search with query doksos
	rows, err := db.Query("select Kode_Barang, Nama_Barang, Kondisi, Lokasi from barang where Kode_Barang like '%DS%'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//scan data from table
	for rows.Next() {
		err := rows.Scan(&barang.Kode_Barang, &barang.Nama_Barang, &barang.Kondisi, &barang.Lokasi)
		if err != nil{
			log.Fatal(err)
		}
		json.NewEncoder(rw).Encode(&barang)
	}

}
func GetBarangDanus(rw http.ResponseWriter, rq *http.Request) {
	db, err := sql.Open("mysql",
		"root:14july98@tcp(127.0.0.1:3306)/inventaris")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	barang := barang{}
	
	//search with query danus
	rows, err := db.Query("select Kode_Barang, Nama_Barang, Kondisi, Lokasi from barang where Kode_Barang like '%DU%'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//scan data from table
	for rows.Next() {
		err := rows.Scan(&barang.Kode_Barang, &barang.Nama_Barang, &barang.Kondisi, &barang.Lokasi)
		if err != nil{
			log.Fatal(err)
		}
		json.NewEncoder(rw).Encode(&barang)
	}
}

func GetBarangBK(rw http.ResponseWriter, rq *http.Request) {
	db, err := sql.Open("mysql",
		"root:14july98@tcp(127.0.0.1:3306)/inventaris")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	barang := barang{}

	//search with query bk
	rows, err := db.Query("select Kode_Barang, Nama_Barang, Kondisi, Lokasi from barang where Kode_Barang like '%BK%'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//scan data from table
	for rows.Next() {
		err := rows.Scan(&barang.Kode_Barang, &barang.Nama_Barang, &barang.Kondisi, &barang.Lokasi)
		if err != nil{
			log.Fatal(err)
		}
		json.NewEncoder(rw).Encode(&barang)
	}
}
func GetBarangFotografi(rw http.ResponseWriter, rq *http.Request) {
	db, err := sql.Open("mysql",
		"root:14july98@tcp(127.0.0.1:3306)/inventaris")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	barang := barang{}

	//search with query fotografi
	rows, err := db.Query("select Kode_Barang, Nama_Barang, Kondisi, Lokasi from barang where Kode_Barang like '%FT%'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//scan data from table
	for rows.Next() {
		err := rows.Scan(&barang.Kode_Barang, &barang.Nama_Barang, &barang.Kondisi, &barang.Lokasi)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(rw).Encode(&barang)
	}
}

func GetBarangBRT(rw http.ResponseWriter, rq *http.Request) {
	db, err := sql.Open("mysql",
		"root:14july98@tcp(127.0.0.1:3306)/inventaris")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	barang := barang{}

	//search with query brt
	rows, err := db.Query("select Kode_Barang, Nama_Barang, Kondisi, Lokasi from barang where Kode_Barang like '%RT%'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//scan data from table
	for rows.Next() {
		err := rows.Scan(&barang.Kode_Barang, &barang.Nama_Barang, &barang.Kondisi, &barang.Lokasi)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(rw).Encode(&barang)
	}
}
//=== Get Kode Barang ===
//Fungsi untuk menampilkan barang inventaris sesuai dengan kode barang
func GetKodeBarang(rw http.ResponseWriter, rq *http.Request, kode string) {
	db , err := sql.Open("mysql",
		"root:14july98@tcp(127.0.0.1:3306)/inventaris")
	if err != nil {
		log.Fatal(err)
	}
	defer rq.Body.Close()

	barang := barang{}

	rows, err := db.Query("select Kode_Barang, Nama_Barang, Kondisi, Lokasi from barang where Kode_Barang like?",kode)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//scan data from table
	for rows.Next() {
		err := rows.Scan(&barang.Kode_Barang, &barang.Nama_Barang, &barang.Kondisi, &barang.Lokasi)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(rw).Encode(&barang)
	}
}

//=== Post Barang ===
//Fungsi untuk memasukkan barang baru ke dalam inventaris
func PostBarang (rw http.ResponseWriter, rq *http.Request) {
	var brg barang
	brgDecoder := json.NewDecoder(rq.Body)
	err := brgDecoder.Decode(&brg)

	if err != nil{
		log.Fatal(err)
	}
	defer rq.Body.Close()

	db , err := sql.Open("mysql",
		"root:14july98@tcp(127.0.0.1:3306)/inventaris")
	if err != nil{
		log.Fatal(err)
	}
	
	st, err := db.Prepare("INSERT INTO barang (Kode_Barang, Nama_Barang, Kondisi, Lokasi) VALUES (?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = st.Exec(brg.Kode_Barang, brg.Nama_Barang, brg.Kondisi, brg.Lokasi)
}

//=== Update Kondisi Barang ===
//Fungsi untuk mengupdate kondisi dari barang inventaris.
func UpdateBarang(rw http.ResponseWriter, rq *http.Request, Kode_Barang string)  {
	var brg barang
	kb := Kode_Barang
	brgDecoder := json.NewDecoder(rq.Body)
	err := brgDecoder.Decode(&brg)

	if err != nil{
		log.Fatal(err)
	}
	defer rq.Body.Close()

	db ,  err := sql.Open("mysql",
		"root:14july98@tcp(127.0.0.1:3306)/inventaris")
	if err != nil{
		log.Fatal(err)
	}

	st, err := db.Prepare("UPDATE barang SET Kondisi = ? where Kode_Barang like ?")
	if err != nil{
		log.Fatal(err)
	}

	_, err = st.Exec(brg.Kondisi, kb)
}

//=== Delete Barang ===
//Fungsi untuk menghapus sebuah barang dari inventaris
func DeleteBarang(rw http.ResponseWriter, rq *http.Request, Kode_Barang string) {
	 
	db , err := sql.Open("mysql",
		"root:14july98@tcp(127.0.0.1:3306)/inventaris")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("DELETE from barang where Kode_Barang like ?",Kode_Barang)
	defer rows.Close()
}
	
