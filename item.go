package main

type barang struct {
	Kode_Barang	string
	Nama_Barang	string
	Kondisi		string
	Lokasi		string
}

type bidang struct {
	Nama_Bidang 	string
	ID_Bidang	string
}

type inventarisbarang struct {
	ID_Bidang	string
	Kode_Barang	string
}
