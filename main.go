package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

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
	var chicken = map[string]int{
		"januari": 1,
		"febuari": 2,
		"maret":   3,
		"april":   4,
	}

	// tambah item baru
	chicken["mei"] = 5

	// tampilkan
	for key, val := range chicken {
		fmt.Println(key, "\t", val)
	}

	// delete item
	delete(chicken, "mei")
	fmt.Println(len(chicken))

	// Deteksi ada array
	var newchicken = map[string]int{"januari": 50, "februari": 40}
	var value, isExist = newchicken["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	// kombinasi map & slice
	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	// Fungsi
	radomNumber := randomWithRange(10, 20)
	fmt.Println("Angka random ", radomNumber)

	area, circumference := hitungSupayaBenar(15)
	fmt.Println("Area ", area)
	fmt.Println("Circumference ", circumference)

	var orang2 = []string{"wick", "jason", "ethan"}
	var dataContains0 = filter(orang2, func(each string) bool {
		return strings.Contains(each, "o")
	})
	var dataLenght5 = filter(orang2, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli \t\t:", orang2)
	// data asli : [wick jason ethan]

	fmt.Println("filter ada huruf \"o\"\t:", dataContains0)
	// filter ada huruf "o" : [jason]

	fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)
	// filter jumlah huruf "5" : [jason ethan]

	// Penerapan strunct
	// inisaisi di ada dibawah
	var s1 = student{}
	s1.name = "mukidi"
	s1.grade = 18

	fmt.Println("name  :", s1.name)
	fmt.Println("grade :", s1.grade)

	var s2 = student{"ethan", 2}

	var s3 = student{name: "jason"}

	fmt.Println("student 1 :", s1.name)
	fmt.Println("student 2 :", s2.name)
	fmt.Println("student 3 :", s3.name)

}

// contoh fungsi
func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}

// multiple return
func hitungSupayaBenar(d float64) (float64, float64) {
	// hitung luas
	var area = math.Pi * math.Pow(d/2, 2)
	// hitung keliling
	var circumference = math.Pi * d

	// kembalikan 2 nilai
	return area, circumference
}

// fungsi sebagai parameter
func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

// contoh struct
type student struct {
	name  string
	grade int
}
