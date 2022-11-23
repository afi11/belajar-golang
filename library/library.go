// NOTED
// fungsi dikatakan public bila nama fungsi diawali dengan huruf kapital, bila huruf kecil berarti private

package library

import "fmt"

type muridPrivate struct {
	name string
	age  int
}

type MuridPublic struct {
	name string
	age  int
}

// fungsi public
func KatakanHallo() {
	fmt.Println("hello")
}

// fungsi private
func introduce(name string) {
	fmt.Println("nama saya", name)
}

// fungsi public
func UcapkanSelamatMalam(name string) {
	introduce(name)
}
