package main

import (
	"fmt"
	"strings"
)

const NMAX = 100

type TabdaftarAwal struct {
	nama, email, pass, noTelp string
}

type DaftarUser [NMAX]TabdaftarAwal
type tabelPelatihan [NMAX]struct {
	nama         string
	deskripsi    string
	tanggal      string
	kuotaPeserta int
}

type Pendaftaran struct {
	id              int
	nama            string
	email           string
	pekerjaan       string
	alasanMengikuti string
	nilai           int
	status          string // "Lulus" or "Tidak Lulus"
}

var data DaftarUser               // Global data array
var pelatihan tabelPelatihan      // Global array untuk menyimpan daftar pelatihan
var nextIndex int = 0             // Global variable to track next available index for users
var nextIdPendaftaran int = 100   // Global variable to track next available id for pendaftaran
var pendaftaran [NMAX]Pendaftaran // Array untuk menyimpan data pendaftaran
var jumlahPendaftaran int = 0     // Mencatat jumlah pendaftaran
var loggedInUserIndex int = -1    // Index of the currently logged-in user
var jumlahPelatihan int = 0       // Mencatat jumlah pelatihan

func menuUtama() {
	var nomorMenu int
	fmt.Printf("------------------------------------------------------------\n                        MENU UTAMA\n------------------------------------------------------------\n1. ADMIN\n2. PESERTA\n\n")
	var pass string
	fmt.Print("PILIH NOMOR : ")
	fmt.Scan(&nomorMenu)
	if nomorMenu == 1 {
		fmt.Print("Masukan password admin : ")
		fmt.Scan(&pass)
		if pass == "123" {
			adminMenu()
		} else {
			fmt.Print("Password salah!\n")
			menuUtama()
		}
	} else if nomorMenu == 2 {
		menuPeserta()
	} else {
		menuUtama()
	}
}

func adminMenu() {
	var nomorMenu int
	fmt.Printf("------------------------------------------------------------\n                        MENU ADMIN\n------------------------------------------------------------\n1. TAMPILKAN DATA PESERTA\n2. TAMBAH LIST PELATIHAN\n3. LIHAT PENDAFTARAN\n4. RANKING NILAI\n5. DAFTAR KELULUSAN\n6. UBAH DATA PESERTA\n7. HAPUS DATA PESERTA\n8. KEMBALI KE MENU UTAMA\n\n")
	fmt.Print("PILIH NOMOR : ")
	fmt.Scan(&nomorMenu)
	if nomorMenu == 1 {
		tampilkanDataPeserta()
	} else if nomorMenu == 2 {
		tambahListPelatihan()
	} else if nomorMenu == 3 {
		lihatPendaftaran()
	} else if nomorMenu == 4 {
		rankingNilai()
	} else if nomorMenu == 5 {
		daftarKelulusan()
	} else if nomorMenu == 6 {
		ubahDataPeserta()
	} else if nomorMenu == 7 {
		hapusDataPeserta()
	} else if nomorMenu == 8 {
		menuUtama()
	} else {
		adminMenu()
	}
}

func ubahDataPeserta() {
	fmt.Printf("------------------------------------------------------------\n                        UBAH DATA PESERTA\n------------------------------------------------------------\n")
	var email string
	fmt.Print("Masukkan email peserta yang ingin diubah: ")
	fmt.Scan(&email)

	for i := 0; i < nextIndex; i++ {
		if data[i].email == email {
			fmt.Print("Nama baru: ")
			fmt.Scan(&data[i].nama)
			fmt.Print("Email baru: ")
			fmt.Scan(&data[i].email)
			fmt.Print("Password baru: ")
			fmt.Scan(&data[i].pass)
			fmt.Println("Data peserta berhasil diubah!")
			adminMenu()
			return
		}
	}

	fmt.Println("Email peserta tidak ditemukan.")
	ubahDataPeserta()
}

