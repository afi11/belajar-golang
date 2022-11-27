package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var jsonString = `{"Name": "john wick", "Age": 27}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user :", data.FullName)
	fmt.Println("age  :", data.Age)

	// Decode Array JSON ke Array Objek
	var jsonString2 = `[
		{"Name": "john wick", "Age": 27},
		{"Name": "ethan hunt", "Age": 32}
	]`

	var data2 []User

	var err2 = json.Unmarshal([]byte(jsonString2), &data2)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}

	fmt.Println("user 1:", data2[0].FullName)
	fmt.Println("user 2:", data2[1].FullName)

	// Encode Objek Ke JSON String
	var object = []User{{"john wick", 27}, {"ethan hunt", 32}}
	var jsonData2, err3 = json.Marshal(object)
	if err3 != nil {
		fmt.Println(err3.Error())
		return
	}

	var jsonString3 = string(jsonData2)
	fmt.Println(jsonString3)
}
