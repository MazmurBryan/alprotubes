package main

import (
	"fmt"
)

const NMAX = 100

type TabdaftarAwal struct {
	nama, email, pass, noTelp string
}

type DaftarUser [NMAX]TabdaftarAwal

var data DaftarUser   // Global data array
var nextIndex int = 0 // Global variable to track next available index

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

	// 	fmt.Print("NOMOR TELEPON : ")
	// 	fmt.Scan(&data[nextIndex].noTelp)

	// 	if !isValidPhoneNumber(data[nextIndex].noTelp) {
	// 		fmt.Print("--------NOMOR TELEPON TIDAK VALID (HARUS 11-12 DIGIT)!!--------\n")
	// 		menuSignUp()
	// 		return
	// 	}

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
		dasboard()
	} else {
		fmt.Print("EMAIL ATAU PASSWORD SALAH\n")
		menuLogin()
	}
}

func dasboard() {
	fmt.Printf("------------------------------------------------------------\n                        DASBOARD\n------------------------------------------------------------\n")
	var nomorMenu int
	fmt.Printf("1. LIST PELATIHAN\n2. STATUS/EVALUASI\n3. LOG OUT\n\n")
	fmt.Print("PILIH NOMOR: ")
	fmt.Scan(&nomorMenu)
	if nomorMenu == 1 {
		listPelatihan()
	} else if nomorMenu == 2 {
		statusEvaluasi()
	} else if nomorMenu == 3 {
		menuUtama()
	} else {
		dasboard()
	}
}

func listPelatihan() {
	fmt.Print("LIST PELATIHAN")
}

func statusEvaluasi() {
	fmt.Print("------------------------------------------------------------\n                        EVALUASI\n------------------------------------------------------------\n\n")
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
