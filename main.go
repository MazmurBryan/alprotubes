package main

import "fmt"

const NMAX = 100

type TabdaftarAwal struct {
	nama, email, pass, noTelp string
}

// Alias for the data type
type DaftarUser [NMAX]TabdaftarAwal

var data DaftarUser   // Global data array
var nextIndex int = 0 // Global variable to track next available index

func menuUtama() {
	//kamus
	var nomorMenu int

	//kamus
	fmt.Printf("------------------------------------------------------------\n                        MENU UTAMA\n------------------------------------------------------------\n1. ADMIN\n2. PESERTA\n\n")

	fmt.Print("PILIH NOMOR : ")
	fmt.Scan(&nomorMenu)
	if nomorMenu == 1 {
		menuAdmin()
	} else if nomorMenu == 2 {
		menuPeserta()
	} else {
		menuUtama() // Pass data to menuUtama if invalid input
	}
}

func menuAdmin() {
	var nomorMenu int
	//fmt.Print("Anda Di Menu Peserta")
	fmt.Printf("------------------------------------------------------------\n                        MENU ADMIN\n------------------------------------------------------------\n1. TAMPILKAN DATA\n2. -----\n3. KEMBALI KE MENU UTAMA\n\n")
	fmt.Print("PILIH NOMOR : ")
	fmt.Scan(&nomorMenu)
	if nomorMenu == 1 {
		tampilkanData()
	} else if nomorMenu == 3 {
		menuUtama()
	}
}

func menuPeserta() {
	var nomorMenu int
	//fmt.Print("Anda Di Menu Peserta")
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
	fmt.Print("NAMA : ")
	fmt.Scan(&data[nextIndex].nama)
	fmt.Print("EMAIL : ")
	fmt.Scan(&data[nextIndex].email)
	fmt.Print("PASSWORD : ")
	fmt.Scan(&data[nextIndex].pass)
	fmt.Print("NOMOR TELEPON : ")
	fmt.Scan(&data[nextIndex].noTelp)

	nextIndex++ // ++ agar bergerak beda tabel
	menuPeserta()
}

func menuLogin() {
	var emailLogin, passLogin string
	fmt.Print("EMAIL : ")
	fmt.Scan(&emailLogin)
	fmt.Print("PASSWORD: ")
	fmt.Scan(&passLogin)
	cekLogin(emailLogin, passLogin)
	if cekLogin(emailLogin, passLogin) == true {
		menuUtama()
	} else {
		fmt.Print("EMAIL ATAU PASSWORD SALAH")
		menuLogin()
	}

}
func cekLogin(emailLogin, passLogin string) bool { //sequencial search
	var status bool
	var i int
	i = 0
	status = false

	for i < NMAX && emailLogin == data[i].email && passLogin == data[i].pass { //&& emailLogin==data[i].email && passLogin == data[i].pass
		if emailLogin == data[i].email && passLogin == data[i].pass {
			status = true

		}
		i++

	}
	return status
}

//menu Admin
func tampilkanData() { //menu admin
	fmt.Printf("------------------------------------------------------------\n                        DATA PESERTA\n------------------------------------------------------------\n")
	for i := 0; i < nextIndex; i++ {
		if data[i].nama != "" {
			fmt.Printf("Nama: %s\nEmail: %s\nPassword: %s\nNo Telp: %s\n\n", data[i].nama, data[i].email, data[i].pass, data[i].noTelp)
		}
	}
}

func main() {
	menuUtama()
}
