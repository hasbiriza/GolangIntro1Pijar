// package main

// import (
// 	"fmt"
// 	"reflect"
// 	"strconv"
// )

// func main() {
// 	var name string = "Aguero"
// 	var age int8 = 30 //Kalau saya ubah int8 maka di inttosring . Itoa error , hanya bisa int saja
// 	var height float32 = 170.5
// 	// var IsMaried bool = true
// 	// var interestInCoding string = "true"

// 	// string to int
// 	stringtoint, _ := strconv.Atoi(name)
// 	// int to string
// 	inttostring := strconv.Itoa(int(age))
// 	// float to string
// 	floattostring := strconv.FormatFloat(float64(height), 'f', 2, 64)
// 	// string to float
// 	stringtofloat, _ := strconv.ParseFloat(name, 64)

// 	// fmt.Println(IsMaried, "tipe data:", reflect.TypeOf(IsMaried))
// 	// fmt.Println(interestInCoding, "tipe data:", reflect.TypeOf(interestInCoding))
// 	fmt.Println(stringtoint, "tipe data:", reflect.TypeOf(stringtoint))
// 	fmt.Println(inttostring, "tipe data:", reflect.TypeOf(inttostring))
// 	fmt.Println(floattostring, "tipe data:", reflect.TypeOf(floattostring))
// 	fmt.Println(stringtofloat, "tipe data:", reflect.TypeOf(stringtofloat))
// }