func hapusDataPeserta() {
	fmt.Printf("------------------------------------------------------------\n                        HAPUS DATA PESERTA\n------------------------------------------------------------\n")
	var email string
	fmt.Print("Masukkan email peserta yang ingin dihapus: ")
	fmt.Scan(&email)

	// Cari email peserta dalam data
	indexToDelete := -1
	for i := 0; i < nextIndex; i++ {
		if data[i].email == email {
			indexToDelete = i
			break
		}
	}

	if indexToDelete == -1 {
		fmt.Println("Email peserta tidak ditemukan.")
		hapusDataPeserta() // Kembali ke menu hapus data peserta jika email tidak ditemukan
		return
	}

	// Geser data ke kiri untuk mengisi slot yang kosong
	for i := indexToDelete; i < nextIndex-1; i++ {
		data[i] = data[i+1]
	}
	nextIndex--

	fmt.Println("Data peserta berhasil dihapus!")
	adminMenu() // Kembali ke menu admin setelah menghapus data peserta
}

func tambahListPelatihan() {
	var n int
	fmt.Print("BANYAK PELATIHAN YANG MAU DITAMBAH : ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("PELATIHAN KE-%d : \n", i+1)
		fmt.Print("Nama Pelatihan: ")
		fmt.Scan(&pelatihan[jumlahPelatihan].nama)
		fmt.Print("Deskripsi: ")
		fmt.Scan(&pelatihan[jumlahPelatihan].deskripsi)
		fmt.Print("Tanggal Pelatihan: ")
		fmt.Scan(&pelatihan[jumlahPelatihan].tanggal)
		fmt.Print("Kuota Peserta: ")
		fmt.Scan(&pelatihan[jumlahPelatihan].kuotaPeserta)
		fmt.Println()
		jumlahPelatihan++
	}
	fmt.Print("BERHASIL DITAMBAHKAN")
	fmt.Println()
	adminMenu()
}

func rankingNilai() {
	fmt.Printf("------------------------------------------------------------\n                        RANKING NILAI\n------------------------------------------------------------\n")

	// Buat array untuk menyimpan data peserta dan nilainya
	var ranking [NMAX]struct {
		email string
		nilai int
	}

	// Copy data peserta dan nilainya ke array ranking
	for i := 0; i < jumlahPendaftaran; i++ {
		ranking[i].email = pendaftaran[i].email
		ranking[i].nilai = pendaftaran[i].nilai
	}

	// Urutkan array ranking menggunakan insertion sort
	for i := 1; i < jumlahPendaftaran; i++ {
		key := ranking[i]
		j := i - 1
		for j >= 0 && ranking[j].nilai < key.nilai {
			ranking[j+1] = ranking[j]
			j--
		}
		ranking[j+1] = key
	}

	// Tampilkan ranking nilai
	for i := 0; i < jumlahPendaftaran; i++ {
		fmt.Printf("Rank %d: %s - Nilai: %d\n", i+1, ranking[i].email, ranking[i].nilai)
	}

	adminMenu() // Kembali ke menu admin setelah menampilkan ranking nilai
}

func daftarKelulusan() {
	fmt.Printf("------------------------------------------------------------\n                        DAFTAR KELULUSAN\n------------------------------------------------------------\n")
	for i := 0; i < jumlahPendaftaran; i++ {
		if pendaftaran[i].status == "Lulus" {
			fmt.Printf("\nNo. %d\nID: %d\nNama: %s\nStatus: %s\n", i+1, pendaftaran[i].id, pendaftaran[i].nama, pendaftaran[i].status)
		}
	}

	fmt.Println()
	fmt.Print("Ketik 0 untuk kembali ke menu admin: ")
	var kembali int
	fmt.Scan(&kembali)
	if kembali == 0 {
		adminMenu() // Kembali ke menu admin setelah menampilkan daftar kelulusan
	} else {
		daftarKelulusan() // Kembali ke fungsi daftarKelulusan jika nomor yang dimasukkan bukan 0
	}
}

func menuPeserta() {
	var nomorMenu int
	fmt.Printf("------------------------------------------------------------\n                        MENU PESERTA\n------------------------------------------------------------\n1. MASUK\n2. DAFTAR AKUN\n3. KEMBALI KE MENU UTAMA\n\n")
	fmt.Print("PILIH NOMOR : ")
	fmt.Scan(&nomorMenu)
	if nomorMenu == 1 {
		menuLogin()
	} else if nomorMenu == 2 {
		menuSignUp()
	} else if nomorMenu == 3 {
		menuUtama()
	} else {
		menuPeserta()
	}
}

