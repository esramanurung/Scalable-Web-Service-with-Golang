package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func newListBiodata(student ...Biodata) []Biodata {
	var listBiodata = []Biodata{}
	for _, biodataStudent := range student {
		listBiodata = append(listBiodata, biodataStudent)
	}
	return listBiodata
}

func printBiodata(biodata Biodata) {
	fmt.Println("Nama : ", biodata.Nama)
	fmt.Println("Alamat : ", biodata.Alamat)
	fmt.Println("Pekerjaan : ", biodata.Pekerjaan)
	fmt.Println("Alasan  : ", biodata.Alasan)
}

func main() {
	nim1 := Biodata{
		Nama:      "M. Arsyad Ramadhan",
		Alamat:    "Jakarta",
		Pekerjaan: "Backend Engineer-I",
		Alasan:    "Mempelajari Technology Baru Di Company saat ini",
	}

	nim2 := Biodata{
		Nama:      "Tiara Dewangga",
		Alamat:    "Jakarta",
		Pekerjaan: "Backend Engineer-I",
		Alasan:    "Mempelajari Technology Baru Di Company saat ini",
	}
	nim3 := Biodata{
		Nama:      "Muhamad Irsyad Rafi Sudirjo",
		Alamat:    "Jakarta",
		Pekerjaan: "Backend Engineer-I",
		Alasan:    "Mempelajari Technology Baru Di Company saat ini",
	}
	nim4 := Biodata{
		Nama:      "Juni Dio Kasandra",
		Alamat:    "Jakarta",
		Pekerjaan: "Backend Engineer-I",
		Alasan:    "Mempelajari Technology Baru Di Company saat ini",
	}
	nim5 := Biodata{
		Nama:      "Tasrifin",
		Alamat:    "Jakarta",
		Pekerjaan: "Backend Engineer-I",
		Alasan:    "Mempelajari Technology Baru Di Company saat ini",
	}
	nim6 := Biodata{
		Nama:      "Adhitya Febhiakbar",
		Alamat:    "Jakarta",
		Pekerjaan: "Backend Engineer-I",
		Alasan:    "Mempelajari Technology Baru Di Company saat ini",
	}
	nim7 := Biodata{
		Nama:      "Esra Delima Manurung",
		Alamat:    "Jakarta",
		Pekerjaan: "Backend Engineer-I",
		Alasan:    "Mempelajari Technology Baru Di Company saat ini",
	}
	nim8 := Biodata{
		Nama:      "Muhammad Avtara Khrisna",
		Alamat:    "Jakarta",
		Pekerjaan: "Backend Engineer-I",
		Alasan:    "Mempelajari Technology Baru Di Company saat ini",
	}
	nim9 := Biodata{
		Nama:      "Hamonangan Sitorus",
		Alamat:    "Jakarta",
		Pekerjaan: "Backend Engineer-I",
		Alasan:    "Mempelajari Technology Baru Di Company saat ini",
	}
	nim10 := Biodata{
		Nama:      "Julius Martogi Hamonangan Samosir",
		Alamat:    "Jakarta",
		Pekerjaan: "Backend Engineer-I",
		Alasan:    "Mempelajari Technology Baru Di Company saat ini",
	}
	nim11 := Biodata{
		Nama:      "Indra Bayu Sudirman",
		Alamat:    "Jakarta",
		Pekerjaan: "Backend Engineer-I",
		Alasan:    "Mempelajari Technology Baru Di Company saat ini",
	}
	nim12 := Biodata{
		Nama:      "Phillip Bryan Halim",
		Alamat:    "Jakarta",
		Pekerjaan: "Backend Engineer-I",
		Alasan:    "Mempelajari Technology Baru Di Company saat ini",
	}
	nim13 := Biodata{
		Nama:      "Teguh Ainul Drajat",
		Alamat:    "Jakarta",
		Pekerjaan: "Backend Engineer-I",
		Alasan:    "Mempelajari Technology Baru Di Company saat ini",
	}
	nim14 := Biodata{
		Nama:      "Saut Raja Marihot Tua Sihotang",
		Alamat:    "Jakarta",
		Pekerjaan: "Backend Engineer-I",
		Alasan:    "Mempelajari Technology Baru Di Company saat ini",
	}

	listBiodata := newListBiodata(nim1, nim2, nim3, nim4, nim5, nim6, nim7, nim8, nim9, nim10, nim11, nim12, nim13, nim14)
	fmt.Println("\t\t==========Assignment - 1==========")
	getPresensi := os.Args[1]
	idPresensi, _ := strconv.ParseInt(getPresensi, 10, 64)
	printBiodata(listBiodata[idPresensi-1])
	fmt.Println("\t\tThank You:)")
}
