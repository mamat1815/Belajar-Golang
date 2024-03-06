package main

import (
	"fmt"
	
	"strconv"// untuk casting ke string
)

func main () {
	
	//Var String
	//Var Manual
	var name string
	name = "Afsar Tambawang"
	fmt.Println("Nama Saya : " +name)
	// Var String Otomatis
	alamat := "Suropati"
	fmt.Println("Alamat Lama : " + alamat)
	alamat = "Suropati 6" 
	fmt.Println("Alamat Baru : " + alamat)
	//Var Multiple
	var (
		firstName = "Muhammad"
		lastName = "Afsar"
		address = "Suropati"
		addressNumber = 9

	)
	fmt.Println("Nama Saya " + firstName + " " + lastName)
	fmt.Println("Alamat Tinggal Saya Jalan :" + address +"Nomor " + strconv.Itoa(addressNumber) )
	/*Var merupakan variabel yang bisa dirubah dan const merupakan data yang tidak dapat dirubah*/

	//mencoba const
	const jenisKelamin = "Laki - Laki"
	const tanggalLahir string = "09-08-1973"
	fmt.Println("Jenis Kelamin " + jenisKelamin + "\nTanggal Lahir Saya : " + tanggalLahir)
	//Const dapat dibuat multiple seperti di var
}