package peserta

import "fmt"

type Person struct {
	noUrut     int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

var Murid = []Person{
	{noUrut: 1, nama: "Ciaran Cote", alamat: "Jakarta", pekerjaan: "Back-end engineer", alasan: "Tertarik dengan golang"},
	{noUrut: 2, nama: "Neville Pitts", alamat: "Tangerang", pekerjaan: "UI/UX Engineer", alasan: "Coba Coba ajh"},
	{noUrut: 3, nama: "Tobias O'brien", alamat: "Bogor", pekerjaan: "Front End Engineer", alasan: "Karena memang mau belajar semua materi BE"},
	{noUrut: 4, nama: "Adele Benjamin", alamat: "Bekasi", pekerjaan: "Back-end engineer", alasan: "Mengikuti trend"},
	{noUrut: 5, nama: "Lysandra Morrow", alamat: "Depok", pekerjaan: "UI/UX End Engineer", alasan: "golang is easy"},
	{noUrut: 6, nama: "Murphy Roth", alamat: "Serang", pekerjaan: "Front End Engineer", alasan: "Powerfull"},
}

func (p Person) Cetak() {
	fmt.Printf("Nama: %v\n", p.nama)
	fmt.Printf("Alamat: %v\n", p.alamat)
	fmt.Printf("Pekerjaan: %v\n", p.pekerjaan)
	fmt.Printf("Alasan memilih kelas golang: %v\n", p.alasan)
}