func tampilkanDataPeserta() {
	fmt.Printf("------------------------------------------------------------\n                        DATA PESERTA\n------------------------------------------------------------\n")

	// Cek apakah ada data peserta yang terdaftar
	if nextIndex == 0 {
		fmt.Println("Belum ada data peserta yang terdaftar.")
	} else {
		// Tampilkan data peserta jika ada
		for i := 0; i < nextIndex; i++ {
			fmt.Printf("No. %d\nNama: %s\nEmail: %s\n", i+1, data[i].nama, data[i].email)
		}
	}

	adminMenu()
}

func menuSignUp() {
	fmt.Printf("------------------------------------------------------------\n                        MENU SIGN UP\n------------------------------------------------------------\n")
	fmt.Print("NAMA : ")
	fmt.Scan(&data[nextIndex].nama)
	fmt.Print("EMAIL : ")
	fmt.Scan(&data[nextIndex].email)

	if !isValidEmail(data[nextIndex].email) {
		fmt.Print("--------EMAIL TIDAK VALID (HARUS @gmail.com)!!--------\n")
		menuSignUp()
		return
	}

	if cekLogin(data[nextIndex].email, "") {
		fmt.Print("--------EMAIL SUDAH TERDAFTAR !!--------\n")
		menuSignUp()
		return
	}

	fmt.Print("PASSWORD : ")
	fmt.Scan(&data[nextIndex].pass)

	nextIndex++
	menuPeserta()
}

func menuLogin() {
	var emailLogin, passLogin string
	var i int
	fmt.Print("EMAIL : ")
	fmt.Scan(&emailLogin)
	if !isValidEmail(emailLogin) {
		fmt.Print("--------EMAIL TIDAK VALID (HARUS @gmail.com)!!--------\n")
		menuLogin()
		return
	}
	fmt.Print("PASSWORD: ")
	fmt.Scan(&passLogin)
	if cekLogin(emailLogin, passLogin) {
		// Set the index of the currently logged-in user
		for i < nextIndex && loggedInUserIndex != i {
			if data[i].email == emailLogin {
				loggedInUserIndex = i

			}
			i++
		}
		dashboard()
	} else {
		fmt.Print("EMAIL ATAU PASSWORD SALAH\n")
		menuLogin()
	}
}

func dashboard() {
	fmt.Printf("------------------------------------------------------------\n                        DASBOARD\n------------------------------------------------------------\n")
	var nomorMenu int
	fmt.Printf("1. LIST PELATIHAN\n2. STATUS\n3. LOG OUT\n\n")

	fmt.Print("PILIH NOMOR: ")
	fmt.Scan(&nomorMenu)
	if nomorMenu == 1 {
		listPelatihan()
	} else if nomorMenu == 2 {
		status()
	} else if nomorMenu == 3 {
		loggedInUserIndex = -1 // Reset index pengguna yang login
		menuUtama()
	} else {
		dashboard()
	}

}

func listPelatihan() {
	fmt.Printf("------------------------------------------------------------\n                        LIST PELATIHAN\n------------------------------------------------------------\n")
	if jumlahPelatihan == 0 {
		fmt.Println("Belum ada pelatihan.")
		dashboard()
	} else {
		for i := 0; i < jumlahPelatihan; i++ {
			if pelatihan[i].nama != "" {
				fmt.Printf("%d. %s\n", i+1, pelatihan[i].nama)
			}
		}
	}
	fmt.Println()
	var nomor int
	fmt.Print("PILIH NOMOR: ")
	fmt.Scan(&nomor)
	menuPelatihan(nomor) // Panggil fungsi menuPelatihan dengan nomor pilihan sebagai argumen
}

