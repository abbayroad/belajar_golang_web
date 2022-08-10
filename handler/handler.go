package handler

import (
	"fmt"
	"golang_web/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles(path.Join("views", "home.html"), path.Join("views", "layout.html"))

	if err != nil {
		log.Println(err)
		return 
	}

	data := map[string]interface{}{
		"nama" : "Sukron",
		"nim" : "213040095",
	}

	tmpl.Execute(w,data)
	if err != nil {
		log.Println(err)
		return
	}
}


func ProdukHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("views", "produk.html"), path.Join("views", "layout.html"))

	if err != nil {
		log.Println(err)
		return 
	}

	

	method := r.Method;

	switch method {
	case "POST" :
		
		data := []map[string]interface{}{
			{
				"NAMA" : r.PostFormValue("nama"),
				"HARGA" : r.PostFormValue("harga"),
				"StatusStok" : "",
			},
		}
		tmpl.Execute(w,data)
	case "GET" :
		data := []entity.Produk{
			{ID: 1, NAMA: "Google Pixel", HARGA: 1500000, STOK: 3},
			{ID: 2, NAMA: "Google Pixel 2a", HARGA: 1800000, STOK: 4},
			{ID: 3, NAMA: "Google Pixel 3a", HARGA: 2500000, STOK: 1},
			{ID: 4, NAMA: "Google Pixel 4a", HARGA: 5000000, STOK: 2},
			{ID: 5, NAMA: "Google Pixel 5a", HARGA: 9000000, STOK: 5},
			{ID: 6, NAMA: "Google Pixel 6a", HARGA: 15000000, STOK: 10},
		}
		tmpl.Execute(w,data)
	default: 
		
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNum, err := strconv.Atoi(id)

	if err != nil || idNum < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Hello World : %d", idNum)
}