package main

import "fmt"

func main() {
	
	// Deklarasi variabel
	numbers := [6]int{1, 2, 3, 4, 5, 6}
	var sum int16

	// Proses looping sum semua
	for i := 0; i < len(numbers); i++ {
		sum += int16(numbers[i])
	}

	// Output / Tampilan
	fmt.Print("Hasil sum: ", sum)
}