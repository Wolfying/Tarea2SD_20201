package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func makeRange(min, max int32) []int32 {
	a := make([]int32, max-min)
	for i := range a {
		a[i] = min + int32(i)
	}
	return a
}

func main() {
	// var number int32 = 1
	// a := makeRange(0, int32(math.Floor(float64(number)/float64(3))))
	// b := makeRange(int32(math.Floor(float64(number)/float64(3))), int32(math.Floor(2*float64(number)/float64(3))))
	// c := makeRange(int32(math.Floor(2*float64(number)/float64(3))), number)
	// fmt.Println(a)
	// fmt.Printf("largo de a es %d \n", len(a))
	// fmt.Println(b)
	// fmt.Printf("largo de b es %d \n", len(b))
	// fmt.Println(c)
	// fmt.Printf("largo de c es %d \n", len(c))

	// var response datanode.DataNodeHandler_UploadFileClient
	// fmt.Println("1) Funcionamiento algoritmo Centralizado")
	// fmt.Println("2) Funcionamiento algoritmo Distribuido")
	var reader *bufio.Reader
	exito := 0
	for exito == 0 {
		fmt.Println("1) Funcionamiento algoritmo Centralizado")
		fmt.Println("2) Funcionamiento algoritmo Distribuido")
		reader = bufio.NewReader(os.Stdin)
		fmt.Printf("Ingrese opcion a utilizar: ")
		opcion, _ := reader.ReadString('\n')
		opcion = strings.TrimSuffix(opcion, "\n")
		opcion = strings.TrimSuffix(opcion, "\r")
		if opcion == "1" {
			// response, err = link.UploadFile(ctx)
			fmt.Println("UNO")
			exito = 1
		} else if opcion == "2" {
			// response, err = link.DistUploadFile(ctx)
			fmt.Println("DOS")
			exito = 1
		} else {
			fmt.Println("Ingrese opcion una opción válida: ")
		}
	}

	// if opcion == "1" {
	// 	// response, err = link.UploadFile(ctx)
	// 	fmt.Printf("UNO")
	// } else if opcion == "2" {
	// 	// response, err = link.DistUploadFile(ctx)
	// 	fmt.Printf("DOS")
	// } else {
	// 	var exito = 0
	// 	for exito == 0 {
	// 		fmt.Printf("Ingrese opcion una opción válida: ")
	// 		opcion, _ = reader.ReadString('\n')
	// 		if opcion == "1" {
	// 			// response, err = link.UploadFile(ctx)
	// 			fmt.Printf("UNO")
	// 			exito = 1
	// 		} else if opcion == "2" {
	// 			// response, err = link.DistUploadFile(ctx)
	// 			fmt.Printf("DOS")
	// 			exito = 1
	// 		}
	// 	}
	// }
}
