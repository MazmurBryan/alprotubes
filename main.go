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
	nama            string
	email           string
	pekerjaan       string
	alasanMengikuti string
	nilai           int
	status          string // "Lulus" or "Tidak Lulus"
}

var data DaftarUser               // Global data array
var pelatihan tabelPelatihan      // Global array untuk menyimpan daftar pelatihan
var nextIndex int = 0             // Global variable to track next available index
var pendaftaran [NMAX]Pendaftaran // Array untuk menyimpan data pendaftaran
var jumlahPendaftaran int = 0     // Mencatat jumlah pendaftaran

func menuUtama() {
	var nomorMenu int
	fmt.Printf("------------------------------------------------------------\n                        MENU UTAMA\n------------------------------------------------------------\n1. ADMIN\n2. PESERTA\n\n")

	fmt.Print("PILIH NOMOR : ")
	fmt.Scan(&nomorMenu)
	if nomorMenu == 1 {
		menuAdmin()
	} else if nomorMenu == 2 {
		menuPeserta()
	} else {
		menuUtama()
	}
}

func menuAdmin() {
	var nomorMenu int
	fmt.Printf("------------------------------------------------------------\n                        MENU ADMIN\n------------------------------------------------------------\n1. TAMPILKAN DATA PESERTA\n2. TAMBAH LIST PELATIHAN\n3. LIHAT PENDAFTARAN\n4. KEMBALI KE MENU UTAMA\n\n")
	fmt.Print("PILIH NOMOR : ")
	fmt.Scan(&nomorMenu)
	if nomorMenu == 1 {
		tampilkanData()
	} else if nomorMenu == 3 {
		lihatPendaftaran()
	} else if nomorMenu == 2 {
		tambahListPelatihan()
	} else if nomorMenu == 4 {
		menuUtama()
	}
}

func tambahListPelatihan() {
	var n int
	fmt.Print("BANYAK PELATIHAN YANG MAU DITAMBAH : ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("PELATIHAN KE-%d : \n", i+1)
		fmt.Print("Nama Pelatihan: ")
		fmt.Scan(&pelatihan[i].nama)
		fmt.Print("Deskripsi: ")
		fmt.Scan(&pelatihan[i].deskripsi)
		fmt.Print("Tanggal Pelatihan: ")
		fmt.Scan(&pelatihan[i].tanggal)
		fmt.Print("Kuota Peserta: ")
		fmt.Scan(&pelatihan[i].kuotaPeserta)
		fmt.Println()
	}
	fmt.Print("BERHASIL DITAMBAHKAN")
	fmt.Println()
	menuAdmin()
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
	}
}

func menuSignUp() {
	fmt.Printf("------------------------------------------------------------\n                        MENU SIGN UP\n------------------------------------------------------------\n")
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
	switch nomorMenu {
	case 1:
		listPelatihan()
	case 2:
		status()
	case 3:
		menuUtama()
	default:
		dashboard()
	}
}

func listPelatihan() {
	fmt.Printf("------------------------------------------------------------\n                        LIST PELATIHAN\n------------------------------------------------------------\n")
	for i := 0; i < len(pelatihan); i++ {
		if pelatihan[i].nama != "" {
			fmt.Printf("%d. %s\n", i+1, pelatihan[i].nama)
		}
	}
	fmt.Println()
	var nomor int
	fmt.Print("PILIH NOMOR: ")
	fmt.Scan(&nomor)
	menuPelatihan(nomor) // Panggil fungsi menuPelatihan dengan nomor pilihan sebagai argumen
}

func menuPelatihan(nomor int) {
	if nomor < 1 || nomor > len(pelatihan) {
		fmt.Println("Nomor pelatihan tidak valid.")
		menuAdmin() // Kembali ke menu admin jika nomor tidak valid
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
		menuAdmin() // Kembali ke menu admin
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
	fmt.Print("Pekerjaan: ")
	fmt.Scan(&p.pekerjaan)
	fmt.Print("Alasan Mengikuti: ")
	fmt.Scan(&p.alasanMengikuti)
	fmt.Println()

	// Tambahkan data pendaftaran ke array (tanpa append)
	pendaftaran[jumlahPendaftaran] = p
	jumlahPendaftaran++

	fmt.Print("Pendaftaran Berhasil")
	fmt.Println()
	menuPeserta()
}

func status() {
	fmt.Printf("------------------------------------------------------------\n                        STATUS\n------------------------------------------------------------\n")
	for i := 0; i < jumlahPendaftaran; i++ {
		fmt.Printf("\nNama: %s\n", pendaftaran[i].nama)
		fmt.Printf("Email: %s\n", pendaftaran[i].email)
		fmt.Printf("Pekerjaan: %s\n", pendaftaran[i].pekerjaan)
		fmt.Printf("Alasan Mengikuti: %s\n", pendaftaran[i].alasanMengikuti)

		if pendaftaran[i].nilai != 0 {
			fmt.Printf("Nilai: %d\n", pendaftaran[i].nilai)
			fmt.Printf("Status: %s\n", pendaftaran[i].status)
		} else {
			fmt.Printf("Nilai: Belum Dinilai\n")
			fmt.Printf("Status: Belum Diputuskan\n")
		}
	}

	fmt.Println()
	menuPeserta()
}

func lihatPendaftaran() {
	fmt.Printf("------------------------------------------------------------\n                        DAFTAR PENDAFTARAN\n------------------------------------------------------------\n")
	for i := 0; i < jumlahPendaftaran; i++ {
		fmt.Printf("\nNama: %s\n", pendaftaran[i].nama)
		fmt.Printf("Email: %s\n", pendaftaran[i].email)
		fmt.Printf("Pekerjaan: %s\n", pendaftaran[i].pekerjaan)
		fmt.Printf("Alasan Mengikuti: %s\n", pendaftaran[i].alasanMengikuti)
		fmt.Printf("Nilai: %d\n", pendaftaran[i].nilai)
		fmt.Printf("Status: %s\n", pendaftaran[i].status)
	}

	fmt.Println()
	var nomor, nilai int
	var status string
	fmt.Print("Masukkan nomor pendaftaran untuk menilai: ")
	fmt.Scan(&nomor)

	if nomor >= 0 && nomor < jumlahPendaftaran {
		fmt.Print("Masukkan nilai (0-100): ")
		fmt.Scan(&nilai)
		fmt.Print("Masukkan status (Lulus/Tidak Lulus): ")
		fmt.Scan(&status)

		pendaftaran[nomor].nilai = nilai
		pendaftaran[nomor].status = status

		fmt.Println("Data dinilai!")
	} else {
		fmt.Println("Nomor pendaftaran tidak valid.")
	}
	menuAdmin()
}

func cekLogin(emailLogin, passLogin string) bool {
	for i := 0; i < nextIndex; i++ {
		if emailLogin == data[i].email && (passLogin == data[i].pass || passLogin == "") {
			return true
		}
	}
	return false
}

func tampilkanData() {
	fmt.Printf("------------------------------------------------------------\n                        DATA PESERTA\n------------------------------------------------------------\n")
	for i := 0; i < nextIndex; i++ {
		fmt.Printf("Nama: %s\nEmail: %s\nPassword: %s\n", data[i].nama, data[i].email, data[i].pass)
	}
	menuAdmin()
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
