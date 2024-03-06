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

}