package logic_test

import "fmt"

// Leap_year is a function for showing leap years that occurs between two input year
func Leap_year() {
	var fYear, lYear int
	fmt.Println("Masukkan tahun awal: ")
	fmt.Scan(&fYear)
	fmt.Println("Masukkan tahun akhir: ")
	fmt.Scan(&lYear)
	fmt.Println("Output: ")

	if fYear <= lYear {

		for i := fYear + 1; i <= lYear; i++ {
			if i%4 == 0 {
				fmt.Println(i)
			}
		}
	} else {
		fmt.Println("Masukan salah")
	}
}