func menuPelatihan(nomor int) {
	if nomor < 1 || nomor > jumlahPelatihan {
		fmt.Println("Nomor pelatihan tidak valid.")
		dashboard() // Kembali ke dashboard jika nomor tidak valid
		return
	}

	namaPelatihan := pelatihan[nomor-1].nama
	namaPelatihan = strings.ToUpper(namaPelatihan) // Mengonversi teks menjadi huruf kapital
	padding := (50 - len(namaPelatihan)) / 2       // Menghitung jumlah padding untuk menengahkan teks
	fmt.Printf("------------------------------------------------------------\n")
	fmt.Printf("%*s%s%*s\n", padding, "", namaPelatihan, padding, "")
	fmt.Printf("------------------------------------------------------------\n")
	fmt.Printf("Deskripsi: %s\n", pelatihan[nomor-1].deskripsi)
	fmt.Printf("Tanggal Pelatihan: %s\n", pelatihan[nomor-1].tanggal)
	fmt.Printf("Kuota Peserta: %d\n", pelatihan[nomor-1].kuotaPeserta)

	// Tambahkan pilihan untuk booking
	fmt.Println("\n1. Booking Pelatihan")
	fmt.Println("2. Kembali")

	var pilihan int
	fmt.Print("Pilih nomor: ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		formDaftarPelatihan(nomor) // Panggil fungsi formDaftarPelatihan
	} else if pilihan == 2 {
		dashboard() // Kembali ke dashboard
	} else {
		menuPelatihan(nomor) // Kembali ke menu pelatihan
	}
}

func formDaftarPelatihan(nomor int) {
	var p Pendaftaran
	fmt.Printf("------------------------------------------------------------\n                        FORMULIR DAFTAR PELATIHAN\n------------------------------------------------------------\n")
	fmt.Print("Nama: ")
	fmt.Scan(&p.nama)
	fmt.Print("Email: ")
	fmt.Scan(&p.email)

	// Check if email is valid (ends with @gmail.com)
	if !isValidEmail(p.email) {
		fmt.Println("Email tidak valid. Pastikan email diakhiri dengan @gmail.com")
		menuPelatihan(nomor) // Return to the pelatihan menu
		return
	}

	fmt.Print("Pekerjaan: ")
	fmt.Scan(&p.pekerjaan)
	fmt.Print("Alasan Mengikuti: ")
	fmt.Scan(&p.alasanMengikuti)
	fmt.Println()

	// Check if the number of participants exceeds the quota
	count := 0
	for i := 0; i < jumlahPendaftaran; i++ {
		if pendaftaran[i].email == p.email {
			count++
		}
	}
	if count >= pelatihan[nomor-1].kuotaPeserta {
		fmt.Println("Kuota peserta telah penuh.")
		dashboard()
		return
	}

	// Tambahkan ID baru ke dalam struct Pendaftaran
	p.id = nextIdPendaftaran
	nextIdPendaftaran++ // Tambahkan nomor ID untuk pendaftar berikutnya

	// Tambahkan data pendaftaran ke array (tanpa append)
	pendaftaran[jumlahPendaftaran] = p
	jumlahPendaftaran++

	fmt.Print("Pendaftaran Berhasil")
	fmt.Println()
	dashboard()
}

func status() {
	fmt.Printf("------------------------------------------------------------\n                        STATUS\n------------------------------------------------------------\n")

	// 	Tampilkan nilai hasil evaluasi dari email peserta yang sedang login
	fmt.Println("\nNilai Anda:")
	for i := 0; i < jumlahPendaftaran; i++ {
		if pendaftaran[i].email == data[loggedInUserIndex].email {
			fmt.Printf("Nilai: %d\n", pendaftaran[i].nilai)
			fmt.Printf("Status: %s\n", pendaftaran[i].status)
			break
		}
	}

	fmt.Println()
	dashboard()
}

