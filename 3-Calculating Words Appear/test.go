package main

import (
    "fmt"
    "sort"
)

func main() {
    kalimat := "selamat malam"
    huruf := []rune(kalimat)
    kemunculan := make(map[rune]int)

    for _, h := range huruf {
		fmt.Println(string(h))
        kemunculan[h]++
    }

    // konversi map menjadi slice untuk diurutkan
    var kemunculanSlice []KemunculanHuruf
    for k, v := range kemunculan {
        kemunculanSlice = append(kemunculanSlice, KemunculanHuruf{k, v})
    }

    // mengurutkan slice berdasarkan jumlah kemunculan huruf
    sort.Slice(kemunculanSlice, func(i, j int) bool {
        return kemunculanSlice[i].Jumlah > kemunculanSlice[j].Jumlah
    })

    fmt.Println("Huruf terbanyak:")
    for _, k := range kemunculanSlice {
        fmt.Printf("%c: %d\n", k.Huruf, k.Jumlah)
    }
}

type KemunculanHuruf struct {
    Huruf   rune
    Jumlah  int
}