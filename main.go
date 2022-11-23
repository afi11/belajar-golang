package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// Variable
	var firstName string = "mukidi"
	var lastName string
	lastName = "forever"
	// deklarasi tanpa tipe data
	status := "beriman"
	// deklarasi multi variable
	first, second, third := 1, 2, 3
	// variable bila tdk digunakan tambahkan underscore
	_ = "anya_geraldine"

	fmt.Println("Hello!\n", firstName, lastName, status)
	fmt.Println("Angka ", first, second, third)

	// Tipe data
	_ = 2.62                  // decimal
	_ = true                  // boolean
	_ = "Nama saya John Wick" // string

	// konstanta
	const phi = 3.14
	// multi konstanta
	const (
		square         = "kotak"
		isToday  bool  = true
		numeric  uint8 = 1
		floatNum       = 2.2
	)

	// Percabangan
	var point = 8
	// IF
	if point >= 8 && point <= 10 {
		fmt.Println("Sangat Bagus")
	} else if point >= 6 && point < 8 {
		fmt.Println("Lumayan")
	} else {
		fmt.Println("Remidi")
	}
	// Switch
	switch point {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}
	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// Perulangan
	// FOR
	for i := 1; i < 10; i++ {
		fmt.Println("Angkan ", i)
	}

	// Break
	j := 1
	for {
		fmt.Println("Angkan ", j)
		j++
		if j == 5 {
			break
		}
	}

	// Break & Continue
	for k := 1; k <= 10; k++ {
		if k%2 == 1 {
			continue
		}

		if k > 8 {
			break
		}

		// Hasil angka genap
		fmt.Println("Angka ", k)
	}

	// Pemanfaatan label
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}

	// Array
	// inisiasi awal array
	var fruits = [4]string{"banana", "grape", "watermelon", "pineaple"}
	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)
	for i := 0; i < len(fruits); i++ {
		fmt.Println("buah :", fruits[i])
	}

	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	// inisiasi tanpa jumlah
	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	// Alokasi array
	var buah = make([]string, 2)
	buah[0] = "manggo"
	buah[1] = "melon"

	fmt.Println(buah)

	// Slice
	var footBallClubs = []string{"arsenal", "man city", "barcelona", "psg"}

	fmt.Println(footBallClubs[1:3])
	// tambah item baru
	var newFootBallClubs = append(footBallClubs, "Inter")
	fmt.Println(newFootBallClubs[1:3:4])

	// copy array
	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	arrayBaru := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src)
	fmt.Println(arrayBaru)

	// map

}
