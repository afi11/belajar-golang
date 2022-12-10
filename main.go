package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"project-pertama/library"
	"runtime"
	"strconv"
	"strings"
	"time"

	gubrak "github.com/novalagung/gubrak/v2"
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

	// Method
	// Penerapan dari struct student
	var s4 = student{"john wick", 21}
	s4.sayHello()

	var name = s4.getNameAt(2)
	fmt.Println("nama panggilan :", name)

	var s5 = student{"john wick", 21}
	fmt.Println("s1 before", s5.name)
	// john wick

	s5.changeName1("jason bourne")
	fmt.Println("s5 after changeName1", s5.name)
	// john wick

	s5.changeName2("ethan hunt")
	fmt.Println("s5 after changeName2", s5.name)
	// ethan hunt

	// Property public & private ada folder libary
	library.KatakanHallo()
	library.UcapkanSelamatMalam("Malam Mukidi")

	// Inteface
	var bangunDatar hitung

	bangunDatar = persegi{10.0}
	fmt.Println("===== persegi")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())

	bangunDatar = lingkaran{14.0}
	fmt.Println("===== lingkaran")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())
	fmt.Println("jari-jari :", bangunDatar.(lingkaran).jariJari())

	// Goroutine = mini thread
	runtime.GOMAXPROCS(2)

	go print(5, "hallo")
	print(5, "apa kabar")

	var input string
	fmt.Scanln(&input)

	// Konversi Tipe Data
	// strconv.Atoi() = string ke int
	var str = "1234"
	var num, err = strconv.Atoi(str)

	if err == nil {
		fmt.Println(num)
	}

	// strconv.Itoa() = int ke string
	var num2 = 124
	var str2 = strconv.Itoa(num2)

	fmt.Println(str2)

	// strconv.ParseInt() = string ke int (non desimal)
	var str3 = "1234"
	var num3, err3 = strconv.ParseInt(str3, 10, 64)

	if err3 == nil {
		fmt.Println(num3)
	}

	// strconv.ParseFloat() = string ke float
	var str4 = "24.12"
	var num4, err4 = strconv.ParseFloat(str4, 32)

	if err4 == nil {
		fmt.Println(num4) // 24.1200008392334
	}

	// Konversi Data Menggunakan Teknik Casting
	var a float64 = float64(24)
	fmt.Println(a) // 24

	var b int32 = int32(24.00)
	fmt.Println(b) // 24

	// Encode - Decode Base64
	var contohEnDe = "mukidi sayang kamu"

	var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(contohEnDe)))
	base64.StdEncoding.Encode(encoded, []byte(contohEnDe))
	var encodedString = string(encoded)
	fmt.Println(encodedString)

	var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	var _, err5 = base64.StdEncoding.Decode(decoded, encoded)
	if err5 != nil {
		fmt.Println(err.Error())
	}
	var decodedString = string(decoded)
	fmt.Println(decodedString)

	var contohURL = "https://kalipare.com/"
	var encodedStrigURL = base64.URLEncoding.EncodeToString([]byte(contohURL))
	fmt.Println(encodedStrigURL)

	var decodedByte, _ = base64.URLEncoding.DecodeString(encodedStrigURL)
	var decodedString2 = string(decodedByte)
	fmt.Println(decodedString2)

	// Hash SHA1
	var textHash = "Aku Sayang Kamu"
	var sha = sha1.New()
	sha.Write([]byte(textHash))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)
	fmt.Println(encryptedString)

	// Metode Salting Pada Hash SHA1
	var hashed1, salt1 = doHashUsingSalt(textHash)
	fmt.Printf("hashed 1 : %s\n\n", hashed1)

	var hashed2, salt2 = doHashUsingSalt(textHash)
	fmt.Printf("hashed 2 : %s\n\n", hashed2)

	var hashed3, salt3 = doHashUsingSalt(textHash)
	fmt.Printf("hashed 3 : %s\n\n", hashed3)

	_, _, _ = salt1, salt2, salt3

	// File
	var path = "/home/afi/Documents/test.txt"
	createFile(path)
	writeFile(path)
	readFile(path)
	deleteFile(path)

	fmt.Println(gubrak.RandomInt(10, 20))
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

// contoh method, penerapan dari struct student
func (s student) sayHello() {
	fmt.Println("halo", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

// method pointer
func (s student) changeName1(name string) {
	fmt.Println("---> on changeName1, name changed to", name)
	s.name = name
}

func (s *student) changeName2(name string) {
	fmt.Println("---> on changeName2, name changed to", name)
	s.name = name
}

// Interface
type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

// fungsi menapilkan Goroutine
func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println("Hallo ", message)
	}
}

// fungsi Metode Salting Pada Hash SHA1
func doHashUsingSalt(text string) (string, string) {
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	var saltedText = fmt.Sprintf("text: '%s', salt: %s", text, salt)
	fmt.Println(saltedText)
	var sha = sha1.New()
	sha.Write([]byte(saltedText))
	var encrypted = sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}

// fungsi file
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func createFile(path string) {
	// deteksi apakah file sudah ada
	var _, err = os.Stat(path)

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("==> file berhasil dibuat", path)
}

func writeFile(path string) {
	// buka file dengan level akses READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// tulis data ke file
	_, err = file.WriteString("halo\n")
	if isError(err) {
		return
	}
	_, err = file.WriteString("mari belajar golang\n")
	if isError(err) {
		return
	}

	// simpan perubahan
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("==> file berhasil di isi")
}

func readFile(path string) {
	// buka file
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// baca file
	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if isError(err) {
				break
			}
		}
		if n == 0 {
			break
		}
	}
	if isError(err) {
		return
	}

	fmt.Println("==> file berhasil dibaca")
	fmt.Println(string(text))
}

func deleteFile(path string) {
	var err = os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("==> file berhasil di delete")
}
