package entity

type Produk struct {
	ID    int
	NAMA  string
	HARGA int
	STOK  int
}

func (p Produk) StatusStok() string {
	var status string
	if p.STOK < 3 {
		status = "Stok Hampir Habis"
	}
	return status
}