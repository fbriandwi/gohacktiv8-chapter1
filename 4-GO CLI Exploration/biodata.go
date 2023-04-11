package main

import (
    "fmt"
    "os"
)

// struct untuk menyimpan data teman-teman di kelas
type Teman struct {
    Nama      string
    Alamat    string
    Pekerjaan string
    Alasan string
}

// slice untuk menyimpan data teman-teman
var daftarTeman = []Teman{
    {"MUHAMAD SAGGY", "JAWA BARAT", "IT DEV", "-"},
	{"Viki Wahyudi", "JAWA TIMUR", "IT DEV", "-"},
	{"ANDREAN DWI ANDARU", "JAWA TIMUR", "IT DEV", "-"},
	{"Saluri Alam", "JAWA BARAT", "IT DEV", "-"},
	{"FAKHRUL KURNIAWAN", "BANTEN", "IT DEV", "-"},
	{"JIHAN CAMILLIA RITONGA", "SUMATERA UTARA", "IT DEV", "-"},
	{"Ikbal yaduar taupik", "JAWA BARAT", "IT DEV", "-"},
	{"AMIR MAHMUD", "JAWA TENGAH", "IT DEV", "-"},
	{"FEBRIAN DWI PUTRA", "JAWA BARAT", "IT DEV", "-"},
	{"Rizky Saputra", "JAWA BARAT", "IT DEV", "-"},
	{"ANITA LASMARIA SIAGIAN", "SUMATERA UTARA", "IT DEV", "-"},
	{"Ricky Surya Adiputra", "KALIMANTAN BARAT", "IT DEV", "-"},
	{"YUDA KURNIA NURUL FIKRI", "BANTEN", "IT DEV", "-"},
	{"ADITYA SURYADI", "JAWA BARAT", "IT DEV", "-"},
	{"RAFLI BIMA PRATANDRA", "JAWA BARAT", "IT DEV", "-"},
	{"Alvin Martin Djong", "DKI JAKARTA", "IT DEV", "-"},
	{"M FIRMANSYAH ARROZI", "JAWA TIMUR", "IT DEV", "-"},
	{"ADAM ARISTAMA", "JAWA BARAT", "IT DEV", "-"},
	{"Muhammad Anwar", "SULAWESI SELATAN", "IT DEV", "-"},
	{"Mohamad Zaelani", "JAWA TIMUR", "IT DEV", "-"},
	{"I PUTU WIRA PRATAMA PUTRA", "BALI", "IT DEV", "-"},
	{"IBRAM MUHARAM BACHRI", "DKI JAKARTA", "IT DEV", "-"},
	{"NANDA PRABU ANGGANATA", "JAWA TIMUR", "IT DEV", "-"},
	{"Yazid Ridwan", "KALIMANTAN BARAT", "IT DEV", "-"},
	{"Ahmad Anshari", "JAWA BARAT", "IT DEV", "-"},
	{"RIAN FEBRIANSYAH", "JAWA BARAT", "IT DEV", "-"},
	{"TRISNA BAYU PERMADI", "JAWA TIMUR", "IT DEV", "-"},
	{"TEZA ALFIAN", "DKI JAKARTA", "IT DEV", "-"},
	{"HENDRA ANTONIUS", "JAWA TIMUR", "IT DEV", "-"},
	{"DITHA LOZERA DEVI", "JAWA TIMUR", "IT DEV", "-"},
	{"NURHALISA MADUKUBAH", "SULAWESI TENGGARA", "IT DEV", "-"},
	{"ARGUMELAR PAMUNGKAS", "JAWA BARAT", "IT DEV", "-"},
	{"RIZKI CAHYO SASONGKO", "JAWA TIMUR", "IT DEV", "-"},
	{"Hamim Yusuf", "JAWA BARAT", "IT DEV", "-"},
	{"BENI SAFANGAT", "JAWA BARAT", "IT DEV", "-"},
	{"Ayu Putry Regita Sari", "JAWA TENGAH", "IT DEV", "-"},
	{"MUHAMMAD FAISAL PANGESTU", "BANTEN", "IT DEV", "-"},
	{"FAIZ ROFI HENCYA", "SUMATERA SELATAN", "IT DEV", "-"},
	{"Findryan Kurnia Pradana", "JAWA TIMUR", "IT DEV", "-"},
	{"APRILIA KHOIRUNNISA", "JAWA TENGAH", "IT DEV", "-"},
	{"RAMADINA AINIRIZQI GARNIZAR", "JAWA BARAT", "IT DEV", "-"},
	{"Akram Firmansyah", "SULAWESI SELATAN", "IT DEV", "-"},
	{"ONAI NADAPDAP", "SUMATERA UTARA", "IT DEV", "-"},
	{"FX. Mario Jevta Prasetio", "JAWA BARAT", "IT DEV", "-"},
	{"Muhammad Fakhryan Zulfahmi", "JAWA TIMUR", "IT DEV", "-"},
	{"MOHAMAD IQBAL AMRULLAH", "JAWA TIMUR", "IT DEV", "-"},
	{"MUHAMMAD JAUHARI", "SUMATERA SELATAN", "IT DEV", "-"},
	{"Heldi Tio Pratama", "JAWA BARAT", "IT DEV", "-"},
	{"WAHYU HAUZAN RAFI", "JAWA TENGAH", "IT DEV", "-"},
	{"M FARID ALGHOZALI}", "LAMPUNG", "IT DEV", "-"},
}

// function untuk menampilkan data teman berdasarkan nomor absen
func tampilTeman(absen int) {
    // pastikan nomor absen dalam range yang valid
    if absen < 1 || absen > len(daftarTeman) {
        fmt.Println("Nomor absen tidak valid")
        return
    }

    // tampilkan data teman
    teman := daftarTeman[absen-1]
    fmt.Println("Nama     :", teman.Nama)
    fmt.Println("Alamat   :", teman.Alamat)
    fmt.Println("Pekerjaan:", teman.Pekerjaan)
    fmt.Println("Alasan:", teman.Alasan)
}

func main() {
    // baca argument dari terminal
    args := os.Args

    // pastikan argument yang dibaca valid
    if len(args) != 2 {
        fmt.Println("Gunakan: go run biodata.go <nomor_absen>")
        return
    }

    // konversi argument ke dalam tipe data int
    absen := 0
    _, err := fmt.Sscan(args[1], &absen)
    if err != nil {
        fmt.Println("Nomor absen tidak valid")
        return
    }

    // tampilkan data teman berdasarkan nomor absen
    tampilTeman(absen)
}
