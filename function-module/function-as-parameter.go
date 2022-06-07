package main

import "fmt"

// todo: type declaration di gunakan sebagai alias

type Filter func(string) string
type FilterNomor func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	namedFilter := filter(name)
	fmt.Println("Hello", namedFilter)
}

func nomorKartu(nomorkartu string, filter FilterNomor) {
	nomorkartunasabah := filter(nomorkartu)
	fmt.Println("Nomor Kartu : ", nomorkartunasabah)
}

func spamFilter(name string) string {
	if name == "...hmm kasar banget ya kamu ngomongnya" {
		return "hmm kasar"
	} else {
		return name
	}
}

func spamKartu(nomorkartu string) string {
	if nomorkartu == "ups salah" {
		return "silahkan input ulang"
	} else {
		return nomorkartu
	}
}

func main() {
	sayHelloWithFilter("Dimas", spamFilter)
	sayHelloWithFilter("Ngentod", spamFilter)
	nomorKartu("silahkan masukan no kartu anda : ", spamKartu)

}