func lihatPendaftaran() {
	fmt.Printf("------------------------------------------------------------\n                        DAFTAR PENDAFTARAN\n------------------------------------------------------------\n")

	// Cek apakah ada data pendaftaran
	if jumlahPendaftaran == 0 {
		fmt.Println("Belum ada yang daftar pelatihan.")
		adminMenu()
		return
	}

	// Tampilkan data pendaftaran
	for i := 0; i < jumlahPendaftaran; i++ {
		fmt.Printf("\nID: %d\n", pendaftaran[i].id)
		fmt.Printf("Nama: %s\n", pendaftaran[i].nama)
		fmt.Printf("Email: %s\n", pendaftaran[i].email)
		fmt.Printf("Pekerjaan: %s\n", pendaftaran[i].pekerjaan)
		fmt.Printf("Alasan Mengikuti: %s\n", pendaftaran[i].alasanMengikuti)
		fmt.Printf("Nilai: %d\n", pendaftaran[i].nilai)
		fmt.Printf("Status: %s\n", pendaftaran[i].status)
	}

	fmt.Println()
	var idPeserta int
	fmt.Print("Masukkan ID Peserta untuk menilai: ")
	fmt.Scan(&idPeserta)

	// Cari ID Peserta
	for i := 0; i < jumlahPendaftaran; i++ {
		if pendaftaran[i].id == idPeserta {
			fmt.Print("Masukkan nilai (0-100): ")
			var nilai int
			fmt.Scan(&nilai)
			fmt.Print("Masukkan status (Lulus/Tidak Lulus): ")
			var status string
			fmt.Scan(&status)
			pendaftaran[i].nilai = nilai
			pendaftaran[i].status = status
			fmt.Println("Data dinilai!")
			adminMenu()
			return
		}
	}

	fmt.Println("ID Peserta tidak valid. Sila coba lagi.")
	lihatPendaftaran() // Kembali ke menu lihat pendaftaran jika ID peserta tidak valid
}

func rankingNilaiPeserta() {
	fmt.Printf("------------------------------------------------------------\n                        RANKING NILAI\n------------------------------------------------------------\n")

	// Buat array untuk menyimpan data peserta dan nilainya
	var ranking [NMAX]struct {
		email string
		nilai int
	}

	// Copy data peserta dan nilainya ke array ranking
	for i := 0; i < jumlahPendaftaran; i++ {
		ranking[i].email = pendaftaran[i].email
		ranking[i].nilai = pendaftaran[i].nilai
	}

	// Urutkan array ranking menggunakan insertion sort
	for i := 1; i < jumlahPendaftaran; i++ {
		key := ranking[i]
		j := i - 1
		for j >= 0 && ranking[j].nilai < key.nilai {
			ranking[j+1] = ranking[j]
			j--
		}
		ranking[j+1] = key
	}

	// Tampilkan ranking nilai
	for i := 0; i < jumlahPendaftaran; i++ {
		fmt.Printf("Rank %d: %s - Nilai: %d\n", i+1, ranking[i].email, ranking[i].nilai)
	}

	fmt.Println()
	fmt.Print("Ketik 0 untuk kembali ke menu admin: ")
	var kembali int
	fmt.Scan(&kembali)
	if kembali == 0 {
		adminMenu() // Kembali ke menu admin setelah menampilkan ranking nilai
	} else {
		rankingNilai() // Kembali ke fungsi rankingNilai jika nomor yang dimasukkan bukan 0
	}
}

func cekLogin(emailLogin, passLogin string) bool {
	for i := 0; i < nextIndex; i++ {
		if emailLogin == data[i].email && (passLogin == data[i].pass || passLogin == "") {
			return true
		}
	}
	return false
}

func isValidEmail(email string) bool {
	kodeEmail := "@gmail.com"
	if len(email) < len(kodeEmail) {
		return false
	}

	// Membandingkan bagian akhir dari email dengan
	for i := 0; i < len(kodeEmail); i++ {
		if email[len(email)-len(kodeEmail)+i] != kodeEmail[i] {
			return false
		}
	}
	return true
}

func isValidPhoneNumber(phoneNumber string) bool {
	if len(phoneNumber) < 11 || len(phoneNumber) > 12 {
		return false
	}

	for i := 0; i < len(phoneNumber); i++ {
		if phoneNumber[i] < '0' || phoneNumber[i] > '9' {
			return false
		}
	}

	return true
}

func main() {
	menuUtama()
}
